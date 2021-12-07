package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Saepudin")

	if result != "Hello Agung" {
		// panic("Result is not Hello Agung")
		t.Error("Result must be 'Hello Agung'")
	}

	fmt.Println("Done!")

}

func TestHelloWorld1(t *testing.T) {
	result := HelloWorld("Agung")

	if result != "Hello Saepudin" {
		t.Fatal("Result must be 'Hello Saepudin'")
	}

	fmt.Println("Done!!!!")
}
