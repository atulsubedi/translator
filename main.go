package main

import(
	"fmt"
	"flag"
	"strings"
	"sync"
	"os"
)

var wg sync.WaitGroup

var sourceLang string
var targetLang string
var sourceText string

func init(){
	flag.StringVar(&sourceLang, "s", "en", "Source languauge[en]")

	flag.StringVar(&targetLang, "t", "fr", "target languauge[fr]")

	flag.StringVar(&sourceText, "st", "", " text To Translate")
}

func main(){
	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	reqBody := &cli.RequestBody{
		SourceLang: sourceLang,
		TargetLang: targetLang,
		SourceText: sourceText,
	}

	cli.RequestTranslate(reqBody)
}