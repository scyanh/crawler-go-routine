package entities

import (
	"fmt"
	"strings"
)

type Link struct {
	URL   string
	Links []string
}

func (l Link) String() string {
	return fmt.Sprintf("URL: %s, Links: [%s] \n", l.URL, strings.Join(l.Links, ", "))
}
