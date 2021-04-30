package images

import (
	"bytes"
	_ "embed"
	"image"
)

//go:embed files/character.png
var characterData []byte

//go:embed files/clean_button.png
var cleanButtonData []byte

//go:embed files/feed_button.png
var feedButtonData []byte

var Character, _, _ = image.Decode(bytes.NewReader(characterData))
var CleanButton, _, _ = image.Decode(bytes.NewReader(cleanButtonData))
var FeedButton, _, _ = image.Decode(bytes.NewReader(feedButtonData))
