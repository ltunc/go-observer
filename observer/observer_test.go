package observer

import (
	"reflect"
	"testing"
)

type testEvent struct {
	name string
}

type mockObserver struct {
	name          string
	notifications []testEvent
}

func (o *mockObserver) Notify(ev testEvent) {
	o.notifications = append(o.notifications, ev)
}

func TestSubject_Fire(t *testing.T) {
	type fields struct {
		observers []Observer[testEvent]
	}
	type args struct {
		ev testEvent
	}
	observers := []*mockObserver{{name: "M5"}, {name: "M1"}, {name: "M3"}}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantCalls []testEvent
	}{
		{
			"basic",
			fields{[]Observer[testEvent]{observers[0], observers[1], observers[2]}},
			args{testEvent{"Test1"}},
			[]testEvent{{"Test1"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Subject[testEvent]{
				observers: tt.fields.observers,
			}
			s.Fire(tt.args.ev)
			for _, o := range observers {
				if !reflect.DeepEqual(o.notifications, tt.wantCalls) {
					t.Errorf("Fire() got calls %v, want %v", o.notifications, tt.wantCalls)
				}
			}
		})
	}
}

func TestSubject_Subscribe(t *testing.T) {
	type fields struct {
		observers []Observer[testEvent]
	}
	type args struct {
		obs []Observer[testEvent]
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Observer[testEvent]
	}{
		{
			"empty",
			fields{
				[]Observer[testEvent]{},
			},
			args{[]Observer[testEvent]{&mockObserver{name: "M1"}}},
			[]Observer[testEvent]{&mockObserver{name: "M1"}},
		},
		{
			"not empty",
			fields{
				[]Observer[testEvent]{&mockObserver{name: "M5"}},
			},
			args{[]Observer[testEvent]{&mockObserver{name: "M10"}}},
			[]Observer[testEvent]{&mockObserver{name: "M5"}, &mockObserver{name: "M10"}},
		},
		{
			"multiple at once",
			fields{
				[]Observer[testEvent]{&mockObserver{name: "M5"}},
			},
			args{[]Observer[testEvent]{&mockObserver{name: "M10"}, &mockObserver{name: "M7"}}},
			[]Observer[testEvent]{&mockObserver{name: "M5"}, &mockObserver{name: "M10"}, &mockObserver{name: "M7"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Subject[testEvent]{
				observers: tt.fields.observers,
			}
			s.Subscribe(tt.args.obs...)
			if got := s.observers; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subscribe() got observers %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubject_Unsubscribe(t *testing.T) {
	type fields struct {
		observers []Observer[testEvent]
	}
	type args struct {
		obs Observer[testEvent]
	}
	observers := []Observer[testEvent]{&mockObserver{name: "M5"}, &mockObserver{name: "M1"}, &mockObserver{name: "M3"}}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Observer[testEvent]
	}{
		{
			"not empty",
			fields{
				[]Observer[testEvent]{observers[0], observers[1], observers[2]},
			},
			args{observers[1]},
			[]Observer[testEvent]{observers[0], observers[2]},
		},
		{
			"empty",
			fields{
				[]Observer[testEvent]{},
			},
			args{&mockObserver{name: "M10"}},
			[]Observer[testEvent]{},
		},
		{
			"from the end",
			fields{
				[]Observer[testEvent]{observers[0], observers[1], observers[2]},
			},
			args{observers[2]},
			[]Observer[testEvent]{observers[0], observers[1]},
		},
		{
			"last",
			fields{
				[]Observer[testEvent]{observers[1]},
			},
			args{observers[1]},
			[]Observer[testEvent]{},
		},
		{
			"notfound",
			fields{
				[]Observer[testEvent]{observers[0], observers[1]},
			},
			args{observers[2]},
			[]Observer[testEvent]{observers[0], observers[1]},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Subject[testEvent]{
				observers: tt.fields.observers,
			}
			s.Unsubscribe(tt.args.obs)
			if got := s.observers; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unsubscribe() got observers %v, want %v", got, tt.want)
			}
		})
	}
}
