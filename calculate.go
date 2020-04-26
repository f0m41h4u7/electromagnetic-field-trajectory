package main

import (
	"fmt"
	"math"
)

var (
	B     = 0.1
	q     = 1.6 * math.Pow10(-19)
	m     = 9.1 * math.Pow10(-31)
	dt    = math.Pow10(-5)
	omega = q * B / m
	v0    = math.Pow10(7)
	r0    = m * v0 * q * B
	t0    = 1 / omega
	alpha = 0.785

	// Coordinates
	x = make([]float64, 100)
	y = make([]float64, 100)
	z = make([]float64, 100)

	// Velocity components
	v_x = make([]float64, 100)
	v_y = make([]float64, 100)
	v_z = v0 * math.Cos(alpha)
)

func calculate() {
	x[0] = 1
	v_y[0] = 1
	i := 1
	for t := dt / t0; t < math.Pow10(-3)/t0-dt/t0; t += dt / t0 {
		v_x[i] = v_x[i-1] + v0*math.Sin(alpha)*dt
		v_y[i] = v_y[i-1] - v0*math.Sin(alpha)*dt

		x[i] = x[i-1] + v_x[i-1]*dt
		y[i] = y[i-1] + v_y[i-1]*dt
		z[i] = z[i-1] + v_z*dt

		i++
	}
	fmt.Printf("%v\n%v\n", v_x, v_y)
}
