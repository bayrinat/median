package structs

import (
	"container/heap"
	"reflect"
	"testing"
)

func TestMinHeap_Len(t *testing.T) {
	tests := []struct {
		name string
		h    MinHeap
		want int
	}{
		{"Len", MinHeap{1, 2, 3, 4, 5}, 5},
		{"Len", MinHeap{}, 0},
		{"Len", MinHeap{1}, 1},
		{"Len", MinHeap{-1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Len(); got != tt.want {
				t.Errorf("MinHeap.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinHeap_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		h    MinHeap
		args args
		want bool
	}{
		{"Less", MinHeap{1, 2, 3}, args{0, 1}, true},
		{"Less", MinHeap{2, 1, 3}, args{0, 1}, false},
		{"Less", MinHeap{1, 2, 3, 4}, args{1, 2}, true},
		{"Less", MinHeap{1, 2, 3, 4}, args{2, 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("MinHeap.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinHeap_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		h    MinHeap
		args args
		want MinHeap
	}{
		{"Swap", MinHeap{1, 2, 3, 4, 5}, args{0, 1}, MinHeap{2, 1, 3, 4, 5}},
		{"Swap", MinHeap{1, 2, 3, 4, 5}, args{1, 0}, MinHeap{2, 1, 3, 4, 5}},
		{"Swap", MinHeap{1, 2, 3, 4, 5}, args{0, 3}, MinHeap{4, 2, 3, 1, 5}},
		{"Swap", MinHeap{1, 2, 3, 4, 5}, args{3, 0}, MinHeap{4, 2, 3, 1, 5}},
		{"Swap", MinHeap{1, 2, 3, 4, 5}, args{3, 4}, MinHeap{1, 2, 3, 5, 4}},
		{"Swap", MinHeap{1, 2, 3, 4, 5}, args{4, 3}, MinHeap{1, 2, 3, 5, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Swap(tt.args.i, tt.args.j)
			if !reflect.DeepEqual(tt.h, tt.want) {
				t.Errorf("After MinHeap.Swap() got %v, want %v", tt.h, tt.want)
			}
		})
	}
}

func TestMinHeap_Push(t *testing.T) {
	tests := []struct {
		name string
		h    *MinHeap
		args int
		want int
	}{
		{"Add", &MinHeap{0}, 1, 0},
		{"Add", &MinHeap{}, 1, 1},
		{"Add", &MinHeap{1, 2, 3, 5}, 4, 1},
		{"Add", &MinHeap{2, 4, 6, 8}, 1, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap.Init(tt.h)
			heap.Push(tt.h, tt.args)
			if got := (*tt.h)[0]; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinHeap.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinHeap_Pop(t *testing.T) {
	tests := []struct {
		name string
		h    *MinHeap
		want interface{}
	}{
		{"Pop", &MinHeap{0}, 0},
		{"Pop", &MinHeap{0, 1, 2}, 0},
		{"Pop", &MinHeap{2, 1, 0}, 0},
		{"Pop", &MinHeap{-2, 1, 0}, -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap.Init(tt.h)
			if got := heap.Pop(tt.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinHeap.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinHeap_IndexOf(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		h    *MinHeap
		args args
		want int
	}{
		{"IndexOf", &MinHeap{0}, args{0}, 0},
		{"IndexOf", &MinHeap{0, 1, 2}, args{0}, 0},
		{"IndexOf", &MinHeap{0, 1, 2, 3}, args{1}, 1},
		{"IndexOf", &MinHeap{0, 1, 2, 3, 4}, args{4}, 4},
		{"IndexOf", &MinHeap{0, 1, 2, 3, 9}, args{9}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.IndexOf(tt.args.x); got != tt.want {
				t.Errorf("MinHeap.IndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
