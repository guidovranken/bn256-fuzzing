/* Functions implemented in Rust */
extern void rustbnadd(const uint8_t* data, size_t len);
extern void rustbnmul(const uint8_t* data, size_t len);
extern void rustbnpairing(const uint8_t* data, size_t len);

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

