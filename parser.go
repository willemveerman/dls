package main

import (
	"fmt"
	//"encoding/xml"
	"os"
	//"path/filepath"
)

func main() {

	xml, err := os.Open("/Users/willemveerman/Documents/FSM/process-conf.xml")

	if err != nil {
		fmt.Println(err)
	} else {
	fmt.Println("Successfully opened XML file.")
	}

	defer xml.Close()
}
