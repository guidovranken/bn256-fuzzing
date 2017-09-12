package main

import "C"

import (
    "crypto/rand"
)

const (
	CoverSize       = 64 << 10
)

var CoverTab    [CoverSize]byte

func resetCoverage() {
    for i := 0; i < CoverSize; i++ {
        CoverTab[i] = 0
    }
}

func calcCoverage() int {
    coverage := 0

    for i := 0; i < CoverSize; i++ {
        if CoverTab[i] != 0 {
            coverage += 1
        }
	}

    return coverage
}

//export Fuzz
func Fuzz(data []byte) int {
    resetCoverage()
    FuzzInner(data)
    return calcCoverage()
}

func FuzzInner(data []byte) {
    a, _ := rand.Int(rand.Reader, Order)
    //b, _ := rand.Int(rand.Reader, Order)
    //c, _ := rand.Int(rand.Reader, Order)

    // Then each party calculates g₁ and g₂ times their private value.
    pa := new(G1).ScalarBaseMult(a)
    _ = pa
/*
    qa := new(G2).ScalarBaseMult(a)

    pb := new(G1).ScalarBaseMult(b)
    qb := new(G2).ScalarBaseMult(b)

    pc := new(G1).ScalarBaseMult(c)
    qc := new(G2).ScalarBaseMult(c)

    // Now each party exchanges its public values with the other two and
    // all parties can calculate the shared key.
    k1 := Pair(pb, qc)
    k1.ScalarMult(k1, a)

    k2 := Pair(pc, qa)
    k2.ScalarMult(k2, b)

    k3 := Pair(pa, qb)
    k3.ScalarMult(k3, c)
*/
}

func main() { }
