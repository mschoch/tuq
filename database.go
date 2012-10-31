package main

import (
	"encoding/json"
	"fmt"
	"github.com/dustin/gomemcached"
	"log"

	// Alias this because we call our connection couchbase
	cb "github.com/couchbaselabs/go-couchbase"
)

var bucketCache = make(map[string]*cb.Bucket, 0)

func getCachedBucket(couchbaseBucket string) (*cb.Bucket, error) {
	bucket, ok := bucketCache[couchbaseBucket]
	if ok {
		return bucket, nil
	} else {
		buck, err := dbConnect(couchbaseBucket)
		if err != nil {
			return nil, err
		} else {
			bucketCache[couchbaseBucket] = buck
			bucket = buck
		}
	}
	return bucket, nil
}

func dbConnect(couchbaseBucket string) (*cb.Bucket, error) {

	//	cb.HttpClient = &http.Client{
	//		Transport: TimeoutTransport(*viewTimeout),
	//	}

	if *debugCouchbase {
		log.Printf("Connecting to couchbase bucket %v at %v",
			couchbaseBucket, *couchbaseServer)
	}
	rv, err := cb.GetBucket(*couchbaseServer, "default", couchbaseBucket)
	if err != nil {
		return nil, err
	}
	return rv, nil
}

func DocsFromMeta(bucketName string, docMeta []CouchbaseDocMeta) ([]interface{}, error) {

	//get a connection to the bucket
	bucket, err := getCachedBucket(bucketName)
	if err != nil {
		return nil, err
	}

	result := make([]interface{}, 0)
	for i := 0; i < len(docMeta); i += *couchbaseBatchSize {
		batchEnd := (i + (*couchbaseBatchSize))
		if batchEnd > len(docMeta) {
			batchEnd = len(docMeta)
		}

		batch := docMeta[i:batchEnd]
		batchIds := make([]string, 0, len(batch))
		for _, meta := range batch {
			batchIds = append(batchIds, meta.Id)
		}

		bulkResponse := bucket.GetBulk(batchIds)

		if *debugCouchbase {
			log.Printf("Couchbase bulk response is: %#v", bulkResponse)
		}

		// response is a map, we need to walk through the original list
		// in order to preserve order of final result
		for _, meta := range batch {
			responseItem := bulkResponse[meta.Id]

			if responseItem == nil {
				return nil, fmt.Errorf("Couchbase does not contain the document %v", meta.Id)
			}

			wrappedDoc, err := WrappedDocFromMcResponse(meta, responseItem)
			if err != nil {
				return nil, err
			}
			result = append(result, wrappedDoc)
		}

	}

	return result, nil
}

func WrappedDocFromMcResponse(meta CouchbaseDocMeta, mcResponse *gomemcached.MCResponse) (interface{}, error) {
	var doc interface{}

	// marshall into json
	jsonErr := json.Unmarshal(mcResponse.Body, &doc)
	if jsonErr != nil {
		return nil, jsonErr
	}

	result := map[string]interface{}{
		"meta": meta,
		"doc":  doc}

	return result, nil
}
