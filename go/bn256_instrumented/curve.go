package main

import fuzz_helper "github.com/guidovranken/go-coverage-instrumentation/helper"

import (
	"math/big"
)

// curvePoint implements the elliptic curve y²=x³+3. Points are kept in
// Jacobian form and t=z² when valid. G₁ is the set of points of this curve on
// GF(p).
type curvePoint struct {
	x, y, z, t *big.Int
}

var curveB = new(big.Int).SetInt64(3)

// curveGen is the generator of G₁.
var curveGen = &curvePoint{
	new(big.Int).SetInt64(1),
	new(big.Int).SetInt64(-2),
	new(big.Int).SetInt64(1),
	new(big.Int).SetInt64(1),
}

func newCurvePoint(pool *bnPool) *curvePoint {
	fuzz_helper.AddCoverage(22588)
	return &curvePoint{
		pool.Get(),
		pool.Get(),
		pool.Get(),
		pool.Get(),
	}
}

func (c *curvePoint) String() string {
	fuzz_helper.AddCoverage(44810)
	c.MakeAffine(new(bnPool))
	return "(" + c.x.String() + ", " + c.y.String() + ")"
}

func (c *curvePoint) Put(pool *bnPool) {
	fuzz_helper.AddCoverage(5262)
	pool.Put(c.x)
	pool.Put(c.y)
	pool.Put(c.z)
	pool.Put(c.t)
}

func (c *curvePoint) Set(a *curvePoint) {
	fuzz_helper.AddCoverage(17878)
	c.x.Set(a.x)
	c.y.Set(a.y)
	c.z.Set(a.z)
	c.t.Set(a.t)
}

// IsOnCurve returns true iff c is on the curve where c must be in affine form.
func (c *curvePoint) IsOnCurve() bool {
	fuzz_helper.AddCoverage(45021)
	yy := new(big.Int).Mul(c.y, c.y)
	xxx := new(big.Int).Mul(c.x, c.x)
	xxx.Mul(xxx, c.x)
	yy.Sub(yy, xxx)
	yy.Sub(yy, curveB)
	if yy.Sign() < 0 || yy.Cmp(P) >= 0 {
		fuzz_helper.AddCoverage(2095)
		yy.Mod(yy, P)
	} else {
		fuzz_helper.AddCoverage(21668)
	}
	fuzz_helper.AddCoverage(39040)
	return yy.Sign() == 0
}

func (c *curvePoint) SetInfinity() {
	fuzz_helper.AddCoverage(45213)
	c.z.SetInt64(0)
}

func (c *curvePoint) IsInfinity() bool {
	fuzz_helper.AddCoverage(16619)
	return c.z.Sign() == 0
}

func (c *curvePoint) Add(a, b *curvePoint, pool *bnPool) {
	fuzz_helper.AddCoverage(12692)
	if a.IsInfinity() {
		fuzz_helper.AddCoverage(64174)
		c.Set(b)
		return
	} else {
		fuzz_helper.AddCoverage(38740)
	}
	fuzz_helper.AddCoverage(42483)
	if b.IsInfinity() {
		fuzz_helper.AddCoverage(35657)
		c.Set(a)
		return
	} else {
		fuzz_helper.AddCoverage(30358)
	}
	fuzz_helper.AddCoverage(6577)

	z1z1 := pool.Get().Mul(a.z, a.z)
	z1z1.Mod(z1z1, P)
	z2z2 := pool.Get().Mul(b.z, b.z)
	z2z2.Mod(z2z2, P)
	u1 := pool.Get().Mul(a.x, z2z2)
	u1.Mod(u1, P)
	u2 := pool.Get().Mul(b.x, z1z1)
	u2.Mod(u2, P)

	t := pool.Get().Mul(b.z, z2z2)
	t.Mod(t, P)
	s1 := pool.Get().Mul(a.y, t)
	s1.Mod(s1, P)

	t.Mul(a.z, z1z1)
	t.Mod(t, P)
	s2 := pool.Get().Mul(b.y, t)
	s2.Mod(s2, P)

	h := pool.Get().Sub(u2, u1)
	xEqual := h.Sign() == 0

	t.Add(h, h)

	i := pool.Get().Mul(t, t)
	i.Mod(i, P)

	j := pool.Get().Mul(h, i)
	j.Mod(j, P)

	t.Sub(s2, s1)
	yEqual := t.Sign() == 0
	if xEqual && yEqual {
		fuzz_helper.AddCoverage(23294)
		c.Double(a, pool)
		return
	} else {
		fuzz_helper.AddCoverage(61639)
	}
	fuzz_helper.AddCoverage(17393)
	r := pool.Get().Add(t, t)

	v := pool.Get().Mul(u1, i)
	v.Mod(v, P)

	t4 := pool.Get().Mul(r, r)
	t4.Mod(t4, P)
	t.Add(v, v)
	t6 := pool.Get().Sub(t4, j)
	c.x.Sub(t6, t)

	t.Sub(v, c.x)
	t4.Mul(s1, j)
	t4.Mod(t4, P)
	t6.Add(t4, t4)
	t4.Mul(r, t)
	t4.Mod(t4, P)
	c.y.Sub(t4, t6)

	t.Add(a.z, b.z)
	t4.Mul(t, t)
	t4.Mod(t4, P)
	t.Sub(t4, z1z1)
	t4.Sub(t, z2z2)
	c.z.Mul(t4, h)
	c.z.Mod(c.z, P)

	pool.Put(z1z1)
	pool.Put(z2z2)
	pool.Put(u1)
	pool.Put(u2)
	pool.Put(t)
	pool.Put(s1)
	pool.Put(s2)
	pool.Put(h)
	pool.Put(i)
	pool.Put(j)
	pool.Put(r)
	pool.Put(v)
	pool.Put(t4)
	pool.Put(t6)
}

