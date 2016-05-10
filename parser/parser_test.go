package parser

import (
	"testing"

	"github.com/willis7/Alice/clipping"
)

// Tips
// use the t.Log and/or t.Logf methods if you need to print information in a test.

func TestParseAuthor(t *testing.T) {
	// Given: a long string
	s := "The 7 Habits of Highly Effective People_ Powerful Lessons in Personal Change (Stephen R. Covey)"

	// When: we run the parse function
	author := extractAuthor(s)

	// Then: the correct name is found
	if got, expected := author, "Stephen R. Covey"; got != expected {
		t.Errorf("The authors do not match. Got: %s, Expected: %s", got, expected)
	}
}

func TestParseTitle(t *testing.T) {
	// Given
	s := "The 7 Habits of Highly Effective People_ Powerful Lessons in Personal Change (Stephen R. Covey)"

	// When: we run the parse function
	title := extractTitle(s)

	// Then: the correct title is found
	if got, expected := title, "The 7 Habits of Highly Effective People_ Powerful Lessons in Personal Change"; got != expected {
		t.Errorf("The titles do not match. Got: %s, Expected: %s", got, expected)
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
	if got, expected := len(clips), 3; got != expected {
		t.Errorf("spitClippings(%q) returned %d clippings, expected %d", s, got, expected)
	}
}

func TestParseClipping(t *testing.T) {
	// Given: a clipping which is correctly formatted and an empty clipping object
	s := `The 7 Habits of Highly Effective People_ Powerful Lessons in Personal Change (Stephen R. Covey)
	- Your Highlight at location 1763-1763 | Added on Sunday, 4 January 2015 22:25:46

	Habit 1 says “You are the programmer.” Habit 2, then, says, “Write the program.”
	==========`
	got := clipping.Clipping{}

	// When: I parse the clipping
	parseClipping(s, &got)

	// Then: expect a correctly populated structure
	expected := clipping.Clipping{Book: `The 7 Habits of Highly Effective People_ Powerful Lessons in Personal Change`,
		Author:  `Stephen R. Covey`,
		Content: `Habit 1 says “You are the programmer.” Habit 2, then, says, “Write the program.”`}
	if got != expected {
		t.Error(`parseClipping didn't return expected object`)
	}
}

// Integration test
func TestParse(t *testing.T) {
	// Given: a file with 4 clippings
	s := `clippings-test.txt`

	// When: I give a file to the Parse function
	clips := Parse(s)

	// Then: an array of Clippings is created
	if got, expected := len(clips), 3; got != expected {
		t.Errorf("Parse didnt return the correct number of objects. Got %d, Expected %d", got, expected)
	}

}
