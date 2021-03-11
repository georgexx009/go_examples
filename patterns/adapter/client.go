package main

type client struct {
}

func (c *client) connectLightin(com computer) {
	com.insertIntoLightningPort()
}
