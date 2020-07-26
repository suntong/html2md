////////////////////////////////////////////////////////////////////////////
// Program: html2md
// Purpose: HTML to Markdown
// Authors: Tong Sun (c) 2020, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"regexp"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/JohannesKaufmann/html-to-markdown/plugin"
	"github.com/PuerkitoBio/goquery"
	"github.com/mkideal/cli"
	"github.com/mkideal/cli/clis"
)

//==========================================================================
// root handler

func html2md(ctx *cli.Context) error {
	rootArgv = ctx.RootArgv().(*rootT)
	// https://pkg.go.dev/github.com/mkideal/cli@v0.2.2/clis?tab=doc
	clis.Setup(progname, rootArgv.Verbose.Value())
	clis.Verbose(1, "%#v\n", rootArgv)
	clis.Verbose(2, "%s\n", rootArgv.Filei.Name())

	// Options handling
	opt := &md.Options{}
	if rootArgv.OptHeadingStyle != "" {
		opt.HeadingStyle = rootArgv.OptHeadingStyle
	}
	if rootArgv.OptHorizontalRule != "" {
		opt.HorizontalRule = rootArgv.OptHorizontalRule
	}
	if rootArgv.OptBulletListMarker != "" {
		opt.BulletListMarker = rootArgv.OptBulletListMarker
	}
	if rootArgv.OptCodeBlockStyle != "" {
		opt.CodeBlockStyle = rootArgv.OptCodeBlockStyle
	}
	if rootArgv.OptFence != "" {
		opt.Fence = rootArgv.OptFence
	}
	if rootArgv.OptEmDelimiter != "" {
		opt.EmDelimiter = rootArgv.OptEmDelimiter
	}
	if rootArgv.OptStrongDelimiter != "" {
		opt.StrongDelimiter = rootArgv.OptStrongDelimiter
	}
	if rootArgv.OptLinkStyle != "" {
		opt.LinkStyle = rootArgv.OptLinkStyle
	}
	if rootArgv.OptLinkReferenceStyle != "" {
		opt.LinkReferenceStyle = rootArgv.OptLinkReferenceStyle
	}
	clis.Verbose(1, "%#v\n", opt)

	doc, err := goquery.NewDocumentFromReader(rootArgv.Filei)
	clis.AbortOn("Reading file with goquery", err)
	content := doc.Find(rootArgv.Sel)

	domain, url := rootArgv.Domain, rootArgv.Filei.Name()
	if domain == "" && regexp.MustCompile(`(?i)^http`).MatchString(url) {
		domain = md.DomainFromURL(url)
	}
	clis.Verbose(2, "domain='%s'\n", domain)
	conv := md.NewConverter(domain, true, opt)

	// Plugin handling
	if rootArgv.PluginConfluenceAttachments {
		conv.Use(plugin.ConfluenceAttachments())
	}
	if rootArgv.PluginConfluenceCodeBlock {
		conv.Use(plugin.ConfluenceCodeBlock())
	}
	// if rootArgv.PluginFrontMatter {
	// 	conv.Use(plugin.FrontMatter())
	// }
	if rootArgv.PluginGitHubFlavored {
		conv.Use(plugin.GitHubFlavored())
	}
	if rootArgv.PluginStrikethrough {
		conv.Use(plugin.Strikethrough(""))
	}
	// if rootArgv.PluginTable {
	// 	conv.Use(plugin.Table())
	// }
	if rootArgv.PluginTaskListItems {
		conv.Use(plugin.TaskListItems())
	}
	// if rootArgv.PluginVimeoEmbed {
	// 	conv.Use(plugin.VimeoEmbed())
	// }
	// if rootArgv.PluginYoutubeEmbed {
	// 	conv.Use(plugin.YoutubeEmbed())
	// }

	markdown := conv.Convert(content)

	fmt.Println(markdown)

	return nil
}
