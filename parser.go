package main

import (
	"fmt"
	"encoding/xml"
	"os"
	"io/ioutil"
)

type Bean struct {
	XMLName xml.Name `xml:"beans"`
	Bean []string `xml"bean"`
	Id []string `xml:"id,attr"`
	//Value []string `xml:value,attr`
}

func main() {

	xmlfile, err := os.Open("/Users/willemveerman/Documents/FSM/process-conf.xml")

	if err != nil {
		fmt.Println(err)
	} else {
	fmt.Println("Successfully opened XML file.")
	}

	byteValue, _ := ioutil.ReadAll(xmlfile)

	var b Bean

	xml.Unmarshal(byteValue, &b)

	fmt.Println(b)

	defer xmlfile.Close()
}
