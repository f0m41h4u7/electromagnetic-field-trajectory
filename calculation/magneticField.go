package calculation

import "math"

var M_data = [][3]float64{} // x, y, z coordinates

func CalculateMagnetic() {
	const N = 200

	phi := make([]float64, N)
	M_data = append(M_data, [3]float64{r, 0, 0})

	i := 1
	for t := dt; t < N*dt-dt; t += dt {
		phi[i] = phi[i-1] + omega*dt
		M_data = append(M_data,
			[3]float64{
				r * math.Cos(phi[i]),
				r * math.Sin(phi[i]),
				M_data[i-1][2] + v*dt,
			},
		)
		i++
	}
}
