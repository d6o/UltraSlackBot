package plugin

import (
	"io/ioutil"
	"log"
	"path"
	"plugin"
	"strings"

	"github.com/disiqueira/ultraslackbot/pkg/bot"
)

type (
	Plugin struct {
	}
)

const (
	dirName = "./plugins"
)

func New() *Plugin {
	return &Plugin{}
}

func (p *Plugin) Load(logger *log.Logger, specs bot.Specs) ([]bot.Handler, error) {
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		return nil, err
	}

	var handlers []bot.Handler
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileDetails := strings.Split(file.Name(), ".")
		if fileDetails[1] != "so" {
			continue
		}
		logger.Printf("Loading: %s \n", file.Name())
		fileName := path.Join("./plugins/", file.Name())
		plug, err := plugin.Open(fileName)
		if err != nil {
			logger.Printf("%s error: %s", fileName, err.Error())
			continue
		}
		symCustomPlugin, err := plug.Lookup("CustomPlugin")
		if err != nil {
			logger.Printf("%s error: %s", fileName, err.Error())
			continue
		}

		var customPlugin bot.Handler
		customPlugin, ok := symCustomPlugin.(bot.Handler)
		if !ok {
			logger.Printf("%s error: unexpected type from module symbol", fileName)
			continue
		}

		if err := customPlugin.Start(specs); err != nil {
			logger.Printf("%s error starting: %s", fileName, err.Error())
			continue
		}
		handlers = append(handlers, customPlugin)
	}

	return handlers, nil
}

