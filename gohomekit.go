package main

import (
	"flag"
	"github.com/bingbaba/util/logs"
	"github.com/xuebing1110/gohomekit/bridge"

	_ "github.com/xuebing1110/gohomekit-aqara"
)

var (
	file          *string
	log           *string
	debug, plugin *bool
)

func init() {
	debug = flag.Bool("debug", false, "the debug model")
	plugin = flag.Bool("plugin", false, "the plugin model")
	file = flag.String("f", "./conf/app.json", "the json configure file")
	log = flag.String("l", "./gohomekit.log", "the log file")
}

func main() {
	flag.Parse()
	if *debug {
		logs.SetDebug(*debug)
	} else {
		logs.InitByString(`{"filename":"` + *log + `"}`)
	}

	LOGGER := logs.GetBlogger()

	//read configure
	b, err := bridge.NewByFile(*file)
	if err != nil {
		panic(err)
	}

	/*
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
	*/

	//add plugin
	for _, pf := range bridge.PlatForms {
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
