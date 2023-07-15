package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M){
	//before
	fmt.Println("---- UNIT TEST START ----")

	m.Run()

	//after
	fmt.Println("---- UNIT TEST END----")
	
}

func TestHelloWorld(t *testing.T){
	name := HelloWorld("World")
	
	require.Equal(t, name, "Hello World", "Result must be Hello World")

	fmt.Println("TestHelloWorld With Require Done")
}

func TestHelloMichael(t *testing.T){
	name := HelloWorld("Michael")

	assert.Equal(t, name, "Hello Michael", "Result must be Hello Michael")

	fmt.Println(("TestHelloMichael With Assert Done"))
}

func TestSkip(t *testing.T){
	if runtime.GOOS == "darwin" {
		t.Skip("Ooopss. Can't run on MacOS")
	}
}