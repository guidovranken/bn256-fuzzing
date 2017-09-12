#include <stdlib.h>
#include <stdint.h>
#include "bn256_instrumented.h"

extern void runbn256(const uint8_t* data, size_t len);

int LLVMFuzzerTestOneInput(const uint8_t *data, size_t size)
{
    runbn256(data, size);
    GoSlice p1 = {(void*)data, size, size};
    GoInt ret = Fuzz(p1);
    return (int)ret;
}
