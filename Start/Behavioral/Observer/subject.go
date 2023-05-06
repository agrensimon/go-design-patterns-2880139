package main

// The DataSubject will have a list of listeners
// and a field that gets changed, triggering them
type DataSubject struct {
	observers []observer
	field     string
}

// The ChangeItem function will cause the Listeners to be called
func (ds *DataSubject) ChangeItem(data string) {
	ds.field = data
	ds.notifyAll()
}

// This function adds an observer to the list
func (ds *DataSubject) registerObserver(o observer) {
	ds.observers = append(ds.observers, o)
}

// This function removes an observer from the list
func (ds *DataSubject) unregisterObserver(o observer) {
	var newObservers []observer
	for _, dl := range ds.observers {
		if o.name() == dl.name() {
			continue
		}
		newObservers = append(newObservers, dl)
	}
	ds.observers = newObservers
}

// The notifyAll function calls all the listeners with the changed data
func (ds *DataSubject) notifyAll() {
	for _, d := range ds.observers {
		d.onUpdate(ds.field)
	}
}
