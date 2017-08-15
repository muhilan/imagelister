package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Compose struct {
	Service map[string]Services `yaml:"services"`
}
type Services struct {
	Image string `yaml:"image"`
}



func main() {
	var filename string
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	if filename == "" {
		filename = "docker-compose-dev.yml"
	}

	data , err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var c Compose

	err = yaml.Unmarshal(data, &c)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	outputfilename := "images.txt"
	f, err :=os.OpenFile(outputfilename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	for k := range c.Service {
		fmt.Println(k)
		fmt.Fprintln(f,c.Service[k].Image)
	}
}
