package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Panjul",
			request: "Panjul",
		},
		{
			name:    "Eko",
			request: "Koko",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Randa", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Randa")
		}
	})
	b.Run("Konate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Konate")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Razzi")
	}
}

func BenchmarkHelloWorldRazzi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Razzi")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Momot",
			request:  "Momot",
			expected: "Holla Momot",
		},
		{
			name:     "Mozza",
			request:  "Mozza",
			expected: "Holla Mozza",
		},
		{
			name:     "Oyen",
			request:  "Oyen",
			expected: "Holla Oyen",
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
	t.Run("Razzi", func(t *testing.T) {
		result := HelloWorld("Razzi")
		require.Equal(t, "Hellow Razzi", result, "Result must be 'Hellow Razzi'")
	})
	t.Run("Catur", func(t *testing.T) {
		result := HelloWorld("Catur")
		require.Equal(t, "Hellow Catur", result, "Result must be 'Hellow Razzi'")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("Before Unit Test")

	m.Run()

	fmt.Println("After Unit Test")
}

// Skip Test
func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac Os")
	}

	result := HelloWorld("Razzi")
	require.Equal(t, "Hellow Razzi", result, "Result must be 'Hellow Razzi'")
}

// Test Assertion
func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Catur")
	assert.Equal(t, "Hellow Catur", result, "Result must be 'Hellow Catur'")
	fmt.Println("TestHelloWorld with Assert Done")
}

// Test Require
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Sofian")
	require.Equal(t, "Hellow Sofian", result, "Result must be 'Hellow Sofian'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Dunia")
	if result != "Hellow Dunia" {
		t.Error("Result must be 'Helow Dunia '")
	}

	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldRazzi(t *testing.T) {
	result := HelloWorld("Razzi")
	if result != "Hellow Razzi" {
		t.Fatal("Result must be 'Hellow Razzi'")
	}

	fmt.Println("TestHelloWorldRazzi Done")
}
