package main

import (
	"flag"
	"github.com/bingbaba/util/logs"
	"github.com/xuebing1110/gohomekit/aqara"
	"github.com/xuebing1110/gohomekit/bridge"
)

var (
	file  *string
	debug *bool
)

func init() {
	debug = flag.Bool("debug", false, "the debug module")
	file = flag.String("f", "./conf/app.json", "the json configure file")
}

func main() {
	flag.Parse()
	logs.SetDebug(*debug)

	b, err := bridge.NewByFile(*file)
	if err != nil {
		panic(err)
	}

	err = b.Add(&aqara.Aqara{})
	if err != nil {
		panic(err)
	}

	err = b.Start()
	if err != nil {
		panic(err)
	}

}
