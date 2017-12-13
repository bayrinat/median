package structs

import (
	"container/heap"
	"reflect"
	"testing"
)

func TestMaxHeap_Len(t *testing.T) {
	tests := []struct {
		name string
		h    MaxHeap
		want int
	}{
		{"Len", MaxHeap{1, 2, 3, 4, 5}, 5},
		{"Len", MaxHeap{}, 0},
		{"Len", MaxHeap{1}, 1},
		{"Len", MaxHeap{-1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Len(); got != tt.want {
				t.Errorf("MaxHeap.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxHeap_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		h    MaxHeap
		args args
		want bool
	}{
		{"Less", MaxHeap{1, 2, 3}, args{0, 1}, false},
		{"Less", MaxHeap{2, 1, 3}, args{0, 1}, true},
		{"Less", MaxHeap{1, 2, 3, 4}, args{1, 2}, false},
		{"Less", MaxHeap{1, 2, 3, 4}, args{2, 1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("MaxHeap.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxHeap_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		h    MaxHeap
		args args
		want MaxHeap
	}{
		{"Swap", MaxHeap{1, 2, 3, 4, 5}, args{0, 1}, MaxHeap{2, 1, 3, 4, 5}},
		{"Swap", MaxHeap{1, 2, 3, 4, 5}, args{1, 0}, MaxHeap{2, 1, 3, 4, 5}},
		{"Swap", MaxHeap{1, 2, 3, 4, 5}, args{0, 3}, MaxHeap{4, 2, 3, 1, 5}},
		{"Swap", MaxHeap{1, 2, 3, 4, 5}, args{3, 0}, MaxHeap{4, 2, 3, 1, 5}},
		{"Swap", MaxHeap{1, 2, 3, 4, 5}, args{3, 4}, MaxHeap{1, 2, 3, 5, 4}},
		{"Swap", MaxHeap{1, 2, 3, 4, 5}, args{4, 3}, MaxHeap{1, 2, 3, 5, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Swap(tt.args.i, tt.args.j)
			if !reflect.DeepEqual(tt.h, tt.want) {
				t.Errorf("After MaxHeap.Swap() got %v, want %v", tt.h, tt.want)
			}
		})
	}
}

func TestMaxHeap_Push(t *testing.T) {
	tests := []struct {
		name string
		h    *MaxHeap
		args int
		want int
	}{
		{"Add", &MaxHeap{0}, 1, 1},
		{"Add", &MaxHeap{}, 1, 1},
		{"Add", &MaxHeap{1, 2, 3, 5}, 4, 5},
		{"Add", &MaxHeap{1, 3, 5, 7}, 9, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap.Init(tt.h)
			heap.Push(tt.h, tt.args)
			if got := (*tt.h)[0]; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxHeap.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxHeap_Pop(t *testing.T) {
	tests := []struct {
		name string
		h    *MaxHeap
		want interface{}
	}{
		{"Pop", &MaxHeap{0}, 0},
		{"Pop", &MaxHeap{0, 1, 2}, 2},
		{"Pop", &MaxHeap{2, 1, 0}, 2},
		{"Pop", &MaxHeap{-2, 1, 0}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap.Init(tt.h)
			if got := heap.Pop(tt.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxHeap.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxHeap_IndexOf(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		h    *MaxHeap
		args args
		want int
	}{
		{"IndexOf", &MaxHeap{0}, args{0}, 0},
		{"IndexOf", &MaxHeap{0, 1, 2}, args{0}, 0},
		{"IndexOf", &MaxHeap{0, 1, 2, 3}, args{1}, 1},
		{"IndexOf", &MaxHeap{0, 1, 2, 3, 4}, args{4}, 4},
		{"IndexOf", &MaxHeap{0, 1, 2, 3, 9}, args{9}, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.IndexOf(tt.args.x); got != tt.want {
				t.Errorf("MaxHeap.IndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
