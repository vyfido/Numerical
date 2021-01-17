package str

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	assert.Equal(t, Reverse("hola"), "aloh", "The text not is the expected")
	assert.Equal(t, Reverse("amor"), "roma", "The text not is the expected")
}

func TestCapitalize(t *testing.T) {
	assert.Equal(t, Capitalize("gato"), "Gato", "The text not is the expected")
	assert.Equal(t, Capitalize("MELON"), "Melon", "The text not is the expected")
}

func TestInvertir(t *testing.T) {
	assert.Equal(t, Invertir("maya"), "MAYA", "The text not is the expected")
	assert.Equal(t, Invertir("MAYA"), "maya", "The text not is the expected")
	assert.Equal(t, Invertir("mAyA"), "MaYa", "The text not is the expected")
}

func TestRotate(t *testing.T) {
	assert.Equal(t, Rotate("maya", 1), "ayam", "The text not is the expected")
	assert.Equal(t, Rotate("maya", 2), "yama", "The text not is the expected")
	assert.Equal(t, Rotate("maya", 3), "amay", "The text not is the expected")
	assert.Equal(t, Rotate("maya", 4), "maya", "The text not is the expected")
}
