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
	flag.StringVar(&sourceLang, "s", "en", "Source languauge")
}