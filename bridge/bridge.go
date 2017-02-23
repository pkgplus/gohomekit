package bridge

import (
	"encoding/json"
	"errors"
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"io/ioutil"
	"os"
	"strings"
)

type Bridge struct {
	conf  *Configure
	plats []PlatForm
}

func NewByFile(path string) (b *Bridge, err error) {
	b = &Bridge{
		plats: make([]PlatForm, 0),
	}

	var fi *os.File
	fi, err = os.Open(path)
	if err != nil {
		return
	}
	defer fi.Close()

	var content []byte
	content, err = ioutil.ReadAll(fi)
	if err != nil {
		return
	}

	b.conf = &Configure{}
	err = json.Unmarshal(content, b.conf)
	if err != nil {
		return
	}

	b.conf.Bridge.PIN = strings.Replace(
		b.conf.Bridge.PIN,
		"-",
		"",
		-1,
	)
	return
}

func New(conf *Configure) (b *Bridge) {
	b = &Bridge{
		conf:  conf,
		plats: make([]PlatForm, 0),
	}
	b.conf.Bridge.PIN = strings.Replace(
		b.conf.Bridge.PIN,
		"-",
		"",
		-1,
	)
	return
}

func (this *Bridge) Add(plat PlatForm) (err error) {
	var count int = 0
	for _, c := range this.conf.PlatForms {
		if c.PlatForm != plat.GetName() {
			continue
		}

		for i := 0; i < len(c.Password); i++ {
			pf := plat.New(c.Sid[i], c.Password[i])
			err = pf.Init()
			if err != nil {
				return
			}

			this.plats = append(this.plats, pf)
			count++
		}
	}

	if count == 0 {
		return errors.New("can't found the platform configure:" + plat.GetName())
	}
	return
}

func (this *Bridge) Start() (err error) {
	for _, plat := range this.plats {
		err = plat.Start()
		if err != nil {
			return
		}
	}

	accs := this.GetAccessorys()
	t, err := hc.NewIPTransport(
		hc.Config{Pin: this.conf.Bridge.PIN},
		accs[0],
		accs[1:]...,
	)
	if err != nil {
		return
	}

	hc.OnTermination(func() {
		this.Stop()
		t.Stop()
	})

	t.Start()

	return
}

func (this *Bridge) Stop() (err error) {
	for _, plat := range this.plats {
		err = plat.Stop()
		if err != nil {
			return
		}
	}
	return
}

func (this *Bridge) GetAccessorys() []*accessory.Accessory {
	accs := make([]*accessory.Accessory, 0)
	for _, plat := range this.plats {
		accs = append(accs, plat.GetAccessorys()...)
	}
	return accs
}
