package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var config_directory_path = fmt.Sprintf("%s/.kcm-test", get_user_home_direcotry())
var full_config_directory_path = fmt.Sprintf("%s/.kcm-test/config.json", get_user_home_direcotry())

func main() {
	if !config_directory_present() {
		create_config_directory()
	}

	create_config_file_if_needed()
	config := read_config_file()

	switch command := os.Args[1]; command {
	case "list":
		list_configs(config)
		fmt.Println("list")
	case "status":
		fmt.Println("status")
	case "create":
		name := os.Args[2]
		path := os.Args[3]
		active := "false"
		configuration := add_config(config, name, path, active)
		write_config(configuration)
		fmt.Println("create")
	case "remove":
		//name := os.Args[2]
		fmt.Println("remove")
	case "activate":
		//name := os.Args[2]
		fmt.Println("activate")
	default:
		fmt.Println("KCM - Kubectl Configuration Manager")
		fmt.Println(" ")
		fmt.Println("Unknown command provided: ", command)
		fmt.Println(" ")
		fmt.Println("Help info should go here?")
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
		return true
	} else {
		return false
	}
}

func config_file_present() bool {
	if _, err := os.Stat(full_config_directory_path); err == nil {
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

		ioutil.WriteFile(full_config_directory_path, config_to_write, 0644)
	}
}

func read_config_file() config_container {
	var config_obj config_container

	jsonFile, err := os.Open(full_config_directory_path)
	check(err)

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &config_obj)

	return config_obj
}

func write_config(configuration config_container) {
	config_to_write, _ := json.MarshalIndent(configuration, "", " ")
	ioutil.WriteFile(full_config_directory_path, config_to_write, 0644)
}

func add_config(configuration config_container, name string, path string, active string) config_container {
	new_config := config{Name: name, Path: path, Active: active}
	existing_configs := configuration.Configs
	fmt.Println(existing_configs)
	configuration.Configs = append(existing_configs, new_config)
	fmt.Println(configuration)
	return configuration
}

func list_configs(configuration config_container) {
	for i := 0; i < len(configuration.Configs); i++ {
		fmt.Println(configuration.Configs[i].Name + " - " + configuration.Configs[i].Path)
	}
}
