package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type options struct {
	desc       string
	cmd        string
	nextnode   *storynode
	nextoption *options
}

type storynode struct {
	text   string
	option *options
}

func (node *storynode) addoptions(cmd string, desc string, n *storynode) {
	option := &options{desc, cmd, n, nil}
	if node.option == nil {
		node.option = option
	} else {
		currentoption := node.option
		for currentoption.nextoption != nil {
			currentoption = currentoption.nextoption
		}
		currentoption.nextoption = option
	}

}

func (node *storynode) render() {
	fmt.Println(node.text)
	current := node.option
	for current != nil {
		fmt.Println(current.cmd, ":", current.desc)
		current = current.nextoption
	}
}

func (node *storynode) executecmd(cmd string) *storynode {
	current := node.option
	for current != nil {
		if strings.ToLower(current.cmd) == strings.ToLower(cmd) {
			return current.nextnode
		}
		current = current.nextoption
	}
	fmt.Printf("Invalid command to the given node....")
	return node
}

var scanner *bufio.Scanner

func (node *storynode) display() {
	node.render()
	if node.option != nil {
		scanner.Scan()
		node.executecmd(scanner.Text()).display()

	}

}

func main() {
	scanner = bufio.NewScanner(os.Stdin)

	start := storynode{text: "My name is bobby Axelrod"}
	second := storynode{text: "I am a billionaire, hedge-fund manager"}
	third := storynode{text: "My arch rivals are Chuck Rhoades and Taylor Mason"}
	start.addoptions("next", "next option", &second)
	second.addoptions("next", "nextoption", &third)

	start.display()

	fmt.Println("The END...")

}
