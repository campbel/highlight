# Highlight

Highlight code for pleasent viewing.

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

### Rich text

Useful if the intention is to copy output into document software like Google docs or word.

```sh
# Output rich text to copy into a doc
highlight -f rtf ruby.rb | pbcopy
```
