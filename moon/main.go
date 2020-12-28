package main

import (
	"flag"
	"io"
	"os"

	"github.com/hajimehoshi/file2byteslice"
)

var (
	inputFilename  = flag.String("input", "", "input filename")
	outputFilename = flag.String("output", "", "output filename")
	packageName    = flag.String("package", "main", "package name")
	varName        = flag.String("var", "_", "variable name")
	compress       = flag.Bool("compress", false, "use gzip compression")
	buildTags      = flag.String("buildtags", "", "build tags")
)

func run() error {
	var out io.Writer
	if *outputFilename != "" {
		f, err := os.Create(*outputFilename)
		if err != nil {
			return err
		}
		defer f.Close()
		out = f
	} else {
		out = os.Stdout
	}

	var in io.Reader
	if *inputFilename != "" {
		f, err := os.Open(*inputFilename)
		if err != nil {
			return err
		}
		defer f.Close()
		in = f
	} else {
		in = os.Stdin
	}

	if err := file2byteslice.Write(out, in, *compress, *buildTags, *packageName, *varName); err != nil {
		return err
	}

	return nil
}

// 启动参数program  添加如下  路径暂时为全路径
//-package=images -input=./resource/01/base.png -output="C:\Users\39495\Desktop\work\go\idea_go_test_git\moon\images\base.go" -var=Base_png
func main() {
	flag.Parse()
	if err := run(); err != nil {
		panic(err)
	}
}
