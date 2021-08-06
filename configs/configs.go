package configs

import (
	"encoding/json"
	"fmt"
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
	UDPSender struct {
		Addr    string `json:"Addr"`
		BufSize int    `json:"BufSize"`
	} `json:"udp_sender"`
	Web struct {
		Addr string `json:"Addr"`
		Tmpl string `json:"Tmpl"`
	} `json:"web"`
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

func rootPath4Test() error {
	root, err := os.Getwd()
	fmt.Println(root)
	if err != nil {
		return err
	}
	ps := strings.Split(root, ProjectName)
	n := 0
	if runtime.GOOS == "windows" {
		n = strings.Count(ps[1], "\\")
	} else {
		n = strings.Count(ps[1], "/")
	}
	for i := 0; i < n; i++ {
		V.RootPath = filepath.Join("../", V.RootPath)
	}
	return nil
}
