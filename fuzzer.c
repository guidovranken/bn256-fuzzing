#include <stdlib.h>
#include <stdio.h>
#include <stdint.h>
#include <string.h>
#include "fuzzer.h"
#include "go/bn256_instrumented.h"

/* If there is sufficient input data:
 *      - set 'slice' to range data..data+size_wanted
 *      - return size_wanted
 * else:
 *      - return 0
 */
static inline size_t make_goslice(
        GoSlice* slice,
        const uint8_t* data,
        size_t size,
        const size_t size_wanted)
{
    if ( size_wanted > size ) {
        /* Insufficient input data */
        return 0;
    }

    slice->data = (void*)data;
    slice->len = size_wanted;
    slice->cap = size_wanted;

    return size_wanted;
}

static inline size_t operation_go(
        const operation_t op,
        const uint8_t* data,
        size_t size,
        uint8_t* output)
{
    size_t bytes_consumed = 0;
    GoSlice p = {};
    GoSlice p_output = {(void*)output, 64, 64};

    switch ( op ) {
        case    BN_ADD:
            if ( (bytes_consumed = make_goslice(&p, data, size, 128)) == 0 ) {
                break;
            }
            GoBNAdd(p, p_output);
            break;
        case    BN_SCALARMUL:
            if ( (bytes_consumed = make_goslice(&p, data, size, 96)) == 0 ) {
                break;
            }
            GoBNScalarMul(p, p_output);
            break;
        case    BN_PAIRING:
            {
                size_t size_wanted;

                if ( size < 1 ) {
                    break;
                }

                size_wanted = 192 * (data[0]+1);
                data++;
                size--;

                if ( (bytes_consumed = make_goslice(&p, data, size, size_wanted)) == 0 ) {
                    break;
                }
                GoBNPairing(p, p_output);
            }
            break;
        default:
            /* Shouldn't happen */
            abort();
            break;
    }

    return bytes_consumed;
}

static inline size_t operation_rust(
        const operation_t op,
        const uint8_t* data,
        size_t size,
        uint8_t* output)
{
    size_t bytes_consumed = 0;

    switch ( op ) {
        case    BN_ADD:
            bytes_consumed = size >= 128 ? 128 : 0;
            if ( bytes_consumed ) {
                rustbnadd(data, 128, output);
                C_alt_bn128_G1_add(data, 128, output);
            }
            break;
        case    BN_SCALARMUL:
            bytes_consumed = size >= 96 ? 96 : 0;
            if ( bytes_consumed ) {
                rustbnmul(data, 96, output);
            }
            break;
        case    BN_PAIRING:
            {
                size_t size_wanted;

                if ( size < 1 ) {
                    break;
                }

                size_wanted = 192 * (data[0]+1);
                data++;
                size--;

                bytes_consumed = size >= size_wanted ? size_wanted : 0;
                if ( bytes_consumed ) {
                    rustbnpairing(data, size_wanted, output);
                }
            }
            break;
        default:
            /* Shouldn't happen */
            abort();
            break;
    }

    return bytes_consumed;
}

static inline size_t operation_cpp(
        const operation_t op,
        const uint8_t* data,
        size_t size,
        uint8_t* output)
{
    size_t bytes_consumed = 0;

    switch ( op ) {
        case    BN_ADD:
            bytes_consumed = size >= 128 ? 128 : 0;
            if ( bytes_consumed ) {
                C_alt_bn128_G1_add(data, 128, output);
            }
            break;
        case    BN_SCALARMUL:
            bytes_consumed = size >= 96 ? 96 : 0;
            if ( bytes_consumed ) {
                C_alt_bn128_G1_mul(data, 96, output);
            }
            break;
        case    BN_PAIRING:
            {
                size_t size_wanted;

                if ( size < 1 ) {
                    break;
                }

                size_wanted = 192 * (data[0]+1);
                data++;
                size--;

                bytes_consumed = size >= size_wanted ? size_wanted : 0;
                if ( bytes_consumed ) {
                    C_alt_bn128_pairing_product(data, size_wanted, output);
                }
            }
            break;
        default:
            /* Shouldn't happen */
            abort();
            break;
    }

    return bytes_consumed;
}

static void print_hex(const uint8_t* data, size_t size)
{
    size_t i;

    for (i = 0; i < size; i += 16) {
        printf(
                "    %02X %02X %02X %02X %02X %02X %02X %02X",
                data[i+0],
                data[i+1],
                data[i+2],
                data[i+3],
                data[i+4],
                data[i+5],
                data[i+6],
                data[i+7]);
        printf(
                "    %02X %02X %02X %02X %02X %02X %02X %02X\n",
                data[i+8],
                data[i+9],
                data[i+10],
                data[i+11],
                data[i+12],
                data[i+13],
                data[i+14],
                data[i+15]);

    }
    printf("\n");
}

static void mismatch(
        const operation_t op,
        const uint8_t* input,
        uint8_t* output_go,
        uint8_t* output_rust,
        uint8_t* output_cpp)
{
    char* op_str;
    size_t op_size;

    switch ( op )
    {
        case    BN_ADD:
            op_str = "ADD";
            op_size = 128;
            break;
        case    BN_SCALARMUL:
            op_str = "SCALARMUL";
            op_size = 96;
            break;
        case    BN_PAIRING:
            op_str = "PAIRING";
            op_size = 192 * input[0];
            input++;
            break;
        default:
            abort();
    }

    printf("Mismatch detected\n\n");
    printf("Operation: %s\n", op_str);
    printf("Input (%zu bytes):\n", op_size);
    print_hex(input, op_size);
    printf("Output (Go):\n");
    print_hex(output_go, 64);
    printf("Output (Rust):\n");
    print_hex(output_rust, 64);
    printf("Output (cpp):\n");
    print_hex(output_cpp, 64);

    fflush(stdout);

    abort();
}

int LLVMFuzzerTestOneInput(const uint8_t *data, size_t size)
{
    operation_t op;
    int go_coverage;
    uint8_t output_go[64], output_rust[64], output_cpp[64];;

    /* Need 1 byte for determining operation, and at least MIN_OP_SIZE for
     * performing the operation */
    if ( size < 1 + MIN_OP_SIZE ) {
        return 0;
    }

    /* Set op to a valid operation */
    op = data[0] % BN_MAX;

    /* One byte has been consumed for the operation */
    data++; size--;

    memset(output_go, 0, sizeof(output_go));
    memset(output_rust, 0, sizeof(output_rust));
    memset(output_cpp, 0, sizeof(output_cpp));

    /* Perform the operation in Go */
    {
        size_t bytes_consumed;

        GoResetCoverage();

        bytes_consumed = operation_go(op, data, size, output_go);

        go_coverage = (int)GoCalcCoverage();

        if ( bytes_consumed == 0 ) {
            goto end;
        }
    }

    /* Perform the operation in Rust */
    {
        size_t bytes_consumed;

        bytes_consumed = operation_rust(op, data, size, output_rust);

        if ( bytes_consumed == 0 ) {
            goto end;
        }
    }

    /* Perform the operation in cpp */
    {
        size_t bytes_consumed;

        bytes_consumed = operation_cpp(op, data, size, output_cpp);

        if ( bytes_consumed == 0 ) {
            goto end;
        }
    }

    if ( memcmp(output_go, output_rust, 64) ) {
        mismatch(op, data, output_go, output_rust, output_cpp);
    }

    if ( memcmp(output_go, output_cpp, 64) ) {
        mismatch(op, data, output_go, output_rust, output_cpp);
    }

    /* At this point it is guaranteed that:
     * output_go == output_rust == output_cpp
     */

end:
    return go_coverage;
}
