// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:5
package main

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:5
//import  "go-fuzz-dep"
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:12

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:11
import (
	"math/big"
)

// gfP6 implements the field of size p⁶ as a cubic extension of gfP2 where τ³=ξ
// and ξ=i+3.
type gfP6 struct {
	x, y, z *gfP2	// value is xτ² + yτ + z
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:22

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:21
func newGFp6(pool *bnPool) *gfP6 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:21
	CoverTab[12059]++
											return &gfP6{newGFp2(pool), newGFp2(pool), newGFp2(pool)}
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:26

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:25
func (e *gfP6) String() string {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:25
	CoverTab[16121]++
											return "(" + e.x.String() + "," + e.y.String() + "," + e.z.String() + ")"
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:30

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:29
func (e *gfP6) Put(pool *bnPool) {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:29
	CoverTab[48007]++
											e.x.Put(pool)
											e.y.Put(pool)
											e.z.Put(pool)
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:36

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:35
func (e *gfP6) Set(a *gfP6) *gfP6 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:35
	CoverTab[45705]++
											e.x.Set(a.x)
											e.y.Set(a.y)
											e.z.Set(a.z)
											return e
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:43

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:42
func (e *gfP6) SetZero() *gfP6 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:42
	CoverTab[1992]++
											e.x.SetZero()
											e.y.SetZero()
											e.z.SetZero()
											return e
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:50

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:49
func (e *gfP6) SetOne() *gfP6 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:49
	CoverTab[41832]++
											e.x.SetZero()
											e.y.SetZero()
											e.z.SetOne()
											return e
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:57

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:56
func (e *gfP6) Minimal() {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:56
	CoverTab[45150]++
											e.x.Minimal()
											e.y.Minimal()
											e.z.Minimal()
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:63

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:62
func (e *gfP6) IsZero() bool {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:62
	CoverTab[50563]++
											return e.x.IsZero() && func() bool {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:63
		CoverTab[15229]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:63
		return e.y.IsZero()
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:63
	}() && func() bool {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:63
		CoverTab[8372]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:63
		return e.z.IsZero()
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:63
	}()
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:67

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:66
func (e *gfP6) IsOne() bool {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:66
	CoverTab[4386]++
											return e.x.IsZero() && func() bool {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:67
		CoverTab[30568]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:67
		return e.y.IsZero()
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:67
	}() && func() bool {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:67
		CoverTab[36092]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:67
		return e.z.IsOne()
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:67
	}()
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:71

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:70
func (e *gfP6) Negative(a *gfP6) *gfP6 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:70
	CoverTab[46626]++
											e.x.Negative(a.x)
											e.y.Negative(a.y)
											e.z.Negative(a.z)
											return e
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:78

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:77
func (e *gfP6) Frobenius(a *gfP6, pool *bnPool) *gfP6 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:77
	CoverTab[37125]++
											e.x.Conjugate(a.x)
											e.y.Conjugate(a.y)
											e.z.Conjugate(a.z)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:83

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:82
	e.x.Mul(e.x, xiTo2PMinus2Over3, pool)
											e.y.Mul(e.y, xiToPMinus1Over3, pool)
											return e
}

// FrobeniusP2 computes (xτ²+yτ+z)^(p²) = xτ^(2p²) + yτ^(p²) + z
func (e *gfP6) FrobeniusP2(a *gfP6) *gfP6 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:88
	CoverTab[48397]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:91

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:90
	e.x.MulScalar(a.x, xiTo2PSquaredMinus2Over3)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:93

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:92
	e.y.MulScalar(a.y, xiToPSquaredMinus1Over3)
											e.z.Set(a.z)
											return e
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:98

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:97
func (e *gfP6) Add(a, b *gfP6) *gfP6 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:97
	CoverTab[18881]++
											e.x.Add(a.x, b.x)
											e.y.Add(a.y, b.y)
											e.z.Add(a.z, b.z)
											return e
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:105

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:104
func (e *gfP6) Sub(a, b *gfP6) *gfP6 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:104
	CoverTab[41043]++
											e.x.Sub(a.x, b.x)
											e.y.Sub(a.y, b.y)
											e.z.Sub(a.z, b.z)
											return e
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:112

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:111
func (e *gfP6) Double(a *gfP6) *gfP6 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:111
	CoverTab[14110]++
											e.x.Double(a.x)
											e.y.Double(a.y)
											e.z.Double(a.z)
											return e
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:119

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:118
func (e *gfP6) Mul(a, b *gfP6, pool *bnPool) *gfP6 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:118
	CoverTab[60549]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:124

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:123
	v0 := newGFp2(pool)
											v0.Mul(a.z, b.z, pool)
											v1 := newGFp2(pool)
											v1.Mul(a.y, b.y, pool)
											v2 := newGFp2(pool)
											v2.Mul(a.x, b.x, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:131

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:130
	t0 := newGFp2(pool)
											t0.Add(a.x, a.y)
											t1 := newGFp2(pool)
											t1.Add(b.x, b.y)
											tz := newGFp2(pool)
											tz.Mul(t0, t1, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:138

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:137
	tz.Sub(tz, v1)
											tz.Sub(tz, v2)
											tz.MulXi(tz, pool)
											tz.Add(tz, v0)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:143

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:142
	t0.Add(a.y, a.z)
											t1.Add(b.y, b.z)
											ty := newGFp2(pool)
											ty.Mul(t0, t1, pool)
											ty.Sub(ty, v0)
											ty.Sub(ty, v1)
											t0.MulXi(v2, pool)
											ty.Add(ty, t0)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:152

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:151
	t0.Add(a.x, a.z)
											t1.Add(b.x, b.z)
											tx := newGFp2(pool)
											tx.Mul(t0, t1, pool)
											tx.Sub(tx, v0)
											tx.Add(tx, v1)
											tx.Sub(tx, v2)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:160

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:159
	e.x.Set(tx)
											e.y.Set(ty)
											e.z.Set(tz)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:164

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:163
	t0.Put(pool)
											t1.Put(pool)
											tx.Put(pool)
											ty.Put(pool)
											tz.Put(pool)
											v0.Put(pool)
											v1.Put(pool)
											v2.Put(pool)
											return e
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:175

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:174
func (e *gfP6) MulScalar(a *gfP6, b *gfP2, pool *bnPool) *gfP6 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:174
	CoverTab[31166]++
											e.x.Mul(a.x, b, pool)
											e.y.Mul(a.y, b, pool)
											e.z.Mul(a.z, b, pool)
											return e
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:182

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:181
func (e *gfP6) MulGFP(a *gfP6, b *big.Int) *gfP6 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:181
	CoverTab[53175]++
											e.x.MulScalar(a.x, b)
											e.y.MulScalar(a.y, b)
											e.z.MulScalar(a.z, b)
											return e
}

// MulTau computes τ·(aτ²+bτ+c) = bτ²+cτ+aξ
func (e *gfP6) MulTau(a *gfP6, pool *bnPool) {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:189
	CoverTab[12792]++
											tz := newGFp2(pool)
											tz.MulXi(a.x, pool)
											ty := newGFp2(pool)
											ty.Set(a.y)
											e.y.Set(a.z)
											e.x.Set(ty)
											e.z.Set(tz)
											tz.Put(pool)
											ty.Put(pool)
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:202

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:201
func (e *gfP6) Square(a *gfP6, pool *bnPool) *gfP6 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:201
	CoverTab[54090]++
											v0 := newGFp2(pool).Square(a.z, pool)
											v1 := newGFp2(pool).Square(a.y, pool)
											v2 := newGFp2(pool).Square(a.x, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:207

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:206
	c0 := newGFp2(pool).Add(a.x, a.y)
											c0.Square(c0, pool)
											c0.Sub(c0, v1)
											c0.Sub(c0, v2)
											c0.MulXi(c0, pool)
											c0.Add(c0, v0)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:214

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:213
	c1 := newGFp2(pool).Add(a.y, a.z)
											c1.Square(c1, pool)
											c1.Sub(c1, v0)
											c1.Sub(c1, v1)
											xiV2 := newGFp2(pool).MulXi(v2, pool)
											c1.Add(c1, xiV2)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:221

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:220
	c2 := newGFp2(pool).Add(a.x, a.z)
											c2.Square(c2, pool)
											c2.Sub(c2, v0)
											c2.Add(c2, v1)
											c2.Sub(c2, v2)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:227

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:226
	e.x.Set(c2)
											e.y.Set(c1)
											e.z.Set(c0)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:231

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:230
	v0.Put(pool)
											v1.Put(pool)
											v2.Put(pool)
											c0.Put(pool)
											c1.Put(pool)
											c2.Put(pool)
											xiV2.Put(pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:239

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:238
	return e
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:242

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:241
func (e *gfP6) Invert(a *gfP6, pool *bnPool) *gfP6 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:241
	CoverTab[4051]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:256

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:255
	t1 := newGFp2(pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:258

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:257
	A := newGFp2(pool)
											A.Square(a.z, pool)
											t1.Mul(a.x, a.y, pool)
											t1.MulXi(t1, pool)
											A.Sub(A, t1)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:264

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:263
	B := newGFp2(pool)
											B.Square(a.x, pool)
											B.MulXi(B, pool)
											t1.Mul(a.y, a.z, pool)
											B.Sub(B, t1)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:270

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:269
	C := newGFp2(pool)
											C.Square(a.y, pool)
											t1.Mul(a.x, a.z, pool)
											C.Sub(C, t1)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:275

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:274
	F := newGFp2(pool)
											F.Mul(C, a.y, pool)
											F.MulXi(F, pool)
											t1.Mul(A, a.z, pool)
											F.Add(F, t1)
											t1.Mul(B, a.x, pool)
											t1.MulXi(t1, pool)
											F.Add(F, t1)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:284

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:283
	F.Invert(F, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:286

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:285
	e.x.Mul(C, F, pool)
											e.y.Mul(B, F, pool)
											e.z.Mul(A, F, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:290

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:289
	t1.Put(pool)
											A.Put(pool)
											B.Put(pool)
											C.Put(pool)
											F.Put(pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:296

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:295
	return e
}

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp6.go:296
//var _ = .Main
