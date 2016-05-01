package parser

import (
	"testing"

	"github.com/willis7/clippings-lib/clipping"
)

func TestParseAuthor(t *testing.T) {
	// Given: a long string
	s := "The 7 Habits of Highly Effective People_ Powerful Lessons in Personal Change (Stephen R. Covey)"

	// When: we run the parse function
	author := parseAuthor(s)

	// Then: the correct name is found
	if got, expected := author, "Stephen R. Covey"; got != expected {
		t.Errorf("The authors do not match. Got: %s, Expected: %s", got, expected)
	}
}

func TestParseTitle(t *testing.T) {
	// Given
	s := "The 7 Habits of Highly Effective People_ Powerful Lessons in Personal Change (Stephen R. Covey)"

	// When: we run the parse function
	title := parseTitle(s)

	// Then: the correct title is found
	if got, expected := title, "The 7 Habits of Highly Effective People_ Powerful Lessons in Personal Change"; got != expected {
		t.Errorf("The titles do not match. Got: %s, Expected: %s", got, expected)
	}
}

func TestSplitClippings(t *testing.T) {
	// Given: a string with a number of clippings
	s := "1 ========== 2 ========== 3 ========== 4 ========== 5"

	// When: I split on clippingSeparator
	clips := splitClippings(s)

	// Then: expect the correct number of clippings
	if got, expected := len(clips), 5; got != expected {
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

	// Then: expect a correctly populated sturcture
	expected := clipping.Clipping{Book: `The 7 Habits of Highly Effective People_ Powerful Lessons in Personal Change`,
		Author:  `Stephen R. Covey`,
		Content: `Habit 1 says “You are the programmer.” Habit 2, then, says, “Write the program.”`}
	if got != expected {
		t.Error(`parseClipping didn't return expected object`)
	}

}
