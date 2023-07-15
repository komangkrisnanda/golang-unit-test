package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T){
	name := HelloWorld("World")
	
	assert.Equal(t, name, "Hello World", "Result must be Hello World")

	fmt.Println("TestHelloWorld With Assert Done")
}

func TestHelloMichael(t *testing.T){
	name := HelloWorld("Michael")

	assert.Equal(t, name, "Hello Michael", "Result must be Hello Michael")

	fmt.Println(("TestHelloMichael With Assert Done"))
}