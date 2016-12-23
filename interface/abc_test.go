package abc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Car struct {
	Name string
	T    *testing.T
}

func (c *Car) Compare(v interface{}) {
	assert.Equal(c.T, c, v, "should be equal")
}

func TestCompare(t *testing.T) {
	a := &Car{
		Name: "Benz",
		T:    t,
	}

	// b := a
	b := &Car{
		Name: "BMW",
		T:    t,
	}
	a.Compare(b)
}
