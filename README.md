# bn256-fuzzing

## Instructions

Clone this repository with --recursive to retrieve the submodules as well.

Be sure to have clang and a recent version of Go installed.

Install Rust nightly:

```
curl -s https://static.rust-lang.org/rustup.sh | sh -s -- --channel=nightly
```

Now type ```make```.

The fuzzer can be started with:

```./fuzzer -custom_guided=1```
