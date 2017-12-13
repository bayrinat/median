package structs

import (
	"container/heap"
	"reflect"
	"testing"
)

func TestNewMedianHeap(t *testing.T) {
	tests := []struct {
		name string
		want *MedianHeap
	}{
		{"NewMedianHeap", &MedianHeap{
			minHeap: MinHeap{},
			maxHeap: MaxHeap{},
			median:  0,
		},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMedianHeap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMedianHeap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMedianHeap_Len(t *testing.T) {
	type fields struct {
		minHeap MinHeap
		maxHeap MaxHeap
		median  float64
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"Len", fields{
			minHeap: MinHeap{1, 2, 3},
			maxHeap: MaxHeap{4, 5, 6},
			median:  0,
		}, 6},

		{"Len", fields{
			minHeap: MinHeap{1, 2, 3},
			maxHeap: MaxHeap{},
			median:  1,
		}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &MedianHeap{
				minHeap: tt.fields.minHeap,
				maxHeap: tt.fields.maxHeap,
				median:  tt.fields.median,
			}
			if got := h.Len(); got != tt.want {
				t.Errorf("MedianHeap.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMedianHeap_Push(t *testing.T) {
	type fields struct {
		minHeap MinHeap
		maxHeap MaxHeap
		median  float64
	}
	type args struct {
		x int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"Push", fields{
			minHeap: MinHeap{},
			maxHeap: MaxHeap{},
			median:  0,
		}, args{6}},

		{"Push", fields{
			minHeap: MinHeap{},
			maxHeap: MaxHeap{},
			median:  1,
		}, args{3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &MedianHeap{
				minHeap: tt.fields.minHeap,
				maxHeap: tt.fields.maxHeap,
				median:  tt.fields.median,
			}
			h.Add(tt.args.x)
		})
	}
}

func TestMedianHeap_Median(t *testing.T) {
	// TODO: add tests
}

func TestMedianHeap_Remove(t *testing.T) {
	tests := []struct {
		name        string
		iterations  int
		removeValue int
		wantLen     int
	}{
		{"remove", 10, 1, 9},
		{"remove", 10, 5, 9},
		{"remove", 10, 4, 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewMedianHeap()

			heap.Init(&h.minHeap)
			heap.Init(&h.maxHeap)

			for i := 0; i < tt.iterations; i++ {
				h.Add(i)
			}

			h.Remove(tt.removeValue)

			if got := h.Len(); got != tt.wantLen {
				t.Errorf("After remove() minHeap= %v, want %v", got, tt.wantLen)
			}
		})
	}
}

func TestMedianHeap_balance(t *testing.T) {
	// TODO: add tests
}

func TestMedianHeap_setMedian(t *testing.T) {
	// TODO: add tests
}
