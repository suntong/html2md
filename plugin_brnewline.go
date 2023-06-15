package main

import (
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

// TaskListItems converts checkboxes into task list items.
func BrToNewline() md.Plugin {
	return func(c *md.Converter) []md.Rule {
		return []md.Rule{
			{
				Filter: []string{"br"},
				Replacement: func(content string, selec *goquery.Selection, opt *md.Options) *string {
					return md.String("\n")
				},
			},
		}
	}
}