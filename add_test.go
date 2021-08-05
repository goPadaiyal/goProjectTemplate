package gopopper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	. "testing"
)

func TestSum(t *T) {
	actualResult := Add(19, -9)
	assert.Equal(t, 10, actualResult)
}

func TestMain(m *M) {
	// call flag.Parse() here if TestMain uses flags
	fmt.Println("Before each test.")
	os.Exit(m.Run())
	fmt.Println("After each test.")
}
