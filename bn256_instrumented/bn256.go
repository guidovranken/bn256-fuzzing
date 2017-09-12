// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:5
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

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:17
//import  "go-fuzz-dep"
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:20

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:19
import (
	"crypto/rand"
	"io"
	"math/big"
)

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:28
// G1 is an abstract cyclic group. The zero value is suitable for use as the
// output of an operation, but cannot be used as an input.
type G1 struct {
	p *curvePoint
}

// RandomG1 returns x and g₁ˣ where x is a random, non-zero number read from r.
func RandomG1(r io.Reader) (*big.Int, *G1, error) {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:35
	CoverTab[12319]++
											var k *big.Int
											var err error
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:40

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:39
	for {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:39
		CoverTab[54119]++
												k, err = rand.Int(r, Order)
												if err != nil {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:41
			CoverTab[46365]++
													return nil, nil, err
		} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:43
			CoverTab[41321]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:43
		}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:43
		CoverTab[19735]++
												if k.Sign() > 0 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:44
			CoverTab[16291]++
													break
		} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:46
			CoverTab[397]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:46
		}
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:47
	CoverTab[20619]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:50

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:49
	return k, new(G1).ScalarBaseMult(k), nil
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:53

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:52
func (g *G1) String() string {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:52
	CoverTab[36676]++
											return "bn256.G1" + g.p.String()
}

// ScalarBaseMult sets e to g*k where g is the generator of the group and
// then returns e.
func (e *G1) ScalarBaseMult(k *big.Int) *G1 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:58
	CoverTab[17069]++
											if e.p == nil {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:59
		CoverTab[37386]++
												e.p = newCurvePoint(nil)
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:61
		CoverTab[1547]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:61
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:61
	CoverTab[36024]++
											e.p.Mul(curveGen, k, new(bnPool))
											return e
}

// ScalarMult sets e to a*k and then returns e.
func (e *G1) ScalarMult(a *G1, k *big.Int) *G1 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:67
	CoverTab[1439]++
											if e.p == nil {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:68
		CoverTab[21132]++
												e.p = newCurvePoint(nil)
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:70
		CoverTab[26538]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:70
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:70
	CoverTab[53499]++
											e.p.Mul(a.p, k, new(bnPool))
											return e
}

// Add sets e to a+b and then returns e.
// BUG(agl): this function is not complete: a==b fails.
func (e *G1) Add(a, b *G1) *G1 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:77
	CoverTab[43261]++
											if e.p == nil {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:78
		CoverTab[62871]++
												e.p = newCurvePoint(nil)
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:80
		CoverTab[19107]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:80
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:80
	CoverTab[2792]++
											e.p.Add(a.p, b.p, new(bnPool))
											return e
}

// Neg sets e to -a and then returns e.
func (e *G1) Neg(a *G1) *G1 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:86
	CoverTab[63019]++
											if e.p == nil {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:87
		CoverTab[29613]++
												e.p = newCurvePoint(nil)
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:89
		CoverTab[62064]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:89
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:89
	CoverTab[19780]++
											e.p.Negative(a.p)
											return e
}

// Marshal converts n to a byte slice.
func (n *G1) Marshal() []byte {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:95
	CoverTab[36216]++
											n.p.MakeAffine(nil)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:99

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:98
	xBytes := new(big.Int).Mod(n.p.x, p).Bytes()
											yBytes := new(big.Int).Mod(n.p.y, p).Bytes()

											// Each value is a 256-bit number.
											const numBytes = 256 / 8
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:105

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:104
	ret := make([]byte, numBytes*2)
											copy(ret[1*numBytes-len(xBytes):], xBytes)
											copy(ret[2*numBytes-len(yBytes):], yBytes)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:109

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:108
	return ret
}

