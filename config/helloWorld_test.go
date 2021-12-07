package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// Before
	fmt.Println("Before Test")

	m.Run()

	// After
	fmt.Println("After Test")
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Saepudin")
	assert.Equal(t, "Hi Saepudin", result, "result must be 'Result Saepudin'")
	fmt.Println("Test Done!")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Saepudin")
	require.Equal(t, "Hii Saepudin", result, "result must be 'Result Saepudin'")

	fmt.Println("Test Done!")
}
