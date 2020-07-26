////////////////////////////////////////////////////////////////////////////
// Program: html2md
// Purpose: HTML to Markdown
// Authors: Tong Sun (c) 2020, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/mkideal/cli"
	"github.com/mkideal/cli/clis"
)

//==========================================================================
// root handler

func html2md(ctx *cli.Context) error {
	ctx.JSON(ctx.RootArgv())
	rootArgv = ctx.RootArgv().(*rootT)
	// https://pkg.go.dev/github.com/mkideal/cli@v0.2.2/clis?tab=doc
	clis.Setup(progname, rootArgv.Verbose.Value())
	clis.Verbose(2, "%s\n", rootArgv.Filei.Name())

	doc, err := goquery.NewDocumentFromReader(rootArgv.Filei)
	clis.AbortOn("Reading file with goquery", err)
	content := doc.Find(rootArgv.Sel)

	conv := md.NewConverter(md.DomainFromURL(rootArgv.Filei.Name()), true, nil)
	markdown := conv.Convert(content)

	fmt.Println(markdown)

	return nil
}
