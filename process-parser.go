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

type Objects struct {
	objects map[string]map[string]string
}

type P_file struct {
	total int
	beans Objects
}

func (p P_file) capitalise(s string) string {
	var a string
	return a
}

func getElements_dict(b Beans) map[string]map[string]string {

	objects_dict := make(map[string]map[string]string)

	for _,v := range(b.BeanList) {
		objects_dict[v.Id] = make(map[string]string)
		for _,c := range(v.EntryList) {
				objects_dict[v.Id][c.Key] = c.Value
		}
	}

	return objects_dict
}

func getElements(b Beans) Objects {

	var objects_struct Objects

	objects_struct.objects = make(map[string]map[string]string)

	for _,v := range(b.BeanList) {
		objects_struct.objects[v.Id] = make(map[string]string)
		for _,c := range(v.EntryList) {
			objects_struct.objects[v.Id][c.Key] = c.Value
		}
	}

	return objects_struct
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

	c := b

	xml.Marshal(c)

	d := c

	objects := getElements(b)

	var pfile P_file

	pfile.total = len(objects.objects)

	pfile.beans = objects

	//var address Object

	//address.Attribute = element_dict["Address"]

	//fmt.Println(b.BeanList[0].Id)

	//fmt.Println(address.Attribute["sfdc.proxyPort"])

	//var address Object

	//for key,value := range(element_dict["Address"]) {
	//	fmt.Println(key, value)
	//}

	fmt.Println(objects.objects["Address"]["sfdc.proxyPort"])

	fmt.Println(len(objects.objects))

	fmt.Println(pfile.beans.objects["Location"])

	fmt.Println(d)

	defer xmlfile.Close()
}
