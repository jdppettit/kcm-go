package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var config_directory_path = "~/.kcm"
var default_configuration = "{\"configs\": []}"

func main() {
	config_directory_present()
	create_config_file_if_needed()
	switch command := os.Args[1]; command {
	case "list":
		fmt.Println("list")
	case "status":
		fmt.Println("status")
	case "create":
		fmt.Println("create")
	case "remove":
		fmt.Println("remove")
	case "activate":
		fmt.Println("activate")
	default:
		fmt.Println("Unknown command provided: ", command)
	}
}

type Config struct {
	name   string
	path   string
	active string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func config_directory_present() bool {
	if _, err := os.Stat(config_directory_path); err != nil {
		return true
	} else {
		return false
	}
}

func create_config_directory() {
	error := os.MkdirAll("~/.kcm", 0755)
	check(error)
}

func create_config_file_if_needed() {
	//f, err := os.Create("~/.kcm/config.json")
	//check(err)

	//defer f.Close()

	config_to_write, _ := json.Marshal(default_configuration)
	fmt.Println(config_to_write)
}
