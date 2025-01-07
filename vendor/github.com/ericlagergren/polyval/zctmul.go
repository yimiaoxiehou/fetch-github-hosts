// Code generated by gen.go. DO NOT EDIT.

package polyval

import "math/bits"

// ctmulGeneric returns the constant time 128-bit product of
// x and y in GF(2^128).
//
// The idea comes from Thomas Pornin's constant-time blog post
// with 64-bit fixes from Tim Taubert's blog post on formally
// verified GHASH.
//
// See https://www.bearssl.org/constanttime.html
// See https://timtaubert.de/blog/2017/06/verified-binary-multiplication-for-ghash/
func ctmulGeneric(x, y uint64) (z1, z0 uint64) {
	// Split both x and y into 5 words with four-bit holes.
	x0 := x & 0x1084210842108421
	y0 := y & 0x1084210842108421
	x1 := x & 0x2108421084210842
	y1 := y & 0x2108421084210842
	x2 := x & 0x4210842108421084
	y2 := y & 0x4210842108421084
	x3 := x & 0x8421084210842108
	y3 := y & 0x8421084210842108
	x4 := x & 0x0842108421084210
	y4 := y & 0x0842108421084210

	// t0 := (x0*y0) ^ (x1*y0) ^ (x2*y0) ^ (x3*y0) ^ (x4*y0)
	// z |= t0 & 0x21084210842108421084210842108421
	x0y0hi, x0y0lo := bits.Mul64(x0, y0)
	x1y4hi, x1y4lo := bits.Mul64(x1, y4)
	x2y3hi, x2y3lo := bits.Mul64(x2, y3)
	x3y2hi, x3y2lo := bits.Mul64(x3, y2)
	x4y1hi, x4y1lo := bits.Mul64(x4, y1)
	z1 |= (x0y0hi ^ x1y4hi ^ x2y3hi ^ x3y2hi ^ x4y1hi) & 0x2108421084210842
	z0 |= (x0y0lo ^ x1y4lo ^ x2y3lo ^ x3y2lo ^ x4y1lo) & 0x1084210842108421

	// t1 := (x0*y1) ^ (x1*y1) ^ (x2*y1) ^ (x3*y1) ^ (x4*y1)
	// z |= t1 & 0x42108421084210842108421084210842
	x0y1hi, x0y1lo := bits.Mul64(x0, y1)
	x1y0hi, x1y0lo := bits.Mul64(x1, y0)
	x2y4hi, x2y4lo := bits.Mul64(x2, y4)
	x3y3hi, x3y3lo := bits.Mul64(x3, y3)
	x4y2hi, x4y2lo := bits.Mul64(x4, y2)
	z1 |= (x0y1hi ^ x1y0hi ^ x2y4hi ^ x3y3hi ^ x4y2hi) & 0x4210842108421084
	z0 |= (x0y1lo ^ x1y0lo ^ x2y4lo ^ x3y3lo ^ x4y2lo) & 0x2108421084210842

	// t2 := (x0*y2) ^ (x1*y2) ^ (x2*y2) ^ (x3*y2) ^ (x4*y2)
	// z |= t2 & 0x84210842108421084210842108421084
	x0y2hi, x0y2lo := bits.Mul64(x0, y2)
	x1y1hi, x1y1lo := bits.Mul64(x1, y1)
	x2y0hi, x2y0lo := bits.Mul64(x2, y0)
	x3y4hi, x3y4lo := bits.Mul64(x3, y4)
	x4y3hi, x4y3lo := bits.Mul64(x4, y3)
	z1 |= (x0y2hi ^ x1y1hi ^ x2y0hi ^ x3y4hi ^ x4y3hi) & 0x8421084210842108
	z0 |= (x0y2lo ^ x1y1lo ^ x2y0lo ^ x3y4lo ^ x4y3lo) & 0x4210842108421084

	// t3 := (x0*y3) ^ (x1*y3) ^ (x2*y3) ^ (x3*y3) ^ (x4*y3)
	// z |= t3 & 0x8421084210842108421084210842108
	x0y3hi, x0y3lo := bits.Mul64(x0, y3)
	x1y2hi, x1y2lo := bits.Mul64(x1, y2)
	x2y1hi, x2y1lo := bits.Mul64(x2, y1)
	x3y0hi, x3y0lo := bits.Mul64(x3, y0)
	x4y4hi, x4y4lo := bits.Mul64(x4, y4)
	z1 |= (x0y3hi ^ x1y2hi ^ x2y1hi ^ x3y0hi ^ x4y4hi) & 0x0842108421084210
	z0 |= (x0y3lo ^ x1y2lo ^ x2y1lo ^ x3y0lo ^ x4y4lo) & 0x8421084210842108

	// t4 := (x0*y4) ^ (x1*y4) ^ (x2*y4) ^ (x3*y4) ^ (x4*y4)
	// z |= t4 & 0x1084210842108421842108421084210
	x0y4hi, x0y4lo := bits.Mul64(x0, y4)
	x1y3hi, x1y3lo := bits.Mul64(x1, y3)
	x2y2hi, x2y2lo := bits.Mul64(x2, y2)
	x3y1hi, x3y1lo := bits.Mul64(x3, y1)
	x4y0hi, x4y0lo := bits.Mul64(x4, y0)
	z1 |= (x0y4hi ^ x1y3hi ^ x2y2hi ^ x3y1hi ^ x4y0hi) & 0x1084210842108421
	z0 |= (x0y4lo ^ x1y3lo ^ x2y2lo ^ x3y1lo ^ x4y0lo) & 0x0842108421084210

	return
}
