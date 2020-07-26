
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
- [Debian package](#debian-package)
- [Install Source](#install-source)
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
Version 0.1.01 built on 2020-07-26
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

# Debian package

Will be available once `github.com/JohannesKaufmann/html-to-markdown` has a release version.

# Install Source

To install the source code instead:

```
go get github.com/suntong/html2md
```


## Credits

- [Johannes Kaufmann's html-to-markdown](github.com/JohannesKaufmann/html-to-markdown) that does the heavy lifting behind the scene.

## Similar Projects

- [turndown (js)](https://github.com/domchristie/turndown), a very good library written in javascript.
- [lunny/html2md](https://github.com/lunny/html2md), which is using [regex instead of goquery](https://stackoverflow.com/a/1732454), which exhibits a few edge cases which prompted `github.com/JohannesKaufmann/html-to-markdown`
- [jaytaylor/html2text](https://github.com/jaytaylor/html2text), which is not converting to markdown but plain text.

## Author(s) & Contributor(s)

Tong SUN  
![suntong from cpan.org](https://img.shields.io/badge/suntong-%40cpan.org-lightgrey.svg "suntong from cpan.org")

_Powered by_ [**WireFrame**](https://github.com/go-easygen/wireframe),  [![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-Y.svg)](http://godoc.org/github.com/go-easygen/wireframe), the _one-stop wire-framing solution_ for Go cli based projects, from start to deploy.

All patches welcome. 
