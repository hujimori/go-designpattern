package weather

import "v0/observer"

type WeatherData struct {
	observers   []observer.Observer
	temperature float64
	humidity    float64
	pressure    float64
}

func New() *WeatherData {
	w := &WeatherData{}
	w.observers = make([]observer.Observer, 10)
	return w
}

func (w *WeatherData) RegisterObserver(o observer.Observer) {
	w.observers = append(w.observers, o)
}

func (w *WeatherData) RemoveObserver(o observer.Observer) {
	w.RemoveObserver(o)
}

func (w *WeatherData) removeFromslice(removeO observer.Observer) {
	for i, o := range w.observers {
		if o.GetId() == removeO.GetId() {
			w.observers = append(w.observers[:i], w.observers[i+1:]...)
		}
	}
}

func (w *WeatherData) NotifyObservers() {
	for _, o := range w.observers {
		o.Update(w.temperature, w.humidity, w.pressure)
	}
}

func (w *WeatherData) MeasurementsChanged() {
	w.NotifyObservers()
}

func (w *WeatherData) SetMeasurements(temperature float64, humidity float64, pressure float64) {
	w.temperature = temperature
	w.humidity = humidity
	w.pressure = pressure
	w.MeasurementsChanged()
}
