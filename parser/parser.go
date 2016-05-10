package parser

import (
	"bufio"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/willis7/Alice/clipping"
	"errors"
)

const clippingSeparator string = "==========\n"
const authorRegxp string = `\((.*?)\)` // find a bracket pair and all characters inbetween
const titleRegxp string = `^(.*?)\(`   // find everything up to, and including a "("

// parseTitle
// uses a regular expression to find the title in a Kindle formatted string.
// The Kindle always puts the title before the author.
// TODO: rename to extractTitle
func parseTitle(input string) string {
	r, _ := regexp.Compile(titleRegxp)
	match := r.FindString(input)

	if len(match) < 2 {
		panic(errors.New("match: less than 2"))
	}
	title := match[: len(match) - 2]
	return title
}

// parseAuthor
// uses a regular expression to find the author name in a Kindle formatted string.
// The Kindle always puts the author name between braces (<<author>>).
// TODO: rename to extractAuthor
func parseAuthor(input string) string {
	r, _ := regexp.Compile(authorRegxp)
	match := r.FindString(input)

	if len(match) < 1 {
		panic(errors.New("match: less than 1"))
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
			parseTitleAndAuthor(scanner.Text(), c)
		case 4:
			parseContent(scanner.Text(), c)
		default:
			break
		}
		lineNumber++
	}
}

// parseTitleAndAuthor
// the first line in a Kindle clipping chunk includes the title and author
// this function takes the
// TODO: rename to extractTitleAndAuthor
func parseTitleAndAuthor(input string, c *clipping.Clipping) {
	c.Book = parseTitle(input)
	c.Author = parseAuthor(input)
}

// parseContent
// the content is always the last line in a record
// TODO: rename to extractContent
func parseContent(input string, c *clipping.Clipping) {
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
		if len(ct) != 0 {
			parseClipping(ct, &temp)
			temp.ToString()
			clippings = append(clippings, temp)
		}
	}

	return clippings
}
