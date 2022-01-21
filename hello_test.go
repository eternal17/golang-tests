package main

// It needs to be in a file with a name like xxx_test.go
// The test function must start with the word Test
// The test function takes one argument only t *testing.T

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		// Helper is needed to tell the test suite that this method is a helper.
		// By doing this when it fails the line number reported will be in our function call
		//  rather than outside rather than inside our test helper.
		// This will help other developers track down problems easier.
		t.Helper()
		// if test fails
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In French", func(t *testing.T) {
		got := Hello("Thiery", "French")
		want := "Bonjour, Thiery"
		assertCorrectMessage(t, got, want)
	})
}
