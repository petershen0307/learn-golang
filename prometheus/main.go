package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/common/expfmt"
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func main() {
	// http.Handle("/metrics", promhttp.Handler())
	// http.ListenAndServe(":2112", nil)
	recordMetrics()
	time.Sleep(5 * time.Second)
	printMetrics()
}

func printMetrics() {
	buffer := bytes.NewBuffer([]byte{})
	enc := expfmt.NewEncoder(buffer, expfmt.NewFormat(expfmt.TypeTextPlain))
	mfs, err := prometheus.DefaultGatherer.Gather()
	if err != nil {
		panic(err)
	}
	for _, mf := range mfs {
		if err := enc.Encode(mf); err != nil {
			fmt.Printf("error=%v\n", err)
			break
		}
	}
	fmt.Println(buffer.String())
}
