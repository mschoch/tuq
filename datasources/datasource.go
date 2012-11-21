package datasources

import (
	"encoding/json"
	"github.com/mschoch/tuq/planner"
	"io/ioutil"
	"log"
	"os/user"
)

// global variable containing definition of all data sources
var dataSources = map[string]interface{}{}
var dataSourceImpls = map[string]func(map[string]interface{}) planner.DataSource{}

func LoadDataSources() {
	currentUser, err := user.Current()
	if err != nil {
		log.Printf("Unable to determine home directory, no datasources will be loaded")
		return
	}
	dataSourceBytes, err := ioutil.ReadFile(currentUser.HomeDir + "/.tuq_datasources")
	if err != nil {
		log.Printf("Error loading data sources %v", err)
		return
	}
	err = json.Unmarshal(dataSourceBytes, &dataSources)
	if err != nil {
		log.Printf("Error loading data sources %v", err)
		return
	}
}

func NewDataSourceWithName(name string) planner.DataSource {
	dsDef, ok := dataSources[name].(map[string]interface{})
	if !ok {
		return nil
	}
	var ds planner.DataSource
	typ := dsDef["type"].(string)
	f := dataSourceImpls[typ]
	if f != nil {
		ds = f(dsDef)
	} else {
		log.Fatalf("Unsupported datasource type %v", typ)
	}

	return ds
}

func RegisterDataSourceImpl(t string, f func(map[string]interface{}) planner.DataSource) {
	dataSourceImpls[t] = f
}
