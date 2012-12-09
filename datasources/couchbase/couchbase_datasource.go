package couchbase

import (
	"fmt"
	"github.com/mschoch/tuq/datasources"
	"github.com/mschoch/tuq/parser"
	"github.com/mschoch/tuq/planner"
	"log"
	// Alias this because we call our connection couchbase
	"encoding/json"
	cb "github.com/couchbaselabs/go-couchbase"
	"github.com/dustin/gomemcached"
)

const defaultCouchbaseBatchSize = 10000
const defaultDebugCouchbase = false

// FIXME this global needs to be reconsidered
// if we have 2 buckets with same name on different servers
// it will break 
var bucketCache = make(map[string]*cb.Bucket, 0)

type CouchbaseDataSource struct {
	Name            string
	As              string
	OutputChannel   planner.DocumentChannel
	cancelled       bool
	bucketName      string
	couchbaseServer string
	batchSize       int
	debug           bool
	ddoc            string
	view            string
	startkey        interface{}
	endkey          interface{}
}

func init() {
	datasources.RegisterDataSourceImpl("couchbase", NewCouchbaseDataSource)
}

func NewCouchbaseDataSource(config map[string]interface{}) planner.DataSource {
	result := &CouchbaseDataSource{
		OutputChannel:   make(planner.DocumentChannel),
		cancelled:       false,
		bucketName:      config["bucket"].(string),
		couchbaseServer: config["couchbase"].(string),
		ddoc:            "",
		view:            "_all_docs",
		startkey:        []interface{}{},
		endkey:          []interface{}{map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}, map[string]interface{}{}}}

	if config["batch_size"] != nil {
		result.batchSize = config["batch_size"].(int)
	} else {
		result.batchSize = defaultCouchbaseBatchSize
	}

	if config["debug"] != nil {
		result.debug = config["debug"].(bool)
	} else {
		result.debug = defaultDebugCouchbase
	}

	return result
}

// PlanPipelineComponent interface

func (ds *CouchbaseDataSource) SetSource(source planner.PlanPipelineComponent) {
	log.Fatalf("SetSource called on DataSource")
}

func (ds *CouchbaseDataSource) GetSource() planner.PlanPipelineComponent {
	return nil
}

func (ds *CouchbaseDataSource) GetDocumentChannel() planner.DocumentChannel {
	return ds.OutputChannel
}

func (ds *CouchbaseDataSource) Run() {
	defer close(ds.OutputChannel)

	//get a connection to the bucket
	bucket, err := ds.getCachedBucket(ds.bucketName)
	if err != nil {
		log.Printf("Error getting bucket: %v", err)
		return
	}

	skey := ds.startkey

	// FIXME is include_docs=false and mget faster than include docs=true?
	vres, err := bucket.View(ds.ddoc, ds.view, map[string]interface{}{
		"limit":        ds.batchSize + 1,
		"include_docs": true,
		"startkey":     skey,
		"endkey":       ds.endkey})

	if err != nil {
		log.Printf("Error accessing view: %v", err)
		return
	}

	for i, row := range vres.Rows {
		if i < ds.batchSize {
			// dont process the last row, its just used to see if we
			// need to continue processing
			rowdoc := (*row.Doc).(map[string]interface{})
			rowdoc["doc"] = rowdoc["json"]
			delete(rowdoc, "json")

			rowdoc = datasources.WrapDocWithDatasourceAs(ds.As, rowdoc)

			ds.OutputChannel <- rowdoc
		}
	}

	// as long as we continue to get batchSize + 1 results back we have to keep going
	for len(vres.Rows) > ds.batchSize {
		skey = vres.Rows[ds.batchSize].Key
		skeydocid := vres.Rows[ds.batchSize].ID

		vres, err = bucket.View(ds.ddoc, ds.view, map[string]interface{}{
			"limit":          ds.batchSize + 1,
			"include_docs":   true,
			"startkey":       skey,
			"startkey_docid": skeydocid,
			"endkey":         ds.endkey})

		if err != nil {
			log.Printf("Error accessing view: %v", err)
			return
		}

		for i, row := range vres.Rows {
			if i < ds.batchSize {
				// dont process the last row, its just used to see if we
				// need to continue processing
				rowdoc := (*row.Doc).(map[string]interface{})
				rowdoc["doc"] = rowdoc["json"]
				delete(rowdoc, "json")

				rowdoc = datasources.WrapDocWithDatasourceAs(ds.As, rowdoc)

				ds.OutputChannel <- rowdoc
			}
		}

	}

}

