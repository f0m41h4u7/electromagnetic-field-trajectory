package calculation

import "math"

var Em_data = [][3]float64{} // x, y, z coordinates

func CalculateElectromagnetic() {
	const N = 200

	phi := make([]float64, N)
	v_z := make([]float64, N)
	Em_data = append(Em_data, [3]float64{r, 0, 0})

	i := 1
	for t := dt; t < N*dt-dt; t += dt {
		phi[i] = phi[i-1] + omega*dt
		v_z[i] = v_z[i-1] + q*E*dt/m
		Em_data = append(Em_data,
			[3]float64{
				r * math.Cos(phi[i]),
				r * math.Sin(phi[i]),
				Em_data[i-1][2] + v_z[i-1]*dt,
			},
		)
		i++
	}
}
