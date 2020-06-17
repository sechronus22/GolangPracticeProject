package main

import "mockproject/internals/container"

func main() {
	c, err := container.NewContainers()
	if err != nil {
		panic(err)
	}

	if err := c.Start(); err != nil {
		panic(err)
	}
}
