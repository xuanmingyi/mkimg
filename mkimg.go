package main

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type config struct {
	Output    string `yaml:"output"`
	Boot      string `yaml:"boot"`
	BootBytes []byte `yaml:"-"`
}

var Config config

// 1.44mb
var size = int64(1024 * 1440)

func main() {
	content, err := ioutil.ReadFile("mkimg.yml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(content, &Config)
	if err != nil {
		panic(err)
	}

	content, err = ioutil.ReadFile(Config.Boot)
	if err != nil {
		panic(err)
	}

	Config.BootBytes = content

	output, err := os.Create(Config.Output)
	if err != nil {
		panic(err)
	}

	err = output.Truncate(size)
	if err != nil {
		panic(err)
	}

	_, err = output.Seek(0, 0)
	if err != nil {
		panic(err)
	}
	_, err = output.Write(Config.BootBytes)
	if err != nil {
		panic(err)
	}
}
