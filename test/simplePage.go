package sadness

import (
	"github.com/Mark7631/plugins/test/i18n"
	"github.com/Mark7631/incubator-answer/pkg/checker"
	"github.com/Mark7631/incubator-answer/plugin"
)

type FormulaPlugin struct {
}

func init() {
	plugin.Register(&FormulaPlugin{})
}

func (d FormulaPlugin) Info() plugin.Info {
	return plugin.Info{
		Name:        plugin.MakeTranslator(i18n.InfoName),
		SlugName:    "karma",
		Description: plugin.MakeTranslator(i18n.InfoDescription),
		Author:      "stradalez",
		Version:     "0.0.1",
	}
}