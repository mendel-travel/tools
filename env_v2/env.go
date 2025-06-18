package env

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// LoadProperties : function that receives a forlder name where the config files are located and a interface to map the file.
// We decide to use a struct to manage de properties because is easy to access values after in the code.
func LoadProperties(v string, f string, p interface{}) {
	env := os.Getenv(v)
	if env == "" {
		env = "develop"
	}

	file, err := ioutil.ReadFile(fmt.Sprintf("%s/%s/properties.yaml", f, env))
	if err != nil {
		panic("I can't find the env properties file. Please put the file and compile the project again!")
	}

	if err := yaml.Unmarshal(file, p); err != nil {
		panic("There is an error on .yaml properties. Please check the file and try again!")
	}
}
