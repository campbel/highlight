# Highlight

Highlight code for easy viewing.

![Demo of Highlight](tape/demo.gif)

## Install

```sh
go install github.com/campbel/highlight@main
```

## Examples

### Other languages

```sh
highlight -l ruby my_file.rb
```

### From stdin

```sh
cat main.go | highlight
```

### To Rich Text and Copy (Google Docs / Word)

```sh
# Output HTML and convert to rich text with textutil
cat my_file.go | highlight -f html | textutil -stdin -format html -convert rtf -stdout | pbcopy
```

This will copy the html output and convert to rich text. You can then paste this into a document program like Word or Google Docs.
