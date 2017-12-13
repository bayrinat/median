package metric

import (
	"median/structs"
	"reflect"
	"testing"
)

func TestNewDelay(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name    string
		args    args
		want    *DelayBase
		wantErr bool
	}{
		{"NewDelay", args{16}, &DelayBase{
			structs.NewMedianHeap(), &structs.Queue{}, 16}, false},
		{"NewDelay", args{0}, nil, true},
		{"NewDelay", args{-10}, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDelay(tt.args.capacity)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDelay() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDelay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelayBase_AddDelay(t *testing.T) {
	tests := []struct {
		name       string
		capacity   int
		iterations int
		wantLen    int
	}{
		{"AddDelay", 16, 10, 10},
		{"AddDelay", 16, 32, 16},
		{"AddDelay", 1, 1000, 1},
		{"AddDelay", 1000, 5, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDelay(tt.capacity)
			if err != nil {
				t.Errorf("NewDelay() error = %v", err)
				return
			}

			for i := 0; i < tt.iterations; i++ {
				got.AddDelay(i)
			}

			if got.medianHeap.Len() != tt.wantLen {
				t.Errorf("length of the heap is wrong, got=%v, but want=%v", got.medianHeap.Len(), tt.wantLen)
			}

			if got.queue.Len() != tt.wantLen {
				t.Errorf("length of the queue is wrong, got=%v, but want=%v", got.medianHeap.Len(), tt.wantLen)
			}
		})
	}
}

func TestDelayBase_GetMedian(t *testing.T) {
	tests := []struct {
		name       string
		capacity   int
		stream     []int
		wantMedian float64
	}{
		{"GetMedian", 16, []int{1, 2, 3, 4}, 2.5},
		{"GetMedian", 16, []int{1, 2, 3, 4, 5}, 3},
		{"GetMedian", 3, []int{1, 2, 3, 4, 5, 6, 7}, 6},
		{"GetMedian", 8, []int{1}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDelay(tt.capacity)
			if err != nil {
				t.Errorf("NewDelay() error = %v", err)
				return
			}

			for _, v := range tt.stream {
				got.AddDelay(v)
			}

			if got.GetMedian() != tt.wantMedian {
				t.Errorf("median is wrong, got=%v, but want=%v", got.medianHeap.Len(), tt.wantMedian)
			}
		})
	}
}
