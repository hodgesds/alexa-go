package api

type Speech string

type CardType string

const (
	Plaintext Speech = "Plaintext"
	SSML      Speech = "SSML"

	Simple   CardType = "Simple"
	Standard CardType = "Standard"
	Linked   CardType = "LinkAccount"
)

type Output struct {
	Ssml string `json:"ssml,omitempty"`
	Text string `json:"text,omitempty"`
	Type Speech `json:"type,"`
}

type Image struct {
	LargeURI string `json:"largeImageUrl,omitempty"`
	SmallURI string `json:"smallImageUrl,omitempty"`
}

type Card struct {
	Content string   `json:"content,omitempty"`
	Image   *Image   `json:"image,omitempty"`
	Text    string   `json:"text,omitempty"`
	Title   string   `json:"title,omitempty"`
	Type    CardType `json:"type,"`
}

type AlexaResponse struct {
	Card       *Card   `json:"card,omitempty"`
	EndSession bool    `json:"shouldEndSession,"`
	Output     *Output `json:"outputSpeech,omitempty"`
	Reprompt   *Output `json:"reprompt,omitempty"`
}

type Response struct {
	AlexaResponse *AlexaResponse         `json:"response,"`
	Session       map[string]interface{} `json:"sessionAtrributes,omitempty"`
	Version       string                 `json:"version,"`
}
