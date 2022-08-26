package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestHandler(t *testing.T) {
	require := require.New(t)

	testCases := map[string]struct{
		url   string
		req    *Data
		hndlr  func(http.ResponseWriter, *http.Request)
		status int
	}{
		"min handler": {
			url:    "/min",
			req:    &Data{Qualifier: 3, Nums: []float64{9,8,7,6,5,4,3,2,1}},
			hndlr:  minHandler,
			status: http.StatusOK,
		},
		"max handler": {
			url:    "/max",
			req:    &Data{Qualifier: 3, Nums: []float64{9,8,7,6,5,4,3,2,1}},
			hndlr:  maxHandler,
			status: http.StatusOK,
		},
		"avg handler": {
			url:    "/avg",
			req:    &Data{Nums: []float64{9,8,7,6,5,4,3,2,1}},
			hndlr:  avgHandler,
			status: http.StatusOK,
		},
		"median handler": {
			url:    "/median",
			req:    &Data{Nums: []float64{9,8,7,6,5,4,3,2,1}},
			hndlr:  medianHandler,
			status: http.StatusOK,
		},
		"percentile handler": {
			url:    "/percentile",
			req:    &Data{Qualifier: 80, Nums: []float64{9,8,7,6,5,4,3,2,1}},
			hndlr:  percentileHandler,
			status: http.StatusOK,
		},
	}

	for tn, tc := range testCases {
		t.Run(tn, func(t *testing.T) {
			rr := httptest.NewRecorder()
			erj, err := json.Marshal(tc.req)
			require.NoError(err, "should marshal the payload")

			req, err := http.NewRequest(http.MethodPost, tc.url, bytes.NewBuffer(erj))
			require.NoError(err, "should setup the new request")

			router := mux.NewRouter()
			router.HandleFunc(tc.url, tc.hndlr)
			router.ServeHTTP(rr, req)

			require.Equal(tc.status, rr.Code, "Create response status code is expected")

		})
	}
}
