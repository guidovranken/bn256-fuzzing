// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bn256 implements a particular bilinear group at the 128-bit security level.
//
// Bilinear groups are the basis of many of the new cryptographic protocols
// that have been proposed over the past decade. They consist of a triplet of
// groups (G₁, G₂ and GT) such that there exists a function e(g₁ˣ,g₂ʸ)=gTˣʸ
// (where gₓ is a generator of the respective group). That function is called
// a pairing function.
//
// This package specifically implements the Optimal Ate pairing over a 256-bit
// Barreto-Naehrig curve as described in
// http://cryptojedi.org/papers/dclxvi-20100714.pdf. Its output is compatible
// with the implementation described in that paper.
package main

import fuzz_helper "github.com/guidovranken/go-coverage-instrumentation/helper"

import (
	"crypto/rand"
	"io"
	"math/big"
)

// BUG(agl): this implementation is not constant time.
// TODO(agl): keep GF(p²) elements in Mongomery form.

// G1 is an abstract cyclic group. The zero value is suitable for use as the
// output of an operation, but cannot be used as an input.
type G1 struct {
	p *curvePoint
}

// RandomG1 returns x and g₁ˣ where x is a random, non-zero number read from r.
func RandomG1(r io.Reader) (*big.Int, *G1, error) {
	fuzz_helper.AddCoverage(1)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	var k *big.Int
	var err error

	for {
		fuzz_helper.AddCoverage(3)
		k, err = rand.Int(r, Order)
		if err != nil {
			fuzz_helper.AddCoverage(5)
			return nil, nil, err
		} else {
			fuzz_helper.AddCoverage(6)
		}
		fuzz_helper.AddCoverage(4)
		if k.Sign() > 0 {
			fuzz_helper.AddCoverage(7)
			break
		} else {
			fuzz_helper.AddCoverage(8)
		}
	}
	fuzz_helper.AddCoverage(2)

	return k, new(G1).ScalarBaseMult(k), nil
}

func (g *G1) String() string {
	fuzz_helper.AddCoverage(9)
	fuzz_helper.IncrementStack(

	// CurvePoints returns p's curve points in big integer
	)
	defer fuzz_helper.DecrementStack()
	return "bn256.G1" + g.p.String()
}

func (e *G1) CurvePoints() (*big.Int, *big.Int, *big.Int, *big.Int) {
	fuzz_helper.AddCoverage(10)
	fuzz_helper.IncrementStack(

	// ScalarBaseMult sets e to g*k where g is the generator of the group and
	// then returns e.
	)
	defer fuzz_helper.DecrementStack()
	return e.p.x, e.p.y, e.p.z, e.p.t
}

func (e *G1) ScalarBaseMult(k *big.Int) *G1 {
	fuzz_helper.AddCoverage(11)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	if e.p == nil {
		fuzz_helper.AddCoverage(13)
		e.p = newCurvePoint(nil)
	} else {
		fuzz_helper.AddCoverage(14)
	}
	fuzz_helper.AddCoverage(12)
	e.p.Mul(curveGen, k, new(bnPool))
	return e
}

// ScalarMult sets e to a*k and then returns e.
func (e *G1) ScalarMult(a *G1, k *big.Int) *G1 {
	fuzz_helper.AddCoverage(15)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	if e.p == nil {
		fuzz_helper.AddCoverage(17)
		e.p = newCurvePoint(nil)
	} else {
		fuzz_helper.AddCoverage(18)
	}
	fuzz_helper.AddCoverage(16)
	e.p.Mul(a.p, k, new(bnPool))
	return e
}

// Add sets e to a+b and then returns e.
// BUG(agl): this function is not complete: a==b fails.
func (e *G1) Add(a, b *G1) *G1 {
	fuzz_helper.AddCoverage(19)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	if e.p == nil {
		fuzz_helper.AddCoverage(21)
		e.p = newCurvePoint(nil)
	} else {
		fuzz_helper.AddCoverage(22)
	}
	fuzz_helper.AddCoverage(20)
	e.p.Add(a.p, b.p, new(bnPool))
	return e
}

