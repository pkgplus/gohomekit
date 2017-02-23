package myaccessory

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
)

type BridgeStatus struct {
	*accessory.Accessory
	BridgingState *service.BridgingState
}

func NewBridgeStatus(info accessory.Info) *BridgeStatus {
	acc := BridgeStatus{}
	acc.Accessory = accessory.New(info, accessory.TypeBridge)
	acc.BridgingState = service.NewBridgingState()
	acc.BridgingState.Reachable.SetValue(true)
	acc.AddService(acc.BridgingState.Service)

	return &acc
}
