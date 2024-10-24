package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	dto "github.com/prometheus/client_model/go"
)

var (
	myRegister    *prometheus.Registry
	txtFileGauge  prometheus.Gauge
	docxFileGauge prometheus.Gauge
)

func init() {
	myRegister = prometheus.NewRegistry()
	txtFileGauge = promauto.With(myRegister).NewGauge(prometheus.GaugeOpts{
		Name: "supported_file_type",
		ConstLabels: prometheus.Labels{
			"file_ext": "txt",
		},
	})
	docxFileGauge = promauto.With(myRegister).NewGauge(prometheus.GaugeOpts{
		Name: "supported_file_type",
		ConstLabels: prometheus.Labels{
			"file_ext": "docx",
		},
	})
}

func addRandomNumToGauge() {
	go func() {
		for range []int{123} {
			txt := float64(rand.Intn(10))
			txtFileGauge.Add(txt)
			docx := float64(rand.Intn(10))
			docxFileGauge.Add(docx)
			fmt.Println(txt, docx)
		}
	}()
}

func runSample() {
	addRandomNumToGauge()
	time.Sleep(1 * time.Second)
	dumpSample()
}

func dumpSample() {
	txtm, _ := getMetricValue(txtFileGauge)
	txtl := txtm.GetLabel()
	for _, l := range txtl {
		fmt.Println("name:", *(l.Name), "value:", *(l.Value))
	}
	fmt.Println("txt", txtFileGauge.Desc().String(), txtm.GetGauge().GetValue())
	docxm, _ := getMetricValue(docxFileGauge)
	docxl := docxm.GetLabel()
	for _, l := range docxl {
		fmt.Println("name:", *(l.Name), "value:", *(l.Value))
	}
	fmt.Println("docx", docxFileGauge.Desc().String(), docxm.GetGauge().GetValue())

}

func getMetricValue(c prometheus.Metric) (*dto.Metric, error) {
	metric := &dto.Metric{}
	err := c.Write(metric)
	return metric, err
}
