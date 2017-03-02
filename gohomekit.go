package main

import (
	"flag"
	"github.com/bingbaba/util/logs"
	"github.com/xuebing1110/gohomekit-aqara"
	"github.com/xuebing1110/gohomekit/bridge"
)

var (
	file          *string
	debug, plugin *bool
)

func init() {
	debug = flag.Bool("debug", false, "the debug model")
	plugin = flag.Bool("plugin", false, "the plugin model")
	file = flag.String("f", "./conf/app.json", "the json configure file")
}

func main() {
	flag.Parse()
	logs.SetDebug(*debug)
	LOGGER := logs.GetBlogger()

	//read configure
	b, err := bridge.NewByFile(*file)
	if err != nil {
		panic(err)
	}

	var pfs []bridge.PlatForm

	//load plugin
	if *plugin {
		pfs, err = bridge.LoadPlugins()
		if err != nil {
			panic(err)
		}
	} else {
		pfs = []bridge.PlatForm{
			&aqara.Aqara{},
		}
	}

	//add plugin
	for _, pf := range pfs {
		err = b.Add(pf)
		if err != nil {
			LOGGER.Error("add plugin %s error: %v", pf.GetName(), err)
		}
	}

	//start
	err = b.Start()
	if err != nil {
		panic(err)
	}

}
