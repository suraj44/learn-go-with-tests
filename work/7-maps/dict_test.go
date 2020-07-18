package dict

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})

}

func assertError(t *testing.T, got, want error) {
	if got != want {
		t.Errorf("got %q want %q, given, %q", got, want, "test")
	}
}

func assertStrings(t *testing.T, got, want string) {
	if got != want {
		t.Errorf("got %q want %q, given, %q", got, want, "test")
	}
}
