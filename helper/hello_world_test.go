package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T){
	name := HelloWorld("World")

	if name != "Hello World"{
		t.Fail()
	}

	fmt.Println("TestHelloWorld Done")
}

func TestHelloMichael(t *testing.T){
	name := HelloWorld("Michael")

	if name != "Hello Michael"{
		t.FailNow()
	}

	fmt.Println(("TestHelloMichael Done"))
}