package maps

import "testing"

const testString = "this is just a test"

func TestSearch(t *testing.T) {
	d := map[string]string{
		"test": testString,
	}
	testKey := "test"
	got := Search(d, testKey)
	want := testString
	if got != want {
		t.Errorf("got %q want %q given %q", got, want, testKey)
	}
}
