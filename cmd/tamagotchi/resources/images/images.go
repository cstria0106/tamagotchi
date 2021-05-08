package images

import (
	"bufio"
	"image"
	_ "image/png"
	"os"
)

var Character image.Image
var CleanButton image.Image
var FeedButton image.Image
var Poo image.Image
var Food image.Image
var Cursor image.Image

func load(filename string) (image.Image, error) {
	var file *os.File
	var err error

	if file, err = os.Open(filename); err != nil {
		return nil, err
	}

	var i image.Image
	if i, _, err = image.Decode(bufio.NewReader(file)); err != nil {
		return nil, err
	}

	return i, nil
}

func Load() error {
	var err error

	Character, err = load("resources/images/cat.png")
	if err != nil {
		return err
	}

	CleanButton, err = load("resources/images/clean_button.png")
	if err != nil {
		return err
	}

	FeedButton, err = load("resources/images/feed_button.png")
	if err != nil {
		return err
	}

	Poo, err = load("resources/images/poo.png")
	if err != nil {
		return err
	}

	Food, err = load("resources/images/rice.png")
	if err != nil {
		return err
	}

	Cursor, err = load("resources/images/cursor.png")
	if err != nil {
		return err
	}

	return nil
}
