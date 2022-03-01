package maps

import "testing"

const (
	testKey    = "test"
	testString = "this is just a test"
)

func TestSearch(t *testing.T) {
	d := Dictionary{testKey: testString}

	t.Run("known word", func(t *testing.T) {
		got, _ := d.Search(testKey)
		assertStrings(t, got, testString, testKey)
	})
	t.Run("unknown word", func(t *testing.T) {
		unknownTerm := "unknownTerm"
		_, err := d.Search(unknownTerm)
		assertError(t, err, ErrNotFound, unknownTerm)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		d := Dictionary{}
		err := d.Add(testKey, testString)
		assertError(t, err, nil, "NA(new word test)")
		assertDefinition(t, d, testKey, testString)
	})
	t.Run("existing word", func(t *testing.T) {
		d := Dictionary{testKey: testString}
		err := d.Add(testKey, "some other string")
		assertError(t, err, ErrKeyExists, "NA(existing word test)")
		assertDefinition(t, d, testKey, testString)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		d := Dictionary{testKey: testString}
		newDefinition := "some new string"
		err := d.Update(testKey, newDefinition)
		assertError(t, err, nil, "NA(existing word test)")
		assertDefinition(t, d, testKey, newDefinition)
	})
	t.Run("new word", func(t *testing.T) {
		d := Dictionary{}
		err := d.Update(testKey, testString)
		assertError(t, err, ErrWordDoesNotExist, "NA(new word test)")
	})
}

func assertError(t testing.TB, got, want error, given string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given %q", got, want, given)
	}
}

func assertDefinition(t testing.TB, d Dictionary, k, v string) {
	t.Helper()

	got, err := d.Search(k)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if v != got {
		t.Errorf("got %q want %q", got, v)
	}
}

func assertStrings(t testing.TB, got, want, given string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given %q", got, want, given)
	}
}
