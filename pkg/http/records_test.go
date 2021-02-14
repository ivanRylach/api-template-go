package http

import (
	"api.ivanrylach.github.io/v1/pkg/mongodb"
	"api.ivanrylach.github.io/v1/pkg/records"
	"api.ivanrylach.github.io/v1/pkg/util"
	"bytes"
	"encoding/json"
	"github.com/benweissmann/memongo"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecords_Fetch(t *testing.T) {
	util.ConfigureLogging()

	mongoServer := startMongo(t)
	t.Cleanup(func() {
		mongoServer.Stop()
	})

	mongo := mongodb.NewClient(mongoServer.URI())
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

func startMongo(t *testing.T) *memongo.Server {
	mongoServer, err := memongo.Start("4.2.0")
	if err != nil {
		zap.S().Error(err)
		t.FailNow()
	}
	return mongoServer
}
