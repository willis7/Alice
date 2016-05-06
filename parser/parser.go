package parser

import (
	"bufio"
	"regexp"
	"strings"

	"github.com/willis7/Alice/clipping"
)

const clippingSeparator string = "=========="
const authorRegxp string = `\((.*?)\)` // this regexp will find a bracket pair and all characters inbetween
const titleRegxp string = `^(.*?)\(`   // this regexp will find everything up to, and including a "("

// parseTitle
// uses a regular expression to find the title in a Kindle formatted string.
// The Kindle always puts the title before the author.
func parseTitle(input string) string {
	r, _ := regexp.Compile(titleRegxp)
	match := r.FindString(input)

	title := match[0 : len(match)-2]

	return title
}

// parseAuthor
// uses a regular expression to find the author name in a Kindle formatted string.
// The Kindle always puts the author name between braces (<<author>>).
func parseAuthor(input string) string {
	r, _ := regexp.Compile(authorRegxp)
	match := r.FindString(input)

	// trim the brackets. Unfortunately "look behind" and "look ahead" are not supported in Go regexp
	name := match[1 : len(match)-1]

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
func parseTitleAndAuthor(s string, c *clipping.Clipping) {
	c.Book = parseTitle(s)
	c.Author = parseAuthor(s)
}

// parseContent
// the content is always the
func parseContent(input string, c *clipping.Clipping) {
	c.Content = strings.TrimSpace(input)
}

// Parse takes a path to a Kindle My Clippings.txt file and returns an array of Clipping objects
// func Parse(path string) []clipping.Clipping {
//
// 	log.Printf("File path: %s", path)
//
// 	// Convert the file to bytes
// 	fileStream, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	// Convert []byte -> string
// 	s := string(fileStream)
//
// 	// get array of clippings
// 	clippingsTxt := splitClippings(s)
//
// 	clips := []clipping.Clipping{}
// 	temp := clipping.Clipping{}
//
// 	for _, clippingText := range clippingsTxt {
// 		parseClipping(clippingText, &temp)
// 		clips = append(clips, temp)
// 	}
//
// 	return clips
// }

// // FileToStructs takes a clippings.txt file and returns an array of Clipping objects
// func FileToStructs(filename string) ([]*clipping.Clipping, error) {
//
// 	// Convert the file to bytes
// 	fileStream, err := ioutil.ReadFile(filename); if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	// Convert []byte -> string
// 	s := string(fileStream)
//
// 	// Split the string into individual clippings
// 	[]clippings := String.Split(s, "==========")
//
// 	// Add to array
//
// 	// Return array
// }
//
// // processClipping takes a buffer representing a single clipping
// func processClipping(*buf Buffer)
