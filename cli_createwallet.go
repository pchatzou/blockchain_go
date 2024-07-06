package main

import "fmt"

func (cli *CLI) createUser(nodeID string) {
	users, _ := NewUsers(nodeID)
	address := users.CreateUser()
	users.SaveToFile(nodeID)

	fmt.Printf("Your new address: %s\n", address)
}
