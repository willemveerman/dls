package main

import (
	"fmt"
	//"encoding/xml"
	"os"
)

func main() {

	xml, err := os.Open("~/Documents/FSM/process-conf.xml")

	if err != nil {
		fmt.Println(err)
	} else {
	fmt.Println("Successfully opened XML file.")
	}

	defer xml.Close()
}
