package configs

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var ProjectName = "udp2mysql"

type configuration struct {
	RootPath string
	API      struct {
		GRPC api `json:"grpc"`
		HTTP api `json:"http"`
	} `json:"api"`
	Database struct {
		Driver string `json:"Driver"`
		Source string `json:"Source"`
	} `json:"database"`
}

type api struct {
	Addr, Timeout string
}

var V = &configuration{}

func setRootPath() error {
	if strings.Contains(os.Args[0], ".test") {
		rootPath4Test()
		return nil
	}
	root, err := os.Getwd()
	if err != nil {
		return err
	}
	V.RootPath = root
	return nil
}

func load() error {
	cf := filepath.Join(V.RootPath, "configs/configs.json")
	f, err := os.ReadFile(cf)
	if err != nil {
		return err
	}
	return json.Unmarshal(f, V)
}

func init() {
	if err := setRootPath(); err != nil {
		log.Printf("configs init error: %v", err)
	}
	if err := load(); err != nil {
		log.Printf("configs load error: %v", err)
	}
}

func rootPath4Test() int {
	ps := strings.Split(os.Args[0], ProjectName)
	n := 0
	if len(ps) == 1 { // go test
		if runtime.GOOS == "windows" {
			n = strings.Count(ps[0], "\\") - 4
		} else {
			n = strings.Count(ps[0], "/") - 4
		}
	} else { // dlv
		if runtime.GOOS == "windows" {
			n = strings.Count(ps[1], "\\") - 1
		} else {
			n = strings.Count(ps[1], "/") - 1
		}
	}
	for i := 0; i < n; i++ {
		V.RootPath = filepath.Join("../", V.RootPath)
	}
	return n
}
