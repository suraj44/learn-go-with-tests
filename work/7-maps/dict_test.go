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

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dict := Dictionary{word: definition}

		assertDefinition(t, dict, word, definition)

	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		newDefinition := "new definition"

		dict := Dictionary{word: definition}

		err := dict.Add(word, newDefinition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dict, word, definition)

	})
}

func TestUpdate(t *testing.T) {

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		newDefinition := "new definition"

		dict := Dictionary{word: definition}

		err := dict.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dict, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		// definition := "this is just a test"
		newDefinition := "new definition"

		dict := Dictionary{}

		err := dict.Update(word, newDefinition)

		assertError(t, err, ErrWordDoesntExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	definition := "test definition"
	dict := Dictionary{word: definition}

	dict.Delete(word)

	_, err := dict.Search(word)

	assertError(t, err, ErrNotFound)
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

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}
