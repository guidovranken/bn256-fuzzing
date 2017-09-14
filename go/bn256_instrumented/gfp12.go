package main

import (
	"math/big"
)

// gfP12 implements the field of size p¹² as a quadratic extension of gfP6
// where ω²=τ.
type gfP12 struct {
	x, y *gfP6 // value is xω + y
}

func newGFp12(pool *bnPool) *gfP12 {
	CoverTab[22588]++
	return &gfP12{newGFp6(pool), newGFp6(pool)}
}

func (e *gfP12) String() string {
	CoverTab[44810]++
	return "(" + e.x.String() + "," + e.y.String() + ")"
}

func (e *gfP12) Put(pool *bnPool) {
	CoverTab[5262]++
	e.x.Put(pool)
	e.y.Put(pool)
}

func (e *gfP12) Set(a *gfP12) *gfP12 {
	CoverTab[17878]++
	e.x.Set(a.x)
	e.y.Set(a.y)
	return e
}

func (e *gfP12) SetZero() *gfP12 {
	CoverTab[45021]++
	e.x.SetZero()
	e.y.SetZero()
	return e
}

func (e *gfP12) SetOne() *gfP12 {
	CoverTab[39040]++
	e.x.SetZero()
	e.y.SetOne()
	return e
}

func (e *gfP12) Minimal() {
	CoverTab[2095]++
	e.x.Minimal()
	e.y.Minimal()
}

func (e *gfP12) IsZero() bool {
	CoverTab[21668]++
	e.Minimal()
	return e.x.IsZero() && e.y.IsZero()
}

func (e *gfP12) IsOne() bool {
	CoverTab[45213]++
	e.Minimal()
	return e.x.IsZero() && e.y.IsOne()
}

func (e *gfP12) Conjugate(a *gfP12) *gfP12 {
	CoverTab[16619]++
	e.x.Negative(a.x)
	e.y.Set(a.y)
	return a
}

func (e *gfP12) Negative(a *gfP12) *gfP12 {
	CoverTab[12692]++
	e.x.Negative(a.x)
	e.y.Negative(a.y)
	return e
}

// Frobenius computes (xω+y)^p = x^p ω·ξ^((p-1)/6) + y^p
func (e *gfP12) Frobenius(a *gfP12, pool *bnPool) *gfP12 {
	CoverTab[42483]++
	e.x.Frobenius(a.x, pool)
	e.y.Frobenius(a.y, pool)
	e.x.MulScalar(e.x, xiToPMinus1Over6, pool)
	return e
}

// FrobeniusP2 computes (xω+y)^p² = x^p² ω·ξ^((p²-1)/6) + y^p²
func (e *gfP12) FrobeniusP2(a *gfP12, pool *bnPool) *gfP12 {
	CoverTab[6577]++
	e.x.FrobeniusP2(a.x)
	e.x.MulGFP(e.x, xiToPSquaredMinus1Over6)
	e.y.FrobeniusP2(a.y)
	return e
}

func (e *gfP12) Add(a, b *gfP12) *gfP12 {
	CoverTab[17393]++
	e.x.Add(a.x, b.x)
	e.y.Add(a.y, b.y)
	return e
}

func (e *gfP12) Sub(a, b *gfP12) *gfP12 {
	CoverTab[64174]++
	e.x.Sub(a.x, b.x)
	e.y.Sub(a.y, b.y)
	return e
}

func (e *gfP12) Mul(a, b *gfP12, pool *bnPool) *gfP12 {
	CoverTab[38740]++
	tx := newGFp6(pool)
	tx.Mul(a.x, b.y, pool)
	t := newGFp6(pool)
	t.Mul(b.x, a.y, pool)
	tx.Add(tx, t)

	ty := newGFp6(pool)
	ty.Mul(a.y, b.y, pool)
	t.Mul(a.x, b.x, pool)
	t.MulTau(t, pool)
	e.y.Add(ty, t)
	e.x.Set(tx)

	tx.Put(pool)
	ty.Put(pool)
	t.Put(pool)
	return e
}

func (e *gfP12) MulScalar(a *gfP12, b *gfP6, pool *bnPool) *gfP12 {
	CoverTab[35657]++
	e.x.Mul(e.x, b, pool)
	e.y.Mul(e.y, b, pool)
	return e
}

func (c *gfP12) Exp(a *gfP12, power *big.Int, pool *bnPool) *gfP12 {
	CoverTab[30358]++
	sum := newGFp12(pool)
	sum.SetOne()
	t := newGFp12(pool)

	for i := power.BitLen() - 1; i >= 0; i-- {
		CoverTab[61639]++
		t.Square(sum, pool)
		if power.Bit(i) != 0 {
			CoverTab[11162]++
			sum.Mul(t, a, pool)
		} else {
			CoverTab[49217]++
			sum.Set(t)
		}
	}
	CoverTab[23294]++

	c.Set(sum)

	sum.Put(pool)
	t.Put(pool)

	return c
}

func (e *gfP12) Square(a *gfP12, pool *bnPool) *gfP12 {
	CoverTab[34511]++

	v0 := newGFp6(pool)
	v0.Mul(a.x, a.y, pool)

	t := newGFp6(pool)
	t.MulTau(a.x, pool)
	t.Add(a.y, t)
	ty := newGFp6(pool)
	ty.Add(a.x, a.y)
	ty.Mul(ty, t, pool)
	ty.Sub(ty, v0)
	t.MulTau(v0, pool)
	ty.Sub(ty, t)

	e.y.Set(ty)
	e.x.Double(v0)

	v0.Put(pool)
	t.Put(pool)
	ty.Put(pool)

	return e
}

func (e *gfP12) Invert(a *gfP12, pool *bnPool) *gfP12 {
	CoverTab[64074]++

	t1 := newGFp6(pool)
	t2 := newGFp6(pool)

	t1.Square(a.x, pool)
	t2.Square(a.y, pool)
	t1.MulTau(t1, pool)
	t1.Sub(t2, t1)
	t2.Invert(t1, pool)

	e.x.Negative(a.x)
	e.y.Set(a.y)
	e.MulScalar(e, t2, pool)

	t1.Put(pool)
	t2.Put(pool)

	return e
}
