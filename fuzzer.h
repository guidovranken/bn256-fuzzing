#include <stdlib.h>
#include <stdint.h>
#include "go/bn256_instrumented.h"

/* Functions implemented in Rust */
extern void rustbnadd(const uint8_t* data, size_t len, uint8_t* output);
extern void rustbnmul(const uint8_t* data, size_t len, uint8_t* output);
extern void rustbnpairing(const uint8_t* data, size_t len, uint8_t* output);

/* Functions implemented in cpp */
void C_alt_bn128_pairing_product(const uint8_t* data, size_t size, uint8_t* output);
void C_alt_bn128_G1_add(const uint8_t* data, size_t size, uint8_t* output);
void C_alt_bn128_G1_mul(const uint8_t* data, size_t size, uint8_t* output);

typedef enum {
    BN_ADD = 0,
    BN_SCALARMUL = 1,
    BN_PAIRING = 2,
    BN_MAX = 3
} operation_t;

typedef void (*go_fn)(
        GoSlice input,
        GoSlice output);
typedef void (*rust_fn)(
        const uint8_t* data,
        size_t len,
        uint8_t* output);
typedef void (*cpp_fn)(
        const uint8_t* data,
        size_t len,
        uint8_t* output);
