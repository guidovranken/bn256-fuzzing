# bn256-fuzzing

## Instructions

Be sure to have clang installed.

Install Rust nightly:

```
curl -s https://static.rust-lang.org/rustup.sh | sh -s -- --channel=nightly
```

Put libFuzzer.a in the root directory of this project. It can be obtained and created by:

```
svn co http://llvm.org/svn/llvm-project/llvm/trunk/lib/Fuzzer
cd Fuzzer
clang++ -c -g -O2 -std=c++11 *.cpp
ar r libFuzzer.a *.o
ranlib libFuzzer.a
```

Now type ```make```.

The fuzzer can be started with:

```./fuzzer```
