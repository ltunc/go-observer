# Go Observer

an implementation of Observer pattern on go using generics.

## Usage

```
type Event string

type PrintObserver struct {
    //...
}
func (o *EventObserver) Notify(ev Event) {
    fmt.Println(ev)
}
func process() {
    subj := &Subject[Event]{}
    priter := &PrintObserver{}
    subj.Subscribe(priter)
    subj.Fire(Event("ev1"))
    subj.Fire(Event("ev2"))
    subj.Unsubscribe(priter)
}
```