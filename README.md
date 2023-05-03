# html2md
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-7-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/suntong/html2md?status.svg)](http://godoc.org/github.com/suntong/html2md)
[![Go Report Card](https://goreportcard.com/badge/github.com/suntong/html2md)](https://goreportcard.com/report/github.com/suntong/html2md)
[![Build Status](https://github.com/suntong/html2md/actions/workflows/go-release-build.yml/badge.svg?branch=master)](https://github.com/suntong/html2md/actions/workflows/go-release-build.yml)
[![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-B.svg)](http://godoc.org/github.com/go-easygen/wireframe)



## TOC
- [html2md - HTML to Markdown converter](#html2md---html-to-markdown-converter)
- [Usage](#usage)
  - [$ html2md](#-html2md)
  - [Examples](#examples)
    - [Simplest form](#simplest-form)
    - [Using goquery](#using-goquery)
  - [The options and plugins](#the-options-and-plugins)
    - [Testing the new table plugins](#testing-the-new-table-plugins)
  - [Credits](#credits)
    - [Credits](#credits-1)
    - [Similar Projects](#similar-projects)

- [Install Debian/Ubuntu package](#install-debianubuntu-package)
- [Download/install binaries](#downloadinstall-binaries)
  - [The binary executables](#the-binary-executables)
  - [Distro package](#distro-package)
  - [Debian package](#debian-package)
- [Install Source](#install-source)
- [Author](#author)
- [Contributors](#contributors-)

## html2md - HTML to Markdown converter

The `html2md` makes use of `github.com/JohannesKaufmann/html-to-markdown`
to convert HTML into Markdown, which is using an [HTML Parser](https://github.com/PuerkitoBio/goquery) to avoid the use of `regexp` as much as possible, which can prevent some [weird cases](https://stackoverflow.com/a/1732454) and allows it to be used for cases where the input is totally unknown.

![gopher stading on top of a machine that converts a box of html to blocks of markdown](https://github.com/JohannesKaufmann/html-to-markdown/raw/master/logo.png)


## Usage

### $ html2md
```sh
HTML to Markdown
Version 1.1.0 built on 2023-05-03
Copyright (C) 2020-2023, Tong Sun

HTML to Markdown converter on command line

Usage:
  html2md [Options...]

Options:

  -h, --help                       display help information 
  -i, --in                        *The html/xml file to read from (or stdin) 
  -d, --domain                     Domain of the web page, needed for links when --in is not url 
  -s, --sel                        CSS/goquery selectors [=body]
  -x, --excl                       Excluding CSS/goquery selectors 
      --xc                         Excluding all children nodes 
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
      --opt-escape-mode            Option EscapeMode 

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

### Examples

#### Simplest form

```md
$ html2md -i https://github.com/suntong/html2md | head -3
[Skip to content](#start-of-content)

[Homepage](https://github.com/)
```

#### Using goquery

The most useful feature is to use and pass a [goquery](https://github.com/PuerkitoBio/goquery) selection to filter for the content you want. 

```md
$ html2md -i https://github.com/JohannesKaufmann/html-to-markdown -s "div.my-3"
[go](http://github.com/topics/go "Topic: go") [html](http://github.com/topics/html "Topic: html") [markdown](http://github.com/topics/markdown "Topic: markdown") [golang](http://github.com/topics/golang "Topic: golang") [converter](http://github.com/topics/converter "Topic: converter") [html-to-markdown](http://github.com/topics/html-to-markdown "Topic: html-to-markdown") [goquery](http://github.com/topics/goquery "Topic: goquery")

```


### The options and plugins

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
Only ~~blue ones~~ ~~left~~
```

#### Testing the new table plugins

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

## Credits


### Credits

- [Johannes Kaufmann's html-to-markdown](https://github.com/JohannesKaufmann/html-to-markdown) that does the heavy lifting behind the scene.

### Similar Projects

- [turndown (js)](https://github.com/domchristie/turndown), a very good library written in javascript.
- [lunny/html2md](https://github.com/lunny/html2md), which is using [regex instead of goquery](https://stackoverflow.com/a/1732454), which exhibits a few edge cases which prompted `github.com/JohannesKaufmann/html-to-markdown`
- [jaytaylor/html2text](https://github.com/jaytaylor/html2text), which is not converting to markdown but plain text.

## Install Debian/Ubuntu package

    sudo apt install -y html2md

## Download/install binaries

- The latest binary executables are available 
as the result of the Continuous-Integration (CI) process.
- I.e., they are built automatically right from the source code at every git release by [GitHub Actions](https://docs.github.com/en/actions).
- There are two ways to get/install such binary executables
  * Using the **binary executables** directly, or
  * Using **packages** for your distro

### The binary executables

- The latest binary executables are directly available under  
https://github.com/suntong/html2md/releases/latest 
- Pick & choose the one that suits your OS and its architecture. E.g., for Linux, it would be the `html2md_verxx_linux_amd64.tar.gz` file. 
- Available OS for binary executables are
  * Linux
  * Mac OS (darwin)
  * Windows
- If your OS and its architecture is not available in the download list, please let me know and I'll add it.
- The manual installation is just to unpack it and move/copy the binary executable to somewhere in `PATH`. For example,

``` sh
tar -xvf html2md_*_linux_amd64.tar.gz
sudo mv -v html2md_*_linux_amd64/html2md /usr/local/bin/
rmdir -v html2md_*_linux_amd64
```


### Distro package

- [Packages available for Linux distros](https://cloudsmith.io/~suntong/repos/repo/packages/) are
  * [Alpine Linux](https://cloudsmith.io/~suntong/repos/repo/setup/#formats-alpine)
  * [Debian](https://cloudsmith.io/~suntong/repos/repo/setup/#formats-deb)
  * [RedHat](https://cloudsmith.io/~suntong/repos/repo/setup/#formats-rpm)

The repo setup instruction url has been given above.
For example, for [Debian](https://cloudsmith.io/~suntong/repos/repo/setup/#formats-deb) --

### Debian package


```sh
curl -1sLf \
  'https://dl.cloudsmith.io/public/suntong/repo/setup.deb.sh' \
  | sudo -E bash

# That's it. You then can do your normal operations, like

sudo apt update
apt-cache policy html2md

sudo apt install -y html2md
```

## Install Source

To install the source code instead:

```
go install github.com/suntong/html2md@latest
```

## Author

Tong SUN  
![suntong from cpan.org](https://img.shields.io/badge/suntong-%40cpan.org-lightgrey.svg "suntong from cpan.org")

_Powered by_ [**WireFrame**](https://github.com/go-easygen/wireframe)  
[![PoweredBy WireFrame](https://github.com/go-easygen/wireframe/blob/master/PoweredBy-WireFrame-Y.svg)](http://godoc.org/github.com/go-easygen/wireframe)  
the _one-stop wire-framing solution_ for Go cli based projects, from _init_ to _deploy_.

## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tbody>
    <tr>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/suntong"><img src="https://avatars.githubusercontent.com/u/422244?v=4?s=100" width="100px;" alt="suntong"/><br /><sub><b>suntong</b></sub></a><br /><a href="https://github.com/suntong/html2md/commits?author=suntong" title="Code">💻</a> <a href="#ideas-suntong" title="Ideas, Planning, & Feedback">🤔</a> <a href="#design-suntong" title="Design">🎨</a> <a href="#data-suntong" title="Data">🔣</a> <a href="https://github.com/suntong/html2md/commits?author=suntong" title="Tests">⚠️</a> <a href="https://github.com/suntong/html2md/issues?q=author%3Asuntong" title="Bug reports">🐛</a> <a href="https://github.com/suntong/html2md/commits?author=suntong" title="Documentation">📖</a> <a href="#blog-suntong" title="Blogposts">📝</a> <a href="#example-suntong" title="Examples">💡</a> <a href="#tutorial-suntong" title="Tutorials">✅</a> <a href="#tool-suntong" title="Tools">🔧</a> <a href="#platform-suntong" title="Packaging/porting to new platform">📦</a> <a href="https://github.com/suntong/html2md/pulls?q=is%3Apr+reviewed-by%3Asuntong" title="Reviewed Pull Requests">👀</a> <a href="#question-suntong" title="Answering Questions">💬</a> <a href="#maintenance-suntong" title="Maintenance">🚧</a> <a href="#infra-suntong" title="Infrastructure (Hosting, Build-Tools, etc)">🚇</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/VPanteleev-S7"><img src="https://avatars.githubusercontent.com/u/108007295?v=4?s=100" width="100px;" alt="VPanteleev-S7"/><br /><sub><b>VPanteleev-S7</b></sub></a><br /><a href="https://github.com/suntong/html2md/commits?author=VPanteleev-S7" title="Code">💻</a> <a href="https://github.com/suntong/html2md/issues?q=author%3AVPanteleev-S7" title="Bug reports">🐛</a> <a href="#userTesting-VPanteleev-S7" title="User Testing">📓</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/itdoginfo"><img src="https://avatars.githubusercontent.com/u/11143109?v=4?s=100" width="100px;" alt="itdoginfo"/><br /><sub><b>itdoginfo</b></sub></a><br /><a href="https://github.com/suntong/html2md/issues?q=author%3Aitdoginfo" title="Bug reports">🐛</a> <a href="#userTesting-itdoginfo" title="User Testing">📓</a></td>
      <td align="center" valign="top" width="14.28%"><a href="http://someurl1.comm"><img src="https://avatars.githubusercontent.com/u/8886?v=4?s=100" width="100px;" alt="somename123"/><br /><sub><b>somename123</b></sub></a><br /><a href="https://github.com/suntong/html2md/issues?q=author%3Am040601" title="Bug reports">🐛</a> <a href="#ideas-m040601" title="Ideas, Planning, & Feedback">🤔</a> <a href="#userTesting-m040601" title="User Testing">📓</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/vivook"><img src="https://avatars.githubusercontent.com/u/24224102?v=4?s=100" width="100px;" alt="vivook"/><br /><sub><b>vivook</b></sub></a><br /><a href="https://github.com/suntong/html2md/issues?q=author%3Avivook" title="Bug reports">🐛</a> <a href="#userTesting-vivook" title="User Testing">📓</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/097115"><img src="https://avatars.githubusercontent.com/u/1415155?v=4?s=100" width="100px;" alt="097115"/><br /><sub><b>097115</b></sub></a><br /><a href="https://github.com/suntong/html2md/issues?q=author%3A097115" title="Bug reports">🐛</a> <a href="#ideas-097115" title="Ideas, Planning, & Feedback">🤔</a> <a href="#userTesting-097115" title="User Testing">📓</a></td>
      <td align="center" valign="top" width="14.28%"><a href="http://magnusviri.com"><img src="https://avatars.githubusercontent.com/u/711269?v=4?s=100" width="100px;" alt="James Reynolds"/><br /><sub><b>James Reynolds</b></sub></a><br /><a href="https://github.com/suntong/html2md/pulls?q=is%3Apr+reviewed-by%3Amagnusviri" title="Reviewed Pull Requests">👀</a> <a href="#talk-magnusviri" title="Talks">📢</a> <a href="#userTesting-magnusviri" title="User Testing">📓</a></td>
    </tr>
  </tbody>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!
