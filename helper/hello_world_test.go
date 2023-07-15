package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T){
	name := HelloWorld("World")

	if name != "Hello World"{
		t.Fatal("Result must be Hello World")
	}

	fmt.Println("TestHelloWorld Done")
}

func TestHelloMichael(t *testing.T){
	name := HelloWorld("Michael")

	if name != "Hello Michael"{
		t.Fatal("Result must be Hello Michael")
	}

	fmt.Println(("TestHelloMichael Done"))
}