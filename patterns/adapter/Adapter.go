package main

type adapter struct {
	windowsMachine *windows
}

func (adapter *adapter) insertIntoLightningPort() {
	adapter.windowsMachine.connectUsb()
}
