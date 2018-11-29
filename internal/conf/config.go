package conf

import (
	"fmt"
	"path"

	"strings"

	"github.com/micro/go-config"
	"github.com/micro/go-config/source"
	"github.com/micro/go-config/source/envvar"
	"github.com/micro/go-config/source/file"
)

type (
	spec struct {
		data   map[string]interface{}
		config Config
	}

	Config interface {
		Load(source ...source.Source) error
		Map() map[string]interface{}
	}
)

const (
	configNameBasic    = "config"
	configNameAdvanced = "ultraslackbot"
	envVarPrefix       = "USB"
)

var (
	extensions = [...]string{
		"json",
		"xml",
		"yaml",
		"hcl",
		"toml",
	}
)

func Load() *spec {
	s := newSpec()
	s.load()
	return s
}

func newSpec() *spec {
	return &spec{
		config: config.NewConfig(),
	}
}

func (s *spec) load() {
	s.config.Load(
		s.sources()...,
	)

	s.data = s.config.Map()
}

func (s *spec) sources() []source.Source {
	sources := make([]source.Source, 0)

	sources = append(sources, s.fileSources("./", configNameBasic)...)
	sources = append(sources, envvar.NewSource(envvar.WithStrippedPrefix(envVarPrefix)))
	sources = append(sources, s.fileSources("~/", configNameAdvanced)...)

	return sources
}

func (s *spec) fileSources(basePath, name string) []source.Source {
	sources := make([]source.Source, 0, len(extensions))

	for _, ext := range extensions {
		f := fmt.Sprintf("%s.%s", name, ext)
		fullPath := path.Join(basePath, f)

		sources = append(sources, file.NewSource(file.WithPath(fullPath)))
	}
	return sources
}

func (s *spec) Get(k string) (interface{}, bool) {
	v, ok := s.data[strings.ToLower(k)]
	return v, ok
}

func (s *spec) Set(k, v string) {
	s.data[k] = v
}

func (s *spec) All() map[string]interface{} {
	return s.data
}
