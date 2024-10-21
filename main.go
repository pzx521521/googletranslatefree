package googletranslatefree

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func Translate(text string, sourceLangue string, targetLangue string) (string, error) {
	ret := ""
	client := &http.Client{}
	if _, ok := CODES_TO_GOOGLE_LANGUAGES[sourceLangue]; !ok {
		sourceLangue = ""
	}
	if _, ok := CODES_TO_GOOGLE_LANGUAGES[targetLangue]; !ok {
		return ret, errors.New("invalid target language")
	}
	if text == "" {
		return ret, errors.New("empty text")
	}
	googleUrl := fmt.Sprintf("%s&sl=%s&tl=%s&q=%s", GOOGLE_BASEURL, sourceLangue, targetLangue, url.QueryEscape(text))
	req, err := http.NewRequest("GET", googleUrl, nil)
	if err != nil {
		return ret, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return ret, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusTooManyRequests {
		return ret, errors.New("too many requests")
	} else if resp.StatusCode != http.StatusOK {
		return ret, errors.New("RequestError")
	}

	bodyText, err := io.ReadAll(resp.Body)
	ret, err = getTextFromDiv(string(bodyText), "result-container")
	if err != nil {
		return ret, err
	}
	return ret, nil
}

// getTextFromDiv 根据类名提取 <div> 标签中的文本内容
func getTextFromDiv(htmlContent, className string) (string, error) {
	// 构造搜索字符串
	startTag := fmt.Sprintf(`<div class="%s">`, className)
	endTag := `</div>`

	// 查找开始和结束标签的位置
	startIndex := strings.Index(htmlContent, startTag)
	if startIndex == -1 {
		return "", fmt.Errorf("no content found for class: %s", className)
	}

	startIndex += len(startTag) // 移动到内容的开始位置
	endIndex := strings.Index(htmlContent[startIndex:], endTag)
	if endIndex == -1 {
		return "", fmt.Errorf("no closing tag found for class: %s", className)
	}

	// 提取内容
	content := htmlContent[startIndex : startIndex+endIndex]
	return content, nil
}