// Neg sets e to -a and then returns e.
func (e *G1) Neg(a *G1) *G1 {
	fuzz_helper.AddCoverage(23)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack(

	// Marshal converts n to a byte slice.
	)
	if e.p == nil {
		fuzz_helper.AddCoverage(25)
		e.p = newCurvePoint(nil)
	} else {
		fuzz_helper.AddCoverage(26)
	}
	fuzz_helper.AddCoverage(24)
	e.p.Negative(a.p)
	return e
}

func (n *G1) Marshal() []byte {
	fuzz_helper.AddCoverage(27)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	n.p.MakeAffine(nil)

	xBytes := new(big.Int).Mod(n.p.x, P).Bytes()
	yBytes := new(big.Int).Mod(n.p.y, P).Bytes()

	// Each value is a 256-bit number.
	const numBytes = 256 / 8

	ret := make([]byte, numBytes*2)
	copy(ret[1*numBytes-len(xBytes):], xBytes)
	copy(ret[2*numBytes-len(yBytes):], yBytes)

	return ret
}

// Unmarshal sets e to the result of converting the output of Marshal back into
// a group element and then returns e.
func (e *G1) Unmarshal(m []byte) (*G1, bool) {
	fuzz_helper.
		// Each value is a 256-bit number.
		AddCoverage(28)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()

	const numBytes = 256 / 8

	if len(m) != 2*numBytes {
		fuzz_helper.AddCoverage(32)
		return nil, false
	} else {
		fuzz_helper.AddCoverage(33)
	}
	fuzz_helper.AddCoverage(29)

	if e.p == nil {
		fuzz_helper.AddCoverage(34)
		e.p = newCurvePoint(nil)
	} else {
		fuzz_helper.AddCoverage(35)
	}
	fuzz_helper.AddCoverage(30)

	e.p.x.SetBytes(m[0*numBytes : 1*numBytes])
	e.p.y.SetBytes(m[1*numBytes : 2*numBytes])

	if e.p.x.Sign() == 0 && e.p.y.Sign() == 0 {
		fuzz_helper.
			// This is the point at infinity.
			AddCoverage(36)

		e.p.y.SetInt64(1)
		e.p.z.SetInt64(0)
		e.p.t.SetInt64(0)
	} else {
		fuzz_helper.AddCoverage(37)
		e.p.z.SetInt64(1)
		e.p.t.SetInt64(1)

		if !e.p.IsOnCurve() {
			fuzz_helper.AddCoverage(38)
			return nil, false
		} else {
			fuzz_helper.AddCoverage(

				// G2 is an abstract cyclic group. The zero value is suitable for use as the
				// output of an operation, but cannot be used as an input.
				39)
		}
	}
	fuzz_helper.AddCoverage(31)

	return e, true
}

type G2 struct {
	p *twistPoint
}

// RandomG1 returns x and g₂ˣ where x is a random, non-zero number read from r.
func RandomG2(r io.Reader) (*big.Int, *G2, error) {
	fuzz_helper.AddCoverage(40)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	var k *big.Int
	var err error

	for {
		fuzz_helper.AddCoverage(42)
		k, err = rand.Int(r, Order)
		if err != nil {
			fuzz_helper.AddCoverage(44)
			return nil, nil, err
		} else {
			fuzz_helper.AddCoverage(45)
		}
		fuzz_helper.AddCoverage(43)
		if k.Sign() > 0 {
			fuzz_helper.AddCoverage(46)
			break
		} else {
			fuzz_helper.AddCoverage(47)
		}
	}
	fuzz_helper.AddCoverage(41)

	return k, new(G2).ScalarBaseMult(k), nil
}

func (g *G2) String() string {
	fuzz_helper.AddCoverage(48)
	fuzz_helper.

		// CurvePoints returns the curve points of p which includes the real
		// and imaginary parts of the curve point.
		IncrementStack()
	defer fuzz_helper.DecrementStack()
	return "bn256.G2" + g.p.String()
}

func (e *G2) CurvePoints() (*gfP2, *gfP2, *gfP2, *gfP2) {
	fuzz_helper.AddCoverage(49)
	fuzz_helper.IncrementStack(

	// ScalarBaseMult sets e to g*k where g is the generator of the group and
	// then returns out.
	)
	defer fuzz_helper.DecrementStack()
	return e.p.x, e.p.y, e.p.z, e.p.t
}

