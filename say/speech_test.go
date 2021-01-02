package say

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKorean(t *testing.T) {
	expected := `{
	  "type":  "PlainText",
	  "lang":  "ko",
	  "value": "안녕하세요"
  	}`
	b, _ := json.Marshal(Korean("안녕하세요"))
	assert.JSONEq(t, expected, string(b), "should be equal")
}

func TestEnglish(t *testing.T) {
	expected := `{
	  "type":  "PlainText",
	  "lang":  "en",
	  "value": "Hello"
  	}`
	b, _ := json.Marshal(English("Hello"))
	assert.JSONEq(t, expected, string(b), "should be equal")
}

func TestJapanese(t *testing.T) {
	expected := `{
	  "type":  "PlainText",
	  "lang":  "ja",
	  "value": "こんにちは"
  	}`
	b, _ := json.Marshal(Japanese("こんにちは"))
	assert.JSONEq(t, expected, string(b), "should be equal")
}

func TestLink(t *testing.T) {
	expected := `{
	  "type":  "URL",
	  "lang":  "",
	  "value": "http://example.mp3"
  	}`
	b, _ := json.Marshal(Link("http://example.mp3"))
	assert.JSONEq(t, expected, string(b), "should be equal")
}

func TestHLS(t *testing.T) {
	expected := `{
      "contentType": "application/vnd.apple.mpegurl",
	  "type":        "URL",
	  "lang":        "",
	  "value":       "http://example.mp3"
  	}`
	b, _ := json.Marshal(HLS("http://example.mp3"))
	assert.JSONEq(t, expected, string(b), "should be equal")
}
