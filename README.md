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
highlight my_file.rb
```

### From stdin

```sh
cat main.go | highlight
```

### To Rich Text and Copy (Google Docs / Word)

```sh
# Output rich text to copy into a doc
highlight -f rtf ruby.rb | pbcopy
```
