package sexpr

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int      `sexpr:"year"`
		Oscars          []string `sexpr:"oscars"`
		Sequel          *string  `sexpr:"sequel"`
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	// Encode it
	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// Check it
	want := `((Title "Dr. Strangelove")` +
		` (Subtitle "How I Learned to Stop Worrying and Love the Bomb")` +
		` (year 1964)` +
		` (oscars` +
		` ("Best Actor (Nomin.)"` +
		` "Best Adapted Screenplay (Nomin.)"` +
		` "Best Director (Nomin.)"` +
		` "Best Picture (Nomin.)"))` +
		` (sequel nil))`
	got := string(data)
	if want != got {
		t.Fatalf("Marshal failed: want: %s\ngot: %s\n", want, got)
	}

	// Decode it
	var movie Movie
	if err := Unmarshal(data, &movie); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	// Check equality.
	if !reflect.DeepEqual(movie, strangelove) {
		t.Fatal("not equal")
	}
}
