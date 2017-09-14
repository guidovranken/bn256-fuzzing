package main

import "C"

const (
	CoverSize       = 64 << 10
)

var CoverTab    [CoverSize]byte

//export GoResetCoverage
func GoResetCoverage() {
    for i := 0; i < CoverSize; i++ {
        CoverTab[i] = 0
    }
}

//export GoCalcCoverage
func GoCalcCoverage() int {
    coverage := 0

    for i := 0; i < CoverSize; i++ {
        if CoverTab[i] != 0 {
            coverage += 1
        }
	}

    return coverage
}

//export GoBNAdd
func GoBNAdd(input []byte) {
    A := bn256Add{}
    res, err := A.Run(input)

    /* Silence compiler */
    _, _ = res, err
}

//export GoBNScalarMul
func GoBNScalarMul(input []byte) {
    M := bn256ScalarMul{}
    res, err := M.Run(input)

    /* Silence compiler */
    _, _ = res, err
}

//export GoBNPairing
func GoBNPairing(input []byte) {
    P := bn256Pairing{}
    res, err := P.Run(input)

    /* Silence compiler */
    _, _ = res, err
}

func main() { }
