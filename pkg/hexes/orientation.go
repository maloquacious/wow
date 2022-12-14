/*
 * wars of warp - an implementation of warpwar
 *
 * Copyright (c) 2022 Michael D Henderson
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package hexes

// Orientation stores the forward matrix (the fN variables) and backward matrix
// (the bN variables), plus the start angle. The start angle determines if we
// have a "flat top" (which is 0°) or "pointy top" (which is 60°) hex.
type Orientation struct {
	f0, f1, f2, f3 float64
	b0, b1, b2, b3 float64
	// The starting angle should be 0.0 for 0° (flat top) or 0.5 for 60° (pointy top).
	start_angle float64 // in multiples of 60°
}

// NewOrientation returns an initialized Orientation.
func NewOrientation(f0, f1, f2, f3, b0, b1, b2, b3 float64, hexLayout HEXLAYOUT) Orientation {
	switch hexLayout {
	case FLATHEX:
		return Orientation{
			f0: f0, f1: f1, f2: f2, f3: f3,
			b0: b0, b1: b1, b2: b2, b3: b3,
			start_angle: 0.0,
		}
	case POINTYHEX:
		return Orientation{
			f0: f0, f1: f1, f2: f2, f3: f3,
			b0: b0, b1: b1, b2: b2, b3: b3,
			start_angle: 0.5,
		}
	}
	panic("assert(layoutType in (FLAT, POINTY))")
}