func (e *G2) ScalarBaseMult(k *big.Int) *G2 {
	fuzz_helper.AddCoverage(50)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	if e.p == nil {
		fuzz_helper.AddCoverage(52)
		e.p = newTwistPoint(nil)
	} else {
		fuzz_helper.AddCoverage(53)
	}
	fuzz_helper.AddCoverage(51)
	e.p.Mul(twistGen, k, new(bnPool))
	return e
}

// ScalarMult sets e to a*k and then returns e.
func (e *G2) ScalarMult(a *G2, k *big.Int) *G2 {
	fuzz_helper.AddCoverage(54)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	if e.p == nil {
		fuzz_helper.AddCoverage(56)
		e.p = newTwistPoint(nil)
	} else {
		fuzz_helper.AddCoverage(57)
	}
	fuzz_helper.AddCoverage(55)
	e.p.Mul(a.p, k, new(bnPool))
	return e
}

// Add sets e to a+b and then returns e.
// BUG(agl): this function is not complete: a==b fails.
func (e *G2) Add(a, b *G2) *G2 {
	fuzz_helper.AddCoverage(58)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	if e.p == nil {
		fuzz_helper.AddCoverage(60)
		e.p = newTwistPoint(nil)
	} else {
		fuzz_helper.AddCoverage(61)
	}
	fuzz_helper.AddCoverage(59)
	e.p.Add(a.p, b.p, new(bnPool))
	return e
}

// Marshal converts n into a byte slice.
func (n *G2) Marshal() []byte {
	fuzz_helper.AddCoverage(62)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	n.p.MakeAffine(nil)

	xxBytes := new(big.Int).Mod(n.p.x.x, P).Bytes()
	xyBytes := new(big.Int).Mod(n.p.x.y, P).Bytes()
	yxBytes := new(big.Int).Mod(n.p.y.x, P).Bytes()
	yyBytes := new(big.Int).Mod(n.p.y.y, P).Bytes()

	// Each value is a 256-bit number.
	const numBytes = 256 / 8

	ret := make([]byte, numBytes*4)
	copy(ret[1*numBytes-len(xxBytes):], xxBytes)
	copy(ret[2*numBytes-len(xyBytes):], xyBytes)
	copy(ret[3*numBytes-len(yxBytes):], yxBytes)
	copy(ret[4*numBytes-len(yyBytes):], yyBytes)

	return ret
}

// Unmarshal sets e to the result of converting the output of Marshal back into
// a group element and then returns e.
func (e *G2) Unmarshal(m []byte) (*G2, bool) {
	fuzz_helper.
		// Each value is a 256-bit number.
		AddCoverage(63)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()

	const numBytes = 256 / 8

	if len(m) != 4*numBytes {
		fuzz_helper.AddCoverage(67)
		return nil, false
	} else {
		fuzz_helper.AddCoverage(68)
	}
	fuzz_helper.AddCoverage(64)

	if e.p == nil {
		fuzz_helper.AddCoverage(69)
		e.p = newTwistPoint(nil)
	} else {
		fuzz_helper.AddCoverage(70)
	}
	fuzz_helper.AddCoverage(65)

	e.p.x.x.SetBytes(m[0*numBytes : 1*numBytes])
	e.p.x.y.SetBytes(m[1*numBytes : 2*numBytes])
	e.p.y.x.SetBytes(m[2*numBytes : 3*numBytes])
	e.p.y.y.SetBytes(m[3*numBytes : 4*numBytes])

	if e.p.x.x.Sign() == 0 &&
		e.p.x.y.Sign() == 0 &&
		e.p.y.x.Sign() == 0 &&
		e.p.y.y.Sign() == 0 {
		fuzz_helper.
			// This is the point at infinity.
			AddCoverage(71)

		e.p.y.SetOne()
		e.p.z.SetZero()
		e.p.t.SetZero()
	} else {
		fuzz_helper.AddCoverage(72)
		e.p.z.SetOne()
		e.p.t.SetOne()

		if !e.p.IsOnCurve() {
			fuzz_helper.AddCoverage(73)
			return nil, false
		} else {
			fuzz_helper.AddCoverage(

				// GT is an abstract cyclic group. The zero value is suitable for use as the
				// output of an operation, but cannot be used as an input.
				74)
		}
	}
	fuzz_helper.AddCoverage(66)

	return e, true
}

