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
	object map[string]map[string]string
}

func unpackElements(packed Beans) Objects {

	var objects Objects

	objects.object = make(map[string]map[string]string)

	for _,v := range packed.BeanList {
		objects.object[v.Id] = make(map[string]string)
		objects.object[v.Id]["description"] = v.Description
		objects.object[v.Id]["class"] = v.Class
		objects.object[v.Id]["singleton"] = v.Singleton
		for _,c := range v.EntryList {
			objects.object[v.Id][c.Key] = c.Value
		}
	}

	return objects
}

func packElements(unpacked P_file) Beans {

	var packedfile Beans

	for k,v := range unpacked.beans.object {
		var bean Bean
		bean.Id = k
		bean.Description = v["description"]
		bean.Class = v["class"]
		bean.Singleton = v["singleton"]
		for key,val := range v {
			if key != "class" && key != "singleton" && key != "description" {
				var e Entry
				e.Key = key
				e.Value = val
				bean.EntryList = append(bean.EntryList, e)
			}
		}

		packedfile.BeanList = append(packedfile.BeanList, bean)
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

	c := b

	xml.Marshal(c)

	objects := unpackElements(b)

	ff := b.BeanList[0].Description

	fmt.Println("Description:",ff)

	var pfile P_file

	pfile.total = len(objects.object)

	pfile.beans = objects

	packed := packElements(pfile)

	marshalled, err := xml.MarshalIndent(packed,"","	")


	ioutil.WriteFile("xmlout.xml",marshalled, 0644)

	fmt.Println(objects.object["Address"])

	for k,_ := range objects.object {
		fmt.Println(k)
	}

	fmt.Println(len(objects.object))

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