func (ds *CouchbaseDataSource) Explain() {
	defer close(ds.OutputChannel)

	thisStep := map[string]interface{}{
		"_type":    "FROM",
		"impl":     "Couchbase",
		"filename": ds.Name,
		"as":       ds.As}

	ds.OutputChannel <- thisStep
}

func (ds *CouchbaseDataSource) SetName(name string) {
	ds.Name = name
}

func (ds *CouchbaseDataSource) SetAs(as string) {
	ds.As = as
}

func (ds *CouchbaseDataSource) GetAs() string {
	return ds.As
}

func (ds *CouchbaseDataSource) SetFilter(filter parser.Expression) error {
	return fmt.Errorf("Couchbase DataSource does not support filter")
}

func (ds *CouchbaseDataSource) GetFilter() parser.Expression {
	return nil
}

func (ds *CouchbaseDataSource) SetOrderBy(sortlist parser.SortList) error {
	return fmt.Errorf("Couchbase DataSource does not support order by")
}

func (ds *CouchbaseDataSource) GetOrderBy() parser.SortList {
	return nil
}

func (ds *CouchbaseDataSource) SetLimit(e parser.Expression) error {
	return fmt.Errorf("Couchbase DataSource does not support limit")
}

func (ds *CouchbaseDataSource) SetOffset(e parser.Expression) error {
	return fmt.Errorf("Couchbase DataSource does not support offset")
}

func (ds *CouchbaseDataSource) SetGroupByWithStatsFields(groupby parser.ExpressionList, stats_fields []string) error {
	return fmt.Errorf("Couchbase DataSource does not support group by")
}

func (ds *CouchbaseDataSource) SetHaving(having parser.Expression) error {
	return fmt.Errorf("Couchbase DataSource does not support having")
}

func (ds *CouchbaseDataSource) Cancel() {
	ds.cancelled = true
}

func (ds *CouchbaseDataSource) DocsFromIds(docIds []string) ([]planner.Document, error) {

	//get a connection to the bucket
	bucket, err := ds.getCachedBucket(ds.bucketName)
	if err != nil {
		return nil, err
	}

	result := make([]planner.Document, 0)
	for i := 0; i < len(docIds); i += ds.batchSize {
		batchEnd := (i + (ds.batchSize))
		if batchEnd > len(docIds) {
			batchEnd = len(docIds)
		}

		batchIds := docIds[i:batchEnd]
		//batchIds := make([]string, 0, len(batch))

		bulkResponse := bucket.GetBulk(batchIds)

		if ds.debug {
			log.Printf("Couchbase bulk response is: %#v", bulkResponse)
		}

		// response is a map, we need to walk through the original list
		// in order to preserve order of final result
		for _, id := range batchIds {
			responseItem := bulkResponse[id]

			if responseItem == nil {
				return nil, fmt.Errorf("Couchbase does not contain the document %v", id)
			}

			doc, err := DocFromMcResponse(responseItem)
			if err != nil {
				return nil, err
			}
			result = append(result, doc)
		}

	}

	return result, nil
}

func (ds *CouchbaseDataSource) getCachedBucket(couchbaseBucket string) (*cb.Bucket, error) {
	bucket, ok := bucketCache[couchbaseBucket]
	if ok {
		return bucket, nil
	} else {
		buck, err := ds.dbConnect(couchbaseBucket)
		if err != nil {
			return nil, err
		} else {
			bucketCache[couchbaseBucket] = buck
			bucket = buck
		}
	}
	return bucket, nil
}

func (ds *CouchbaseDataSource) dbConnect(couchbaseBucket string) (*cb.Bucket, error) {

	if ds.debug {
		log.Printf("Connecting to couchbase bucket %v at %v",
			couchbaseBucket, ds.couchbaseServer)
	}
	rv, err := cb.GetBucket(ds.couchbaseServer, "default", couchbaseBucket)
	if err != nil {
		return nil, err
	}
	return rv, nil
}

func DocFromMcResponse(mcResponse *gomemcached.MCResponse) (planner.Document, error) {
	var doc planner.Document

	// marshall into json
	jsonErr := json.Unmarshal(mcResponse.Body, &doc)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return doc, nil
}
