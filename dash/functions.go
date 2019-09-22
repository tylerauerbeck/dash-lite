package dash

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"

	v3 "github.com/tylerauerbeck/dash/dash/v3"
	yaml "gopkg.in/yaml.v2"
)

// GetInventoryConfig finds the Dash.yml file within a directory path that is provided
func GetInventoryConfig(envDir string) string {
	var dashFile string

	files, err := ioutil.ReadDir(envDir)
	if err != nil {
		log.Fatalln("Unable to read files in " + envDir)
		os.Exit(1)
	}

	for _, file := range files {
		matched, _ := regexp.MatchString(`dash.(yml|yaml)`, file.Name())
		if matched {
			dashFile = file.Name()
			break
		}
	}

	if dashFile == "" {
		log.Println("Unable to find dash.yml inventory. Defaulting to 2.X")
	}

	return dashFile
}

// GetInventoryContent blah
func GetInventoryContent(envDir string, dashFile string) Inventory {
	var data, _ = ioutil.ReadFile(envDir + "/" + dashFile)
	t := Inventory{}

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return t
}

// ExecV3 processes a V3 Dash Inventory
func ExecV3(DashInventory Inventory) {
	v3.RunV3(DashInventory.Namespace, DashInventory.Context, DashInventory.Resources)
}

// ValidateLogin should verify that we have an active/valid session with a cluster
func ValidateLogin() {
 // TODO: We should make sure that we're logged in and that a context is good before we do anything. Otherwise we'll get unauthorized, etc. errorss.
}
