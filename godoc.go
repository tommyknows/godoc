// Package godoc is an example
// on how you can document your
// go code
package godoc

import "time"

// Godoc contains necessary information about
// the documentation.
//
// It is possible to group that into multiple
// paragraphs.
// To generate a new GoDoc, call
//   New("my documentation", "author")
// You can also add examples, though that needs to be
// done in _test.go files.
// Just prefix the example function with "Example[type]",
// meaning this example will be named "ExampleGoDoc",
// and showing up as "Example" at the documentation
// for the GoDoc type. You can add a name to the example by
// adding "_[name]" to your function name, e.g.
// "ExampleGoDoc_cool"
type GoDoc struct {
	// Documentation is the actual content
	Documentation string
	// CreatedAt is the date that the Godoc has been generated
	CreatedAt time.Time
	// Author of the documentation
	Author string
}

// New Godoc from documentation and author.
func New(doc, author string) *GoDoc {
	return &GoDoc{
		Documentation: doc,
		CreatedAt:     time.Now(),
		Author:        author,
	}
}

// String converts GoDoc to a readable string, containing all
// the information you will ever need
func (g *GoDoc) String() string {
	return g.Documentation + ", written at " + g.CreatedAt.String() + " and authored by " + g.Author
}
