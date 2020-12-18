package yaml

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func NewConfig(f string) (cf Config, err error) {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(data, &cf)
	if err != nil {
		return
	}
	return
}

type Config struct{}
