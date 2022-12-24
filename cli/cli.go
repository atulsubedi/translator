package cli

import (

)
 
type reqBody struct{
	SourceLang  string
	TargetLang  string
	SourceText  string
}

const translateUrl = "https://translate.googleapis.com/translate_a/single"

func RequestTranslate(body *reqBody){

}