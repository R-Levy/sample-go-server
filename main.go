package main

import (
    "io"
    "log"
    "net/http"
    // "os"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
  "time"
  "math/rand"
)
var (
    counter = prometheus.NewCounter(
       prometheus.CounterOpts{
          Namespace: "golang",
          Name:      "my_counter",
          Help:      "This is my counter",
       })
  
    gauge = prometheus.NewGauge(
       prometheus.GaugeOpts{
          Namespace: "golang",
          Name:      "my_gauge",
          Help:      "This is my gauge",
       })

  histogram = prometheus.NewHistogram(
    prometheus.HistogramOpts{
       Namespace: "golang",
       Name:      "my_histogram",
       Help:      "This is my histogram",
    })

 summary = prometheus.NewSummary(
    prometheus.SummaryOpts{
       Namespace: "golang",
       Name:      "my_summary",
       Help:      "This is my summary",
    })
)


func main() {
    http.HandleFunc("/", ExampleHandler)

    // port := os.Getenv("PORT")
    // if port == "" {
    //     port = "8080"
    // }

    // log.Println("** Service Started on Port " + port + " **")
    // if err := http.ListenAndServe(":"+port, nil); err != nil {
    //     log.Fatal(err)
    // }

    rand.Seed(time.Now().Unix())

  http.Handle("/metrics", promhttp.Handler())

  prometheus.MustRegister(counter)
  prometheus.MustRegister(gauge)
  prometheus.MustRegister(histogram)
  prometheus.MustRegister(summary)

  go func() {
    for {
       counter.Add(rand.Float64() * 5)
       gauge.Add(rand.Float64()*15 - 5)
       histogram.Observe(rand.Float64() * 10)
       summary.Observe(rand.Float64() * 10)

       time.Sleep(time.Second)
    }
 }()

 log.Fatal(http.ListenAndServe(":8080", nil))
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json")
    io.WriteString(w, `{"status":"ok"}`)
}
