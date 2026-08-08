package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"lesiw.io/errfmt"
)

func main() { singlechecker.Main(errfmt.Analyzer) }
