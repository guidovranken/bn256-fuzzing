// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/constants.go:5
package main

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/constants.go:5
//import  "go-fuzz-dep"
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/constants.go:8

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/constants.go:7
import (
	"math/big"
)
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/constants.go:12

//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/constants.go:11
func bigFromBase10(s string) *big.Int {
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/constants.go:11
	CoverTab[45203]++
												n, _ := new(big.Int).SetString(s, 10)
												return n
}

// u is the BN parameter that determines the prime: 1868033³.
var u = bigFromBase10("6518589491078791937")

// p is a prime over which we form a basic field: 36u⁴+36u³+24u²+6u+1.
var p = bigFromBase10("65000549695646603732796438742359905742825358107623003571877145026864184071783")

// Order is the number of elements in both G₁ and G₂: 36u⁴+36u³+18u²+6u+1.
var Order = bigFromBase10("65000549695646603732796438742359905742570406053903786389881062969044166799969")

// xiToPMinus1Over6 is ξ^((p-1)/6) where ξ = i+3.
var xiToPMinus1Over6 = &gfP2{bigFromBase10("8669379979083712429711189836753509758585994370025260553045152614783263110636"), bigFromBase10("19998038925833620163537568958541907098007303196759855091367510456613536016040")}

// xiToPMinus1Over3 is ξ^((p-1)/3) where ξ = i+3.
var xiToPMinus1Over3 = &gfP2{bigFromBase10("26098034838977895781559542626833399156321265654106457577426020397262786167059"), bigFromBase10("15931493369629630809226283458085260090334794394361662678240713231519278691715")}

// xiToPMinus1Over2 is ξ^((p-1)/2) where ξ = i+3.
var xiToPMinus1Over2 = &gfP2{bigFromBase10("50997318142241922852281555961173165965672272825141804376761836765206060036244"), bigFromBase10("38665955945962842195025998234511023902832543644254935982879660597356748036009")}

// xiToPSquaredMinus1Over3 is ξ^((p²-1)/3) where ξ = i+3.
var xiToPSquaredMinus1Over3 = bigFromBase10("65000549695646603727810655408050771481677621702948236658134783353303381437752")

// xiTo2PSquaredMinus2Over3 is ξ^((2p²-2)/3) where ξ = i+3 (a cubic root of unity, mod p).
var xiTo2PSquaredMinus2Over3 = bigFromBase10("4985783334309134261147736404674766913742361673560802634030")

// xiToPSquaredMinus1Over6 is ξ^((1p²-1)/6) where ξ = i+3 (a cubic root of -1, mod p).
var xiToPSquaredMinus1Over6 = bigFromBase10("65000549695646603727810655408050771481677621702948236658134783353303381437753")

// xiTo2PMinus2Over3 is ξ^((2p-2)/3) where ξ = i+3.
var xiTo2PMinus2Over3 = &gfP2{bigFromBase10("19885131339612776214803633203834694332692106372356013117629940868870585019582"), bigFromBase10("21645619881471562101905880913352894726728173167203616652430647841922248593627")}
//line /tmp/go-fuzz-build854356423/gopath/src/golang.org/x/crypto/bn256/constants.go:44
//var _ = .Main
