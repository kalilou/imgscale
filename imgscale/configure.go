package imgscale

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Configure(config *Config) *Handler {
	for _, ext := range config.Exts {
		if _, ok := supportedExts[ext]; !ok {
			panic(fmt.Sprintf("Extension '%s' not supported", ext))
		}
	}

	prefixes := make([]string, len(config.Formats))
	formats := make(map[string]*Format)
	for i, format := range config.Formats {
		prefixes[i] = format.Prefix
		formats[format.Prefix] = format
	}

	path := fmt.Sprintf("/%s/(?P<format>%s)/(?P<filename>.+)\\.(?P<ext>%s)", config.Prefix, strings.Join(prefixes, "|"), strings.Join(config.Exts, "|"))

	return &Handler{Formats: formats, Path: path, Config: config, regexp: regexp.MustCompile(path), supportedExts: supportedExts}
}

func LoadConfig(filename string) *Config {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := Config{}
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err)
	}
	return &conf
}