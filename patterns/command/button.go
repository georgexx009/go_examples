package main

// invoker
type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}
