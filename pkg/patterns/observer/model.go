package observer

type Observer[T any] interface {
	Update(data T)
}

type Subject[T any] interface {
	RegisterObserver(o Observer[T])
	RemoveObserver(o Observer[T])
	NotifyObservers()
}
