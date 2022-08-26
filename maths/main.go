package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"

	"sojourn/maths/maths"
)

// port
const port = 8338

type Data struct {
	Nums      []float64  `json:"nums"`
	Qualifier int        `json:"qualifier,omitempty"`
}

type Response struct {
	Answer  float64   `json:"answer,omitempty"`
	Answers []float64 `json:"answers,omitempty"`
}

func main() {
	// register monitor handlers
	mux := registerHandlers()

	// http server
	srv := &http.Server{Addr: ":8338", Handler: mux}

	fmt.Printf("Maths server running: %s\n", srv.Addr)
	var err error
	err = srv.ListenAndServe()
	fmt.Println(err)
}

// parseRequest - get data from the request
func parseRequest(r *http.Request) (*Data, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	req := &Data{}
	err = json.Unmarshal([]byte(body), req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// parseResponse - set data for the response
func parseResponse(answer *float64, answers []float64) ([]byte, error) {
	ans := &Response{Answers: answers}
	if answer != nil {
		ans.Answer = *answer
	}
	return json.Marshal(ans)
}

// minHandler - handle min request
func minHandler(w http.ResponseWriter, r *http.Request) {
	data, err := parseRequest(r)
	if err != nil {
		log.Printf("failed parsing request: %+v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}

	answers := maths.Min(data.Nums, data.Qualifier)
	resp, err := parseResponse(nil, answers)
	if err != nil {
		log.Printf("failed parsing response: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error setting response"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// maxHandler - handle max request
func maxHandler(w http.ResponseWriter, r *http.Request) {
	data, err := parseRequest(r)
	if err != nil {
		log.Printf("failed parsing request: %+v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}

	answers := maths.Max(data.Nums, data.Qualifier)
	resp, err := parseResponse(nil, answers)
	if err != nil {
		log.Printf("failed parsing response: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error setting response"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// avgHandler - handle avg request
func avgHandler(w http.ResponseWriter, r *http.Request) {
	data, err := parseRequest(r)
	if err != nil {
		log.Printf("failed parsing request: %+v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}

	answer := maths.Avg(data.Nums)
	resp, err := parseResponse(&answer, nil)
	if err != nil {
		log.Printf("failed parsing response: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error setting response"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// medianHandler - handle median request
func medianHandler(w http.ResponseWriter, r *http.Request) {
	data, err := parseRequest(r)
	if err != nil {
		log.Printf("failed parsing request: %+v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}

	answer := maths.Median(data.Nums)
	resp, err := parseResponse(&answer, nil)
	if err != nil {
		log.Printf("failed parsing response: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error setting response"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// percentileHandler - handle percentile request
func percentileHandler(w http.ResponseWriter, r *http.Request) {
	data, err := parseRequest(r)
	if err != nil {
		log.Printf("failed parsing request: %+v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}

	answer := maths.Percentile(data.Nums, data.Qualifier)
	resp, err := parseResponse(&answer, nil)
	if err != nil {
		log.Printf("failed parsing response: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error setting response"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// reqister the endpoint handlers
func registerHandlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/min", minHandler)
	r.HandleFunc("/max", maxHandler)
	r.HandleFunc("/avg", avgHandler)
	r.HandleFunc("/median", medianHandler)
	r.HandleFunc("/percentile", percentileHandler)

	return r
}
