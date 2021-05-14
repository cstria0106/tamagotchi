package versionutil

import (
	"github.com/cstria0106/tamagotchi/internal/data/version"
	"io/ioutil"
)

func GetLocalVersion() (*version.Version, error) {
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
