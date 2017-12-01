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
	Class string `xml:"class,attr"`
	Singleton string `xml:"singleton,attr"`
	Description string `xml:"description"`
	EntryList []Entry `xml:"property>map>entry"`
}

type Entry struct {
	Key string `xml:"key,attr"`
	Value string `xml:"value,attr"`
}

type P_file struct {
	total int
	beans Objects
}

type Objects struct {
	objects map[string]map[string]string
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
		objects_struct.objects[v.Id]["description"] = v.Description
		objects_struct.objects[v.Id]["class"] = v.Class
		objects_struct.objects[v.Id]["singleton"] = v.Singleton
		for _,c := range(v.EntryList) {
			objects_struct.objects[v.Id][c.Key] = c.Value
		}
	}

	return objects_struct
}

func packElements(f P_file) Beans {

	var packedfile Beans

	for k,v := range f.beans.objects {
		var b Bean
		b.Id = k
		b.Description = v["description"]
		b.Class = v["class"]
		b.Singleton = v["singleton"]
		for key,val := range v {
			if key != "class" && key != "singleton" && key != "description" {
				var e Entry
				e.Key = key
				e.Value = val
				b.EntryList = append(b.EntryList, e)
			}
		}

		packedfile.BeanList = append(packedfile.BeanList, b)
	}

	return packedfile
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

	//fmt.Println("b.BeanList[0]",b.BeanList)

	c := b

	xml.Marshal(c)

	objects := getElements(b)

	ff := b.BeanList[0].Description

	fmt.Println("Description:",ff)

	var pfile P_file

	pfile.total = len(objects.objects)

	pfile.beans = objects

	unpacked := packElements(pfile)

	packed, err := xml.Marshal(unpacked)


	ioutil.WriteFile("xmlout.xml",packed, 0644)

	//var address Object

	//address.Attribute = element_dict["Address"]

	//fmt.Println(b.BeanList[0].Id)

	//fmt.Println(address.Attribute["sfdc.proxyPort"])

	//var address Object

	//for key,value := range(element_dict["Address"]) {
	//	fmt.Println(key, value)
	//}

	fmt.Println(objects.objects["Address"])

	for k,_ := range objects.objects {
		fmt.Println(k)
	}

	fmt.Println(len(objects.objects))

	//fmt.Println(pfile.beans.objects["Location"])

	a := 42

	point := &a

	fmt.Println(*point)

	fmt.Println(&point)

	fmt.Println(point)

	line, _ := fmt.Println(point)

	fmt.Println(line)

	a = 21

	fmt.Println(*point)

	defer xmlfile.Close()
}
