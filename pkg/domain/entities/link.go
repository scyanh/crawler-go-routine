package entities

import (
	"fmt"
	"strings"
)

type Link struct {
	URL   string
	Links []string
	Error bool
}

// String returns a string representation of the Link.
func (l Link) String() string {
	return fmt.Sprintf("URL: %s, Success: %t Links: [%s] \n", l.URL, !l.Error, strings.Join(l.Links, ", "))
}
