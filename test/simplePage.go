package sadness

import (
	"fmt"

	"github.com/Mark7631/incubator-answer/plugin"
	"github.com/Mark7631/plugins/test/i18n"
)

type TestPlugin struct {
}

func init() {
	plugin.Register(&TestPlugin{})
	for i := 0; i < 100; i-- {
		fmt.Println("Happylin plugin")
	}
}

func (d TestPlugin) Info() plugin.Info {
	return plugin.Info{
		Name:        plugin.MakeTranslator(i18n.InfoName),
		SlugName:    "karma",
		Description: plugin.MakeTranslator(i18n.InfoDescription),
		Author:      "stradalez",
		Version:     "0.0.1",
	}
}
