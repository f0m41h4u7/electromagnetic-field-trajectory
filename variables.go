package main

import "math"

var (
	E     = 2300.0                // electric field strength, in v/m
	B     = 0.01                  // induction, in teslas
	q     = 1.6 * math.Pow10(-19) // charge, in coulombs
	m     = 9.1 * math.Pow10(-31) // mass, in kg
	dt    = math.Pow10(-10)       // time step, in s
	v     = math.Pow10(5)         // initial velocity, in m/s
	omega = q * B / m             // Larmor fequency, in herz
	r     = m * v / q / B         // Larmor radius, in m
)
