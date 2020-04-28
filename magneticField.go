package main

import "math"

var m_data = [][3]float64{} // x, y, z coordinates

func calculateMagnetic() {
	const N = 200

	phi := make([]float64, N)
	m_data = append(m_data, [3]float64{r, 0, 0})

	i := 1
	for t := dt; t < N*dt-dt; t += dt {
		phi[i] = phi[i-1] + omega*dt
		m_data = append(m_data,
			[3]float64{
				r * math.Cos(phi[i]),
				r * math.Sin(phi[i]),
				m_data[i-1][2] + v*dt,
			},
		)
		i++
	}
}
