CFLAGS=-fsanitize-coverage=trace-pc-guard -g
CXXFLAGS=-fsanitize-coverage=trace-pc-guard -g
all: rust/bnfuzzer/target/release fuzzer
cpp/build/libdevcrypto/libdevcrypto.a:
	cd cpp; mkdir -p build; cd build; cmake ../cpp-ethereum; cmake --build libdevcrypto
rust/bnfuzzer/target/release:
	cd rust/bnfuzzer; cargo rustc --release -- -Cpasses=sancov -Cllvm-args="-sanitizer-coverage-level=3  -sanitizer-coverage-trace-pc-guard" --crate-type=staticlib
go/bn256_instrumented.a:
	cd go; go build -buildmode=c-archive ./bn256_instrumented/
libfuzzer-gv/libFuzzer.a:
	cd libfuzzer-gv; make -j4
fuzzer: fuzzer.c libfuzzer-gv/libFuzzer.a go/bn256_instrumented.a cpp/build/libdevcrypto/libdevcrypto.a
	clang -fsanitize-coverage=trace-pc-guard fuzzer.c -c -o fuzzer.o
	clang++ fuzzer.o libfuzzer-gv/libFuzzer.a rust/bnfuzzer/target/release/deps/libbnfuzzer*.a go/bn256_instrumented.a  cpp/build/libdevcrypto/libdevcrypto.a ./cpp/build/deps/lib/libff.a ./cpp/build/deps/lib/libmpir.a -lboost_system -ldl -lpthread -o fuzzer
clean:
	rm -rf fuzzer fuzzer.o go/bn256_instrumented.a go/bn256_instrumented.h
	cd rust/bnfuzzer; cargo clean
	rm -rf cpp/build
