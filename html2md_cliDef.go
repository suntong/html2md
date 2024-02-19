////////////////////////////////////////////////////////////////////////////
// Program: html2md
// Purpose: HTML to Markdown
// Authors: Tong Sun (c) 2020-2023, All rights reserved
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
	Domain                      string       `cli:"d,domain" usage:"Domain of the web page, needed for links when --in is not url"`
	Sel                         string       `cli:"s,sel" usage:"CSS/goquery selectors" dft:"body"`
	Excl                        []string     `cli:"x,excl" usage:"Excluding CSS/goquery selectors"`
	ExclChildren                bool         `cli:"xc" usage:"Excluding all children nodes"`
	Verbose                     cli.Counter  `cli:"v,verbose" usage:"Verbose mode (Multiple -v options increase the verbosity.)\n"`
	OptHeadingStyle             string       `cli:"opt-heading-style" usage:"Option HeadingStyle"`
	OptHorizontalRule           string       `cli:"opt-horizontal-rule" usage:"Option HorizontalRule"`
	OptBulletListMarker         string       `cli:"opt-bullet-list-marker" usage:"Option BulletListMarker"`
	OptCodeBlockStyle           string       `cli:"opt-code-block-style" usage:"Option CodeBlockStyle"`
	OptFence                    string       `cli:"opt-fence" usage:"Option Fence"`
	OptEmDelimiter              string       `cli:"opt-em-delimiter" usage:"Option EmDelimiter"`
	OptStrongDelimiter          string       `cli:"opt-strong-delimiter" usage:"Option StrongDelimiter"`
	OptLinkStyle                string       `cli:"opt-link-style" usage:"Option LinkStyle"`
	OptLinkReferenceStyle       string       `cli:"opt-link-reference-style" usage:"Option LinkReferenceStyle"`
	OptEscapeMode               string       `cli:"opt-escape-mode" usage:"Option EscapeMode\n"`
	PluginBrToNewline           bool         `cli:"plugin-br-to-newline" usage:"Plugin BrToNewline"`
	PluginConfluenceAttachments bool         `cli:"A,plugin-conf-attachment" usage:"Plugin ConfluenceAttachments"`
	PluginConfluenceCodeBlock   bool         `cli:"C,plugin-conf-code" usage:"Plugin ConfluenceCodeBlock"`
	PluginFrontMatter           bool         `cli:"F,plugin-frontmatter" usage:"Plugin FrontMatter"`
	PluginGitHubFlavored        bool         `cli:"G,plugin-gfm" usage:"Plugin GitHubFlavored"`
	PluginStrikethrough         bool         `cli:"S,plugin-strikethrough" usage:"Plugin Strikethrough"`
	PluginTable                 bool         `cli:"T,plugin-table" usage:"Plugin Table"`
	PluginTableCompat           bool         `cli:"plugin-table-compat" usage:"Plugin TableCompat"`
	PluginTaskListItems         bool         `cli:"L,plugin-task-list" usage:"Plugin TaskListItems"`
	PluginVimeoEmbed            bool         `cli:"V,plugin-vimeo" usage:"Plugin VimeoEmbed"`
	PluginYoutubeEmbed          bool         `cli:"Y,plugin-youtube" usage:"Plugin YoutubeEmbed"`
}

var root = &cli.Command{
	Name: "html2md",
	Desc: "HTML to Markdown\nVersion " + version + " built on " + date +
		"\nCopyright (C) 2020-2024, Tong Sun",
	Text: "HTML to Markdown converter on command line" +
		"\n\nUsage:\n  html2md [Options...]",
	Argv: func() interface{} { return new(rootT) },
	Fn:   Html2md,

	NumOption: cli.AtLeast(1),
}

// Template for main starts here
////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The OptsT type defines all the configurable options from cli.
//  type OptsT struct {
//  	Filei	*clix.Reader
//  	Domain	string
//  	Sel	string
//  	Excl	[]string
//  	ExclChildren	bool
//  	Verbose	cli.Counter
//  	OptHeadingStyle	string
//  	OptHorizontalRule	string
//  	OptBulletListMarker	string
//  	OptCodeBlockStyle	string
//  	OptFence	string
//  	OptEmDelimiter	string
//  	OptStrongDelimiter	string
//  	OptLinkStyle	string
//  	OptLinkReferenceStyle	string
//  	OptEscapeMode	string
//  	PluginBrToNewline	bool
//  	PluginConfluenceAttachments	bool
//  	PluginConfluenceCodeBlock	bool
//  	PluginFrontMatter	bool
//  	PluginGitHubFlavored	bool
//  	PluginStrikethrough	bool
//  	PluginTable	bool
//  	PluginTableCompat	bool
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
//          date = "2023-05-03"

//  	rootArgv *rootT
//  	// Opts store all the configurable options
//  	Opts OptsT
//  )

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
//  func main() {
//  	cli.SetUsageStyle(cli.DenseNormalStyle)
//  	if err := cli.Root(root,).Run(os.Args[1:]); err != nil {
//  		fmt.Fprintln(os.Stderr, err)
//  		os.Exit(1)
//  	}
//  	fmt.Println("")
//  }

// Template for main dispatcher starts here
//==========================================================================
// Dumb root handler

// Html2md - main dispatcher dumb handler
//  func Html2md(ctx *cli.Context) error {
//  	ctx.JSON(ctx.RootArgv())
//  	ctx.JSON(ctx.Argv())
//  	fmt.Println()

//  	return nil
//  }

// Template for CLI handling starts here
