package images

import (
	"bytes"
	_ "embed"
	"image"
)

//go:embed files/cat.png
var characterData []byte

//go:embed files/clean_button.png
var cleanButtonData []byte

//go:embed files/feed_button.png
var feedButtonData []byte

//go:embed files/poo.png
var pooData []byte

//go:embed files/rice.png
var foodData []byte

var Character, _, _ = image.Decode(bytes.NewReader(characterData))
var CleanButton, _, _ = image.Decode(bytes.NewReader(cleanButtonData))
var FeedButton, _, _ = image.Decode(bytes.NewReader(feedButtonData))
var Poo, _, _ = image.Decode(bytes.NewReader(pooData))
var Food, _, _ = image.Decode(bytes.NewReader(foodData))
