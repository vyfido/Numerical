package conjuntos

import (
	"strings"
)

const (
	//represent the empty set
	EMPTY int = 0
	//represent the one element
	UNARY int = 1
	//numerical value when not exists record
	LOSS int = -1
	//literal value when not exists record
	NULL string = ""
)

// structured by represent the conjuntos the values
type Conjunto struct {
	name     string
	elements []string
}

func (conjunto *Conjunto) Size() int {
	return len(conjunto.elements)
}

func (conjunto *Conjunto) Find(element string) int {
	for ind, elem := range conjunto.elements {
		if elem == element {
			return ind
		}
	}
	return LOSS
}

func (conjunto *Conjunto) Get(pos int) string {
	if pos >= 0 && pos < conjunto.Size() {
		return conjunto.elements[pos]
	} else {
		return NULL
	}
}

func (conjunto *Conjunto) ChangeName(name string) {
	conjunto.name = name
}

func (conjunto *Conjunto) Exists(element string) bool {
	if conjunto.Find(element) != LOSS {
		return true
	} else {
		return false
	}
}

func (conjunto *Conjunto) IsEmpty() bool {
	if conjunto.Size() == EMPTY {
		return true
	} else {
		return false
	}
}

func (conjunto *Conjunto) IsUnary() bool {
	if conjunto.Size() == UNARY {
		return true
	} else {
		return false
	}
}

func (conjunto *Conjunto) Add(element string) {
	var elements []string = strings.Split(element, ",")
	for ind := range elements {
		if !conjunto.Exists(elements[ind]) {
			conjunto.elements = append(conjunto.elements, strings.TrimSpace(elements[ind]))
		}
	}
}

func (conjunto *Conjunto) AddMany(eles ...string) {
	for _, element := range eles {
		if !conjunto.Exists(element) {
			conjunto.Add(element)
		}
	}
}

func (conjunto *Conjunto) String() string {
	var s string = ""
	if conjunto.name == "" {
		s += "UNKWON"
	} else {
		s += conjunto.name
	}

	if conjunto.Size() > 0 {
		var max int = conjunto.Size()
		s += "={"
		for ind := 0; ind < max; ind++ {
			s += conjunto.Get(ind)
			if ind < max-1 {
				s += ","
			}
		}
		s += "}"
	} else {
		s += "={}"
	}

	return strings.TrimSpace(s)
}

func (conjunto *Conjunto) Equals(a Conjunto) bool {
	if conjunto.Size() == a.Size() {
		var cont int = 0
		for _, a.elements := range element {
             if conjunto.Exists(element){
				cont++;
			}
		}
		return cont == conjunto.Size()
	} else {
		return false
	}
}
