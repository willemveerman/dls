package main

import (
	"fmt"
	"encoding/xml"
	"os"
	"io/ioutil"
)

type Beans struct {
	BeanList []Bean `xml:"bean"`
}

type Bean struct {
	Id string `xml:"id,attr"`
	Description string `xml:"description"`
	EntryList []Entry `xml:"property>map>entry"`
}

type Entry struct {
	Key string `xml:"key,attr"`
	Value string `xml:"value,attr"`
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

	c := b.BeanList

	fmt.Println(c[0].EntryList[0])

	defer xmlfile.Close()
}
