package helper

import "testing"

func TestHelloWorld(t *testing.T){
	name := HelloWorld("World")

	if name != "Hello World"{
		panic("Test failed.")
	}
}

func TestHelloMichael(t *testing.T){
	name := HelloWorld("Michael")

	if name != "Hello Michael"{
		panic("Test failed.")
	}
}