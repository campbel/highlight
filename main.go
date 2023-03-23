package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/quick"
	"github.com/campbel/yoshi"
)

type Options struct {
	File     string `yoshi:"FILE;Input file, alternatively pass from stdin;"`
	Language string `yoshi:"-l,--language;Language of the content;"`
	Theme    string `yoshi:"-t,--theme;Theme for the output;dracula"`
	Format   string `yoshi:"-f,--format;Format of the output;terminal256"`
}

func main() {
	yoshi.New("highlight").Run(func(options Options) {
		content := getContent(options.File)
		if content == "" {
			fmt.Fprintln(os.Stderr, "no content to highlight")
			os.Exit(0)
		}
		if options.Language == "" && options.File != "" {
			options.Language = lexers.Match(options.File).Config().Name
		}
		fmt.Println(highlight(string(content), options.Language, options.Format, options.Theme))
	})
}

func getContent(file string) string {
	if file == "" {
		content, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error reading from stdin: ", err)
			os.Exit(1)
		}
		return string(content)
	}
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading file: ", err)
		os.Exit(1)
	}
	return string(content)
}

func highlight(content, language, format, theme string) string {
	content = strings.Replace(content, "\t", "    ", -1)
	var b bytes.Buffer
	err := quick.Highlight(&b, content, language, format, theme)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error highlighting: ", err)
	}
	return b.String()
}
