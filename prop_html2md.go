////////////////////////////////////////////////////////////////////////////
// Program: html2md
// Purpose: HTML to Markdown
// Authors: Tong Sun (c) 2020, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"log"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/mkideal/cli"
)

//==========================================================================
// root handler

func html2md(ctx *cli.Context) error {
	ctx.JSON(ctx.RootArgv())
	ctx.JSON(ctx.Argv())
	fmt.Println()

	url := "https://blog.golang.org/godoc-documenting-go-code"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	content := doc.Find("#content")

	conv := md.NewConverter(md.DomainFromURL(url), true, nil)
	markdown := conv.Convert(content)

	fmt.Println(markdown)

	return nil
}
