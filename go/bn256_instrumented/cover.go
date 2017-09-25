package main

import "C"

import fuzz_helper "github.com/guidovranken/go-coverage-instrumentation/helper"

//export GoResetCoverage
func GoResetCoverage() {
    fuzz_helper.ResetCoverage()
}

//export GoCalcCoverage
func GoCalcCoverage() int {
    return fuzz_helper.CalcCoverage()
}

//export GoBNAdd
func GoBNAdd(input []byte, output []byte) {
    A := bn256Add{}
    res, err := A.Run(input)

    if err == nil {
        copy(output, res)
    }
    /* Silence compiler */
    _, _ = res, err
}

//export GoBNScalarMul
func GoBNScalarMul(input []byte, output []byte) {
    M := bn256ScalarMul{}
    res, err := M.Run(input)

    if err == nil {
        copy(output, res)
    }
    /* Silence compiler */
    _, _ = res, err
}

//export GoBNPairing
func GoBNPairing(input []byte, output []byte) {
    P := bn256Pairing{}
    res, err := P.Run(input)

    if err == nil {
        copy(output, res)
    }
    /* Silence compiler */
    _, _ = res, err
}

func main() { }
