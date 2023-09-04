package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

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
