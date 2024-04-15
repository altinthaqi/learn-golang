package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {

		_, err := dictionary.Search("unknown")

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		k := "test"
		v := "this is just a test"

		dictionary := Dictionary{}
		dictionary.Add(k, v)

		assertDefinition(t, dictionary, k, v)
	})

	t.Run("existing word", func(t *testing.T) {
		k := "test"
		v := "this is just a test"

		dictionary := Dictionary{k: v}
		err := dictionary.Add(k, "duplicate key")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, k, v)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		k := "test"
		v := "this is just a test"

		dictionary := Dictionary{k: v}
		newV := "updated value"

		err := dictionary.Update(k, newV)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, k, newV)
	})

	t.Run("new word", func(t *testing.T) {
		k := "test"
		v := "this is just a test"

		dictionary := Dictionary{}
		err := dictionary.Update(k, v)

		assertError(t, err, ErrWordDoesNotExist)
	})

}

func TestDelete(t *testing.T) {
	k := "test"
	v := "this is just a test"

	dictionary := Dictionary{k: v}
	dictionary.Delete(k)

	_, err := dictionary.Search(k)
	assertError(t, err, ErrNotFound)

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, d Dictionary, k, v string) {
	t.Helper()

	got, err := d.Search(k)
	if err != nil {
		t.Fatal("should have fond the added word: ", err)
	}

	assertStrings(t, got, v)

}
