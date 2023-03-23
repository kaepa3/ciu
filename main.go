package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

//-IC:/Program Files/NVIDIA GPU Computing Toolkit/CUDA/v12.0/include
//-IC:\tools\opencv\include
//--cuda-gpu-arch=sm_86
//--cuda-path=C:/Program Files/NVIDIA GPU Computing Toolkit/CUDA/v12.0/

type Config struct {
	Includes []string
	Options  []string
}

const (
	confFile = ".cu"
	textFile = "compile_flags.txt"
)

func loadEnv() *Config {
	var config Config
	paths := []string{
		filepath.Join(".env"),
		filepath.Join(os.Getenv("HOME"), confFile),
		filepath.Join(os.Getenv("HOMEPATH"), confFile),
	}
	for _, v := range paths {
		_, err := toml.DecodeFile(v, &config)
		if err != nil {
			continue
		}
		fmt.Println("use config", v)
		return &config
	}
	return nil
}

func confInit() *Config {
	return &Config{
		Includes: []string{
			"-IC:/Program Files/NVIDIA GPU Computing Toolkit/CUDA/v12.0/include",
			"-IC:/tools/opencv/include ",
		},
		Options: []string{
			"--cuda-gpu-arch=sm_86",
			"--cuda-path=C:/Program Files/NVIDIA GPU Computing Toolkit/CUDA/v12.0/",
		},
	}
}
func exist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func main() {
	config := loadEnv()
	if config == nil {
		config = confInit()
	}
	if exist(textFile) {
		log.Println("file exist")
		return
	}

	f, err := os.Create(textFile)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	writer := bufio.NewWriter(f)
	writeFile(writer, config.Includes)
	writeFile(writer, config.Options)
	writer.Flush()
}

func writeFile(w *bufio.Writer, str []string) {
	for _, v := range str {
		i, err := w.WriteString(v)
		if err != nil {
			log.Println(err)
		}
		log.Println(v, i)
	}
}
