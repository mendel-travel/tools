package env

import (
	"flag"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// env vars
var e string

func createFlagsForBuild() {
	flag.StringVar(&e, "env", "development", "Specify environment")
	flag.Parse()
}

func LoadEnvProperties(f string, p interface{}) {

	createFlagsForBuild()

	file, err := ioutil.ReadFile(fmt.Sprintf("%s/%s.yaml", f, e))

	if err != nil {
		panic("I can't find the env properties file. Please put the file and compile the project again!")
	} else {
		if err := yaml.Unmarshal(file, &p); err != nil {
			panic("There is an error on .yaml properties. Please check the file and try again!")
		}
	}

}
