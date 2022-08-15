package observer

import "sync"

type Observer[T any] interface {
	Notify(ev T)
}

// Subject is a watcher for events with type E.
// It sends notifications about event E to all subscribed observers.
// Subscribed observers receive the same message as the subject
type Subject[E any] struct {
	observers []Observer[E]
	sync.Mutex
}

// Subscribe adds the observer to the list of subscribers
func (s *Subject[E]) Subscribe(obs Observer[E]) {
	s.Lock()
	defer s.Unlock()
	s.observers = append(s.observers, obs)
}

// Unsubscribe removes the observer from the list of subscribers
func (s *Subject[E]) Unsubscribe(obs Observer[E]) {
	s.Lock()
	defer s.Unlock()
	for k, o := range s.observers {
		if o == obs {
			s.observers = append(s.observers[:k], s.observers[k+1:]...)
			return
		}
	}
}

// Fire sends event ev to all subscribed observers
func (s *Subject[E]) Fire(ev E) {
	s.Lock()
	defer s.Unlock()
	for _, o := range s.observers {
		o.Notify(ev)
	}
}
