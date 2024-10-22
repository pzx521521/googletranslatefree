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


### Previous article
+ Currently, most Google interfaces use `https://translate.google.com/translate_a/single`

But it always pops up verification (Google's reCAPTCHA),
Because of domestic reasons, the IP often changes, so verification pops up frequently. As a result, it cannot be used

+ Found that `https://translate.google.com/m` can be used
Python version [deep-translator](https://github.com/nidhaloff/deep-translator)

+ I looked and found that there is no golang version (because I want to run the binary on arm), so I wrote one

[googletranslatefree](https://github.com/pzx521521/googletranslatefree)
