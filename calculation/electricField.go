package calculation

var E_data = [][3]float64{} // x, y, z coordinates

func CalculateElectric() {
	const N = 1000

	v_z := make([]float64, N)
	E_data = append(E_data, [3]float64{0, 0, 1})

	i := 1
	for t := dt; t < N*dt-dt; t += dt {
		v_z[i] = v_z[i-1] + q*E*dt/m
		E_data = append(E_data,
			[3]float64{
				E_data[i-1][0] + v*dt,
				0,
				E_data[i-1][2] + v_z[i-1]*dt,
			},
		)
		i++
	}
}
