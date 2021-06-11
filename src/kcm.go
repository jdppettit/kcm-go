package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var config_directory_path = fmt.Sprintf("%s/.kcm-test", get_user_home_direcotry())
var full_config_directory_path = fmt.Sprintf("%s/.kcm-test/config.json", get_user_home_direcotry())
var default_configuration = "{\"configs\": []}"

func main() {
	if !config_directory_present() {
		create_config_directory()
	}

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

type config_container struct {
	Configs []config `json:"configs"`
}

type config struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Active string `json:"active"`
}

func get_user_home_direcotry() string {
	dirname, err := os.UserHomeDir()
	check(err)
	return dirname
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func config_directory_present() bool {
	if _, err := os.Stat(config_directory_path); err == nil {
		fmt.Println("true")
		return true
	} else {
		fmt.Println("false")
		return false
	}
}

func config_file_present() bool {
	if _, err := os.Stat(full_config_directory_path); err != nil {
		return true
	} else {
		return false
	}
}

func create_config_directory() {
	error := os.MkdirAll(config_directory_path, 0755)
	fmt.Println(error)
	check(error)
}

func create_config_file_if_needed() {
	if !config_file_present() {
		default_config := &config_container{}

		config_to_write, _ := json.MarshalIndent(default_config, "", " ")
		fmt.Println(string(config_to_write))

		ioutil.WriteFile(full_config_directory_path, config_to_write, 0755)
	}
}
