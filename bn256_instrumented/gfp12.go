// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:5
package main

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:5
//import  "go-fuzz-dep"
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:12

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:11
import (
	"math/big"
)

// gfP12 implements the field of size p¹² as a quadratic extension of gfP6
// where ω²=τ.
type gfP12 struct {
	x, y *gfP6	// value is xω + y
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:22

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:21
func newGFp12(pool *bnPool) *gfP12 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:21
	CoverTab[56865]++
											return &gfP12{newGFp6(pool), newGFp6(pool)}
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:26

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:25
func (e *gfP12) String() string {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:25
	CoverTab[17355]++
											return "(" + e.x.String() + "," + e.y.String() + ")"
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:30

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:29
func (e *gfP12) Put(pool *bnPool) {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:29
	CoverTab[54878]++
											e.x.Put(pool)
											e.y.Put(pool)
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:35

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:34
func (e *gfP12) Set(a *gfP12) *gfP12 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:34
	CoverTab[22787]++
											e.x.Set(a.x)
											e.y.Set(a.y)
											return e
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:41

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:40
func (e *gfP12) SetZero() *gfP12 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:40
	CoverTab[14285]++
											e.x.SetZero()
											e.y.SetZero()
											return e
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:47

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:46
func (e *gfP12) SetOne() *gfP12 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:46
	CoverTab[48556]++
											e.x.SetZero()
											e.y.SetOne()
											return e
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:53

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:52
func (e *gfP12) Minimal() {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:52
	CoverTab[52240]++
											e.x.Minimal()
											e.y.Minimal()
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:58

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:57
func (e *gfP12) IsZero() bool {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:57
	CoverTab[7543]++
											e.Minimal()
											return e.x.IsZero() && func() bool {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:59
		CoverTab[37194]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:59
		return e.y.IsZero()
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:59
	}()
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:63

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:62
func (e *gfP12) IsOne() bool {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:62
	CoverTab[44299]++
											e.Minimal()
											return e.x.IsZero() && func() bool {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:64
		CoverTab[28619]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:64
		return e.y.IsOne()
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:64
	}()
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:68

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:67
func (e *gfP12) Conjugate(a *gfP12) *gfP12 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:67
	CoverTab[31290]++
											e.x.Negative(a.x)
											e.y.Set(a.y)
											return a
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:74

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:73
func (e *gfP12) Negative(a *gfP12) *gfP12 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:73
	CoverTab[20444]++
											e.x.Negative(a.x)
											e.y.Negative(a.y)
											return e
}

// Frobenius computes (xω+y)^p = x^p ω·ξ^((p-1)/6) + y^p
func (e *gfP12) Frobenius(a *gfP12, pool *bnPool) *gfP12 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:80
	CoverTab[50340]++
											e.x.Frobenius(a.x, pool)
											e.y.Frobenius(a.y, pool)
											e.x.MulScalar(e.x, xiToPMinus1Over6, pool)
											return e
}

// FrobeniusP2 computes (xω+y)^p² = x^p² ω·ξ^((p²-1)/6) + y^p²
func (e *gfP12) FrobeniusP2(a *gfP12, pool *bnPool) *gfP12 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:88
	CoverTab[24762]++
											e.x.FrobeniusP2(a.x)
											e.x.MulGFP(e.x, xiToPSquaredMinus1Over6)
											e.y.FrobeniusP2(a.y)
											return e
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:96

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:95
func (e *gfP12) Add(a, b *gfP12) *gfP12 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:95
	CoverTab[4618]++
											e.x.Add(a.x, b.x)
											e.y.Add(a.y, b.y)
											return e
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:102

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:101
func (e *gfP12) Sub(a, b *gfP12) *gfP12 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:101
	CoverTab[41694]++
											e.x.Sub(a.x, b.x)
											e.y.Sub(a.y, b.y)
											return e
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:108

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:107
func (e *gfP12) Mul(a, b *gfP12, pool *bnPool) *gfP12 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:107
	CoverTab[9145]++
											tx := newGFp6(pool)
											tx.Mul(a.x, b.y, pool)
											t := newGFp6(pool)
											t.Mul(b.x, a.y, pool)
											tx.Add(tx, t)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:115

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:114
	ty := newGFp6(pool)
											ty.Mul(a.y, b.y, pool)
											t.Mul(a.x, b.x, pool)
											t.MulTau(t, pool)
											e.y.Add(ty, t)
											e.x.Set(tx)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:122

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:121
	tx.Put(pool)
											ty.Put(pool)
											t.Put(pool)
											return e
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:128

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:127
func (e *gfP12) MulScalar(a *gfP12, b *gfP6, pool *bnPool) *gfP12 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:127
	CoverTab[60144]++
											e.x.Mul(e.x, b, pool)
											e.y.Mul(e.y, b, pool)
											return e
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:134

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:133
func (c *gfP12) Exp(a *gfP12, power *big.Int, pool *bnPool) *gfP12 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:133
	CoverTab[29573]++
											sum := newGFp12(pool)
											sum.SetOne()
											t := newGFp12(pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:139

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:138
	for i := power.BitLen() - 1; i >= 0; i-- {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:138
		CoverTab[33386]++
												t.Square(sum, pool)
												if power.Bit(i) != 0 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:140
			CoverTab[37659]++
													sum.Mul(t, a, pool)
		} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:142
			CoverTab[27008]++
													sum.Set(t)
		}
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:145
	CoverTab[64755]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:148

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:147
	c.Set(sum)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:150

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:149
	sum.Put(pool)
											t.Put(pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:153

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:152
	return c
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:156

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:155
func (e *gfP12) Square(a *gfP12, pool *bnPool) *gfP12 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:155
	CoverTab[11068]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:158

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:157
	v0 := newGFp6(pool)
											v0.Mul(a.x, a.y, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:161

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:160
	t := newGFp6(pool)
											t.MulTau(a.x, pool)
											t.Add(a.y, t)
											ty := newGFp6(pool)
											ty.Add(a.x, a.y)
											ty.Mul(ty, t, pool)
											ty.Sub(ty, v0)
											t.MulTau(v0, pool)
											ty.Sub(ty, t)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:171

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:170
	e.y.Set(ty)
											e.x.Double(v0)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:174

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:173
	v0.Put(pool)
											t.Put(pool)
											ty.Put(pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:178

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:177
	return e
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:181

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:180
func (e *gfP12) Invert(a *gfP12, pool *bnPool) *gfP12 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:180
	CoverTab[14056]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:184

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:183
	t1 := newGFp6(pool)
											t2 := newGFp6(pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:187

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:186
	t1.Square(a.x, pool)
											t2.Square(a.y, pool)
											t1.MulTau(t1, pool)
											t1.Sub(t2, t1)
											t2.Invert(t1, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:193

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:192
	e.x.Negative(a.x)
											e.y.Set(a.y)
											e.MulScalar(e, t2, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:197

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:196
	t1.Put(pool)
											t2.Put(pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:200

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:199
	return e
}

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp12.go:200
//var _ = .Main
