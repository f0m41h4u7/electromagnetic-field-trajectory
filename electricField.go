package main

var e_data = [][3]float64{} // x, y, z coordinates

func calculateElectric() {
	const N = 1000

	v_z := make([]float64, N)
	e_data = append(e_data, [3]float64{0, 0, 1})

	i := 1
	for t := dt; t < N*dt-dt; t += dt {
		v_z[i] = v_z[i-1] + q*E*dt/m
		e_data = append(e_data,
			[3]float64{
				e_data[i-1][0] + v*dt,
				0,
				e_data[i-1][2] + v_z[i-1]*dt,
			},
		)
		i++
	}
}
