# bn256-fuzzing

## Instructions

Clone this repository with --recursive to retrieve the submodules as well.

Be sure to have clang (>= version 4) and a recent version of Go installed.

```
export CC=/path/to/clang
export CXX=/path/to/clang++
```

Install Rust nightly:

```
curl -s https://static.rust-lang.org/rustup.sh | sh -s -- --channel=nightly
```

Now type ```make```.

The fuzzer can be started with:

```./fuzzer -custom_guided=1```

If you want to generate an initial corpus of inputs that comprise valid input to the BN functions, do:

```
mkdir corpus
python gen_corpus.py corpus
```

The fuzzer can now be run on this corpus with:

```
./fuzzer -custom_guided=1 corpus
```
