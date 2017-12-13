package structs

import (
	"reflect"
	"testing"
)

func TestQueue_Len(t *testing.T) {
	tests := []struct {
		name string
		h    Queue
		want int
	}{
		{"Len", Queue{1, 2, 3, 4, 5}, 5},
		{"Len", Queue{}, 0},
		{"Len", Queue{1}, 1},
		{"Len", Queue{-1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Len(); got != tt.want {
				t.Errorf("Queue.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		h    Queue
		args args
		want bool
	}{
		{"Less", Queue{1, 2, 3}, args{0, 1}, true},
		{"Less", Queue{2, 1, 3}, args{0, 1}, true},
		{"Less", Queue{1, 2, 3, 4}, args{1, 2}, true},
		{"Less", Queue{1, 2, 3, 4}, args{2, 1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Queue.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		h    Queue
		args args
		want Queue
	}{
		{"Swap", Queue{1, 2, 3, 4, 5}, args{0, 1}, Queue{1, 2, 3, 4, 5}},
		{"Swap", Queue{1, 2, 3, 4, 5}, args{1, 0}, Queue{1, 2, 3, 4, 5}},
		{"Swap", Queue{1, 2, 3, 4, 5}, args{0, 3}, Queue{1, 2, 3, 4, 5}},
		{"Swap", Queue{1, 2, 3, 4, 5}, args{3, 0}, Queue{1, 2, 3, 4, 5}},
		{"Swap", Queue{1, 2, 3, 4, 5}, args{3, 4}, Queue{1, 2, 3, 4, 5}},
		{"Swap", Queue{1, 2, 3, 4, 5}, args{4, 3}, Queue{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Swap(tt.args.i, tt.args.j)
			if !reflect.DeepEqual(tt.h, tt.want) {
				t.Errorf("After Queue.Swap() got %v, want %v", tt.h, tt.want)
			}
		})
	}
}

func TestQueue_Push(t *testing.T) {
	type args struct {
		x interface{}
	}
	tests := []struct {
		name string
		h    *Queue
		args args
		want *Queue
	}{
		{"Add", &Queue{0}, args{1}, &Queue{0, 1}},
		{"Add", &Queue{}, args{1}, &Queue{1}},
		{"Add", &Queue{1, 2, 3, 5}, args{4}, &Queue{1, 2, 3, 5, 4}},
		{"Add", &Queue{1, 3, 5, 7}, args{3}, &Queue{1, 3, 5, 7, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Push(tt.args.x)
			if !reflect.DeepEqual(tt.h, tt.want) {
				t.Errorf("After Queue.Add() got %v, want %v", tt.h, tt.want)
			}
		})
	}
}

func TestQueue_Pop(t *testing.T) {
	tests := []struct {
		name string
		h    *Queue
		want interface{}
	}{
		{"Pop", &Queue{0}, 0},
		{"Pop", &Queue{0, 1, 2}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queue.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}