type GT struct {
	p *gfP12
}

func (g *GT) String() string {
	fuzz_helper.AddCoverage(75)
	fuzz_helper.

		// ScalarMult sets e to a*k and then returns e.
		IncrementStack()
	defer fuzz_helper.DecrementStack()
	return "bn256.GT" + g.p.String()
}

func (e *GT) ScalarMult(a *GT, k *big.Int) *GT {
	fuzz_helper.AddCoverage(76)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack(

	// Add sets e to a+b and then returns e.
	)
	if e.p == nil {
		fuzz_helper.AddCoverage(78)
		e.p = newGFp12(nil)
	} else {
		fuzz_helper.AddCoverage(79)
	}
	fuzz_helper.AddCoverage(77)
	e.p.Exp(a.p, k, new(bnPool))
	return e
}

func (e *GT) Add(a, b *GT) *GT {
	fuzz_helper.AddCoverage(80)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	if e.p == nil {
		fuzz_helper.AddCoverage(82)
		e.p = newGFp12(nil)
	} else {
		fuzz_helper.AddCoverage(83)
	}
	fuzz_helper.AddCoverage(81)
	e.p.Mul(a.p, b.p, new(bnPool))
	return e
}

// Neg sets e to -a and then returns e.
func (e *GT) Neg(a *GT) *GT {
	fuzz_helper.AddCoverage(84)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack(

	// Marshal converts n into a byte slice.
	)
	if e.p == nil {
		fuzz_helper.AddCoverage(86)
		e.p = newGFp12(nil)
	} else {
		fuzz_helper.AddCoverage(87)
	}
	fuzz_helper.AddCoverage(85)
	e.p.Invert(a.p, new(bnPool))
	return e
}

func (n *GT) Marshal() []byte {
	fuzz_helper.AddCoverage(88)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	n.p.Minimal()

	xxxBytes := n.p.x.x.x.Bytes()
	xxyBytes := n.p.x.x.y.Bytes()
	xyxBytes := n.p.x.y.x.Bytes()
	xyyBytes := n.p.x.y.y.Bytes()
	xzxBytes := n.p.x.z.x.Bytes()
	xzyBytes := n.p.x.z.y.Bytes()
	yxxBytes := n.p.y.x.x.Bytes()
	yxyBytes := n.p.y.x.y.Bytes()
	yyxBytes := n.p.y.y.x.Bytes()
	yyyBytes := n.p.y.y.y.Bytes()
	yzxBytes := n.p.y.z.x.Bytes()
	yzyBytes := n.p.y.z.y.Bytes()

	// Each value is a 256-bit number.
	const numBytes = 256 / 8

	ret := make([]byte, numBytes*12)
	copy(ret[1*numBytes-len(xxxBytes):], xxxBytes)
	copy(ret[2*numBytes-len(xxyBytes):], xxyBytes)
	copy(ret[3*numBytes-len(xyxBytes):], xyxBytes)
	copy(ret[4*numBytes-len(xyyBytes):], xyyBytes)
	copy(ret[5*numBytes-len(xzxBytes):], xzxBytes)
	copy(ret[6*numBytes-len(xzyBytes):], xzyBytes)
	copy(ret[7*numBytes-len(yxxBytes):], yxxBytes)
	copy(ret[8*numBytes-len(yxyBytes):], yxyBytes)
	copy(ret[9*numBytes-len(yyxBytes):], yyxBytes)
	copy(ret[10*numBytes-len(yyyBytes):], yyyBytes)
	copy(ret[11*numBytes-len(yzxBytes):], yzxBytes)
	copy(ret[12*numBytes-len(yzyBytes):], yzyBytes)

	return ret
}

