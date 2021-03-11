package main

func main() {
	mac := &mac{}
	client := &client{}

	w := &windows{}
	adapter := &adapter{
		windowsMachine: w,
	}

	client.connectLightin(mac)
	client.connectLightin(adapter)
}
