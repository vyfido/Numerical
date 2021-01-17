package number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsMultiple(t *testing.T) {
	assert.True(t, IsMultiple(4, 2), "the numbers should be multiples")
	assert.True(t, IsMultiple(56, 7), "the numbers should be multiples")
	assert.False(t, IsMultiple(7, 3), "the numbers not should be multiples")
	assert.False(t, IsMultiple(6, 4), "the numbers should be multiples")
}

func TestIsEven(t *testing.T) {
	assert.True(t, IsEven(12), "the number should be even")
	assert.True(t, IsEven(18), "the number should be even")
	assert.False(t, IsEven(7), "the number not should be even")
	assert.False(t, IsEven(9), "the number not should be even")
}

func TestIsPrime(t *testing.T) {
	assert.True(t, IsPrime(2), "the number should be prime")
	assert.True(t, IsPrime(7), "the number should be prime")
	assert.True(t, IsPrime(11), "the number should be prime")
	assert.False(t, IsPrime(20), "the number not should be prime")
	assert.False(t, IsPrime(100), "the number not should be prime")
	assert.False(t, IsPrime(28), "the number not should be prime")
}

func TestFastPrime(t *testing.T) {
	assert.True(t, FastPrime(2), "the number should be prime")
	assert.True(t, FastPrime(7), "the number should be prime")
	assert.True(t, FastPrime(11), "the number should be prime")
	assert.False(t, FastPrime(20), "the number not should be prime")
	assert.False(t, FastPrime(100), "the number not should be prime")
	assert.False(t, FastPrime(28), "the number not should be prime")
}

func TestDescomposePrime(t *testing.T) {
	assert.Equal(t, len(DescomposePrime(6)), 2, "the number of digits not is correct")
	assert.Equal(t, DescomposePrime(6), []int{2, 3}, "the digits not are equals")
	assert.Equal(t, len(DescomposePrime(100)), 4, "the number of digits not is correct")
	assert.Equal(t, DescomposePrime(100), []int{2, 2, 5, 5}, "the digits not are equals")
}

func TestIsPerfect(t *testing.T) {
	assert.True(t, IsPerfect(6), "the number should be perfect")
	assert.True(t, IsPerfect(28), "the number should be perfect")
	assert.True(t, IsPerfect(8128), "the number should be perfect")
	assert.False(t, IsPerfect(20), "the number not should be perfect")
	assert.False(t, IsPerfect(67), "the number not should be perfect")
}

func TestAreFriend(t *testing.T) {
	assert.True(t, AreFriends(220, 284), "the number should be friends")
	assert.False(t, AreFriends(110, 112), "the number not should be friends")
}

func TestCountDigits(t *testing.T) {
	assert.Equal(t, CountDigits(1), 1, "The digits in the count not are the expected")
	assert.Equal(t, CountDigits(4812), 4, "The digits in the count not are the expected")
}

func TestDescomposeDigits(t *testing.T) {
	assert.Equal(t, DescomposeDigits(21), []int{1, 2}, "The digits are not correct")
	assert.Equal(t, DescomposeDigits(3), []int{3}, "The digits are not correct")
	assert.Equal(t, DescomposeDigits(478), []int{8, 7, 4}, "The digits are not correct")
}

func TestReverseNumber(t *testing.T) {
	assert.Equal(t, ReverseNumber(11), 11, "The number are not expected")
}

func TestIsEmirp(t *testing.T) {
	assert.True(t, IsEmirP(13), "The number are not emirp")
	assert.False(t, IsEmirP(18), "The number are emirp")
}
