package v3

import "log"

// RunV3 blah
func RunV3(Namespace string, Context string, RG []ResourceGroup) {
	// Namespace -> group.Namespace -> item.Namespace
	// NS = "", then NS gets set to group.Namespace or item.Namespace if they differ, get reset at the bottom of the loop
	log.Println("The global namepspace has been set to " + Namespace)
	log.Println("The global context has been set to " + Context)

	for _, group := range RG {
		NS := Namespace
		CTX := Context

		log.Println("Running items for resource group " + group.Name)
		NS = DetermineDefaults(NS, group.Namespace)
		CTX = DetermineDefaults(CTX, group.Context)

		log.Println("Default Namespace for " + group.Name + " is: " + NS)
		log.Println("Default Context for " + group.Name + "is: " + CTX)

		for _, item := range group.Resources {
			object := ""
			NS = DetermineDefaults(NS, item.Namespace)
			CTX = DetermineDefaults(CTX, item.Context)

			if (item.Template != "" && item.File != "") || (item.Template == "" && item.File == "") {
				log.Fatalln("Unable to determine type of object to use. Skipping..")
				break
			} else if item.Template != "" {
				object = "template"
			} else if item.File != "" {
				object = "file"
			}

			log.Println("Now creating " + item.Name)
			log.Println("Creating in " + NS)
			log.Println("Creating using " + CTX)

			switch object {
			case "template":
				log.Println("Processing template: " + item.Template)
				log.Println("Template is of type: " + item.Type)
				log.Println("Template will use params from " + item.Params)
			case "file":
				log.Println("Processing file: " + item.File)
			}
		}
	}
}

// MergeInventories blah
//func MergeInventories() {
// TODO: Write some code here that will take multiple inventories
// and merge them appropriately (according to priority, etc. )

//}

// DetermineDefaults blah
func DetermineDefaults(defaultVar string, subVar string) string {
	var def string

	if subVar != defaultVar && subVar != "" {
		def = subVar
	} else {
		def = defaultVar
	}
	return def
}
