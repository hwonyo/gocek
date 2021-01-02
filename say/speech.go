package say

// ---------------------------------------------------------------------------------------------------------------------
// https://developers.naver.com/console/clova/custom_ext/Develop/References/Custom_Extension_Message.md#CustomExtSpeechInfoObject
type SpeechInfoObject struct {
	ContentType string       `json:"contentType,omitempty"`
	Lang        LanguageType `json:"lang"`
	Token       string       `json:"token,omitempty"`
	Type        SpeechType   `json:"type"`
	Value       string       `json:"value"`
}

// ---------------------------------------------------------------------------------------------------------------------
// Korean function return *SpeechInfoObject with Language type Ko
func Korean(value string) *SpeechInfoObject {
	return &SpeechInfoObject{
		Value: value,
		Lang:  LangKo,
		Type:  SpeechTypeText,
	}
}

// English function return *SpeechInfoObject with Language type En
func English(value string) *SpeechInfoObject {
	return &SpeechInfoObject{
		Value: value,
		Lang:  LangEn,
		Type:  SpeechTypeText,
	}
}

// Japanese function return *SpeechInfoObject with Speech Type PlainText and Language type Ja
func Japanese(value string) *SpeechInfoObject {
	return &SpeechInfoObject{
		Value: value,
		Lang:  LangJa,
		Type:  SpeechTypeText,
	}
}

// Link function return *SpeechInfoObject with Speech Type URL
func Link(value string) *SpeechInfoObject {
	return &SpeechInfoObject{
		Value: value,
		Type:  SpeechTypeURL,
		Lang:  LangNone,
	}
}

func HLS(value string) *SpeechInfoObject {
	return &SpeechInfoObject{
		ContentType: "application/vnd.apple.mpegurl",
		Value:       value,
		Type:        SpeechTypeURL,
		Lang:        LangNone,
	}
}

// ---------------------------------------------------------------------------------------------------------------------
type LanguageType string

const (
	LangKo   LanguageType = "ko"
	LangJa   LanguageType = "ja"
	LangEn   LanguageType = "en"
	LangNone LanguageType = ""
)

type SpeechType string

const (
	SpeechTypeText SpeechType = "PlainText"
	SpeechTypeURL  SpeechType = "URL"
)
