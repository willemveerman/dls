package main

import (
	"fmt"
	"encoding/xml"
	"os"
	"io/ioutil"
)

type Beans struct {
	XMLName xml.Name `xml:"beans"`
	BeanList []Bean `xml:"bean"`
}

type Bean struct {
	Description string `xml:"description"`
/*	Entry string `xml:"description""`
	Description string `xml:"description""`
	Description string `xml:"description""`
	Description string `xml:"description""`
	Description string `xml:"description""`
	Description string `xml:"description""`*/
}


func main() {

	xmlfile, err := os.Open("/Users/willemveerman/Documents/FSM/process-conf.xml")

	if err != nil {
		fmt.Println(err)
	} else {
	fmt.Println("Successfully opened XML file.")
	}

	byteValue, _ := ioutil.ReadAll(xmlfile)

	var b Beans

	xml.Unmarshal(byteValue, &b)

	fmt.Println(b)

	defer xmlfile.Close()
}
