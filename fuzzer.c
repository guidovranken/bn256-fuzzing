#include <stdlib.h>
#include <stdio.h>
#include <stdint.h>
#include <string.h>
#include "fuzzer.h"
#include "go/bn256_instrumented.h"

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
        size_t size,
        uint8_t* output_go,
        uint8_t* output_rust,
        uint8_t* output_cpp)
{
    char* op_str;

    switch ( op )
    {
        case    BN_ADD:
            op_str = "ADD";
            break;
        case    BN_SCALARMUL:
            op_str = "SCALARMUL";
            break;
        case    BN_PAIRING:
            op_str = "PAIRING";
            input++;
            break;
        default:
            abort();
    }

    printf("Mismatch detected\n\n");
    printf("Operation: %s\n", op_str);
    printf("Input (%zu bytes):\n", size);
    print_hex(input, size);
    printf("Output (Go):\n");
    print_hex(output_go, 64);
    printf("Output (Rust):\n");
    print_hex(output_rust, 64);
    printf("Output (cpp):\n");
    print_hex(output_cpp, 64);

    fflush(stdout);

    abort();
}

static inline void perform_go(
        operation_t op,
        const uint8_t* data,
        size_t size,
        uint8_t* output)
{
    static go_fn fn[] = {
        GoBNAdd,
        GoBNScalarMul,
        GoBNPairing,
    };

    GoSlice p = {(void*)data, size, size};
    GoSlice p_output = {(void*)output, 64, 64};

    fn[op](p, p_output);
}

static inline void perform_rust(
        operation_t op,
        const uint8_t* data,
        size_t size,
        uint8_t* output)
{
    static rust_fn fn[] = {
        rustbnadd,
        rustbnmul,
        rustbnpairing
    };

    fn[op](data, size, output);
}

static inline void perform_cpp(
        operation_t op,
        const uint8_t* data,
        size_t size,
        uint8_t* output)
{
    static cpp_fn fn[] = {
        rustbnadd,
        rustbnmul,
        rustbnpairing
    };

    fn[op](data, size, output);
}

int LLVMFuzzerTestOneInput(const uint8_t *data, size_t size)
{
    operation_t op;
    int go_coverage = 0;
    size_t input_size;
    uint8_t output_go[64], output_rust[64], output_cpp[64];

    /* Need 1 byte for determining operation, and at least MIN_OP_SIZE for
     * performing the operation */
    if ( size < 1 + MIN_OP_SIZE ) {
        return 0;
    }

    memset(output_go, 0, sizeof(output_go));
    memset(output_rust, 0, sizeof(output_rust));
    memset(output_cpp, 0, sizeof(output_cpp));

    /* Set op to a valid operation */
    op = data[0] % BN_MAX;

    /* One byte has been consumed for the operation */
    data++; size--;

    /* Determine the size required for the operation,
     * and bail out if there is insufficient input data
     */
    switch ( op ) {
        case    BN_ADD:
            input_size = 128;
            break;
        case    BN_SCALARMUL:
            input_size = 96;
            break;
        case    BN_PAIRING:
            /* Consume one byte for multiplier */
            if ( size < 1 ) {
                break;
            }

            input_size = 192*(data[0]+1);

            data++;
            size--;
            break;
        default:
            abort();
    }

    if ( input_size > size ) {
        goto end;
    }

    GoResetCoverage();

    perform_go(op, data, input_size, output_go);
    perform_rust(op, data, input_size, output_rust);
    perform_cpp(op, data, input_size, output_cpp);

    go_coverage = (int)GoCalcCoverage();


    if ( memcmp(output_go, output_rust, 64) ) {
        mismatch(op, data, input_size, output_go, output_rust, output_cpp);
    }

    if ( memcmp(output_go, output_cpp, 64) ) {
        mismatch(op, data, input_size, output_go, output_rust, output_cpp);
    }

    /* At this point it is guaranteed that:
     * output_go == output_rust == output_cpp
     */

end:
    return go_coverage;
}