// Unmarshal sets e to the result of converting the output of Marshal back into
// a group element and then returns e.
func (e *GT) Unmarshal(m []byte) (*GT, bool) {
	fuzz_helper.
		// Each value is a 256-bit number.
		AddCoverage(89)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()

	const numBytes = 256 / 8

	if len(m) != 12*numBytes {
		fuzz_helper.AddCoverage(92)
		return nil, false
	} else {
		fuzz_helper.AddCoverage(93)
	}
	fuzz_helper.AddCoverage(90)

	if e.p == nil {
		fuzz_helper.AddCoverage(94)
		e.p = newGFp12(nil)
	} else {
		fuzz_helper.AddCoverage(95)
	}
	fuzz_helper.AddCoverage(91)

	e.p.x.x.x.SetBytes(m[0*numBytes : 1*numBytes])
	e.p.x.x.y.SetBytes(m[1*numBytes : 2*numBytes])
	e.p.x.y.x.SetBytes(m[2*numBytes : 3*numBytes])
	e.p.x.y.y.SetBytes(m[3*numBytes : 4*numBytes])
	e.p.x.z.x.SetBytes(m[4*numBytes : 5*numBytes])
	e.p.x.z.y.SetBytes(m[5*numBytes : 6*numBytes])
	e.p.y.x.x.SetBytes(m[6*numBytes : 7*numBytes])
	e.p.y.x.y.SetBytes(m[7*numBytes : 8*numBytes])
	e.p.y.y.x.SetBytes(m[8*numBytes : 9*numBytes])
	e.p.y.y.y.SetBytes(m[9*numBytes : 10*numBytes])
	e.p.y.z.x.SetBytes(m[10*numBytes : 11*numBytes])
	e.p.y.z.y.SetBytes(m[11*numBytes : 12*numBytes])

	return e, true
}

// Pair calculates an Optimal Ate pairing.
func Pair(g1 *G1, g2 *G2) *GT {
	fuzz_helper.AddCoverage(96)
	fuzz_helper.IncrementStack(

	// PairingCheck calculates the Optimal Ate pairing for a set of points.
	)
	defer fuzz_helper.DecrementStack()
	return &GT{optimalAte(g2.p, g1.p, new(bnPool))}
}

func PairingCheck(a []*G1, b []*G2) bool {
	fuzz_helper.AddCoverage(97)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	pool := new(bnPool)

	acc := newGFp12(pool)
	acc.SetOne()

	for i := 0; i < len(a); i++ {
		fuzz_helper.AddCoverage(99)
		if a[i].p.IsInfinity() || b[i].p.IsInfinity() {
			fuzz_helper.AddCoverage(101)
			continue
		} else {
			fuzz_helper.AddCoverage(102)
		}
		fuzz_helper.AddCoverage(100)
		acc.Mul(acc, miller(b[i].p, a[i].p, pool), pool)
	}
	fuzz_helper.AddCoverage(98)
	ret := finalExponentiation(acc, pool)
	acc.Put(pool)

	return ret.IsOne()
}

// bnPool implements a tiny cache of *big.Int objects that's used to reduce the
// number of allocations made during processing.
type bnPool struct {
	bns   []*big.Int
	count int
}

func (pool *bnPool) Get() *big.Int {
	fuzz_helper.AddCoverage(103)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	if pool == nil {
		fuzz_helper.AddCoverage(106)
		return new(big.Int)
	} else {
		fuzz_helper.AddCoverage(107)
	}
	fuzz_helper.AddCoverage(104)

	pool.count++
	l := len(pool.bns)
	if l == 0 {
		fuzz_helper.AddCoverage(108)
		return new(big.Int)
	} else {
		fuzz_helper.AddCoverage(109)
	}
	fuzz_helper.AddCoverage(105)

	bn := pool.bns[l-1]
	pool.bns = pool.bns[:l-1]
	return bn
}

func (pool *bnPool) Put(bn *big.Int) {
	fuzz_helper.AddCoverage(110)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	if pool == nil {
		fuzz_helper.AddCoverage(112)
		return
	} else {
		fuzz_helper.AddCoverage(113)
	}
	fuzz_helper.AddCoverage(111)
	pool.bns = append(pool.bns, bn)
	pool.count--
}

func (pool *bnPool) Count() int {
	fuzz_helper.AddCoverage(114)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	return pool.count
}

var _ = fuzz_helper.AddCoverage
