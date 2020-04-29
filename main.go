package main

import (
	"log"
	"net/http"
	"os"
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

func electromagneticPlotHandler(w http.ResponseWriter, _ *http.Request) {
	page := charts.NewPage(orderRouters("electromagneticFieldPlot")...)
	page.Add(
		EMPlotBase(),
	)
	f, _ := os.Create(getRenderPath("electromagneticPlot.html"))
	page.Render(w, f)
}

func electricPlotHandler(w http.ResponseWriter, _ *http.Request) {
	page := charts.NewPage(orderRouters("electricFieldPlot")...)
	page.Add(
		ElectricPlotBase(),
	)
	f, _ := os.Create(getRenderPath("electricPlot.html"))
	page.Render(w, f)
}

func magneticPlotHandler(w http.ResponseWriter, _ *http.Request) {
	page := charts.NewPage(orderRouters("magneticFieldPlot")...)
	page.Add(
		MagneticPlotBase(),
	)
	f, _ := os.Create(getRenderPath("magneticPlot.html"))
	page.Render(w, f)
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
