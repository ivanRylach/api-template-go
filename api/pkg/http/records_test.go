package http

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"ivanrylach.github.io/api/v1/pkg/records"
	"ivanrylach.github.io/mongodb"
	"ivanrylach.github.io/util"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecords_Fetch(t *testing.T) {
	util.ConfigureLogging()

	// TODO: MongoDB Transactions require a sharded deployment. Does 'memongo' support it?
	//mongoServer := startMongo(t)
	//t.Cleanup(func() {
	//	mongoServer.Stop()
	//})

	mongo := mongodb.NewClient("mongodb://root:password123@localhost:27017")
	recordsRepo := records.NewRepository(mongo)
	router := NewRouter(recordsRepo)

	w := httptest.NewRecorder()

	postBody, _ := json.Marshal(map[string]string{
		"name":        "one",
		"description": "The 1st one",
	})
	body := bytes.NewBuffer(postBody)

	req, _ := http.NewRequest("POST", "/v1/records", body)
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)

	var view records.RecordDTO
	json.Unmarshal(w.Body.Bytes(), &view)
	assert.Equal(t, "one", view.Name)

	w = httptest.NewRecorder()

	req, _ = http.NewRequest("GET", "/v1/record/"+view.Id, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	json.Unmarshal(w.Body.Bytes(), &view)
	assert.Equal(t, "one", view.Name)
}

//func startMongo(t *testing.T) *memongo.Server {
//	mongoServer, err := memongo.Start("4.2.0")
//	if err != nil {
//		zap.S().Error(err)
//		t.FailNow()
//	}
//	return mongoServer
//}
