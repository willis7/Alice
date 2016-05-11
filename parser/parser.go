package parser

import (
	"bufio"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/willis7/Alice/clipping"
)

const clippingSeparator string = "==========\n"
const authorRegxp string = `\((.*?)\)` // find a bracket pair and all characters inbetween
const titleRegxp string = `^(.*?)\(`   // find everything up to, and including a "("
const bookmarkRegxp string = `- Your Bookmark at location`   // find everything up to, and including a "("

// extractTitle
// uses a regular expression to find the title in a Kindle formatted string.
// The Kindle always puts the title before the author.
func extractTitle(input string) string {
	// TODO: switch to MustCompile. Panics when pattern not valid
	r, _ := regexp.Compile(titleRegxp)
	match := r.FindString(input)

	if len(match) < 2 {
		panic("match: less than 2")
	}
	title := match[: len(match) - 2]
	return title
}

// extractAuthor
// uses a regular expression to find the author name in a Kindle formatted string.
// The Kindle always puts the author name between braces (<<author>>).
func extractAuthor(input string) string {
	// TODO: switch to MustCompile. Panics when pattern not valid
	r, _ := regexp.Compile(authorRegxp)
	match := r.FindString(input)

	if len(match) < 1 {
		panic("match: less than 1")
	}

	// trim the brackets. Unfortunately "look behind" and "look ahead" are not supported in Go regexp
	name := match[1 : len(match) - 1]

	return name
}

// splitClippings
// uses a clippingSeparator pattern to break a larger string
// into smaller chunks
func splitClippings(input string) []string {
	clips := strings.Split(input, clippingSeparator)
	return clips
}

// isTypeBookmark
// checks if a string contains the Bookmark regular expression
// signifying it is a Bookmark type of clipping
func isTypeBookmark(input string) bool {
	r := regexp.MustCompile(bookmarkRegxp)
	if r.MatchString(input) {
		return true
	}
	return false
}

// parseClipping
// takes a single clipping string and populates a Clipping object
func parseClipping(input string, c *clipping.Clipping) {
	lineNumber := 1

	// create a scanner object to read the clipping
	scanner := bufio.NewScanner(strings.NewReader(input))
	// iterate over the lines "\n" is the delimiter
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		// from the line number, determine what parsing is required
		switch lineNumber {
		case 1:
			extractTitleAndAuthor(scanner.Text(), c)
		case 4:
			extractContent(scanner.Text(), c)
		default:
			break
		}
		lineNumber++
	}
}

// extractTitleAndAuthor
// the first line in a Kindle clipping chunk includes the title and author
// this function takes the
func extractTitleAndAuthor(input string, c *clipping.Clipping) {
	c.Book = extractTitle(input)
	c.Author = extractAuthor(input)
}

// extractContent
// the content is always the last line in a record
func extractContent(input string, c *clipping.Clipping) {
	c.Content = strings.TrimSpace(input)
}

// Parse takes a path to a Kindle My Clippings.txt file and returns an array of Clipping objects
func Parse(path string) []clipping.Clipping {

	// Convert the file to bytes
	fileStream, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte -> string
	s := string(fileStream)

	// get array of clippings
	clippingsTxt := splitClippings(s)

	clippings := []clipping.Clipping{}
	temp := clipping.Clipping{}

	for _, ct := range clippingsTxt {
		// Should not be empty or a bookmark type clipping
		if len(ct) != 0 && !isTypeBookmark(ct) {
			parseClipping(ct, &temp)
			temp.String()
			clippings = append(clippings, temp)
		}
	}

	return clippings
}
