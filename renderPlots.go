package main

import "github.com/go-echarts/go-echarts/charts"

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
