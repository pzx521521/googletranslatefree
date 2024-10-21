# Google Translate FREE in golang
Note: This is not normal commercial Translator API provided by Google.


## CURL
```curl
curl "https://translate.google.com/m?tl=es&q=Hello%2C+World%21"
```
the reseponse is a html, result in 
```html
<div class="result-container">¡Hola Mundo!</div>
```
## Install:
```
go get github.com/pzx521521/googletranslatefree
```


Example usage:
```
package main

import (
    "fmt"
    gt "github.com/pzx521521/googletranslatefree"
)

func main(){
    const text string = `Hello, World!`
    // you can use "auto" for source language
    // so, translator will detect language
    result, _ := gt.Translate(text, "en", "es")
    fmt.Println(result)
    // Output: "Hola, Mundo!"
}
```
All language codes can be found here: 
[const.go](https://github.com/pzx521521/googletranslatefree/blob/main/const.go)