package googletranslatefree

import (
	"fmt"
	"testing"
)

func TestTranslate(t *testing.T) {
	translate, err := Translate("Hello,World!", "en", "zh-CN")
	if err != nil {
		return
	}
	fmt.Printf("%v\n", translate)
}
