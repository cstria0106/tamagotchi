package versionutil

import (
	"io/ioutil"
	"tamagotchi/internal/data/version"
)

func GetVersion() (*version.Version, error) {
	yaml, err := ioutil.ReadFile("resources/version.yaml")
	if err != nil {
		return nil, err
	}

	v, err := version.FromYML(yaml)
	if err != nil {
		return nil, err
	}

	return v, nil
}