// Unmarshal sets e to the result of converting the output of Marshal back into
// a group element and then returns e.
func (e *G1) Unmarshal(m []byte) (*G1, bool) {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:113
	CoverTab[27608]++
											// Each value is a 256-bit number.
											const numBytes = 256 / 8
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:118

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:117
	if len(m) != 2*numBytes {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:117
		CoverTab[35344]++
												return nil, false
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:119
		CoverTab[10593]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:119
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:119
	CoverTab[5050]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:122

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:121
	if e.p == nil {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:121
		CoverTab[49181]++
												e.p = newCurvePoint(nil)
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:123
		CoverTab[64282]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:123
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:123
	CoverTab[18400]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:126

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:125
	e.p.x.SetBytes(m[0*numBytes : 1*numBytes])
											e.p.y.SetBytes(m[1*numBytes : 2*numBytes])
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:129

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:128
	if e.p.x.Sign() == 0 && func() bool {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:128
		CoverTab[13617]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:128
		return e.p.y.Sign() == 0
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:128
	}() {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:128
		CoverTab[63298]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:131

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:130
		e.p.y.SetInt64(1)
												e.p.z.SetInt64(0)
												e.p.t.SetInt64(0)
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:133
		CoverTab[16177]++
												e.p.z.SetInt64(1)
												e.p.t.SetInt64(1)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:138

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:137
		if !e.p.IsOnCurve() {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:137
			CoverTab[5773]++
													return nil, false
		} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:139
			CoverTab[48244]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:139
		}
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:140
	CoverTab[49027]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:143

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:142
	return e, true
}

// G2 is an abstract cyclic group. The zero value is suitable for use as the
// output of an operation, but cannot be used as an input.
type G2 struct {
	p *twistPoint
}

// RandomG1 returns x and g₂ˣ where x is a random, non-zero number read from r.
func RandomG2(r io.Reader) (*big.Int, *G2, error) {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:152
	CoverTab[45153]++
											var k *big.Int
											var err error
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:157

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:156
	for {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:156
		CoverTab[47082]++
												k, err = rand.Int(r, Order)
												if err != nil {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:158
			CoverTab[35791]++
													return nil, nil, err
		} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:160
			CoverTab[19588]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:160
		}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:160
		CoverTab[40895]++
												if k.Sign() > 0 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:161
			CoverTab[52121]++
													break
		} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:163
			CoverTab[50247]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:163
		}
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:164
	CoverTab[603]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:167

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:166
	return k, new(G2).ScalarBaseMult(k), nil
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:170

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:169
func (g *G2) String() string {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:169
	CoverTab[44866]++
											return "bn256.G2" + g.p.String()
}

// ScalarBaseMult sets e to g*k where g is the generator of the group and
// then returns out.
func (e *G2) ScalarBaseMult(k *big.Int) *G2 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:175
	CoverTab[52073]++
											if e.p == nil {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:176
		CoverTab[16086]++
												e.p = newTwistPoint(nil)
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:178
		CoverTab[13767]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:178
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:178
	CoverTab[47206]++
											e.p.Mul(twistGen, k, new(bnPool))
											return e
}

// ScalarMult sets e to a*k and then returns e.
func (e *G2) ScalarMult(a *G2, k *big.Int) *G2 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:184
	CoverTab[32175]++
											if e.p == nil {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:185
		CoverTab[37250]++
												e.p = newTwistPoint(nil)
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:187
		CoverTab[60722]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:187
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:187
	CoverTab[14386]++
											e.p.Mul(a.p, k, new(bnPool))
											return e
}

// Add sets e to a+b and then returns e.
// BUG(agl): this function is not complete: a==b fails.
func (e *G2) Add(a, b *G2) *G2 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:194
	CoverTab[56788]++
											if e.p == nil {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:195
		CoverTab[18844]++
												e.p = newTwistPoint(nil)
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:197
		CoverTab[54681]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:197
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:197
	CoverTab[58168]++
											e.p.Add(a.p, b.p, new(bnPool))
											return e
}

// Marshal converts n into a byte slice.
func (n *G2) Marshal() []byte {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:203
	CoverTab[21261]++
											n.p.MakeAffine(nil)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:207

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:206
	xxBytes := new(big.Int).Mod(n.p.x.x, p).Bytes()
											xyBytes := new(big.Int).Mod(n.p.x.y, p).Bytes()
											yxBytes := new(big.Int).Mod(n.p.y.x, p).Bytes()
											yyBytes := new(big.Int).Mod(n.p.y.y, p).Bytes()

											// Each value is a 256-bit number.
											const numBytes = 256 / 8
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:215

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:214
	ret := make([]byte, numBytes*4)
											copy(ret[1*numBytes-len(xxBytes):], xxBytes)
											copy(ret[2*numBytes-len(xyBytes):], xyBytes)
											copy(ret[3*numBytes-len(yxBytes):], yxBytes)
											copy(ret[4*numBytes-len(yyBytes):], yyBytes)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:221

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:220
	return ret
}

// Unmarshal sets e to the result of converting the output of Marshal back into
// a group element and then returns e.
func (e *G2) Unmarshal(m []byte) (*G2, bool) {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:225
	CoverTab[39071]++
											// Each value is a 256-bit number.
											const numBytes = 256 / 8
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:230

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:229
	if len(m) != 4*numBytes {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:229
		CoverTab[41775]++
												return nil, false
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:231
		CoverTab[31429]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:231
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:231
	CoverTab[19323]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:234

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:233
	if e.p == nil {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:233
		CoverTab[23615]++
												e.p = newTwistPoint(nil)
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:235
		CoverTab[36435]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:235
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:235
	CoverTab[64422]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:238

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:237
	e.p.x.x.SetBytes(m[0*numBytes : 1*numBytes])
											e.p.x.y.SetBytes(m[1*numBytes : 2*numBytes])
											e.p.y.x.SetBytes(m[2*numBytes : 3*numBytes])
											e.p.y.y.SetBytes(m[3*numBytes : 4*numBytes])
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:243

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:242
	if e.p.x.x.Sign() == 0 && func() bool {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:242
		CoverTab[61591]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:242
		return e.p.x.y.Sign() == 0
	}() && func() bool {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:243
		CoverTab[24810]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:243
		return e.p.y.x.Sign() == 0
	}() && func() bool {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:244
		CoverTab[60512]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:244
		return e.p.y.y.Sign() == 0
	}() {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:245
		CoverTab[31982]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:248

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:247
		e.p.y.SetOne()
												e.p.z.SetZero()
												e.p.t.SetZero()
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:250
		CoverTab[39595]++
												e.p.z.SetOne()
												e.p.t.SetOne()
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:255

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:254
		if !e.p.IsOnCurve() {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:254
			CoverTab[49488]++
													return nil, false
		} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:256
			CoverTab[54972]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:256
		}
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:257
	CoverTab[33006]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:260

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:259
	return e, true
}

// GT is an abstract cyclic group. The zero value is suitable for use as the
// output of an operation, but cannot be used as an input.
type GT struct {
	p *gfP12
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:269

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:268
func (g *GT) String() string {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:268
	CoverTab[56239]++
											return "bn256.GT" + g.p.String()
}

// ScalarMult sets e to a*k and then returns e.
func (e *GT) ScalarMult(a *GT, k *big.Int) *GT {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:273
	CoverTab[12417]++
											if e.p == nil {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:274
		CoverTab[50623]++
												e.p = newGFp12(nil)
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:276
		CoverTab[17328]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:276
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:276
	CoverTab[15278]++
											e.p.Exp(a.p, k, new(bnPool))
											return e
}

// Add sets e to a+b and then returns e.
func (e *GT) Add(a, b *GT) *GT {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:282
	CoverTab[26879]++
											if e.p == nil {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:283
		CoverTab[49054]++
												e.p = newGFp12(nil)
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:285
		CoverTab[4374]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:285
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:285
	CoverTab[23002]++
											e.p.Mul(a.p, b.p, new(bnPool))
											return e
}

// Neg sets e to -a and then returns e.
func (e *GT) Neg(a *GT) *GT {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:291
	CoverTab[42227]++
											if e.p == nil {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:292
		CoverTab[43815]++
												e.p = newGFp12(nil)
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:294
		CoverTab[4530]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:294
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:294
	CoverTab[62186]++
											e.p.Invert(a.p, new(bnPool))
											return e
}

// Marshal converts n into a byte slice.
func (n *GT) Marshal() []byte {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:300
	CoverTab[25358]++
											n.p.Minimal()
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:304

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:303
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
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:320

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:319
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
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:334

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:333
	return ret
}

// Unmarshal sets e to the result of converting the output of Marshal back into
// a group element and then returns e.
func (e *GT) Unmarshal(m []byte) (*GT, bool) {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:338
	CoverTab[65448]++
											// Each value is a 256-bit number.
											const numBytes = 256 / 8
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:343

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:342
	if len(m) != 12*numBytes {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:342
		CoverTab[35215]++
												return nil, false
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:344
		CoverTab[32854]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:344
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:344
	CoverTab[65367]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:347

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:346
	if e.p == nil {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:346
		CoverTab[17664]++
												e.p = newGFp12(nil)
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:348
		CoverTab[54213]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:348
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:348
	CoverTab[65042]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:351

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:350
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
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:364

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:363
	return e, true
}

// Pair calculates an Optimal Ate pairing.
func Pair(g1 *G1, g2 *G2) *GT {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:367
	CoverTab[65301]++
											return &GT{optimalAte(g2.p, g1.p, new(bnPool))}
}

// bnPool implements a tiny cache of *big.Int objects that's used to reduce the
// number of allocations made during processing.
type bnPool struct {
	bns	[]*big.Int
	count	int
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:379

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:378
func (pool *bnPool) Get() *big.Int {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:378
	CoverTab[36131]++
											if pool == nil {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:379
		CoverTab[20298]++
												return new(big.Int)
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:381
		CoverTab[18565]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:381
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:381
	CoverTab[20680]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:384

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:383
	pool.count++
											l := len(pool.bns)
											if l == 0 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:385
		CoverTab[9070]++
												return new(big.Int)
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:387
		CoverTab[7088]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:387
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:387
	CoverTab[45917]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:390

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:389
	bn := pool.bns[l-1]
											pool.bns = pool.bns[:l-1]
											return bn
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:395

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:394
func (pool *bnPool) Put(bn *big.Int) {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:394
	CoverTab[48093]++
											if pool == nil {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:395
		CoverTab[37026]++
												return
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:397
		CoverTab[645]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:397
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:397
	CoverTab[39380]++
											pool.bns = append(pool.bns, bn)
											pool.count--
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:403

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:402
func (pool *bnPool) Count() int {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:402
	CoverTab[43672]++
											return pool.count
}

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/bn256.go:404
//var _ = .Main
