package core

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

// YamlConfig is exported.
type YamlConfig struct {
	Routers []struct {
		Router string `yaml:"index"`
		Url    string `yaml:"url"`
	} `yaml:"Routers"`
}

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "go start")
}

func RunHandle() {

	fmt.Println("Hello")

	configPath := "./cmd/code/adminhtml/etc.yml"

	if configPath != "" {
		fmt.Println("Parsing config: %s", configPath)
		viper.SetConfigFile(configPath)
		err := viper.ReadInConfig()
		if err != nil {
			log.Panicf("Unable to read config file: %s", err)
		}
	} else {
		fmt.Printf("Config file is not specified.")
	}

	router := viper.GetString("router")

	http.HandleFunc(fmt.Sprintf("/%s", router), home_page)
	http.ListenAndServe(":8080", nil)

	fmt.Println(viper.GetString("router"))

}
