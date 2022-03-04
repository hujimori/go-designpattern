package main

import (
	condition "v0/condtion"
	"v0/weather"
)

func main() {
	w := weather.New()
	c := condition.New(w)
	c.GetId()

	w.SetMeasurements(27, 65, 24.4)
}
