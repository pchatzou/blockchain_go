package main

import (
	"fmt"
	"log"
)

func (cli *CLI) listAddresses(nodeID string) {
	users, err := NewUsers(nodeID)
	if err != nil {
		log.Panic(err)
	}
	addresses := users.GetAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}
}
