package main

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type config struct {
	Output     string   `yaml:"output"`
	OutputFile *os.File `yaml:"-"`
	Boot       string   `yaml:"boot"`
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

func (c *config) WriteBoot() (err error) {
	content, err := ioutil.ReadFile(c.Boot)
	if err != nil {
		return err
	}

	_, err = c.OutputFile.Seek(0, 0)
	if err != nil {
		return err
	}

	_, err = c.OutputFile.Write(content)
	if err != nil {
		return err
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

	err = Config.WriteBoot()
	if err != nil {
		panic(err)
	}
}
