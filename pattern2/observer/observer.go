package observer

type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers()
}

type Observer interface {
	Update(temperature float64, humidity float64, pressure float64)
	GetId() int
}

type DisplayElement interface {
	display()
}
