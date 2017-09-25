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

// G1 is an abstract cyclic group. The zero value is suitable for use as the
// output of an operation, but cannot be used as an input.
type G1 struct {
	p *curvePoint
}

// RandomG1 returns x and g₁ˣ where x is a random, non-zero number read from r.
func RandomG1(r io.Reader) (*big.Int, *G1, error) {
	fuzz_helper.CoverTab[22588]++
	var k *big.Int
	var err error

	for {
		fuzz_helper.CoverTab[5262]++
		k, err = rand.Int(r, Order)
		if err != nil {
			fuzz_helper.CoverTab[45021]++
			return nil, nil, err
		} else {
			fuzz_helper.CoverTab[39040]++
		}
		fuzz_helper.CoverTab[17878]++
		if k.Sign() > 0 {
			fuzz_helper.CoverTab[2095]++
			break
		} else {
			fuzz_helper.CoverTab[21668]++
		}
	}
	fuzz_helper.CoverTab[44810]++

	return k, new(G1).ScalarBaseMult(k), nil
}

func (g *G1) String() string {
	fuzz_helper.CoverTab[45213]++
	return "bn256.G1" + g.p.String()
}

// CurvePoints returns p's curve points in big integer
func (e *G1) CurvePoints() (*big.Int, *big.Int, *big.Int, *big.Int) {
	fuzz_helper.CoverTab[16619]++
	return e.p.x, e.p.y, e.p.z, e.p.t
}

// ScalarBaseMult sets e to g*k where g is the generator of the group and
// then returns e.
func (e *G1) ScalarBaseMult(k *big.Int) *G1 {
	fuzz_helper.CoverTab[12692]++
	if e.p == nil {
		fuzz_helper.CoverTab[6577]++
		e.p = newCurvePoint(nil)
	} else {
		fuzz_helper.CoverTab[17393]++
	}
	fuzz_helper.CoverTab[42483]++
	e.p.Mul(curveGen, k, new(bnPool))
	return e
}

// ScalarMult sets e to a*k and then returns e.
func (e *G1) ScalarMult(a *G1, k *big.Int) *G1 {
	fuzz_helper.CoverTab[64174]++
	if e.p == nil {
		fuzz_helper.CoverTab[35657]++
		e.p = newCurvePoint(nil)
	} else {
		fuzz_helper.CoverTab[30358]++
	}
	fuzz_helper.CoverTab[38740]++
	e.p.Mul(a.p, k, new(bnPool))
	return e
}

// Add sets e to a+b and then returns e.
// BUG(agl): this function is not complete: a==b fails.
func (e *G1) Add(a, b *G1) *G1 {
	fuzz_helper.CoverTab[23294]++
	if e.p == nil {
		fuzz_helper.CoverTab[11162]++
		e.p = newCurvePoint(nil)
	} else {
		fuzz_helper.CoverTab[49217]++
	}
	fuzz_helper.CoverTab[61639]++
	e.p.Add(a.p, b.p, new(bnPool))
	return e
}

// Neg sets e to -a and then returns e.
func (e *G1) Neg(a *G1) *G1 {
	fuzz_helper.CoverTab[34511]++
	if e.p == nil {
		fuzz_helper.CoverTab[28614]++
		e.p = newCurvePoint(nil)
	} else {
		fuzz_helper.CoverTab[39226]++
	}
	fuzz_helper.CoverTab[64074]++
	e.p.Negative(a.p)
	return e
}

// Marshal converts n to a byte slice.
func (n *G1) Marshal() []byte {
	fuzz_helper.CoverTab[2297]++
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
	fuzz_helper.CoverTab[40870]++
	// Each value is a 256-bit number.
	const numBytes = 256 / 8

	if len(m) != 2*numBytes {
		fuzz_helper.CoverTab[15638]++
		return nil, false
	} else {
		fuzz_helper.CoverTab[45869]++
	}
	fuzz_helper.CoverTab[52877]++

	if e.p == nil {
		fuzz_helper.CoverTab[23368]++
		e.p = newCurvePoint(nil)
	} else {
		fuzz_helper.CoverTab[12901]++
	}
	fuzz_helper.CoverTab[778]++

	e.p.x.SetBytes(m[0*numBytes : 1*numBytes])
	e.p.y.SetBytes(m[1*numBytes : 2*numBytes])

	if e.p.x.Sign() == 0 && e.p.y.Sign() == 0 {
		fuzz_helper.CoverTab[12499]++

		e.p.y.SetInt64(1)
		e.p.z.SetInt64(0)
		e.p.t.SetInt64(0)
	} else {
		fuzz_helper.CoverTab[42993]++
		e.p.z.SetInt64(1)
		e.p.t.SetInt64(1)

		if !e.p.IsOnCurve() {
			fuzz_helper.CoverTab[30301]++
			return nil, false
		} else {
			fuzz_helper.CoverTab[45210]++
		}
	}
	fuzz_helper.CoverTab[33340]++

	return e, true
}

