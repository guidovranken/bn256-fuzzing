package main

import fuzz_helper "github.com/guidovranken/go-coverage-instrumentation/helper"

import (
	"math/big"
)

// gfP2 implements a field of size p² as a quadratic extension of the base
// field where i²=-1.
type gfP2 struct {
	x, y *big.Int // value is xi+y.
}

func newGFp2(pool *bnPool) *gfP2 {
	fuzz_helper.CoverTab[22588]++
	return &gfP2{pool.Get(), pool.Get()}
}

func (e *gfP2) String() string {
	fuzz_helper.CoverTab[44810]++
	x := new(big.Int).Mod(e.x, P)
	y := new(big.Int).Mod(e.y, P)
	return "(" + x.String() + "," + y.String() + ")"
}

func (e *gfP2) Put(pool *bnPool) {
	fuzz_helper.CoverTab[5262]++
	pool.Put(e.x)
	pool.Put(e.y)
}

func (e *gfP2) Set(a *gfP2) *gfP2 {
	fuzz_helper.CoverTab[17878]++
	e.x.Set(a.x)
	e.y.Set(a.y)
	return e
}

func (e *gfP2) SetZero() *gfP2 {
	fuzz_helper.CoverTab[45021]++
	e.x.SetInt64(0)
	e.y.SetInt64(0)
	return e
}

func (e *gfP2) SetOne() *gfP2 {
	fuzz_helper.CoverTab[39040]++
	e.x.SetInt64(0)
	e.y.SetInt64(1)
	return e
}

func (e *gfP2) Minimal() {
	fuzz_helper.CoverTab[2095]++
	if e.x.Sign() < 0 || e.x.Cmp(P) >= 0 {
		fuzz_helper.CoverTab[45213]++
		e.x.Mod(e.x, P)
	} else {
		fuzz_helper.CoverTab[16619]++
	}
	fuzz_helper.CoverTab[21668]++
	if e.y.Sign() < 0 || e.y.Cmp(P) >= 0 {
		fuzz_helper.CoverTab[12692]++
		e.y.Mod(e.y, P)
	} else {
		fuzz_helper.CoverTab[42483]++
	}
}

func (e *gfP2) IsZero() bool {
	fuzz_helper.CoverTab[6577]++
	return e.x.Sign() == 0 && e.y.Sign() == 0
}

func (e *gfP2) IsOne() bool {
	fuzz_helper.CoverTab[17393]++
	if e.x.Sign() != 0 {
		fuzz_helper.CoverTab[38740]++
		return false
	} else {
		fuzz_helper.CoverTab[35657]++
	}
	fuzz_helper.CoverTab[64174]++
	words := e.y.Bits()
	return len(words) == 1 && words[0] == 1
}

func (e *gfP2) Conjugate(a *gfP2) *gfP2 {
	fuzz_helper.CoverTab[30358]++
	e.y.Set(a.y)
	e.x.Neg(a.x)
	return e
}

func (e *gfP2) Negative(a *gfP2) *gfP2 {
	fuzz_helper.CoverTab[23294]++
	e.x.Neg(a.x)
	e.y.Neg(a.y)
	return e
}

func (e *gfP2) Add(a, b *gfP2) *gfP2 {
	fuzz_helper.CoverTab[61639]++
	e.x.Add(a.x, b.x)
	e.y.Add(a.y, b.y)
	return e
}

func (e *gfP2) Sub(a, b *gfP2) *gfP2 {
	fuzz_helper.CoverTab[11162]++
	e.x.Sub(a.x, b.x)
	e.y.Sub(a.y, b.y)
	return e
}

func (e *gfP2) Double(a *gfP2) *gfP2 {
	fuzz_helper.CoverTab[49217]++
	e.x.Lsh(a.x, 1)
	e.y.Lsh(a.y, 1)
	return e
}

func (c *gfP2) Exp(a *gfP2, power *big.Int, pool *bnPool) *gfP2 {
	fuzz_helper.CoverTab[34511]++
	sum := newGFp2(pool)
	sum.SetOne()
	t := newGFp2(pool)

	for i := power.BitLen() - 1; i >= 0; i-- {
		fuzz_helper.CoverTab[28614]++
		t.Square(sum, pool)
		if power.Bit(i) != 0 {
			fuzz_helper.CoverTab[39226]++
			sum.Mul(t, a, pool)
		} else {
			fuzz_helper.CoverTab[2297]++
			sum.Set(t)
		}
	}
	fuzz_helper.CoverTab[64074]++

	c.Set(sum)

	sum.Put(pool)
	t.Put(pool)

	return c
}

// See "Multiplication and Squaring in Pairing-Friendly Fields",
// http://eprint.iacr.org/2006/471.pdf
func (e *gfP2) Mul(a, b *gfP2, pool *bnPool) *gfP2 {
	fuzz_helper.CoverTab[40870]++
	tx := pool.Get().Mul(a.x, b.y)
	t := pool.Get().Mul(b.x, a.y)
	tx.Add(tx, t)
	tx.Mod(tx, P)

	ty := pool.Get().Mul(a.y, b.y)
	t.Mul(a.x, b.x)
	ty.Sub(ty, t)
	e.y.Mod(ty, P)
	e.x.Set(tx)

	pool.Put(tx)
	pool.Put(ty)
	pool.Put(t)

	return e
}

func (e *gfP2) MulScalar(a *gfP2, b *big.Int) *gfP2 {
	fuzz_helper.CoverTab[52877]++
	e.x.Mul(a.x, b)
	e.y.Mul(a.y, b)
	return e
}

// MulXi sets e=ξa where ξ=i+9 and then returns e.
func (e *gfP2) MulXi(a *gfP2, pool *bnPool) *gfP2 {
	fuzz_helper.CoverTab[778]++

	tx := pool.Get().Lsh(a.x, 3)
	tx.Add(tx, a.x)
	tx.Add(tx, a.y)

	ty := pool.Get().Lsh(a.y, 3)
	ty.Add(ty, a.y)
	ty.Sub(ty, a.x)

	e.x.Set(tx)
	e.y.Set(ty)

	pool.Put(tx)
	pool.Put(ty)

	return e
}

func (e *gfP2) Square(a *gfP2, pool *bnPool) *gfP2 {
	fuzz_helper.CoverTab[33340]++

	t1 := pool.Get().Sub(a.y, a.x)
	t2 := pool.Get().Add(a.x, a.y)
	ty := pool.Get().Mul(t1, t2)
	ty.Mod(ty, P)

	t1.Mul(a.x, a.y)
	t1.Lsh(t1, 1)

	e.x.Mod(t1, P)
	e.y.Set(ty)

	pool.Put(t1)
	pool.Put(t2)
	pool.Put(ty)

	return e
}

func (e *gfP2) Invert(a *gfP2, pool *bnPool) *gfP2 {
	fuzz_helper.CoverTab[15638]++

	t := pool.Get()
	t.Mul(a.y, a.y)
	t2 := pool.Get()
	t2.Mul(a.x, a.x)
	t.Add(t, t2)

	inv := pool.Get()
	inv.ModInverse(t, P)

	e.x.Neg(a.x)
	e.x.Mul(e.x, inv)
	e.x.Mod(e.x, P)

	e.y.Mul(a.y, inv)
	e.y.Mod(e.y, P)

	pool.Put(t)
	pool.Put(t2)
	pool.Put(inv)

	return e
}

func (e *gfP2) Real() *big.Int {
	fuzz_helper.CoverTab[45869]++
	return e.x
}

func (e *gfP2) Imag() *big.Int {
	fuzz_helper.CoverTab[23368]++
	return e.y
}
