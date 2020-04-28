package main

import (
	"net/http"
	"os"

	"github.com/go-echarts/go-echarts/charts"
)

// Render plots

func EMPlotBase() *charts.Line3D {
	line3d := charts.NewLine3D()
	line3d.SetGlobalOptions(
		charts.TitleOpts{Title: "Electromagnetic field"},
		charts.VisualMapOpts{
			Calculable: true,
			Min:        MinMax(em_data, "min"),
			Max:        MinMax(em_data, "max"),
			InRange:    charts.VMInRange{Color: rangeColor},
		},
	)
	line3d.AddZAxis("", em_data)
	return line3d
}

func ElectricPlotBase() *charts.Line3D {
	line3d := charts.NewLine3D()
	line3d.SetGlobalOptions(
		charts.TitleOpts{Title: "Electric field"},
		charts.VisualMapOpts{
			Calculable: true,
			Min:        MinMax(e_data, "min"),
			Max:        MinMax(e_data, "max"),
			InRange:    charts.VMInRange{Color: rangeColor},
		},
	)
	line3d.AddZAxis("", e_data)
	return line3d
}

func MagneticPlotBase() *charts.Line3D {
	line3d := charts.NewLine3D()
	line3d.SetGlobalOptions(
		charts.TitleOpts{Title: "Magnetic field"},
		charts.VisualMapOpts{
			Calculable: true,
			Min:        MinMax(m_data, "min"),
			Max:        MinMax(m_data, "max"),
			InRange:    charts.VMInRange{Color: rangeColor},
		},
	)
	line3d.AddZAxis("", m_data)
	return line3d
}

// Handlers

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
