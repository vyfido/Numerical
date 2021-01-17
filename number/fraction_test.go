package number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {
	var f1 = Fraction{uni: 0, num: 3, den: 8}
	var f2 = Fraction{uni: 1, num: 1, den: 3}
	var f3 = NewFraction(2, 1, 7)
	assert.Equal(t, f1.String(), "0 3/8")
	assert.Equal(t, f2.String(), "1 1/3")
	assert.Equal(t, f3.String(), "2 1/7")
}

func TestReal(t *testing.T) {
	var f1 = Fraction{uni: 0, num: 1, den: 4}
	assert.Equal(t, f1.ToDecimal(), float32(0.25))
}
