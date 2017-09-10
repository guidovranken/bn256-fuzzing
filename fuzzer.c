#include <stdlib.h>
#include <stdint.h>

extern void runbn256(const uint8_t* data, size_t len);

int LLVMFuzzerTestOneInput(const uint8_t *data, size_t size)
{
    runbn256(data, size);
    return 0;
}
