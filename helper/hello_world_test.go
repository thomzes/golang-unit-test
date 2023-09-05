package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Thomas",
			request:  "Thomas",
			expected: "Hello Thomas",
		},
		{
			name:     "Ardiansah",
			request:  "Ardiansah",
			expected: "Hello Ardiansah",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Thomas", func(t *testing.T) {
		result := HelloWorld("Thomas")
		require.Equal(t, "Hello Thomas", result, "Result must be 'Hello Thomas'")
	})

	t.Run("Ardiansah", func(t *testing.T) {
		result := HelloWorld("Ardiansah")
		require.Equal(t, "Hello Ardiansah", result, "Result must be 'Hello Ardiansah'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("before unit test")
	m.Run()

	//	After
	fmt.Println("after unit test")

}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac Os")
	}

	result := HelloWorld("Thomas")
	require.Equal(t, "Hello Thomas", result, "Result must be 'Hello Thomas'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Thomas")
	require.Equal(t, "Hello Thomas", result, "Result must be 'Hello Thomas'")
	fmt.Println("TestHelloWorld with Require done!")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Thomas")
	assert.Equal(t, "Hello Thomas", result, "Result must be 'Hello Thomas'")
	fmt.Println("TestHelloWorld with Assert done!")
}

func TestHelloWorldThomas(t *testing.T) {
	result := HelloWorld("Thomas")

	if result != "Hello Thomas" {
		//		error
		t.Error("result must be 'Hello Thomas'")
	}

	fmt.Println("TestHelloWorldThomas Done")

}

func TestHelloWorldArdiansah(t *testing.T) {
	result := HelloWorld("Ardiansah")

	if result != "Hello Ardiansah" {
		//		error
		t.Fatal("result must be 'Hello Ardiansah'")
	}

	fmt.Println("TestHelloWorldArdiansah Done")

}
