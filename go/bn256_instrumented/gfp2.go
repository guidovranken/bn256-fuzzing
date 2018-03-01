// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import

// For details of the algorithms used, see "Multiplication and Squaring on
// Pairing-Friendly Fields, Devegili et al.
// http://eprint.iacr.org/2006/471.pdf.
fuzz_helper "github.com/guidovranken/go-coverage-instrumentation/helper"

import (
	"math/big"
)

// gfP2 implements a field of size p² as a quadratic extension of the base
// field where i²=-1.
type gfP2 struct {
	x, y *big.Int // value is xi+y.
}

func newGFp2(pool *bnPool) *gfP2 {
	fuzz_helper.AddCoverage(246)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	return &gfP2{pool.Get(), pool.Get()}
}

func (e *gfP2) String() string {
	fuzz_helper.AddCoverage(247)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	x := new(big.Int).Mod(e.x, P)
	y := new(big.Int).Mod(e.y, P)
	return "(" + x.String() + "," + y.String() + ")"
}

func (e *gfP2) Put(pool *bnPool) {
	fuzz_helper.AddCoverage(248)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	pool.Put(e.x)
	pool.Put(e.y)
}

func (e *gfP2) Set(a *gfP2) *gfP2 {
	fuzz_helper.AddCoverage(249)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	e.x.Set(a.x)
	e.y.Set(a.y)
	return e
}

func (e *gfP2) SetZero() *gfP2 {
	fuzz_helper.AddCoverage(250)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	e.x.SetInt64(0)
	e.y.SetInt64(0)
	return e
}

func (e *gfP2) SetOne() *gfP2 {
	fuzz_helper.AddCoverage(251)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	e.x.SetInt64(0)
	e.y.SetInt64(1)
	return e
}

func (e *gfP2) Minimal() {
	fuzz_helper.AddCoverage(252)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	if e.x.Sign() < 0 || e.x.Cmp(P) >= 0 {
		fuzz_helper.AddCoverage(254)
		e.x.Mod(e.x, P)
	} else {
		fuzz_helper.AddCoverage(255)
	}
	fuzz_helper.AddCoverage(253)
	if e.y.Sign() < 0 || e.y.Cmp(P) >= 0 {
		fuzz_helper.AddCoverage(256)
		e.y.Mod(e.y, P)
	} else {
		fuzz_helper.AddCoverage(257)
	}
}

func (e *gfP2) IsZero() bool {
	fuzz_helper.AddCoverage(258)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	return e.x.Sign() == 0 && e.y.Sign() == 0
}

func (e *gfP2) IsOne() bool {
	fuzz_helper.AddCoverage(259)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	if e.x.Sign() != 0 {
		fuzz_helper.AddCoverage(261)
		return false
	} else {
		fuzz_helper.AddCoverage(262)
	}
	fuzz_helper.AddCoverage(260)
	words := e.y.Bits()
	return len(words) == 1 && words[0] == 1
}

func (e *gfP2) Conjugate(a *gfP2) *gfP2 {
	fuzz_helper.AddCoverage(263)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	e.y.Set(a.y)
	e.x.Neg(a.x)
	return e
}

func (e *gfP2) Negative(a *gfP2) *gfP2 {
	fuzz_helper.AddCoverage(264)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	e.x.Neg(a.x)
	e.y.Neg(a.y)
	return e
}

func (e *gfP2) Add(a, b *gfP2) *gfP2 {
	fuzz_helper.AddCoverage(265)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	e.x.Add(a.x, b.x)
	e.y.Add(a.y, b.y)
	return e
}

func (e *gfP2) Sub(a, b *gfP2) *gfP2 {
	fuzz_helper.AddCoverage(266)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	e.x.Sub(a.x, b.x)
	e.y.Sub(a.y, b.y)
	return e
}

func (e *gfP2) Double(a *gfP2) *gfP2 {
	fuzz_helper.AddCoverage(267)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	e.x.Lsh(a.x, 1)
	e.y.Lsh(a.y, 1)
	return e
}

func (c *gfP2) Exp(a *gfP2, power *big.Int, pool *bnPool) *gfP2 {
	fuzz_helper.AddCoverage(268)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	sum := newGFp2(pool)
	sum.SetOne()
	t := newGFp2(pool)

	for i := power.BitLen() - 1; i >= 0; i-- {
		fuzz_helper.AddCoverage(270)
		t.Square(sum, pool)
		if power.Bit(i) != 0 {
			fuzz_helper.AddCoverage(271)
			sum.Mul(t, a, pool)
		} else {
			fuzz_helper.AddCoverage(272)
			sum.Set(t)
		}
	}
	fuzz_helper.AddCoverage(269)

	c.Set(sum)

	sum.Put(pool)
	t.Put(pool)

	return c
}

// See "Multiplication and Squaring in Pairing-Friendly Fields",
// http://eprint.iacr.org/2006/471.pdf
func (e *gfP2) Mul(a, b *gfP2, pool *bnPool) *gfP2 {
	fuzz_helper.AddCoverage(273)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
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
	fuzz_helper.AddCoverage(274)
	fuzz_helper.IncrementStack(

	// MulXi sets e=ξa where ξ=i+9 and then returns e.
	)
	defer fuzz_helper.DecrementStack()
	e.x.Mul(a.x, b)
	e.y.Mul(a.y, b)
	return e
}

func (e *gfP2) MulXi(a *gfP2, pool *bnPool) *gfP2 {
	fuzz_helper.
		// (xi+y)(i+3) = (9x+y)i+(9y-x)
		AddCoverage(275)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()

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
	fuzz_helper.
		// Complex squaring algorithm:
		// (xi+b)² = (x+y)(y-x) + 2*i*x*y
		AddCoverage(276)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()

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
	fuzz_helper.
		// See "Implementing cryptographic pairings", M. Scott, section 3.2.
		// ftp://136.206.11.249/pub/crypto/pairings.pdf
		AddCoverage(277)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()

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
	fuzz_helper.AddCoverage(278)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	return e.x
}

func (e *gfP2) Imag() *big.Int {
	fuzz_helper.AddCoverage(279)
	fuzz_helper.IncrementStack()
	defer fuzz_helper.DecrementStack()
	return e.y
}

var _ = fuzz_helper.AddCoverage
