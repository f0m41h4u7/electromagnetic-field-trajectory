package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-echarts/go-echarts/charts"
)

var rangeColor = []string{
	"#313695", "#4575b4", "#74add1", "#abd9e9", "#e0f3f8",
	"#fee090", "#fdae61", "#f46d43", "#d73027", "#a50026",
}

func MinMax(minmax string) float32 {
	var max float64 = z[0]
	var min float64 = z[0]
	for _, value := range z {
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

func getData() [][3]float64 {
	data := make([][3]float64, 0)
	for i := 0; i < int(N); i++ {
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

func getData2D() [][2]float64 {
	calculate()
	data := make([][2]float64, 0)
	for i := 0; i < int(N); i++ {
		data = append(data,
			[2]float64{
				x[i],
				y[i],
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
			Min:        MinMax("min"),
			Max:        MinMax("max"),
			InRange:    charts.VMInRange{Color: rangeColor},
		},
	)
	line3d.AddZAxis("", getData())
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