// G2 is an abstract cyclic group. The zero value is suitable for use as the
// output of an operation, but cannot be used as an input.
type G2 struct {
	p *twistPoint
}

// RandomG1 returns x and g₂ˣ where x is a random, non-zero number read from r.
func RandomG2(r io.Reader) (*big.Int, *G2, error) {
	fuzz_helper.CoverTab[264]++
	var k *big.Int
	var err error

	for {
		fuzz_helper.CoverTab[47636]++
		k, err = rand.Int(r, Order)
		if err != nil {
			fuzz_helper.CoverTab[20539]++
			return nil, nil, err
		} else {
			fuzz_helper.CoverTab[63931]++
		}
		fuzz_helper.CoverTab[8730]++
		if k.Sign() > 0 {
			fuzz_helper.CoverTab[19009]++
			break
		} else {
			fuzz_helper.CoverTab[64748]++
		}
	}
	fuzz_helper.CoverTab[3566]++

	return k, new(G2).ScalarBaseMult(k), nil
}

func (g *G2) String() string {
	fuzz_helper.CoverTab[50446]++
	return "bn256.G2" + g.p.String()
}

// CurvePoints returns the curve points of p which includes the real
// and imaginary parts of the curve point.
func (e *G2) CurvePoints() (*gfP2, *gfP2, *gfP2, *gfP2) {
	fuzz_helper.CoverTab[18500]++
	return e.p.x, e.p.y, e.p.z, e.p.t
}

// ScalarBaseMult sets e to g*k where g is the generator of the group and
// then returns out.
func (e *G2) ScalarBaseMult(k *big.Int) *G2 {
	fuzz_helper.CoverTab[52152]++
	if e.p == nil {
		fuzz_helper.CoverTab[9670]++
		e.p = newTwistPoint(nil)
	} else {
		fuzz_helper.CoverTab[55848]++
	}
	fuzz_helper.CoverTab[17111]++
	e.p.Mul(twistGen, k, new(bnPool))
	return e
}

// ScalarMult sets e to a*k and then returns e.
func (e *G2) ScalarMult(a *G2, k *big.Int) *G2 {
	fuzz_helper.CoverTab[50755]++
	if e.p == nil {
		fuzz_helper.CoverTab[64631]++
		e.p = newTwistPoint(nil)
	} else {
		fuzz_helper.CoverTab[15513]++
	}
	fuzz_helper.CoverTab[912]++
	e.p.Mul(a.p, k, new(bnPool))
	return e
}

// Add sets e to a+b and then returns e.
// BUG(agl): this function is not complete: a==b fails.
func (e *G2) Add(a, b *G2) *G2 {
	fuzz_helper.CoverTab[17300]++
	if e.p == nil {
		fuzz_helper.CoverTab[40937]++
		e.p = newTwistPoint(nil)
	} else {
		fuzz_helper.CoverTab[33825]++
	}
	fuzz_helper.CoverTab[16403]++
	e.p.Add(a.p, b.p, new(bnPool))
	return e
}

// Marshal converts n into a byte slice.
func (n *G2) Marshal() []byte {
	fuzz_helper.CoverTab[7237]++
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
	fuzz_helper.CoverTab[23248]++
	// Each value is a 256-bit number.
	const numBytes = 256 / 8

	if len(m) != 4*numBytes {
		fuzz_helper.CoverTab[23245]++
		return nil, false
	} else {
		fuzz_helper.CoverTab[5383]++
	}
	fuzz_helper.CoverTab[52715]++

	if e.p == nil {
		fuzz_helper.CoverTab[52957]++
		e.p = newTwistPoint(nil)
	} else {
		fuzz_helper.CoverTab[6211]++
	}
	fuzz_helper.CoverTab[11389]++

	e.p.x.x.SetBytes(m[0*numBytes : 1*numBytes])
	e.p.x.y.SetBytes(m[1*numBytes : 2*numBytes])
	e.p.y.x.SetBytes(m[2*numBytes : 3*numBytes])
	e.p.y.y.SetBytes(m[3*numBytes : 4*numBytes])

	if e.p.x.x.Sign() == 0 &&
		e.p.x.y.Sign() == 0 &&
		e.p.y.x.Sign() == 0 &&
		e.p.y.y.Sign() == 0 {
		fuzz_helper.CoverTab[49245]++

		e.p.y.SetOne()
		e.p.z.SetZero()
		e.p.t.SetZero()
	} else {
		fuzz_helper.CoverTab[15785]++
		e.p.z.SetOne()
		e.p.t.SetOne()

		if !e.p.IsOnCurve() {
			fuzz_helper.CoverTab[9735]++
			return nil, false
		} else {
			fuzz_helper.CoverTab[45823]++
		}
	}
	fuzz_helper.CoverTab[60629]++

	return e, true
}

