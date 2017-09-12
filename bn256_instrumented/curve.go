// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:5
package main

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:5
//import  "go-fuzz-dep"
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:8

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:7
import (
	"math/big"
)

// curvePoint implements the elliptic curve y²=x³+3. Points are kept in
// Jacobian form and t=z² when valid. G₁ is the set of points of this curve on
// GF(p).
type curvePoint struct {
	x, y, z, t *big.Int
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:19

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:18
var curveB = new(big.Int).SetInt64(3)

// curveGen is the generator of G₁.
var curveGen = &curvePoint{
	new(big.Int).SetInt64(1),
	new(big.Int).SetInt64(-2),
	new(big.Int).SetInt64(1),
	new(big.Int).SetInt64(1),
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:29

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:28
func newCurvePoint(pool *bnPool) *curvePoint {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:28
	CoverTab[40199]++
											return &curvePoint{
		pool.Get(),
		pool.Get(),
		pool.Get(),
		pool.Get(),
	}
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:38

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:37
func (c *curvePoint) String() string {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:37
	CoverTab[28204]++
											c.MakeAffine(new(bnPool))
											return "(" + c.x.String() + ", " + c.y.String() + ")"
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:43

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:42
func (c *curvePoint) Put(pool *bnPool) {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:42
	CoverTab[58635]++
											pool.Put(c.x)
											pool.Put(c.y)
											pool.Put(c.z)
											pool.Put(c.t)
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:50

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:49
func (c *curvePoint) Set(a *curvePoint) {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:49
	CoverTab[59486]++
											c.x.Set(a.x)
											c.y.Set(a.y)
											c.z.Set(a.z)
											c.t.Set(a.t)
}

// IsOnCurve returns true iff c is on the curve where c must be in affine form.
func (c *curvePoint) IsOnCurve() bool {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:57
	CoverTab[60594]++
											yy := new(big.Int).Mul(c.y, c.y)
											xxx := new(big.Int).Mul(c.x, c.x)
											xxx.Mul(xxx, c.x)
											yy.Sub(yy, xxx)
											yy.Sub(yy, curveB)
											if yy.Sign() < 0 || func() bool {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:63
		CoverTab[52194]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:63
		return yy.Cmp(p) >= 0
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:63
	}() {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:63
		CoverTab[56853]++
												yy.Mod(yy, p)
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:65
		CoverTab[31275]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:65
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:65
	CoverTab[6574]++
											return yy.Sign() == 0
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:70

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:69
func (c *curvePoint) SetInfinity() {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:69
	CoverTab[17980]++
											c.z.SetInt64(0)
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:74

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:73
func (c *curvePoint) IsInfinity() bool {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:73
	CoverTab[62596]++
											return c.z.Sign() == 0
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:78

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:77
func (c *curvePoint) Add(a, b *curvePoint, pool *bnPool) {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:77
	CoverTab[1362]++
											if a.IsInfinity() {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:78
		CoverTab[23107]++
												c.Set(b)
												return
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:81
		CoverTab[14042]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:81
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:81
	CoverTab[25639]++
											if b.IsInfinity() {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:82
		CoverTab[30949]++
												c.Set(a)
												return
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:85
		CoverTab[43321]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:85
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:85
	CoverTab[51930]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:93

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:92
	z1z1 := pool.Get().Mul(a.z, a.z)
											z1z1.Mod(z1z1, p)
											z2z2 := pool.Get().Mul(b.z, b.z)
											z2z2.Mod(z2z2, p)
											u1 := pool.Get().Mul(a.x, z2z2)
											u1.Mod(u1, p)
											u2 := pool.Get().Mul(b.x, z1z1)
											u2.Mod(u2, p)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:102

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:101
	t := pool.Get().Mul(b.z, z2z2)
											t.Mod(t, p)
											s1 := pool.Get().Mul(a.y, t)
											s1.Mod(s1, p)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:107

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:106
	t.Mul(a.z, z1z1)
											t.Mod(t, p)
											s2 := pool.Get().Mul(b.y, t)
											s2.Mod(s2, p)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:119

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:118
	h := pool.Get().Sub(u2, u1)
											xEqual := h.Sign() == 0
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:122

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:121
	t.Add(h, h)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:124

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:123
	i := pool.Get().Mul(t, t)
											i.Mod(i, p)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:127

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:126
	j := pool.Get().Mul(h, i)
											j.Mod(j, p)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:130

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:129
	t.Sub(s2, s1)
											yEqual := t.Sign() == 0
											if xEqual && func() bool {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:131
		CoverTab[49101]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:131
		return yEqual
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:131
	}() {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:131
		CoverTab[47999]++
												c.Double(a, pool)
												return
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:134
		CoverTab[10087]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:134
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:134
	CoverTab[59800]++
											r := pool.Get().Add(t, t)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:138

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:137
	v := pool.Get().Mul(u1, i)
											v.Mod(v, p)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:142

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:141
	t4 := pool.Get().Mul(r, r)
											t4.Mod(t4, p)
											t.Add(v, v)
											t6 := pool.Get().Sub(t4, j)
											c.x.Sub(t6, t)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:151

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:150
	t.Sub(v, c.x)
											t4.Mul(s1, j)
											t4.Mod(t4, p)
											t6.Add(t4, t4)
											t4.Mul(r, t)
											t4.Mod(t4, p)
											c.y.Sub(t4, t6)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:160

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:159
	t.Add(a.z, b.z)
											t4.Mul(t, t)
											t4.Mod(t4, p)
											t.Sub(t4, z1z1)
											t4.Sub(t, z2z2)
											c.z.Mul(t4, h)
											c.z.Mod(c.z, p)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:168

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:167
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
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:184

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:183
func (c *curvePoint) Double(a *curvePoint, pool *bnPool) {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:183
	CoverTab[39720]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:186

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:185
	A := pool.Get().Mul(a.x, a.x)
											A.Mod(A, p)
											B := pool.Get().Mul(a.y, a.y)
											B.Mod(B, p)
											C := pool.Get().Mul(B, B)
											C.Mod(C, p)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:193

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:192
	t := pool.Get().Add(a.x, B)
											t2 := pool.Get().Mul(t, t)
											t2.Mod(t2, p)
											t.Sub(t2, A)
											t2.Sub(t, C)
											d := pool.Get().Add(t2, t2)
											t.Add(A, A)
											e := pool.Get().Add(t, A)
											f := pool.Get().Mul(e, e)
											f.Mod(f, p)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:204

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:203
	t.Add(d, d)
											c.x.Sub(f, t)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:207

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:206
	t.Add(C, C)
											t2.Add(t, t)
											t.Add(t2, t2)
											c.y.Sub(d, c.x)
											t2.Mul(e, c.y)
											t2.Mod(t2, p)
											c.y.Sub(t2, t)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:215

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:214
	t.Mul(a.y, a.z)
											t.Mod(t, p)
											c.z.Add(t, t)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:219

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:218
	pool.Put(A)
											pool.Put(B)
											pool.Put(C)
											pool.Put(t)
											pool.Put(t2)
											pool.Put(d)
											pool.Put(e)
											pool.Put(f)
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:229

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:228
func (c *curvePoint) Mul(a *curvePoint, scalar *big.Int, pool *bnPool) *curvePoint {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:228
	CoverTab[56066]++
											sum := newCurvePoint(pool)
											sum.SetInfinity()
											t := newCurvePoint(pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:234

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:233
	for i := scalar.BitLen(); i >= 0; i-- {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:233
		CoverTab[1263]++
												t.Double(sum, pool)
												if scalar.Bit(i) != 0 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:235
			CoverTab[18533]++
													sum.Add(t, a, pool)
		} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:237
			CoverTab[59283]++
													sum.Set(t)
		}
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:240
	CoverTab[10764]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:243

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:242
	c.Set(sum)
											sum.Put(pool)
											t.Put(pool)
											return c
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:249

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:248
func (c *curvePoint) MakeAffine(pool *bnPool) *curvePoint {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:248
	CoverTab[34995]++
											if words := c.z.Bits(); len(words) == 1 && func() bool {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:249
		CoverTab[51083]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:249
		return words[0] == 1
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:249
	}() {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:249
		CoverTab[57798]++
												return c
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:251
		CoverTab[2412]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:251
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:251
	CoverTab[9410]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:254

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:253
	zInv := pool.Get().ModInverse(c.z, p)
											t := pool.Get().Mul(c.y, zInv)
											t.Mod(t, p)
											zInv2 := pool.Get().Mul(zInv, zInv)
											zInv2.Mod(zInv2, p)
											c.y.Mul(t, zInv2)
											c.y.Mod(c.y, p)
											t.Mul(c.x, zInv2)
											t.Mod(t, p)
											c.x.Set(t)
											c.z.SetInt64(1)
											c.t.SetInt64(1)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:267

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:266
	pool.Put(zInv)
											pool.Put(t)
											pool.Put(zInv2)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:271

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:270
	return c
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:274

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:273
func (c *curvePoint) Negative(a *curvePoint) {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:273
	CoverTab[55526]++
											c.x.Set(a.x)
											c.y.Neg(a.y)
											c.z.Set(a.z)
											c.t.SetInt64(0)
}

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/curve.go:278
//var _ = .Main
