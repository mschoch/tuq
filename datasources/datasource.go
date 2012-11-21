package datasources

import (
	"encoding/json"
	"github.com/mschoch/tuq/planner"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"syscall"
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

		if e, ok := err.(*os.PathError); ok && (e.Err == syscall.ENOENT) {
			dataSourceBytes, err = json.MarshalIndent(defaultDataSources, "", "  ")
			if err != nil {
				log.Printf("unable to set up default datasources")
				return
			}
			log.Printf("NOTICE: No datasources were found, default datasources have been loaded (these require internet access).")

			err = ioutil.WriteFile(currentUser.HomeDir+"/.tuq_datasources", dataSourceBytes, 0600)
			if err != nil {
				log.Printf("WARNING: Unable to save datasources to %v, check permissions.", currentUser.HomeDir+"/.tuq_datasources")
			}

		} else {
			// some other error, just return
			log.Printf("Error loading data sources", err)
			return
		}

	}
	err = json.Unmarshal(dataSourceBytes, &dataSources)
	if err != nil {
		log.Printf("Error loading data sources2 %v", err)
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

var defaultDataSources = map[string]interface{}{
	"employees": map[string]interface{}{
		"type": "csv",
		"path": "http://raw.github.com/mschoch/tuq/master/datasources/csv/test_csv_datasources/employees.csv"},
	"departments": map[string]interface{}{
		"type": "csv",
		"path": "http://raw.github.com/mschoch/tuq/master/datasources/csv/test_csv_datasources/departments.csv"}}
