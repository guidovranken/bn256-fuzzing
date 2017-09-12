// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:5
package main

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:5
//import  "go-fuzz-dep"
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:8

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:7
import (
	"math/big"
)

// twistPoint implements the elliptic curve y²=x³+3/ξ over GF(p²). Points are
// kept in Jacobian form and t=z² when valid. The group G₂ is the set of
// n-torsion points of this curve over GF(p²) (where n = Order)
type twistPoint struct {
	x, y, z, t *gfP2
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:19

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:18
var twistB = &gfP2{
	bigFromBase10("6500054969564660373279643874235990574282535810762300357187714502686418407178"),
	bigFromBase10("45500384786952622612957507119651934019977750675336102500314001518804928850249"),
}

// twistGen is the generator of group G₂.
var twistGen = &twistPoint{
	&gfP2{
		bigFromBase10("21167961636542580255011770066570541300993051739349375019639421053990175267184"),
		bigFromBase10("64746500191241794695844075326670126197795977525365406531717464316923369116492"),
	},
	&gfP2{
		bigFromBase10("20666913350058776956210519119118544732556678129809273996262322366050359951122"),
		bigFromBase10("17778617556404439934652658462602675281523610326338642107814333856843981424549"),
	},
	&gfP2{
		bigFromBase10("0"),
		bigFromBase10("1"),
	},
	&gfP2{
		bigFromBase10("0"),
		bigFromBase10("1"),
	},
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:44

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:43
func newTwistPoint(pool *bnPool) *twistPoint {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:43
	CoverTab[3129]++
											return &twistPoint{
		newGFp2(pool),
		newGFp2(pool),
		newGFp2(pool),
		newGFp2(pool),
	}
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:53

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:52
func (c *twistPoint) String() string {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:52
	CoverTab[32294]++
											return "(" + c.x.String() + ", " + c.y.String() + ", " + c.z.String() + ")"
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:57

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:56
func (c *twistPoint) Put(pool *bnPool) {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:56
	CoverTab[542]++
											c.x.Put(pool)
											c.y.Put(pool)
											c.z.Put(pool)
											c.t.Put(pool)
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:64

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:63
func (c *twistPoint) Set(a *twistPoint) {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:63
	CoverTab[20743]++
											c.x.Set(a.x)
											c.y.Set(a.y)
											c.z.Set(a.z)
											c.t.Set(a.t)
}

// IsOnCurve returns true iff c is on the curve where c must be in affine form.
func (c *twistPoint) IsOnCurve() bool {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:71
	CoverTab[39989]++
											pool := new(bnPool)
											yy := newGFp2(pool).Square(c.y, pool)
											xxx := newGFp2(pool).Square(c.x, pool)
											xxx.Mul(xxx, c.x, pool)
											yy.Sub(yy, xxx)
											yy.Sub(yy, twistB)
											yy.Minimal()
											return yy.x.Sign() == 0 && func() bool {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:79
		CoverTab[13405]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:79
		return yy.y.Sign() == 0
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:79
	}()
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:83

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:82
func (c *twistPoint) SetInfinity() {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:82
	CoverTab[39183]++
											c.z.SetZero()
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:87

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:86
func (c *twistPoint) IsInfinity() bool {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:86
	CoverTab[30461]++
											return c.z.IsZero()
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:91

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:90
func (c *twistPoint) Add(a, b *twistPoint, pool *bnPool) {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:90
	CoverTab[47593]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:94

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:93
	if a.IsInfinity() {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:93
		CoverTab[30003]++
												c.Set(b)
												return
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:96
		CoverTab[26644]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:96
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:96
	CoverTab[41071]++
											if b.IsInfinity() {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:97
		CoverTab[54004]++
												c.Set(a)
												return
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:100
		CoverTab[10375]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:100
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:100
	CoverTab[51540]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:104

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:103
	z1z1 := newGFp2(pool).Square(a.z, pool)
											z2z2 := newGFp2(pool).Square(b.z, pool)
											u1 := newGFp2(pool).Mul(a.x, z2z2, pool)
											u2 := newGFp2(pool).Mul(b.x, z1z1, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:109

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:108
	t := newGFp2(pool).Mul(b.z, z2z2, pool)
											s1 := newGFp2(pool).Mul(a.y, t, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:112

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:111
	t.Mul(a.z, z1z1, pool)
											s2 := newGFp2(pool).Mul(b.y, t, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:115

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:114
	h := newGFp2(pool).Sub(u2, u1)
											xEqual := h.IsZero()
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:118

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:117
	t.Add(h, h)
											i := newGFp2(pool).Square(t, pool)
											j := newGFp2(pool).Mul(h, i, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:122

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:121
	t.Sub(s2, s1)
											yEqual := t.IsZero()
											if xEqual && func() bool {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:123
		CoverTab[45837]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:123
		return yEqual
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:123
	}() {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:123
		CoverTab[3734]++
												c.Double(a, pool)
												return
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:126
		CoverTab[22712]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:126
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:126
	CoverTab[20409]++
											r := newGFp2(pool).Add(t, t)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:130

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:129
	v := newGFp2(pool).Mul(u1, i, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:132

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:131
	t4 := newGFp2(pool).Square(r, pool)
											t.Add(v, v)
											t6 := newGFp2(pool).Sub(t4, j)
											c.x.Sub(t6, t)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:137

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:136
	t.Sub(v, c.x)
											t4.Mul(s1, j, pool)
											t6.Add(t4, t4)
											t4.Mul(r, t, pool)
											c.y.Sub(t4, t6)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:143

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:142
	t.Add(a.z, b.z)
											t4.Square(t, pool)
											t.Sub(t4, z1z1)
											t4.Sub(t, z2z2)
											c.z.Mul(t4, h, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:149

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:148
	z1z1.Put(pool)
											z2z2.Put(pool)
											u1.Put(pool)
											u2.Put(pool)
											t.Put(pool)
											s1.Put(pool)
											s2.Put(pool)
											h.Put(pool)
											i.Put(pool)
											j.Put(pool)
											r.Put(pool)
											v.Put(pool)
											t4.Put(pool)
											t6.Put(pool)
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:165

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:164
func (c *twistPoint) Double(a *twistPoint, pool *bnPool) {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:164
	CoverTab[58939]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:167

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:166
	A := newGFp2(pool).Square(a.x, pool)
											B := newGFp2(pool).Square(a.y, pool)
											C := newGFp2(pool).Square(B, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:171

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:170
	t := newGFp2(pool).Add(a.x, B)
											t2 := newGFp2(pool).Square(t, pool)
											t.Sub(t2, A)
											t2.Sub(t, C)
											d := newGFp2(pool).Add(t2, t2)
											t.Add(A, A)
											e := newGFp2(pool).Add(t, A)
											f := newGFp2(pool).Square(e, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:180

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:179
	t.Add(d, d)
											c.x.Sub(f, t)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:183

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:182
	t.Add(C, C)
											t2.Add(t, t)
											t.Add(t2, t2)
											c.y.Sub(d, c.x)
											t2.Mul(e, c.y, pool)
											c.y.Sub(t2, t)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:190

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:189
	t.Mul(a.y, a.z, pool)
											c.z.Add(t, t)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:193

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:192
	A.Put(pool)
											B.Put(pool)
											C.Put(pool)
											t.Put(pool)
											t2.Put(pool)
											d.Put(pool)
											e.Put(pool)
											f.Put(pool)
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:203

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:202
func (c *twistPoint) Mul(a *twistPoint, scalar *big.Int, pool *bnPool) *twistPoint {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:202
	CoverTab[33973]++
											sum := newTwistPoint(pool)
											sum.SetInfinity()
											t := newTwistPoint(pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:208

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:207
	for i := scalar.BitLen(); i >= 0; i-- {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:207
		CoverTab[10069]++
												t.Double(sum, pool)
												if scalar.Bit(i) != 0 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:209
			CoverTab[46963]++
													sum.Add(t, a, pool)
		} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:211
			CoverTab[20263]++
													sum.Set(t)
		}
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:214
	CoverTab[33935]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:217

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:216
	c.Set(sum)
											sum.Put(pool)
											t.Put(pool)
											return c
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:223

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:222
func (c *twistPoint) MakeAffine(pool *bnPool) *twistPoint {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:222
	CoverTab[12726]++
											if c.z.IsOne() {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:223
		CoverTab[52277]++
												return c
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:225
		CoverTab[7855]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:225
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:225
	CoverTab[55511]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:228

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:227
	zInv := newGFp2(pool).Invert(c.z, pool)
											t := newGFp2(pool).Mul(c.y, zInv, pool)
											zInv2 := newGFp2(pool).Square(zInv, pool)
											c.y.Mul(t, zInv2, pool)
											t.Mul(c.x, zInv2, pool)
											c.x.Set(t)
											c.z.SetOne()
											c.t.SetOne()
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:237

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:236
	zInv.Put(pool)
											t.Put(pool)
											zInv2.Put(pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:241

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:240
	return c
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:244

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:243
func (c *twistPoint) Negative(a *twistPoint, pool *bnPool) {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:243
	CoverTab[63125]++
											c.x.Set(a.x)
											c.y.SetZero()
											c.y.Sub(c.y, a.y)
											c.z.Set(a.z)
											c.t.SetZero()
}

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/twist.go:249
//var _ = .Main
