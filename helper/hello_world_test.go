package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing.B){
	for i := 0; i < b.N; i++ {
		HelloWorld("Taylor Swift")
	}
}

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

func TestSubTest(t *testing.T){
	t.Run("John", func(t *testing.T) {
		result := HelloWorld("John")

		assert.Equal(t, result, "Hello John", "Result must be Hello John")
	})

	t.Run("Jordan", func(t *testing.T) {
		result := HelloWorld("Jordan")

		assert.Equal(t, result, "Hello Jordan", "Result must be Hello Jordan")
	})
}

func TestTableHelloWorld(t *testing.T){
	tests := []struct {
		name string
		request string
		expected string
	}{
		{
			name: "TestWayan",
			request: "Wayan",
			expected: "Hello Wayan",
		},
		{
			name: "TestMade",
			request: "Made",
			expected: "Hello Made",
		},
		{
			name: "TestKomang",
			request: "Komang",
			expected: "Hello Komang",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)

			require.Equal(t, result, test.expected, "Result must be " + test.expected)
		})
	}
}