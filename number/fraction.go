package number

import (
	"fmt"
)

type Fraction struct {
	uni int
	num int
	den int
}

func NewFraction(u int, n int, d int) Fraction {
	if d == 0 {
		panic("Denominator never can be zero(0)")
	} else {
		return Fraction{uni: u, num: n, den: d}
	}
}

func (f Fraction) String() string {
	return fmt.Sprintf("%d %d/%d", f.uni, f.num, f.den)
}

func (f Fraction) ToDecimal() float32 {
	if f.den != 0 {
		return float32(f.uni) + float32(f.num)/float32(f.den)
	} else {
		panic("The value of denominator is zero(0)")
	}
}

func (f Fraction) SetUnit(unit int) {
	f.uni = unit
}

func (f Fraction) GetUnit() int {
	return f.uni
}

func (f Fraction) GetNumerator() int {
	return f.num
}

func (f Fraction) GetDenominator() int {
	return f.den
}

func (f Fraction) SetNumerator(numerator int) {
	f.num = numerator
}

func (f Fraction) SetDenominator(denominator int) {
	if denominator == 0 {
		panic("Denominator never can be zero(0)")
	} else {
		f.den = denominator
	}

}

func Add(g Fraction, f Fraction) Fraction {
	if f.den == g.den {
		return NewFraction(f.uni+g.uni, f.num+g.num, f.den)
	} else {
		return NewFraction(0, 0, 1)
	}
}

func (f Fraction) Add(g Fraction) {
	if f.den == g.den {
		f.uni += g.uni
		f.num += g.num
	}
}

func (f Fraction) Signo() int {
	var value float32 = f.ToDecimal()
	if value == 0 {
		return 0
	} else if value > 0 {
		return 1
	} else {
		return -1
	}
}

func (f Fraction) Equals(g Fraction) bool {
	return f.uni == g.uni && f.num == g.num && f.den == g.den
}

func (f Fraction) TypeFraction() string {
	if f.num < f.den {
		return "propia"
	} else {
		return "impropia"
	}
}

func Simplify(f Fraction) Fraction {
	var g Fraction = f
	for ind := 2; ind < g.den; {
		if g.num%ind == 0 && g.den%ind == 0 {
			g.num /= ind
			g.den /= ind
		} else {
			ind++
		}
	}
	if g.uni >= g.den {
		g.uni = g.num / g.den
		g.num -= (g.uni * g.den)
	}
	return g
}
