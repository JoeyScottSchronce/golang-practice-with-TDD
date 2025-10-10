package main

import "testing"

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, but want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDef(t testing.TB, dict Dict, word, def string) {
	t.Helper()

	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("should find added word: ", err)
	}
	assertStrings(t, got, def)
}

func TestSearch(t *testing.T) {
	dict := Dict{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dict.Search("unknown")

		if got == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dict{}
		word := "new"
		def := "This is a new word in the dict"

		err := dict.Add(word, def)
		assertError(t, err, nil)
		assertDef(t, dict, word, def)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "existing"
		def := "This def already exists"
		dict := Dict{word: def}
		err := dict.Add(word, "new def")

		assertError(t, err, ErrWordExists)
		assertDef(t, dict, word, def)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "old word"
		def := "This is the old def"
		dict := Dict{word: def}
		newDef := "This is the new def"

		err := dict.Update(word, newDef)
		assertError(t, err, nil)
		assertDef(t, dict, word, newDef)
	})

	t.Run("new word", func(t *testing.T) {
		word := "new word"
		def := "new def"
		dict := Dict{}

		err := dict.Update(word, def)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete existing word", func(t *testing.T) {
		word := "word"
		def := "This word will be deleted"
		dict := Dict{word: def}

		err := dict.Delete(word)
		assertError(t, err, nil)

		_, err = dict.Search(word)
		assertError(t, err, ErrNotFound)
	})

	t.Run("delete non-existing word", func(t *testing.T) {
		word := "word"
		dict := Dict{}

		err := dict.Delete(word)
		assertError(t, err, ErrWordDoesNotExist)
	})
}
