package yaml

import (
	"github.com/belson77/Go-001/Week04/news/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func NewConfig(f string) (config.Config, error) {
	var cf config.Config
	data, err := ioutil.ReadFile(f)
	if err == nil {
		err = yaml.Unmarshal(data, &cf)
	}
	return cf, err
}
