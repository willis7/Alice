package clipping

// Clipping is a struct which will be used to represent a clipping from Kindle.
type Clipping struct {
	Book, Author, Content string
}

// ToString gives the to string output of the struct
func (c *Clipping) String() string {
	return "{ Title: " + c.Book + ", Author: " + c.Author + ", Content: " + c.Content + " }"
}
