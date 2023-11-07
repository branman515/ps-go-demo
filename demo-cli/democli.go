package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("What would you like me to scream?")
	in := bufio.NewReader(os.Stdin) //creates the reader we will be using
	s, _ := in.ReadString('\n')     // reads the string until a newline
	s = strings.TrimSpace(s)        //trim the newline out
	s = strings.ToUpper(s)          //capitalizes the string
	fmt.Println(s + "!")            //prints the scream
}
