package cli

import "net/http"
 
type reqBody struct{
	SourceLang  string
	TargetLang  string
	SourceText  string
}

const translateUrl = "https://translate.googleapis.com/translate_a/single"

func RequestTranslate(body *reqBody){

	client := &http.Client
	req := 

	client.Do(req)
}