// GT is an abstract cyclic group. The zero value is suitable for use as the
// output of an operation, but cannot be used as an input.
type GT struct {
	p *gfP12
}

func (g *GT) String() string {
	fuzz_helper.CoverTab[48647]++
	return "bn256.GT" + g.p.String()
}

// ScalarMult sets e to a*k and then returns e.
func (e *GT) ScalarMult(a *GT, k *big.Int) *GT {
	fuzz_helper.CoverTab[24978]++
	if e.p == nil {
		fuzz_helper.CoverTab[19607]++
		e.p = newGFp12(nil)
	} else {
		fuzz_helper.CoverTab[28743]++
	}
	fuzz_helper.CoverTab[61755]++
	e.p.Exp(a.p, k, new(bnPool))
	return e
}

// Add sets e to a+b and then returns e.
func (e *GT) Add(a, b *GT) *GT {
	fuzz_helper.CoverTab[8832]++
	if e.p == nil {
		fuzz_helper.CoverTab[63449]++
		e.p = newGFp12(nil)
	} else {
		fuzz_helper.CoverTab[47485]++
	}
	fuzz_helper.CoverTab[40052]++
	e.p.Mul(a.p, b.p, new(bnPool))
	return e
}

// Neg sets e to -a and then returns e.
func (e *GT) Neg(a *GT) *GT {
	fuzz_helper.CoverTab[60075]++
	if e.p == nil {
		fuzz_helper.CoverTab[57682]++
		e.p = newGFp12(nil)
	} else {
		fuzz_helper.CoverTab[45496]++
	}
	fuzz_helper.CoverTab[21817]++
	e.p.Invert(a.p, new(bnPool))
	return e
}

// Marshal converts n into a byte slice.
func (n *GT) Marshal() []byte {
	fuzz_helper.CoverTab[3661]++
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
	fuzz_helper.CoverTab[22210]++
	// Each value is a 256-bit number.
	const numBytes = 256 / 8

	if len(m) != 12*numBytes {
		fuzz_helper.CoverTab[56274]++
		return nil, false
	} else {
		fuzz_helper.CoverTab[1404]++
	}
	fuzz_helper.CoverTab[4417]++

	if e.p == nil {
		fuzz_helper.CoverTab[34815]++
		e.p = newGFp12(nil)
	} else {
		fuzz_helper.CoverTab[58334]++
	}
	fuzz_helper.CoverTab[5093]++

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
	fuzz_helper.CoverTab[46899]++
	return &GT{optimalAte(g2.p, g1.p, new(bnPool))}
}

// PairingCheck calculates the Optimal Ate pairing for a set of points.
func PairingCheck(a []*G1, b []*G2) bool {
	fuzz_helper.CoverTab[40738]++
	pool := new(bnPool)

	acc := newGFp12(pool)
	acc.SetOne()

	for i := 0; i < len(a); i++ {
		fuzz_helper.CoverTab[43066]++
		if a[i].p.IsInfinity() || b[i].p.IsInfinity() {
			fuzz_helper.CoverTab[12109]++
			continue
		} else {
			fuzz_helper.CoverTab[29966]++
		}
		fuzz_helper.CoverTab[14816]++
		acc.Mul(acc, miller(b[i].p, a[i].p, pool), pool)
	}
	fuzz_helper.CoverTab[10840]++
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
	fuzz_helper.CoverTab[27153]++
	if pool == nil {
		fuzz_helper.CoverTab[61379]++
		return new(big.Int)
	} else {
		fuzz_helper.CoverTab[16873]++
	}
	fuzz_helper.CoverTab[51386]++

	pool.count++
	l := len(pool.bns)
	if l == 0 {
		fuzz_helper.CoverTab[20099]++
		return new(big.Int)
	} else {
		fuzz_helper.CoverTab[61712]++
	}
	fuzz_helper.CoverTab[4676]++

	bn := pool.bns[l-1]
	pool.bns = pool.bns[:l-1]
	return bn
}

func (pool *bnPool) Put(bn *big.Int) {
	fuzz_helper.CoverTab[12762]++
	if pool == nil {
		fuzz_helper.CoverTab[53729]++
		return
	} else {
		fuzz_helper.CoverTab[64454]++
	}
	fuzz_helper.CoverTab[17951]++
	pool.bns = append(pool.bns, bn)
	pool.count--
}

func (pool *bnPool) Count() int {
	fuzz_helper.CoverTab[7160]++
	return pool.count
}
