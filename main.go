package main

import (
	"log"
	"net/http"
	"path"

	"github.com/go-echarts/go-echarts/charts"
)

const host = "http://0.0.0.0:7272"

type router struct {
	name string
	charts.RouterOpts
}

var routers = []router{
	{"magneticFieldPlot", charts.RouterOpts{URL: host + "/magnetic", Text: "Magnetic field"}},
	{"electricFieldPlot", charts.RouterOpts{URL: host + "/electric", Text: "Electric field"}},
	{"electromagneticFieldPlot", charts.RouterOpts{URL: host + "/electromagnetic", Text: "Electromagnetic field"}},
}

var rangeColor = []string{
	"#313695", "#4575b4", "#74add1", "#abd9e9", "#e0f3f8",
	"#fee090", "#fdae61", "#f46d43", "#d73027", "#a50026",
}

func MinMax(z [][3]float64, minmax string) float32 {
	var max float64 = z[0][2]
	var min float64 = z[0][2]

	for i := 0; i < len(z); i++ {
		value := z[i][2]
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}

	if minmax == "min" {
		return float32(min)
	}
	return float32(max)
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
	calculateMagnetic()
	calculateElectric()
	calculateElectromagnetic()

	http.HandleFunc("/magnetic", logTracing(magneticPlotHandler))
	http.HandleFunc("/electric", logTracing(electricPlotHandler))
	http.HandleFunc("/electromagnetic", logTracing(electromagneticPlotHandler))

	log.Println("Run server at " + host)
	http.ListenAndServe(":7272", nil)
}
