package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Task struct {
	Name    string
	Command string
}

type Config struct {
	Version string
	Tasks   []Task
}

func main() {
	file, err := ioutil.ReadFile("./task.omit")

	if err != nil {
		fmt.Fprintln(os.Stderr, "Can not open file: ", err)
	}

	config := Config{}
	err = yaml.Unmarshal(file, &config)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Faild parse yaml file: ", err)
	}

	for _, task := range config.Tasks {
		fmt.Println(task.Command)
	}

}
