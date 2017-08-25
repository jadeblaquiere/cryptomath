// Copyright (c) 2017, Joseph deBlaquiere <jadeblaquiere@yahoo.com>
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// * Neither the name of ciphrtxt nor the names of its
//   contributors may be used to endorse or promote products derived from
//   this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package ecgo

import (
	"math/big"
	"testing"
)

func TestAdd(t *testing.T) {
	var p big.Int

	p.SetString("1021", 10)

	f := PrimeField(&p)

	for a := 0; a < 1021; a++ {
		for b := 0; b < 1021; b++ {
			var aa, bb, cc big.Int
			aa.SetInt64(int64(a))
			bb.SetInt64(int64(b))

			aaa := f.Int(&aa)
			bbb := f.Int(&bb)

			aaa.Add(aaa, bbb)
			c := (a + b) % 1021
			cc.SetInt64(int64(c))
			tc := f.Int(&cc)

			if tc.Cmp(aaa) != 0 {
				t.Fail()
			}
		}
	}
}

func TestSub(t *testing.T) {
	var p big.Int

	p.SetString("1021", 10)

	f := PrimeField(&p)

	for a := 0; a < 1021; a++ {
		for b := 0; b < 1021; b++ {
			var aa, bb, cc big.Int
			aa.SetInt64(int64(a))
			bb.SetInt64(int64(b))

			aaa := f.Int(&aa)
			bbb := f.Int(&bb)

			aaa.Sub(aaa, bbb)
			c := (a - b) % 1021
			cc.SetInt64(int64(c))
			tc := f.Int(&cc)

			if tc.Cmp(aaa) != 0 {
				t.Fail()
			}
		}
	}
}

func TestMul(t *testing.T) {
	var p big.Int

	p.SetString("1021", 10)

	f := PrimeField(&p)

	for a := 0; a < 1021; a++ {
		for b := 0; b < 1021; b++ {
			var aa, bb, cc big.Int
			aa.SetInt64(int64(a))
			bb.SetInt64(int64(b))

			aaa := f.Int(&aa)
			bbb := f.Int(&bb)

			aaa.Mul(aaa, bbb)
			c := (a * b) % 1021
			cc.SetInt64(int64(c))
			tc := f.Int(&cc)

			if tc.Cmp(aaa) != 0 {
				t.Fail()
			}
		}
	}
}

func TestInv(t *testing.T) {
	var p big.Int

	p.SetString("65521", 10)
	pmin2 := new(big.Int).SetInt64(int64(2))
	pmin2.Sub(&p, pmin2)

	f := PrimeField(&p)
	one := f.Int(new(big.Int).SetInt64(int64(1)))

	for a := 1; a < 65521; a++ {
		var aa big.Int
		aa.SetInt64(int64(a))

		aaa := f.Int(&aa)
		bbb := new(FieldInt).Inv(aaa)
		ccc := new(FieldInt).Mul(aaa, bbb)

		aa.Exp(&aa, pmin2, &p)
		tbb := f.Int(&aa)

		if tbb.Cmp(bbb) != 0 {
			t.Fail()
		}

		if one.Cmp(ccc) != 0 {
			t.Fail()
		}
	}
}

func TestDiv(t *testing.T) {
	var p big.Int

	p.SetString("1021", 10)
	pmin2 := new(big.Int).SetInt64(int64(2))
	pmin2.Sub(&p, pmin2)

	f := PrimeField(&p)

	for a := 0; a < 1021; a++ {
		for b := 1; b < 1021; b++ {
			var aa, bb, cc big.Int
			aa.SetInt64(int64(a))
			bb.SetInt64(int64(b))

			aaa := f.Int(&aa)
			bbb := f.Int(&bb)

			ccc := new(FieldInt).Div(aaa, bbb)
			ddd := new(FieldInt).Mul(ccc, bbb)

			cc.Exp(&bb, pmin2, &p)
			cc.Mul(&cc, &aa)
			cc.Mod(&cc, &p)
			tc := f.Int(&cc)

			if ddd.Cmp(aaa) != 0 {
				t.Fail()
			}

			if ccc.Cmp(tc) != 0 {
				t.Fail()
			}
		}
	}
}

func TestExp(t *testing.T) {
	var p big.Int

	p.SetString("1021", 10)
	pmin2 := new(big.Int).SetInt64(int64(2))
	pmin2.Sub(&p, pmin2)

	f := PrimeField(&p)

	for a := 0; a < 1021; a++ {
		for b := 1; b < 1021; b++ {
			var aa, bb, cc big.Int
			aa.SetInt64(int64(a))
			bb.SetInt64(int64(b))

			aaa := f.Int(&aa)

			ccc := new(FieldInt).Exp(aaa, &bb)

			cc.Exp(&aa, &bb, &p)
			tc := f.Int(&cc)

			if ccc.Cmp(tc) != 0 {
				t.Fail()
			}
		}
	}
}
