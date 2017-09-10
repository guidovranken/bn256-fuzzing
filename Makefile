all: bnfuzzer/target/release fuzzer
bnfuzzer/target/release:
	cd bnfuzzer; cargo rustc --release -- -Cpasses=sancov -Cllvm-args="-sanitizer-coverage-level=3  -sanitizer-coverage-trace-pc-guard" --crate-type=staticlib
fuzzer:
	clang fuzzer.c -c -o fuzzer.o
	clang++ fuzzer.o libFuzzer.a bnfuzzer/target/release/deps/libbnfuzzer*.a -ldl -lpthread -o fuzzer
clean:
	rm -rf fuzzer fuzzer.o
	cd bnfuzzer; cargo clean
