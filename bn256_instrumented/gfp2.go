// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:5
package main

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:5
//import  "go-fuzz-dep"
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:12

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:11
import (
	"math/big"
)

// gfP2 implements a field of size p² as a quadratic extension of the base
// field where i²=-1.
type gfP2 struct {
	x, y *big.Int	// value is xi+y.
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:22

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:21
func newGFp2(pool *bnPool) *gfP2 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:21
	CoverTab[16674]++
											return &gfP2{pool.Get(), pool.Get()}
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:26

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:25
func (e *gfP2) String() string {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:25
	CoverTab[35474]++
											x := new(big.Int).Mod(e.x, p)
											y := new(big.Int).Mod(e.y, p)
											return "(" + x.String() + "," + y.String() + ")"
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:32

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:31
func (e *gfP2) Put(pool *bnPool) {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:31
	CoverTab[7108]++
											pool.Put(e.x)
											pool.Put(e.y)
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:37

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:36
func (e *gfP2) Set(a *gfP2) *gfP2 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:36
	CoverTab[58745]++
											e.x.Set(a.x)
											e.y.Set(a.y)
											return e
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:43

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:42
func (e *gfP2) SetZero() *gfP2 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:42
	CoverTab[14093]++
											e.x.SetInt64(0)
											e.y.SetInt64(0)
											return e
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:49

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:48
func (e *gfP2) SetOne() *gfP2 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:48
	CoverTab[48586]++
											e.x.SetInt64(0)
											e.y.SetInt64(1)
											return e
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:55

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:54
func (e *gfP2) Minimal() {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:54
	CoverTab[40988]++
											if e.x.Sign() < 0 || func() bool {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:55
		CoverTab[41348]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:55
		return e.x.Cmp(p) >= 0
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:55
	}() {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:55
		CoverTab[17010]++
												e.x.Mod(e.x, p)
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:57
		CoverTab[5456]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:57
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:57
	CoverTab[16335]++
											if e.y.Sign() < 0 || func() bool {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:58
		CoverTab[23943]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:58
		return e.y.Cmp(p) >= 0
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:58
	}() {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:58
		CoverTab[26294]++
												e.y.Mod(e.y, p)
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:60
		CoverTab[29263]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:60
	}
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:64

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:63
func (e *gfP2) IsZero() bool {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:63
	CoverTab[49734]++
											return e.x.Sign() == 0 && func() bool {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:64
		CoverTab[44303]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:64
		return e.y.Sign() == 0
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:64
	}()
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:68

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:67
func (e *gfP2) IsOne() bool {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:67
	CoverTab[49785]++
											if e.x.Sign() != 0 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:68
		CoverTab[9539]++
												return false
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:70
		CoverTab[48911]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:70
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:70
	CoverTab[65042]++
											words := e.y.Bits()
											return len(words) == 1 && func() bool {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:72
		CoverTab[40201]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:72
		return words[0] == 1
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:72
	}()
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:76

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:75
func (e *gfP2) Conjugate(a *gfP2) *gfP2 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:75
	CoverTab[63834]++
											e.y.Set(a.y)
											e.x.Neg(a.x)
											return e
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:82

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:81
func (e *gfP2) Negative(a *gfP2) *gfP2 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:81
	CoverTab[14159]++
											e.x.Neg(a.x)
											e.y.Neg(a.y)
											return e
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:88

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:87
func (e *gfP2) Add(a, b *gfP2) *gfP2 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:87
	CoverTab[42304]++
											e.x.Add(a.x, b.x)
											e.y.Add(a.y, b.y)
											return e
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:94

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:93
func (e *gfP2) Sub(a, b *gfP2) *gfP2 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:93
	CoverTab[40669]++
											e.x.Sub(a.x, b.x)
											e.y.Sub(a.y, b.y)
											return e
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:100

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:99
func (e *gfP2) Double(a *gfP2) *gfP2 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:99
	CoverTab[48498]++
											e.x.Lsh(a.x, 1)
											e.y.Lsh(a.y, 1)
											return e
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:106

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:105
func (c *gfP2) Exp(a *gfP2, power *big.Int, pool *bnPool) *gfP2 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:105
	CoverTab[33777]++
											sum := newGFp2(pool)
											sum.SetOne()
											t := newGFp2(pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:111

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:110
	for i := power.BitLen() - 1; i >= 0; i-- {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:110
		CoverTab[9929]++
												t.Square(sum, pool)
												if power.Bit(i) != 0 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:112
			CoverTab[8089]++
													sum.Mul(t, a, pool)
		} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:114
			CoverTab[18427]++
													sum.Set(t)
		}
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:117
	CoverTab[19221]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:120

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:119
	c.Set(sum)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:122

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:121
	sum.Put(pool)
											t.Put(pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:125

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:124
	return c
}

// See "Multiplication and Squaring in Pairing-Friendly Fields",
// http://eprint.iacr.org/2006/471.pdf
func (e *gfP2) Mul(a, b *gfP2, pool *bnPool) *gfP2 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:129
	CoverTab[15432]++
											tx := pool.Get().Mul(a.x, b.y)
											t := pool.Get().Mul(b.x, a.y)
											tx.Add(tx, t)
											tx.Mod(tx, p)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:136

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:135
	ty := pool.Get().Mul(a.y, b.y)
											t.Mul(a.x, b.x)
											ty.Sub(ty, t)
											e.y.Mod(ty, p)
											e.x.Set(tx)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:142

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:141
	pool.Put(tx)
											pool.Put(ty)
											pool.Put(t)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:146

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:145
	return e
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:149

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:148
func (e *gfP2) MulScalar(a *gfP2, b *big.Int) *gfP2 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:148
	CoverTab[8702]++
											e.x.Mul(a.x, b)
											e.y.Mul(a.y, b)
											return e
}

// MulXi sets e=ξa where ξ=i+3 and then returns e.
func (e *gfP2) MulXi(a *gfP2, pool *bnPool) *gfP2 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:155
	CoverTab[61301]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:158

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:157
	tx := pool.Get().Lsh(a.x, 1)
											tx.Add(tx, a.x)
											tx.Add(tx, a.y)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:162

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:161
	ty := pool.Get().Lsh(a.y, 1)
											ty.Add(ty, a.y)
											ty.Sub(ty, a.x)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:166

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:165
	e.x.Set(tx)
											e.y.Set(ty)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:169

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:168
	pool.Put(tx)
											pool.Put(ty)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:172

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:171
	return e
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:175

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:174
func (e *gfP2) Square(a *gfP2, pool *bnPool) *gfP2 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:174
	CoverTab[56052]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:178

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:177
	t1 := pool.Get().Sub(a.y, a.x)
											t2 := pool.Get().Add(a.x, a.y)
											ty := pool.Get().Mul(t1, t2)
											ty.Mod(ty, p)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:183

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:182
	t1.Mul(a.x, a.y)
											t1.Lsh(t1, 1)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:186

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:185
	e.x.Mod(t1, p)
											e.y.Set(ty)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:189

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:188
	pool.Put(t1)
											pool.Put(t2)
											pool.Put(ty)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:193

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:192
	return e
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:196

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:195
func (e *gfP2) Invert(a *gfP2, pool *bnPool) *gfP2 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:195
	CoverTab[9610]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:199

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:198
	t := pool.Get()
											t.Mul(a.y, a.y)
											t2 := pool.Get()
											t2.Mul(a.x, a.x)
											t.Add(t, t2)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:205

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:204
	inv := pool.Get()
											inv.ModInverse(t, p)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:208

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:207
	e.x.Neg(a.x)
											e.x.Mul(e.x, inv)
											e.x.Mod(e.x, p)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:212

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:211
	e.y.Mul(a.y, inv)
											e.y.Mod(e.y, p)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:215

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:214
	pool.Put(t)
											pool.Put(t2)
											pool.Put(inv)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:219

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:218
	return e
}

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/gfp2.go:219
//var _ = .Main
