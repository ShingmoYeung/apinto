package plugin

import "github.com/eolinker/goku/filter"

//Config 普通插件配置，在router、service、upstream的插件格式
type Config struct {
	Disable bool        `json:"disable"`
	Config  interface{} `json:"config"`
}
type IPlugin interface {
	filter.IChain
	Destroy()
}
type IPluginManager interface {
	CreateRouter(id string, conf map[string]*Config) IPlugin
	CreateService(id string, conf map[string]*Config) IPlugin
	CreateUpstream(id string, conf map[string]*Config) IPlugin
}

func MergeConfig(high, low map[string]*Config) map[string]*Config {
	if high == nil && low == nil {
		return make(map[string]*Config)
	}
	if high == nil {
		return clone(low)
	}
	if low == nil {
		return clone(low)
	}

	mv := clone(low)

	for k, hv := range high {
		lv, has := mv[k]
		if has {
			*lv = *hv
		} else {
			c := new(Config)
			*c = *hv
			mv[k] = c
		}
	}
	return mv

}
func clone(v map[string]*Config) map[string]*Config {
	cv := make(map[string]*Config)
	if v == nil {
		return cv
	}
	for k, v := range v {
		c := new(Config)
		*c = *v
		cv[k] = c
	}
	return cv
}
