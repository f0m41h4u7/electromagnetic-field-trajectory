package main

import (
	"log"
	"net/http"
	"path"

	"github.com/go-echarts/go-echarts/charts"
)

const (
	host   = "http://0.0.0.0:7272"
	maxNum = 50
)

type router struct {
	name string
	charts.RouterOpts
}

var routers = []router{
	{"plot", charts.RouterOpts{URL: host + "/plot", Text: "Plot"}},
}

func orderRouters(chartType string) []charts.RouterOpts {
	for i := 0; i < len(routers); i++ {
		if routers[i].name == chartType {
			routers[i], routers[0] = routers[0], routers[i]
			break
		}
	}

	rs := make([]charts.RouterOpts, 0)
	for i := 0; i < len(routers); i++ {
		rs = append(rs, routers[i].RouterOpts)
	}
	return rs
}

func getRenderPath(f string) string {
	return path.Join("html", f)
}

func logTracing(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Tracing request for %s\n", r.RequestURI)
		next.ServeHTTP(w, r)
	}
}

func main() {
	calculate()
	http.HandleFunc("/plot", logTracing(plotHandler))

	log.Println("Run server at " + host)
	http.ListenAndServe(":7272", nil)
}
