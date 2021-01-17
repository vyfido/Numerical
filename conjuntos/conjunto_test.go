package conjuntos

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChangeName(t *testing.T) {
	var conjunto = Conjunto{}
	conjunto.ChangeName("letters")
	assert.Equal(t, conjunto.name, "letters", "the name not is valid")
}

func TestFind(t *testing.T) {
	var conjunto = Conjunto{}
	conjunto.AddMany("1", "3", "5", "7", "9")
	assert.Equal(t, conjunto.Find("3"), 1, "the position is incorrect")
	assert.Equal(t, conjunto.Find("0"), LOSS, "the position is incorrect")
}

func TestGet(t *testing.T) {
	var conjunto = Conjunto{}
	conjunto.name = "letters"
	conjunto.AddMany("a", "e", "i", "o", "u")
	assert.Equal(t, conjunto.Get(2), "i", "the value not is correct")
	assert.Equal(t, conjunto.Get(-4), NULL, "the value not is correct")
	assert.Equal(t, conjunto.Get(6), NULL, "the value not is correct")
}

func TestSize(t *testing.T) {
	var conjunto = Conjunto{}
	assert.Equal(t, conjunto.Size(), EMPTY, "The size not is correct")
	conjunto.Add("a")
	assert.Equal(t, conjunto.Size(), UNARY, "The size not is correct")
	conjunto.Add("a")
	assert.Equal(t, conjunto.Size(), UNARY, "The size not is correct")
}

func TestExists(t *testing.T) {
	var conjunto = Conjunto{}
	conjunto.AddMany("a,e,i,o,u")
	assert.Equal(t, conjunto.Size(), 5, "The size not is correct")
	assert.True(t, conjunto.Exists("i"), "The value not is correct")
	assert.True(t, conjunto.Exists("u"), "The value not is correct")
	assert.True(t, conjunto.Exists("a"), "The value not is correct")
	assert.False(t, conjunto.Exists("s"), "The value not is correct")
}

func TestIsEmpty(t *testing.T) {
	var conjunto = Conjunto{}
	assert.True(t, conjunto.IsEmpty(), "The conjunto should be empty")
	conjunto.Add("a")
	assert.False(t, conjunto.IsEmpty(), "the conjunto not should be empty")
}

func TestIsUnary(t *testing.T) {
	var conjunto = Conjunto{}
	conjunto.Add("1")
	assert.True(t, conjunto.IsUnary(), "The conjunto should be unary")
	conjunto.Add("2")
	assert.False(t, conjunto.IsUnary(), "The conjunto should be unary")
}

func TestAdd(t *testing.T) {
	var conjunto = Conjunto{}
	assert.Equal(t, conjunto.Size(), 0, "The conjunto should be empty")
	conjunto.Add("a")
	conjunto.Add("b")
	conjunto.Add("c")
	assert.Equal(t, conjunto.Size(), 3, "The size should be three")
	conjunto.Add("d")
	assert.Equal(t, conjunto.Size(), 4, "The size should be four")
}

func TestAddMany(t *testing.T) {
	var conjunto = Conjunto{}
	assert.Equal(t, conjunto.Size(), 0, "The conjunto should be empty")
	conjunto.AddMany("a", "b", "c", "d", "e", "f", "g", "h", "i", "j")
	assert.Equal(t, conjunto.Size(), 10, "The size should be ten")
}

func TestString(t *testing.T) {
	var conjunto = Conjunto{}
	conjunto.name = "vocals"
	conjunto.AddMany("a", "e", "i", "o", "u")
	assert.Equal(t, conjunto.String(), "vocals={a,e,i,o,u}", "The convert to String is not equal")
	var numbers = Conjunto{}
	numbers.name = "numbers"
	numbers.AddMany("1", "2", "3", "4", "5")
	assert.Equal(t, numbers.String(), "numbers={1,2,3,4,5}", "The convert to String is not equal")
}

func TestEquals(t *testing.T) {
	var vocals = Conjunto{name: "vocals", elements: []string{"a", "e", "i", "o", "u"}}
	var letters = Conjunto{name: "letters", elements: []string{"a", "b", "c", "d", "e"}}
	var other = Conjunto{name: "other", elements: []string{"a", "e", "i", "o", "u"}}

	assert.True(t, vocals.Equals(other), "The values are incorrect")
	assert.False(t, letters.Equals(vocals), "The values are incorrect")

}
