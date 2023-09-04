package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
