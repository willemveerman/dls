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

type XMLobjects struct {
	Object string
	Fields map[string]string
}

func getElements(b Beans) map[string]map[string]string {

	objects := make(map[string]map[string]string)

	for _,v := range(b.BeanList) {
		objects[v.Id] = make(map[string]string)
		for _,c := range(v.EntryList) {
				objects[v.Id][c.Key] = c.Value
		}
	}

	return objects
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

	a := getElements(b)

	fmt.Println(b.BeanList[0].Id)

	fmt.Println(a["Address"]["dataAccess.type"])

	defer xmlfile.Close()
}
