// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package color

import (
	"testing"
)

func delta(x, y uint8) uint8 {
	if x >= y {
		return x - y
	}
	return y - x
}

// TestYCbCrRoundtrip tests that a subset of RGB space can be converted to YCbCr
// and back to within 1/256 tolerance.
func TestYCbCrRoundtrip(t *testing.T) {
	for r := 0; r < 256; r += 7 {
		for g := 0; g < 256; g += 5 {
			for b := 0; b < 256; b += 3 {
				r0, g0, b0 := uint8(r), uint8(g), uint8(b)
				y, cb, cr := RGBToYCbCr(r0, g0, b0)
				r1, g1, b1 := YCbCrToRGB(y, cb, cr)
				if delta(r0, r1) > 1 || delta(g0, g1) > 1 || delta(b0, b1) > 1 {
					t.Fatalf("\nr0, g0, b0 = %d, %d, %d\nr1, g1, b1 = %d, %d, %d", r0, g0, b0, r1, g1, b1)
				}
			}
		}
	}
}

// TestYCbCrToRGBConsistency tests that calling the RGBA method (16 bit color)
// then truncating to 8 bits is equivalent to calling the YCbCrToRGB function (8
// bit color).
func TestYCbCrToRGBConsistency(t *testing.T) {
	for y := 0; y < 256; y += 7 {
		for cb := 0; cb < 256; cb += 5 {
			for cr := 0; cr < 256; cr += 3 {
				x := YCbCr{uint8(y), uint8(cb), uint8(cr)}
				r0, g0, b0, _ := x.RGBA()
				r1, g1, b1 := uint8(r0>>8), uint8(g0>>8), uint8(b0>>8)
				r2, g2, b2 := YCbCrToRGB(x.Y, x.Cb, x.Cr)
				if r1 != r2 || g1 != g2 || b1 != b2 {
					t.Fatalf("y, cb, cr = %d, %d, %d\nr1, g1, b1 = %d, %d, %d\nr2, g2, b2 = %d, %d, %d",
						y, cb, cr, r1, g1, b1, r2, g2, b2)
				}
			}
		}
	}
}

// TestCMYKRoundtrip tests that a subset of RGB space can be converted to CMYK
// and back to within 1/256 tolerance.
func TestCMYKRoundtrip(t *testing.T) {
	for r := 0; r < 256; r += 7 {
		for g := 0; g < 256; g += 5 {
			for b := 0; b < 256; b += 3 {
				r0, g0, b0 := uint8(r), uint8(g), uint8(b)
				c, m, y, k := RGBToCMYK(r0, g0, b0)
				r1, g1, b1 := CMYKToRGB(c, m, y, k)
				if delta(r0, r1) > 1 || delta(g0, g1) > 1 || delta(b0, b1) > 1 {
					t.Fatalf("\nr0, g0, b0 = %d, %d, %d\nr1, g1, b1 = %d, %d, %d", r0, g0, b0, r1, g1, b1)
				}
			}
		}
	}
}

// TestCMYKToRGBConsistency tests that calling the RGBA method (16 bit color)
// then truncating to 8 bits is equivalent to calling the CMYKToRGB function (8
// bit color).
func TestCMYKToRGBConsistency(t *testing.T) {
	for c := 0; c < 256; c += 7 {
		for m := 0; m < 256; m += 5 {
			for y := 0; y < 256; y += 3 {
				for k := 0; k < 256; k += 11 {
					x := CMYK{uint8(c), uint8(m), uint8(y), uint8(k)}
					r0, g0, b0, _ := x.RGBA()
					r1, g1, b1 := uint8(r0>>8), uint8(g0>>8), uint8(b0>>8)
					r2, g2, b2 := CMYKToRGB(x.C, x.M, x.Y, x.K)
					if r1 != r2 || g1 != g2 || b1 != b2 {
						t.Fatalf("c, m, y, k = %d, %d, %d, %d\nr1, g1, b1 = %d, %d, %d\nr2, g2, b2 = %d, %d, %d",
							c, m, y, k, r1, g1, b1, r2, g2, b2)
					}
				}
			}
		}
	}
}
