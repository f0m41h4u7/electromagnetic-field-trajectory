package main

import "math"

var em_data = [][3]float64{} // x, y, z coordinates

func calculateElectromagnetic() {
	const N = 200

	phi := make([]float64, N)
	v_z := make([]float64, N)
	em_data = append(em_data, [3]float64{r, 0, 0})

	i := 1
	for t := dt; t < N*dt-dt; t += dt {
		phi[i] = phi[i-1] + omega*dt
		v_z[i] = v_z[i-1] + q*E*dt/m
		em_data = append(em_data,
			[3]float64{
				r * math.Cos(phi[i]),
				r * math.Sin(phi[i]),
				em_data[i-1][2] + v_z[i-1]*dt,
			},
		)
		i++
	}
}
