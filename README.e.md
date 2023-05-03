
## {{toc 5}}

## {{.Name}} - HTML to Markdown converter

The `{{.Name}}` makes use of `github.com/JohannesKaufmann/html-to-markdown`
to convert HTML into Markdown, which is using an [HTML Parser](https://github.com/PuerkitoBio/goquery) to avoid the use of `regexp` as much as possible, which can prevent some [weird cases](https://stackoverflow.com/a/1732454) and allows it to be used for cases where the input is totally unknown.

![gopher stading on top of a machine that converts a box of html to blocks of markdown](https://github.com/JohannesKaufmann/html-to-markdown/raw/master/logo.png)


## Usage

### $ {{exec "html2md" | color "sh"}}

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
$ {{.Name}} -i https://github.com/JohannesKaufmann/html-to-markdown -s "div.BorderGrid-row.hide-sm.hide-md > div"
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
Only ~blue ones~ ~left~
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
