package parser

import (
	"testing"

	"github.com/willis7/Alice/clipping"
	"path/filepath"
)

// Tips
// use the t.Log and/or t.Logf methods if you need to print information in a test.

func TestExtractAuthor(t *testing.T) {
	// Given: a long string
	s := "The 7 Habits of Highly Effective People_ Powerful Lessons in Personal Change (Stephen R. Covey)"

	// When: we run the parse function
	author := extractAuthor(s)

	// Then: the correct name is found
	if actual, expected := author, "Stephen R. Covey"; actual != expected {
		t.Errorf("The authors do not match. Actual: %s, Expected: %s", actual, expected)
	}
}

func TestExtractTitle(t *testing.T) {
	// Given
	s := "The 7 Habits of Highly Effective People_ Powerful Lessons in Personal Change (Stephen R. Covey)"

	// When: we run the parse function
	title := extractTitle(s)

	// Then: the correct title is found
	if actual, expected := title, "The 7 Habits of Highly Effective People_ Powerful Lessons in Personal Change"; actual != expected {
		t.Errorf("The titles do not match. Actual: %s, Expected: %s", actual, expected)
	}
}

func TestSplitClippings(t *testing.T) {
	// Given: a string with a number of clippings
	s := `The 7 Habits of Highly Effective People_ Powerful Lessons in Personal Change (Stephen R. Covey)
	- Your Highlight at location 1258-1259 | Added on Saturday, 3 January 2015 23:04:21

	“Management is doing things right; leadership is doing the right things.”
	==========
	The 7 Habits of Highly Effective People_ Powerful Lessons in Personal Change (Stephen R. Covey)
	- Your Highlight at location 1268-1269 | Added on Saturday, 3 January 2015 23:06:54

	We are more in need of a vision or designation and a compass (a set of principles or directions) and less in need of a road map.
	==========
`

	// When: I split on clippingSeparator
	clips := splitClippings(s)

	// Then: expect the correct number of clippings
	if actual, expected := len(clips), 3; actual != expected {
		t.Errorf("spitClippings(%q) returned %d clippings, expected %d", s, actual, expected)
	}
}

func TestParseClipping(t *testing.T) {
	// Given: a clipping which is correctly formatted and an empty clipping object
	s := `The 7 Habits of Highly Effective People_ Powerful Lessons in Personal Change (Stephen R. Covey)
	- Your Highlight at location 1763-1763 | Added on Sunday, 4 January 2015 22:25:46

	Habit 1 says “You are the programmer.” Habit 2, then, says, “Write the program.”
	==========`
	actual := clipping.Clipping{}

	// When: I parse the clipping
	parseClipping(s, &actual)

	// Then: expect a correctly populated structure
	expected := clipping.Clipping{Book: `The 7 Habits of Highly Effective People_ Powerful Lessons in Personal Change`,
		Author:  `Stephen R. Covey`,
		Content: `Habit 1 says “You are the programmer.” Habit 2, then, says, “Write the program.”`}
	if actual != expected {
		t.Error(`parseClipping didn't return expected object`)
	}
}

func TestIsTypeBookmark(t *testing.T) {
	// Given: a set of cases
	var cases = []struct {
		Input    string
		Expected bool
	}{
		{"- Your Highlight at location 1763-1763", false},
		{"Your Bookmark at location 333", false},
		{"- Your Bookmark at location 333", true},
	}

	// When: I iterate over the cases calling the isTypeBookmark function
	for _, tc := range cases {
		actual := isTypeBookmark(tc.Input)

		// Then: I get the correct output for the tests
		if actual != tc.Expected {
			t.Errorf("Actual: %b, Expected: %b", actual, tc.Expected)
		}
	}
}

// Integration test
func TestParse(t *testing.T) {
	// Given: a file with 4 clippings
	s := filepath.Join("test-fixtures", "clippings-test.txt")

	// When: I give a file to the Parse function
	clips := Parse(s)

	// Then: an array of Clippings is created
	if got, expected := len(clips), 3; got != expected {
		t.Errorf("Parse didnt return the correct number of objects. Got %d, Expected %d", got, expected)
	}

}

// TODO: write a test to cover a scenario where entry is a bookmark
