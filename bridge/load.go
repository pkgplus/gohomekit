package bridge

import (
	"errors"
	//"io/ioutil"
	//"log"
	"os"
	//"path/filepath"
	//"plugin"
	//"runtime"
	//"strings"
)

func LoadPlugins() (pfs []PlatForm, err error) {
	pfs = make([]PlatForm, 0)

	/*
		var homeDir string
		if "windows" == runtime.GOOS {
			homeDir, err = homeWindows()
			if err != nil {
				return
			}
		} else {
			homeDir = os.Getenv("HOME")
		}
		homeDir = filepath.Join(homeDir, ".gohomekit")

		var dir_list []os.FileInfo
		dir_list, err = ioutil.ReadDir(homeDir)
		if err != nil {
			return
		}

		for _, v := range dir_list {
			if !strings.HasSuffix(v.Name(), ".so") {
				continue
			}

			var p *plugin.Plugin
			p, err = plugin.Open(v.Name())
			if err != nil {
				log.Println(err)
				continue
			}

			var ps plugin.Symbol
			ps, err = p.Lookup("PlatForm")
			if err != nil {
				log.Println(err)
				continue
			}

			pf_func, ok := ps.(func() (PlatForm, error))
			if !ok {
				log.Println("found a incorrect plugin:" + v.Name())
				continue
			}

			var pf PlatForm
			pf, err := pf_func()
			if err != nil {
				log.Printf("load plugin %s error: %v", v.Name(), err)
				continue
			}

			pfs = append(pfs, pf)
		}
	*/

	return
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}
