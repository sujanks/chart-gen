package process

import (
	"io/ioutil"
	"log"
)

func Read(name string) *[]byte{
	content, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal("error reading file")
	}
	return &content
}
