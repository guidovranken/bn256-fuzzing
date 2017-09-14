#include <stdlib.h>
#include <stdint.h>
#include "bn256_instrumented.h"

extern void runbn256(const uint8_t* data, size_t len);

/* The smallest operation in terms of data consumption is BN_SCALARMUL, which
 * needs 96 bytes of input
 */
#define MIN_OP_SIZE 96

typedef enum {
    BN_ADD = 0,
    BN_SCALARMUL = 1,
    BN_PAIRING = 2,
    BN_MAX = 3
} operation_t;

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
        size_t size)
{
    size_t bytes_consumed = 0;
    GoSlice p = {};

    switch ( op ) {
        case    BN_ADD:
            if ( (bytes_consumed = make_goslice(&p, data, size, 128)) == 0 ) {
                break;
            }
            GoBNAdd(p);
            break;
        case    BN_SCALARMUL:
            if ( (bytes_consumed = make_goslice(&p, data, size, 96)) == 0 ) {
                break;
            }
            GoBNScalarMul(p);
            break;
        case    BN_PAIRING:
            /* TODO any multiple of 192 should be possible */
            if ( (bytes_consumed = make_goslice(&p, data, size, 192)) == 0 ) {
                break;
            }
            GoBNPairing(p);
            break;
        default:
            /* Shouldn't happen */
            abort();
            break;
    }

    return bytes_consumed;
}

static inline size_t operation_rust(const operation_t op, const uint8_t* data, size_t size)
{
    size_t bytes_consumed = 0;

    switch ( op ) {
        case    BN_ADD:
            /* TODO */
            break;
        case    BN_SCALARMUL:
            /* TODO */
            break;
        case    BN_PAIRING:
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

    /* Need 1 byte for determining operation, and at least MIN_OP_SIZE for
     * performing the operation */
    if ( size < 1 + MIN_OP_SIZE ) {
        return 0;
    }

    /* Set op to a valid operation */
    op = data[0] % BN_MAX;

    /* One byte has been consumed for the operation */
    data++; size--;

    /* Perform the operation in Go */
    {
        size_t bytes_consumed;

        GoResetCoverage();

        bytes_consumed = operation_go(op, data, size);

        go_coverage = (int)GoCalcCoverage();
    }

    return go_coverage;
}
