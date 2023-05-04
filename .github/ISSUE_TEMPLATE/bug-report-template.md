---
name: Bug report template
about: Create a report to help us improve
title: ''
labels: ''
assignees: ''

---

**Describe the bug**
A clear and concise description of what the bug is.

**To Reproduce**

```sh
$ echo "foo<br>bar" | html2md -i
foo

bar

$ html2md | head -2
HTML to Markdown
Version x.y.0 built on 2023-...
```

**Expected output**

```md
foo
bar
```

**Additional context**
Add any other context about the problem here.
