#include <stdlib.h>
#include <stdint.h>
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
            /* TODO any multiple of 192 should be possible */
            if ( (bytes_consumed = make_goslice(&p, data, size, 192)) == 0 ) {
                break;
            }
            GoBNPairing(p, p_output);
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
            }
            /* TODO */
            break;
        case    BN_SCALARMUL:
            bytes_consumed = size >= 96 ? 96 : 0;
            if ( bytes_consumed ) {
                rustbnadd(data, 96, output);
            }
            /* TODO */
            break;
        case    BN_PAIRING:
            bytes_consumed = size >= 192 ? 192 : 0;
            if ( bytes_consumed ) {
                rustbnadd(data, 192, output);
            }
            /* TODO */
            break;
        default:
            /* Shouldn't happen */
            abort();
            break;
    }

    return bytes_consumed;
}

int LLVMFuzzerTestOneInput(const uint8_t *data, size_t size)
{
    operation_t op;
    int go_coverage;
    uint8_t output_go[64], output_rust[64];

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

end:
    return go_coverage;
}
