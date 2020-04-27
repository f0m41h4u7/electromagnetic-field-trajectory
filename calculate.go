package main

import (
	"math"
)

var N float64 // number of dots

var (
	B     = 0.01                  // induction, in teslas
	q     = 1.6 * math.Pow10(-19) // charge, in coulombs
	m     = 9.1 * math.Pow10(-31) // mass, in kg
	dt    = math.Pow10(-10)       // time step, in s
	omega = q * B / m             // Larmor fequency, in herz
	v     = math.Pow10(7)         // initial velocity, in m/s
	r     = m * v / q / B         // Larmor radius, in m

	x   = []float64{}
	y   = []float64{}
	z   = []float64{}
	phi = []float64{}
)

func calculate() {
	x = append(x, r)
	y = append(y, 0)
	z = append(z, 0)
	phi = append(phi, 0)

	i := 1
	for t := dt; t < N*dt-dt; t += dt {
		phi = append(phi, phi[i-1]+omega*dt)

		x = append(x, r*math.Cos(phi[i]))
		y = append(y, r*math.Sin(phi[i]))
		z = append(z, z[i-1]+v*dt)
		i++
	}
	N = float64(len(x))
}
