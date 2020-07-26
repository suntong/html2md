////////////////////////////////////////////////////////////////////////////
// Program: html2md
// Purpose: HTML to Markdown
// Authors: Tong Sun (c) 2020, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	//  	"fmt"
	//  	"os"

	"github.com/mkideal/cli"
	//  	"github.com/mkideal/cli/clis"
	clix "github.com/mkideal/cli/ext"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

//==========================================================================
// html2md

type rootT struct {
	cli.Helper
	Filei                       *clix.Reader `cli:"*i,in" usage:"The html/xml file to read from (or stdin)"`
	Sel                         string       `cli:"s,sel" usage:"CSS/goquery selectors\n"`
	OptHeadingStyle             string       `cli:"opt-heading-style" usage:"Option HeadingStyle"`
	OptHorizontalRule           string       `cli:"opt-horizontal-rule" usage:"Option HorizontalRule"`
	OptBulletListMarker         string       `cli:"opt-bullet-list-marker" usage:"Option BulletListMarker"`
	OptCodeBlockStyle           string       `cli:"opt-code-block-style" usage:"Option CodeBlockStyle"`
	OptFence                    string       `cli:"opt-fence" usage:"Option Fence"`
	OptEmDelimiter              string       `cli:"opt-em-delimiter" usage:"Option EmDelimiter"`
	OptStrongDelimiter          string       `cli:"opt-strong-delimiter" usage:"Option StrongDelimiter"`
	OptLinkStyle                string       `cli:"opt-link-style" usage:"Option LinkStyle"`
	OptLinkReferenceStyle       string       `cli:"opt-link-reference-style" usage:"Option LinkReferenceStyle\n"`
	PluginConfluenceAttachments bool         `cli:"plugin-conf-attachment" usage:"Plugin ConfluenceAttachments"`
	PluginConfluenceCodeBlock   bool         `cli:"plugin-conf-code" usage:"Plugin ConfluenceCodeBlock"`
	PluginFrontMatter           bool         `cli:"plugin-frontmatter" usage:"Plugin FrontMatter"`
	PluginGitHubFlavored        bool         `cli:"plugin-gfm" usage:"Plugin GitHubFlavored"`
	PluginStrikethrough         bool         `cli:"plugin-strikethrough" usage:"Plugin Strikethrough"`
	PluginTable                 bool         `cli:"plugin-table" usage:"Plugin Table"`
	PluginTaskListItems         bool         `cli:"plugin-task-list" usage:"Plugin TaskListItems"`
	PluginVimeoEmbed            bool         `cli:"plugin-vimeo" usage:"Plugin VimeoEmbed"`
	PluginYoutubeEmbed          bool         `cli:"plugin-youtube" usage:"Plugin YoutubeEmbed"`
}

var root = &cli.Command{
	Name: "html2md",
	Desc: "HTML to Markdown\nVersion " + version + " built on " + date +
		"\nCopyright (C) 2020, Tong Sun",
	Text: "HTML to Markdown converter on command line" +
		"\n\nUsage:\n  html2md [Options...]",
	Argv: func() interface{} { return new(rootT) },
	Fn:   html2md,

	NumOption: cli.AtLeast(1),
}

// Template for main starts here
////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The OptsT type defines all the configurable options from cli.
//  type OptsT struct {
//  	Filei	*clix.Reader
//  	Sel	string
//  	OptHeadingStyle	string
//  	OptHorizontalRule	string
//  	OptBulletListMarker	string
//  	OptCodeBlockStyle	string
//  	OptFence	string
//  	OptEmDelimiter	string
//  	OptStrongDelimiter	string
//  	OptLinkStyle	string
//  	OptLinkReferenceStyle	string
//  	PluginConfluenceAttachments	bool
//  	PluginConfluenceCodeBlock	bool
//  	PluginFrontMatter	bool
//  	PluginGitHubFlavored	bool
//  	PluginStrikethrough	bool
//  	PluginTable	bool
//  	PluginTaskListItems	bool
//  	PluginVimeoEmbed	bool
//  	PluginYoutubeEmbed	bool
//  	Verbose int
//  }

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

//  var (
//          progname  = "html2md"
//          version   = "0.1.0"
//          date = "2020-07-25"

//  	rootArgv *rootT
//  	// Opts store all the configurable options
//  	Opts OptsT
//  )

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
//  func main() {
//  	cli.SetUsageStyle(cli.DenseNormalStyle) // left-right, for up-down, use ManualStyle
//  	//NOTE: You can set any writer implements io.Writer
//  	// default writer is os.Stdout
//  	if err := cli.Root(root,).Run(os.Args[1:]); err != nil {
//  		fmt.Fprintln(os.Stderr, err)
//  		os.Exit(1)
//  	}
//  	fmt.Println("")
//  }

// Template for main dispatcher starts here
//==========================================================================
// Dumb root handler

//  func html2md(ctx *cli.Context) error {
//  	ctx.JSON(ctx.RootArgv())
//  	ctx.JSON(ctx.Argv())
//  	fmt.Println()

//  	return nil
//  }

// Template for CLI handling starts here
