package configs

import (
	"os"
	"strings"
)

type configuration struct {
	RootPath string
	API      struct {
		GRPC api `json:"grpc"`
		HTTP api `json:"http"`
	} `json:"api"`
	Data struct {
		Driver string
		Source string
	}
}

type api struct {
	Addr, Timeout string
}

var V = &configuration{}

func setRootPath() error {
	if strings.Contains(os.Args[0], ".test") {
		V.RootPath = "../../" // for test
		return nil
	}
	root, err := os.Getwd()
	if err != nil {
		return err
	}
	V.RootPath = root
	return nil
}