func (c *curvePoint) Double(a *curvePoint, pool *bnPool) {
	fuzz_helper.AddCoverage(11162)

	A := pool.Get().Mul(a.x, a.x)
	A.Mod(A, P)
	B := pool.Get().Mul(a.y, a.y)
	B.Mod(B, P)
	C_ := pool.Get().Mul(B, B)
	C_.Mod(C_, P)

	t := pool.Get().Add(a.x, B)
	t2 := pool.Get().Mul(t, t)
	t2.Mod(t2, P)
	t.Sub(t2, A)
	t2.Sub(t, C_)
	d := pool.Get().Add(t2, t2)
	t.Add(A, A)
	e := pool.Get().Add(t, A)
	f := pool.Get().Mul(e, e)
	f.Mod(f, P)

	t.Add(d, d)
	c.x.Sub(f, t)

	t.Add(C_, C_)
	t2.Add(t, t)
	t.Add(t2, t2)
	c.y.Sub(d, c.x)
	t2.Mul(e, c.y)
	t2.Mod(t2, P)
	c.y.Sub(t2, t)

	t.Mul(a.y, a.z)
	t.Mod(t, P)
	c.z.Add(t, t)

	pool.Put(A)
	pool.Put(B)
	pool.Put(C_)
	pool.Put(t)
	pool.Put(t2)
	pool.Put(d)
	pool.Put(e)
	pool.Put(f)
}

func (c *curvePoint) Mul(a *curvePoint, scalar *big.Int, pool *bnPool) *curvePoint {
	fuzz_helper.AddCoverage(49217)
	sum := newCurvePoint(pool)
	sum.SetInfinity()
	t := newCurvePoint(pool)

	for i := scalar.BitLen(); i >= 0; i-- {
		fuzz_helper.AddCoverage(64074)
		t.Double(sum, pool)
		if scalar.Bit(i) != 0 {
			fuzz_helper.AddCoverage(28614)
			sum.Add(t, a, pool)
		} else {
			fuzz_helper.AddCoverage(39226)
			sum.Set(t)
		}
	}
	fuzz_helper.AddCoverage(34511)

	c.Set(sum)
	sum.Put(pool)
	t.Put(pool)
	return c
}

func (c *curvePoint) MakeAffine(pool *bnPool) *curvePoint {
	fuzz_helper.AddCoverage(2297)
	if words := c.z.Bits(); len(words) == 1 && words[0] == 1 {
		fuzz_helper.AddCoverage(52877)
		return c
	} else {
		fuzz_helper.AddCoverage(778)
	}
	fuzz_helper.AddCoverage(40870)

	zInv := pool.Get().ModInverse(c.z, P)
	t := pool.Get().Mul(c.y, zInv)
	t.Mod(t, P)
	zInv2 := pool.Get().Mul(zInv, zInv)
	zInv2.Mod(zInv2, P)
	c.y.Mul(t, zInv2)
	c.y.Mod(c.y, P)
	t.Mul(c.x, zInv2)
	t.Mod(t, P)
	c.x.Set(t)
	c.z.SetInt64(1)
	c.t.SetInt64(1)

	pool.Put(zInv)
	pool.Put(t)
	pool.Put(zInv2)

	return c
}

func (c *curvePoint) Negative(a *curvePoint) {
	fuzz_helper.AddCoverage(33340)
	c.x.Set(a.x)
	c.y.Neg(a.y)
	c.z.Set(a.z)
	c.t.SetInt64(0)
}
