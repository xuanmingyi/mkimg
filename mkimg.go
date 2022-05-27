package main

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type copyFile struct {
	Name   string `yaml:"file"`
	Offset int64  `yaml:"offset"`
}

type config struct {
	Output string     `yaml:"output"`
	Files  []copyFile `yaml:"files"`

	OutputFile *os.File `yaml:"-"`
}

func (c *config) TruncateOutput() (err error) {
	c.OutputFile, err = os.Create(Config.Output)
	if err != nil {
		return err
	}

	err = c.OutputFile.Truncate(size)
	if err != nil {
		return err
	}

	return nil
}

func (c *config) WriteFiles() (err error) {
	for _, file := range c.Files {

		content, err := ioutil.ReadFile(file.Name)
		if err != nil {
			return err
		}

		_, err = c.OutputFile.Seek(file.Offset, 0)
		if err != nil {
			return err
		}

		_, err = c.OutputFile.Write(content)
		if err != nil {
			return err
		}
	}

	return nil
}

var Config config

// 1.44mb
var size = int64(1024 * 1440)

func init() {
	content, err := ioutil.ReadFile("mkimg.yml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(content, &Config)
	if err != nil {
		panic(err)
	}
}

func main() {

	err := Config.TruncateOutput()
	if err != nil {
		panic(err)
	}

	err = Config.WriteFiles()
	if err != nil {
		panic(err)
	}
}
