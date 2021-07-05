
[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/suntong/html2md?status.svg)](http://godoc.org/github.com/suntong/html2md)
[![Go Report Card](https://goreportcard.com/badge/github.com/suntong/html2md)](https://goreportcard.com/report/github.com/suntong/html2md)
[![travis Status](https://travis-ci.org/suntong/html2md.svg?branch=master)](https://travis-ci.org/suntong/html2md)
[![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-R.svg)](http://godoc.org/github.com/go-easygen/wireframe)

# TOC
- [html2md - HTML to Markdown converter](#html2md---html-to-markdown-converter)
- [Usage](#usage)
  - [$ html2md](#-html2md)
- [Examples](#examples)
  - [Simplest form](#simplest-form)
  - [Using goquery](#using-goquery)
  - [The options and plugins](#the-options-and-plugins)
  - [Testing the new table plugins](#testing-the-new-table-plugins)
- [Download/Install](#downloadinstall)
  - [Download binaries](#download-binaries)
  - [Install Source](#install-source)
- [Credits & Authors](#credits-&-authors)
  - [Credits](#credits)
  - [Similar Projects](#similar-projects)
  - [Author(s) & Contributor(s)](#author(s)-&-contributor(s))

# html2md - HTML to Markdown converter

The `html2md` makes use of `github.com/JohannesKaufmann/html-to-markdown`
to convert HTML into Markdown, which is using an [HTML Parser](https://github.com/PuerkitoBio/goquery) to avoid the use of `regexp` as much as possible, which can prevent some [weird cases](https://stackoverflow.com/a/1732454) and allows it to be used for cases where the input is totally unknown.

![gopher stading on top of a machine that converts a box of html to blocks of markdown](https://github.com/JohannesKaufmann/html-to-markdown/raw/master/logo.png)


# Usage

### $ html2md
```sh
HTML to Markdown
Version 0.2.01 built on 2020-08-08
Copyright (C) 2020, Tong Sun

HTML to Markdown converter on command line

Usage:
  html2md [Options...]

Options:

  -h, --help                       display help information 
  -i, --in                        *The html/xml file to read from (or stdin) 
  -d, --domain                     Domain of the web page, needed for links when --in is not url 
  -s, --sel                        CSS/goquery selectors [=body]
  -v, --verbose                    Verbose mode (Multiple -v options increase the verbosity.) 

      --opt-heading-style          Option HeadingStyle 
      --opt-horizontal-rule        Option HorizontalRule 
      --opt-bullet-list-marker     Option BulletListMarker 
      --opt-code-block-style       Option CodeBlockStyle 
      --opt-fence                  Option Fence 
      --opt-em-delimiter           Option EmDelimiter 
      --opt-strong-delimiter       Option StrongDelimiter 
      --opt-link-style             Option LinkStyle 
      --opt-link-reference-style   Option LinkReferenceStyle 

  -A, --plugin-conf-attachment     Plugin ConfluenceAttachments 
  -C, --plugin-conf-code           Plugin ConfluenceCodeBlock 
  -F, --plugin-frontmatter         Plugin FrontMatter 
  -G, --plugin-gfm                 Plugin GitHubFlavored 
  -S, --plugin-strikethrough       Plugin Strikethrough 
  -T, --plugin-table               Plugin Table 
      --plugin-table-compat        Plugin TableCompat 
  -L, --plugin-task-list           Plugin TaskListItems 
  -V, --plugin-vimeo               Plugin VimeoEmbed 
  -Y, --plugin-youtube             Plugin YoutubeEmbed
```

# Examples

## Simplest form

```md
$ html2md -i https://github.com/suntong/html2md | head -3
[Skip to content](#start-of-content)

[Homepage](https://github.com/)
```

## Using goquery

The most useful feature is to use and pass a [goquery](https://github.com/PuerkitoBio/goquery) selection to filter for the content you want. 

```md
$ html2md -i https://github.com/JohannesKaufmann/html-to-markdown -s "div.BorderGrid-row.hide-sm.hide-md > div"
```


## The options and plugins

Works as expected:

```sh
$ echo '<strong>Bold Text</strong>' | html2md -i
**Bold Text**

$ echo '<strong>Bold Text</strong>' | html2md -i --opt-strong-delimiter="__"
__Bold Text__


$ echo '<ul><li><input type=checkbox checked>Checked!</li><li><input type=checkbox>Check Me!</li></ul>' | html2md -i -G
- [x] Checked!
- [ ] Check Me!

$ echo 'Only <del>blue ones</del> <s> left</s>' | html2md -i --plugin-strikethrough
Only ~blue ones~ ~left~
```

## Testing the new table plugins

```sh
$ cat $GOPATH/src/github.com/JohannesKaufmann/html-to-markdown/testdata/TestPlugins/table/input.html | html2md -i -T | head -6
| Firstname | Lastname | Age |
| --- | --- | --- |
| Jill | Smith | 50 |
| Eve | Jackson | 94 |
| Empty |  |  |
| End |

$ cat $GOPATH/src/github.com/JohannesKaufmann/html-to-markdown/testdata/TestPlugins/table/input.html | html2md -i -T --domain example.com | diff -wU 1 $GOPATH/src/github.com/JohannesKaufmann/html-to-markdown/testdata/TestPlugins/table/output.table.golden -
---
@@ -41 +41,2 @@
 | `var` | b | c |
\ No newline at end of file
+

$ cat $GOPATH/src/github.com/JohannesKaufmann/html-to-markdown/testdata/TestPlugins/table/input.html | html2md -i --plugin-table-compat | head -6
Firstname · Lastname · Age

Jill · Smith · 50

Eve · Jackson · 94

$ cat $GOPATH/src/github.com/JohannesKaufmann/html-to-markdown/testdata/TestPlugins/table/input.html | html2md -i --plugin-table-compat --domain example.com | diff -wU 1 $GOPATH/src/github.com/JohannesKaufmann/html-to-markdown/testdata/TestPlugins/table/output.tablecompat.golden -
---
@@ -41 +41,2 @@
 `var` · b · c
\ No newline at end of file
+
```

# Download/Install


## Download binaries

- The latest binary executables are available right under the github release page  
https://github.com/suntong/html2md/releases  
as the result of the Continuous-Integration process.
- I.e., they are built during every git tagged push, automatically by [GitHub Actions](https://github.com/features/actions), right from the source code, truely WYSIWYG.
- The `.deb`, `.rpm` and `.apk` packages are readily available, as well as the executables for other Linux and Windows as well.
- Pick & choose the binary executable that suits your OS and its architecture. E.g., for Linux, it would most probably be the `html2md_ver_linux_amd64.tar.gz` file.
- Unzip it and put the executable somewhere in the PATH, after downloading it. 


## Install Source

To install the source code instead:

```
go get github.com/suntong/html2md
```

# Credits & Authors


## Credits

- [Johannes Kaufmann's html-to-markdown](github.com/JohannesKaufmann/html-to-markdown) that does the heavy lifting behind the scene.

## Similar Projects

- [turndown (js)](https://github.com/domchristie/turndown), a very good library written in javascript.
- [lunny/html2md](https://github.com/lunny/html2md), which is using [regex instead of goquery](https://stackoverflow.com/a/1732454), which exhibits a few edge cases which prompted `github.com/JohannesKaufmann/html-to-markdown`
- [jaytaylor/html2text](https://github.com/jaytaylor/html2text), which is not converting to markdown but plain text.

## Author(s) & Contributor(s)

Tong SUN  
![suntong from cpan.org](https://img.shields.io/badge/suntong-%40cpan.org-lightgrey.svg "suntong from cpan.org")

[![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-Y.svg)](http://godoc.org/github.com/go-easygen/wireframe)  
_Powered by_ [**WireFrame**](https://github.com/go-easygen/wireframe), the _one-stop wire-framing solution_ for Go cli based projects, from start to deploy.

All patches welcome. 
