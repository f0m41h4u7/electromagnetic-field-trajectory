package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-echarts/go-echarts/charts"
)

func genData() [][3]float64 {
	calculate()
	data := make([][3]float64, 0)
	for i := 0; i < 100; i++ {
		data = append(data,
			[3]float64{
				x[i],
				y[i],
				z[i],
			},
		)
	}
	return data
}

func PlotBase() *charts.Line3D {
	line3d := charts.NewLine3D()
	line3d.SetGlobalOptions(
		charts.TitleOpts{Title: "Electromagnetic field"},
		charts.VisualMapOpts{
			Calculable: true,
			InRange:    charts.VMInRange{Color: rangeColor},
			Max:        30,
		},
	)
	line3d.AddZAxis("", genData())
	return line3d
}

func plotHandler(w http.ResponseWriter, _ *http.Request) {
	page := charts.NewPage(orderRouters("plot")...)
	page.Add(
		PlotBase(),
	)
	f, err := os.Create(getRenderPath("plot.html"))
	if err != nil {
		log.Println(err)
	}
	page.Render(w, f)
}
