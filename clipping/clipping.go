package clipping

// Clipping is a struct which will be used to represent a clipping from Kindle.
type Clipping struct {
	Book    string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

// String returns the string representation of the object
func (c *Clipping) String() string {
	return "{ Title: " + c.Book + ", Author: " + c.Author + ", Content: " + c.Content + " }"
}
