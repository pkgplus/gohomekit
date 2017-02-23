package bridge

import (
	"github.com/brutella/hc/accessory"
	"log"
)

type PlatForm interface {
	New(sid, pwd string) PlatForm
	GetName() string
	Init() error
	Start() error
	Stop() error
	OnError(err error)
	GetAccessorys() []*accessory.Accessory
}

type BasePlatForm struct {
	accs []*accessory.Accessory
}

func NewBasePlatForm() *BasePlatForm {
	return &BasePlatForm{
		accs: make([]*accessory.Accessory, 0),
	}
}

func (this *BasePlatForm) Init() error {
	return nil
}

func (this *BasePlatForm) OnError(err error) {
	log.Printf("Error %v\n", err)
}

func (this *BasePlatForm) AddAcc(acc *accessory.Accessory) {
	this.accs = append(this.accs, acc)
}

func (this *BasePlatForm) GetAccessorys() []*accessory.Accessory {
	return this.accs
}
