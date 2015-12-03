package translate

import (
  "net/http"
  "fmt"
  "encoding/json"
)


const GOOGLE_TRANSLATE = "https://www.googleapis.com/language/translate/v2"
const API_KEY = "AIzaSyClT0vRhV12mGtjFuWFqePPqgO7FSaUil4"

type (

  TReply struct {

    Error struct {
       Code int `json:"code"`
       Message string `json:"message"`
    } `json:"error"`

    Data struct {
      Translations []struct {
        TranslatedText string `json:"translatedText"`
        SourceLang string `json:detectedSourceLanguage`
      } `json:"translations"`
    } `json:"data"`

  }
)

// TODO: google орчуулга ашиглах

func Translate(w string, target string) (string, error) {
  uri:=fmt.Sprintf("%s?key=%s&target=%s&q=%s", GOOGLE_TRANSLATE, API_KEY, target, w)
  resp, err := http.Get(uri)
  if err != nil {
    return "", err
  }
  defer resp.Body.Close()

  dec := json.NewDecoder(resp.Body)

  var tran TReply
	if err := dec.Decode(&tran); err != nil {
		return "", err
	}

  if tran.Error.Code != 0 {
    fmt.Println(tran.Error.Message)
  }

  return "", nil
  //return tran.Data.Translations[0].TranslatedText, nil;

}
