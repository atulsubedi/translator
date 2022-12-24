package main

import(
	"fmt"
	"flag"
	"strings"
	"sync"
	"os"
)
var sourceLang string
var targetLang string
var sourceText string

func init(){
	flag.StringVar(&sourceLang, "s", "en", "Source languauge[en]")

	flag.StringVar(&targetLang, "t", "fr", "target languauge[fr]")

	flag.StringVar(&sourceText, "st", "", " text To Translate")
}