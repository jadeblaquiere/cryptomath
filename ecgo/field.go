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

//import "fmt"
import "math/big"

// Field defines a finite (Galois) integer field meeting the defintion as
// explained in https://en.wikipedia.org/wiki/Finite_field
type Field interface {
	// order returns the order of the field
	Order() *big.Int

	// Int creates an element within the field
	Int(*big.Int) *FieldInt
}

type primefield struct {
	p big.Int
}

func (f *primefield) Order() *big.Int {
	return &f.p
}

func (z *primefield) Int(v *big.Int) *FieldInt {
	i := new(FieldInt)
	i.p.Set(&z.p)
	i.v.Mod(v, &i.p)
	return i
}

func PrimeField(p *big.Int) Field {
	f := new(primefield)
	f.p.Set(p)
	return f
}

type FieldInt struct {
	v big.Int
	p big.Int
}

func (z *FieldInt) Add(x, y *FieldInt) *FieldInt {
	if x.p.Cmp(&y.p) != 0 {
		panic("Addition not in the same Field")
	}
	z.p.Set(&x.p)
	z.v.Add(&x.v, &y.v)
	z.v.Mod(&z.v, &x.p)
	return z
}

func (z *FieldInt) Sub(x, y *FieldInt) *FieldInt {
	if x.p.Cmp(&y.p) != 0 {
		panic("Subtraction not in the same Field")
	}
	z.p.Set(&x.p)
	z.v.Sub(&x.v, &y.v)
	z.v.Mod(&z.v, &x.p)
	return z
}

func (z *FieldInt) Mul(x, y *FieldInt) *FieldInt {
	if x.p.Cmp(&y.p) != 0 {
		panic("Multiplication not in the same Field")
	}
	z.p.Set(&x.p)
	z.v.Mul(&x.v, &y.v)
	z.v.Mod(&z.v, &x.p)
	return z
}

func (z *FieldInt) Inv(x *FieldInt) *FieldInt {
	//r := new(fint)
	z.p.Set(&x.p)
	z.v.ModInverse(&x.v, &x.p)
	return z
}

func (z *FieldInt) Div(x, y *FieldInt) *FieldInt {
	if x.p.Cmp(&y.p) != 0 {
		panic("Division not in the same Field")
	}
	z.Inv(y)
	z.v.Mul(&x.v, &z.v)
	z.v.Mod(&z.v, &x.p)
	return z
}

func (f *FieldInt) String() string {
	return f.v.String() + "(mod " + f.p.String() + ")"
}

func (z *FieldInt) Cmp(y *FieldInt) int {
	if z.p.Cmp(&y.p) != 0 {
		panic("Comparison not in the same Field")
	}
	return z.v.Cmp(&y.v)
}

func (z *FieldInt) Exp(x *FieldInt, y *big.Int) *FieldInt {
	z.p.Set(&x.p)
	z.v.Exp(&x.v, y, &x.p)
	return z
}
