package slingshotplugin

import (
	"errors"

	"github.com/extism/extism"
)

// store all your plugins in a normal Go hash map, protected by a Mutex
var plugins = make(map[string]*extism.Plugin)

func StorePlugin(key string, plugin *extism.Plugin) {
	// store all your plugins in a normal Go hash map, protected by a Mutex
	plugins[key] = plugin

}

func GetPlugin(key string) (extism.Plugin, error) {

	if plugin, ok := plugins[key]; ok {
		return *plugin, nil
	} else {
		return extism.Plugin{}, errors.New("🔴 no plugin")
	}
}
