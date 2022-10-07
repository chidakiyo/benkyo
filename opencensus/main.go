package main

import (
	"log"
	"net/http"
	"os"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"contrib.go.opencensus.io/exporter/stackdriver/propagation"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
)

// URL: https://cloud.google.com/trace/docs/setup/go?hl=ja#sample_application_for_go

func main() {
	// Create and register a OpenCensus Stackdriver Trace exporter.
	exporter, err := stackdriver.NewExporter(stackdriver.Options{
		//ProjectID: os.Getenv("GOOGLE_CLOUD_PROJECT"), // GCP 環境で実行する際には指定しなくても良い
	})
	if err != nil {
		log.Fatal(err)
	}
	trace.RegisterExporter(exporter)

	// By default, traces will be sampled relatively rarely. To change the
	// sampling frequency for your entire program, call ApplyConfig. Use a
	// ProbabilitySampler to sample a subset of traces, or use AlwaysSample to
	// collect a trace on every run.
	//
	// Be careful about using trace.AlwaysSample in a production application
	// with significant traffic: a new trace will be started and exported for
	// every request.
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	client := &http.Client{
		Transport: &ochttp.Transport{
			// Use Google Cloud propagation format.
			Propagation: &propagation.HTTPFormat{},
		},
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req, _ := http.NewRequest("GET", "https://www.google.com", nil)

		// The trace ID from the incoming request will be
		// propagated to the outgoing request.
		req = req.WithContext(r.Context())

		// The outgoing request will be traced with r's trace ID.
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		// Because we don't read the resp.Body, need to manually call Close().
		resp.Body.Close()
	})
	http.Handle("/foo", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)

	// Use an ochttp.Handler in order to instrument OpenCensus for incoming
	// requests.
	httpHandler := &ochttp.Handler{
		// Use the Google Cloud propagation format.
		Propagation: &propagation.HTTPFormat{},
	}
	if err := http.ListenAndServe(":"+port, httpHandler); err != nil {
		log.Fatal(err)
	}
}
