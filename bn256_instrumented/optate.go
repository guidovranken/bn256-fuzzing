// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:5
package main

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:5
//import  "go-fuzz-dep"
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:8

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:7
func lineFunctionAdd(r, p *twistPoint, q *curvePoint, r2 *gfP2, pool *bnPool) (a, b, c *gfP2, rOut *twistPoint) {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:7
	CoverTab[27142]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:12

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:11
	B := newGFp2(pool).Mul(p.x, r.t, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:14

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:13
	D := newGFp2(pool).Add(p.y, r.z)
											D.Square(D, pool)
											D.Sub(D, r2)
											D.Sub(D, r.t)
											D.Mul(D, r.t, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:20

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:19
	H := newGFp2(pool).Sub(B, r.x)
											I := newGFp2(pool).Square(H, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:23

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:22
	E := newGFp2(pool).Add(I, I)
											E.Add(E, E)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:26

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:25
	J := newGFp2(pool).Mul(H, E, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:28

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:27
	L1 := newGFp2(pool).Sub(D, r.y)
											L1.Sub(L1, r.y)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:31

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:30
	V := newGFp2(pool).Mul(r.x, E, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:33

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:32
	rOut = newTwistPoint(pool)
											rOut.x.Square(L1, pool)
											rOut.x.Sub(rOut.x, J)
											rOut.x.Sub(rOut.x, V)
											rOut.x.Sub(rOut.x, V)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:39

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:38
	rOut.z.Add(r.z, H)
											rOut.z.Square(rOut.z, pool)
											rOut.z.Sub(rOut.z, r.t)
											rOut.z.Sub(rOut.z, I)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:44

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:43
	t := newGFp2(pool).Sub(V, rOut.x)
											t.Mul(t, L1, pool)
											t2 := newGFp2(pool).Mul(r.y, J, pool)
											t2.Add(t2, t2)
											rOut.y.Sub(t, t2)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:50

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:49
	rOut.t.Square(rOut.z, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:52

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:51
	t.Add(p.y, rOut.z)
											t.Square(t, pool)
											t.Sub(t, r2)
											t.Sub(t, rOut.t)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:57

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:56
	t2.Mul(L1, p.x, pool)
											t2.Add(t2, t2)
											a = newGFp2(pool)
											a.Sub(t2, t)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:62

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:61
	c = newGFp2(pool)
											c.MulScalar(rOut.z, q.y)
											c.Add(c, c)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:66

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:65
	b = newGFp2(pool)
											b.SetZero()
											b.Sub(b, L1)
											b.MulScalar(b, q.x)
											b.Add(b, b)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:72

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:71
	B.Put(pool)
											D.Put(pool)
											H.Put(pool)
											I.Put(pool)
											E.Put(pool)
											J.Put(pool)
											L1.Put(pool)
											V.Put(pool)
											t.Put(pool)
											t2.Put(pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:83

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:82
	return
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:86

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:85
func lineFunctionDouble(r *twistPoint, q *curvePoint, pool *bnPool) (a, b, c *gfP2, rOut *twistPoint) {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:85
	CoverTab[10513]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:90

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:89
	A := newGFp2(pool).Square(r.x, pool)
											B := newGFp2(pool).Square(r.y, pool)
											C := newGFp2(pool).Square(B, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:94

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:93
	D := newGFp2(pool).Add(r.x, B)
											D.Square(D, pool)
											D.Sub(D, A)
											D.Sub(D, C)
											D.Add(D, D)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:100

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:99
	E := newGFp2(pool).Add(A, A)
											E.Add(E, A)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:103

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:102
	G := newGFp2(pool).Square(E, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:105

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:104
	rOut = newTwistPoint(pool)
											rOut.x.Sub(G, D)
											rOut.x.Sub(rOut.x, D)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:109

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:108
	rOut.z.Add(r.y, r.z)
											rOut.z.Square(rOut.z, pool)
											rOut.z.Sub(rOut.z, B)
											rOut.z.Sub(rOut.z, r.t)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:114

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:113
	rOut.y.Sub(D, rOut.x)
											rOut.y.Mul(rOut.y, E, pool)
											t := newGFp2(pool).Add(C, C)
											t.Add(t, t)
											t.Add(t, t)
											rOut.y.Sub(rOut.y, t)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:121

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:120
	rOut.t.Square(rOut.z, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:123

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:122
	t.Mul(E, r.t, pool)
											t.Add(t, t)
											b = newGFp2(pool)
											b.SetZero()
											b.Sub(b, t)
											b.MulScalar(b, q.x)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:130

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:129
	a = newGFp2(pool)
											a.Add(r.x, E)
											a.Square(a, pool)
											a.Sub(a, A)
											a.Sub(a, G)
											t.Add(B, B)
											t.Add(t, t)
											a.Sub(a, t)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:139

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:138
	c = newGFp2(pool)
											c.Mul(rOut.z, r.t, pool)
											c.Add(c, c)
											c.MulScalar(c, q.y)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:144

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:143
	A.Put(pool)
											B.Put(pool)
											C.Put(pool)
											D.Put(pool)
											E.Put(pool)
											G.Put(pool)
											t.Put(pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:152

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:151
	return
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:155

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:154
func mulLine(ret *gfP12, a, b, c *gfP2, pool *bnPool) {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:154
	CoverTab[65003]++
											a2 := newGFp6(pool)
											a2.x.SetZero()
											a2.y.Set(a)
											a2.z.Set(b)
											a2.Mul(a2, ret.x, pool)
											t3 := newGFp6(pool).MulScalar(ret.y, c, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:163

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:162
	t := newGFp2(pool)
											t.Add(b, c)
											t2 := newGFp6(pool)
											t2.x.SetZero()
											t2.y.Set(a)
											t2.z.Set(t)
											ret.x.Add(ret.x, ret.y)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:171

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:170
	ret.y.Set(t3)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:173

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:172
	ret.x.Mul(ret.x, t2, pool)
											ret.x.Sub(ret.x, a2)
											ret.x.Sub(ret.x, ret.y)
											a2.MulTau(a2, pool)
											ret.y.Add(ret.y, a2)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:179

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:178
	a2.Put(pool)
											t3.Put(pool)
											t2.Put(pool)
											t.Put(pool)
}

// sixuPlus2NAF is 6u+2 in non-adjacent form.
var sixuPlus2NAF = []int8{0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, -1, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, -1, 0, 1, 0, 0, 0, 1, 0, -1, 0, 0, 0, -1, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, -1, 0, -1, 0, 0, 0, 0, 1, 0, 0, 0, 1}

// miller implements the Miller loop for calculating the Optimal Ate pairing.
// See algorithm 1 from http://cryptojedi.org/papers/dclxvi-20100714.pdf
func miller(q *twistPoint, p *curvePoint, pool *bnPool) *gfP12 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:189
	CoverTab[26763]++
											ret := newGFp12(pool)
											ret.SetOne()
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:194

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:193
	aAffine := newTwistPoint(pool)
											aAffine.Set(q)
											aAffine.MakeAffine(pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:198

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:197
	bAffine := newCurvePoint(pool)
											bAffine.Set(p)
											bAffine.MakeAffine(pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:202

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:201
	minusA := newTwistPoint(pool)
											minusA.Negative(aAffine, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:205

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:204
	r := newTwistPoint(pool)
											r.Set(aAffine)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:208

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:207
	r2 := newGFp2(pool)
											r2.Square(aAffine.y, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:211

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:210
	for i := len(sixuPlus2NAF) - 1; i > 0; i-- {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:210
		CoverTab[63127]++
												a, b, c, newR := lineFunctionDouble(r, bAffine, pool)
												if i != len(sixuPlus2NAF)-1 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:212
			CoverTab[5809]++
													ret.Square(ret, pool)
		} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:214
			CoverTab[12654]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:214
		}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:214
		CoverTab[63054]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:217

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:216
		mulLine(ret, a, b, c, pool)
												a.Put(pool)
												b.Put(pool)
												c.Put(pool)
												r.Put(pool)
												r = newR
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:224

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:223
		switch sixuPlus2NAF[i-1] {
		case 1:
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:224
			CoverTab[12359]++
													a, b, c, newR = lineFunctionAdd(r, aAffine, bAffine, r2, pool)
		case -1:
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:226
			CoverTab[2137]++
													a, b, c, newR = lineFunctionAdd(r, minusA, bAffine, r2, pool)
		default:
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:228
			CoverTab[27356]++
													continue
		}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:230
		CoverTab[7143]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:233

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:232
		mulLine(ret, a, b, c, pool)
												a.Put(pool)
												b.Put(pool)
												c.Put(pool)
												r.Put(pool)
												r = newR
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:238
	CoverTab[59381]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:256

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:255
	q1 := newTwistPoint(pool)
											q1.x.Conjugate(aAffine.x)
											q1.x.Mul(q1.x, xiToPMinus1Over3, pool)
											q1.y.Conjugate(aAffine.y)
											q1.y.Mul(q1.y, xiToPMinus1Over2, pool)
											q1.z.SetOne()
											q1.t.SetOne()
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:270

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:269
	minusQ2 := newTwistPoint(pool)
											minusQ2.x.MulScalar(aAffine.x, xiToPSquaredMinus1Over3)
											minusQ2.y.Set(aAffine.y)
											minusQ2.z.SetOne()
											minusQ2.t.SetOne()
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:276

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:275
	r2.Square(q1.y, pool)
											a, b, c, newR := lineFunctionAdd(r, q1, bAffine, r2, pool)
											mulLine(ret, a, b, c, pool)
											a.Put(pool)
											b.Put(pool)
											c.Put(pool)
											r.Put(pool)
											r = newR
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:285

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:284
	r2.Square(minusQ2.y, pool)
											a, b, c, newR = lineFunctionAdd(r, minusQ2, bAffine, r2, pool)
											mulLine(ret, a, b, c, pool)
											a.Put(pool)
											b.Put(pool)
											c.Put(pool)
											r.Put(pool)
											r = newR
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:294

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:293
	aAffine.Put(pool)
											bAffine.Put(pool)
											minusA.Put(pool)
											r.Put(pool)
											r2.Put(pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:300

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:299
	return ret
}

// finalExponentiation computes the (p¹²-1)/Order-th power of an element of
// GF(p¹²) to obtain an element of GT (steps 13-15 of algorithm 1 from
// http://cryptojedi.org/papers/dclxvi-20100714.pdf)
func finalExponentiation(in *gfP12, pool *bnPool) *gfP12 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:305
	CoverTab[5297]++
											t1 := newGFp12(pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:310

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:309
	t1.x.Negative(in.x)
											t1.y.Set(in.y)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:313

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:312
	inv := newGFp12(pool)
											inv.Invert(in, pool)
											t1.Mul(t1, inv, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:317

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:316
	t2 := newGFp12(pool).FrobeniusP2(t1, pool)
											t1.Mul(t1, t2, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:320

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:319
	fp := newGFp12(pool).Frobenius(t1, pool)
											fp2 := newGFp12(pool).FrobeniusP2(t1, pool)
											fp3 := newGFp12(pool).Frobenius(fp2, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:324

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:323
	fu, fu2, fu3 := newGFp12(pool), newGFp12(pool), newGFp12(pool)
											fu.Exp(t1, u, pool)
											fu2.Exp(fu, u, pool)
											fu3.Exp(fu2, u, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:329

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:328
	y3 := newGFp12(pool).Frobenius(fu, pool)
											fu2p := newGFp12(pool).Frobenius(fu2, pool)
											fu3p := newGFp12(pool).Frobenius(fu3, pool)
											y2 := newGFp12(pool).FrobeniusP2(fu2, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:334

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:333
	y0 := newGFp12(pool)
											y0.Mul(fp, fp2, pool)
											y0.Mul(y0, fp3, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:338

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:337
	y1, y4, y5 := newGFp12(pool), newGFp12(pool), newGFp12(pool)
											y1.Conjugate(t1)
											y5.Conjugate(fu2)
											y3.Conjugate(y3)
											y4.Mul(fu, fu2p, pool)
											y4.Conjugate(y4)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:345

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:344
	y6 := newGFp12(pool)
											y6.Mul(fu3, fu3p, pool)
											y6.Conjugate(y6)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:349

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:348
	t0 := newGFp12(pool)
											t0.Square(y6, pool)
											t0.Mul(t0, y4, pool)
											t0.Mul(t0, y5, pool)
											t1.Mul(y3, y5, pool)
											t1.Mul(t1, t0, pool)
											t0.Mul(t0, y2, pool)
											t1.Square(t1, pool)
											t1.Mul(t1, t0, pool)
											t1.Square(t1, pool)
											t0.Mul(t1, y1, pool)
											t1.Mul(t1, y0, pool)
											t0.Square(t0, pool)
											t0.Mul(t0, t1, pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:364

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:363
	inv.Put(pool)
											t1.Put(pool)
											t2.Put(pool)
											fp.Put(pool)
											fp2.Put(pool)
											fp3.Put(pool)
											fu.Put(pool)
											fu2.Put(pool)
											fu3.Put(pool)
											fu2p.Put(pool)
											fu3p.Put(pool)
											y0.Put(pool)
											y1.Put(pool)
											y2.Put(pool)
											y3.Put(pool)
											y4.Put(pool)
											y5.Put(pool)
											y6.Put(pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:383

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:382
	return t0
}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:386

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:385
func optimalAte(a *twistPoint, b *curvePoint, pool *bnPool) *gfP12 {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:385
	CoverTab[2838]++
											e := miller(a, b, pool)
											ret := finalExponentiation(e, pool)
											e.Put(pool)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:391

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:390
	if a.IsInfinity() || func() bool {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:390
		CoverTab[62963]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:390
		return b.IsInfinity()
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:390
	}() {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:390
		CoverTab[39709]++
												ret.SetOne()
	} else {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:392
		CoverTab[44266]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:392
	}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:392
	CoverTab[25040]++
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:395

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:394
	return ret
}

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/optate.go:395
//var _ = .Main
