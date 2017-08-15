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
	data , err := ioutil.ReadFile("docker-compose-dev.yml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var c Compose

	err = yaml.Unmarshal(data, &c)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	filename := "images.txt"
	f, err :=os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	for k ,_ := range c.Service {
		fmt.Println(k)
		fmt.Fprintln(f,c.Service[k].Image)
	}
	//fmt.Println(c.Service["api-central"].Image)
}
