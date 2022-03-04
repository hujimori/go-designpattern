package condition

import (
	"fmt"
	"v0/observer"
)

type CurrentCondition struct {
	temperature float64
	humidity    float64
	weatherData *observer.Subject
	id          int64
}

func New(w *observer.Subject) *CurrentCondition {
	c := &CurrentCondition{weatherData: w, id: 1}
	(*w).RegisterObserver(c)
	return c
}

func (c *CurrentCondition) Update(temperature float64, humidity float64, pressure float64) {
	c.temperature = temperature
	c.humidity = humidity
	c.display()
}

func (c *CurrentCondition) display() {
	fmt.Printf("現在の気象状況：%f %f", c.temperature, c.humidity)
}

func (c *CurrentCondition) GetId() int {
	return int(c.id)
}
