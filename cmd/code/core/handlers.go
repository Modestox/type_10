package core

import (
	"fmt"
	"net/http"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)


// YamlConfig is exported.
type YamlConfig struct {
	Routers  []struct{
		Routerd	string `yaml:"index"`
		Urld		string `yaml:"url"`
	} `yaml:"Routers"`
}

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "go start")	
}

func RunHandle() {
	// http.HandleFunc("/", home_page)
	fmt.Println("Hello")
	// http.ListenAndServe(":8080", nil)


	fileName := "./cmd/code/adminhtml/etc.yml"

    // yamlFile, err := ioutil.ReadFile(fileName)
    // if err != nil {
    //     fmt.Printf("Error reading YAML file: %s\n", err)
    //     return
    // }

	// var yamlConfig YamlConfig
	var conf YamlConfig
    reader, _ := os.Open(fileName)
    buf, _ := ioutil.ReadAll(reader)
    yaml.Unmarshal(buf, &conf)
    fmt.Printf("%+v\n", conf)
	
    // err = yaml.Unmarshal(yamlFile, &yamlConfig)
    // if err != nil {
    //     fmt.Printf("Error parsing YAML file: %s\n", err)
    // }

	
    // fmt.Printf("Result: %+v\n", yamlConfig)
}
