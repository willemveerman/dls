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

type Object struct {
	Id string
	Attrs struct {
		externalIdField []string
		outputError []string
		proxyPort []string
		operation []string
		endpoint []string
		timeoutSecs []string
		initialLastRunDate []string
		proxyUsername []string
		loadBulkApi []string
		mappingFile []string
		name []string
		bulkApiCheckStatusInterval []string
		entity []string
		password []string
		proxyHost []string
		proxyPassword []string
		loadBatchSize []string
		bulkApiSerialMode []string
		debugMessages []string
		debugMessagesFile []string
		username []string
		Type []string
		outputSuccess []string
		readUTF8 []string
		encryptionKeyFile []string
		timezone []string
		writeUTF8 []string

	}
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

	for key,_ := range(a["Address"]) {
		fmt.Println(key)
	}

	defer xmlfile.Close()
}
