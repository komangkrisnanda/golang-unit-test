package helper

import "testing"

func TestHelloWorld(t *testing.T){
	name := HelloWorld("John")

	if name != "Hello John"{
		panic("Test failed.")
	}
}