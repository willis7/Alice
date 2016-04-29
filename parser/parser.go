package parser

import (
	"regexp"
	"strings"
)

const clippingSeparator string = "=========="
const authorRegxp string = `\((.*?)\)` // this regexp will find a bracket pair and all characters inbetween
const titleRegxp string = `^(.*?)\(`   // this regexp will find a bracket pair and all characters inbetween

// parseTitle
// uses a regular expression to find the title in a random string.
// The Kindle always puts the title before the author.
func parseTitle(s string) string {
	r, _ := regexp.Compile(titleRegxp)
	match := r.FindString(s)

	title := match[0 : len(match)-2]

	return title
}

// parseAuthor
// uses a regular expression to find the author name in a random string.
// The Kindle always puts the author name between braces (<<author>>).
func parseAuthor(s string) string {
	r, _ := regexp.Compile(authorRegxp)
	match := r.FindString(s)

	// trim the brackets. Unfortunately "look behind" and "look ahead" are not supported in Go regexp
	name := match[1 : len(match)-1]

	return name
}

// splitClippings uses a clippingSeparator pattern to break a larger string
// into smaller chunks
func splitClippings(s string) []string {
	clips := strings.Split(s, clippingSeparator)

	return clips
}

// // FileToStructs takes a clippings.txt file and returns an array of Clipping objects
// func FileToStructs(filename string) ([]*clipping.Clipping, error) {
//
// 	// Convert the file to bytes
// 	fileStream, err := ioutil.ReadFile(filename)
// 	if err != nil {
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
