package queue

import (
	"reflect"
	"testing"
)

func TestNewArrayQueue(t *testing.T) {
	tests := []struct {
		name string
		want *ArrayQueue
	}{
		{
			name: "case",
			want: &ArrayQueue{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewArrayQueue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArrayQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayQueue_Top(t *testing.T) {
	type fields struct {
		elements []interface{}
		size     int
	}
	tests := []struct {
		name      string
		fields    fields
		wantValue interface{}
		wantErr   bool
	}{
		{
			name: "empty queue",
			fields: fields{
				elements: []interface{}{},
				size:     0,
			},
			wantValue: nil,
			wantErr:   true,
		},
		{
			name: "filled queue",
			fields: fields{
				elements: []interface{}{3, 5, 8},
				size:     3,
			},
			wantValue: interface{}(3),
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := &ArrayQueue{
				elements: tt.fields.elements,
				size:     tt.fields.size,
			}
			gotValue, err := queue.Top()
			if (err != nil) != tt.wantErr {
				t.Errorf("ArrayQueue.Top() error = %v, wantErr %v,struct:%v", err, tt.wantErr, queue)
				return
			}
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("ArrayQueue.Top() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestArrayQueue_Push(t *testing.T) {
	type fields struct {
		elements []interface{}
		size     int
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "string",
			fields: fields{
				elements: nil,
				size:     0,
			},
			args:    args{value: "5234234"},
			wantErr: false,
		},
		{
			name: "int",
			fields: fields{
				elements: nil,
				size:     0,
			},
			args:    args{value: 34534},
			wantErr: false,
		},
		{
			name: "inconsistent type",
			fields: fields{
				elements: []interface{}{"123", "453"},
				size:     2,
			},
			args:    args{123},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := &ArrayQueue{
				elements: tt.fields.elements,
				size:     tt.fields.size,
			}
			if err := queue.Push(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ArrayQueue.Push() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestArrayQueue_Pop(t *testing.T) {
	type fields struct {
		elements []interface{}
		size     int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "filled queue",
			fields: fields{
				elements: []interface{}{"3", "66", "test"},
				size:     3,
			},
			wantErr: false,
		},
		{
			name: "empty queue",
			fields: fields{
				elements: nil,
				size:     0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := &ArrayQueue{
				elements: tt.fields.elements,
				size:     tt.fields.size,
			}
			if err := queue.Pop(); (err != nil) != tt.wantErr {
				t.Errorf("ArrayQueue.Pop() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestArrayQueue_Empty(t *testing.T) {
	type fields struct {
		elements []interface{}
		size     int
	}
	tests := []struct {
		name   string
		fields fields
		wantIs bool
	}{
		{
			name: "filled queue",
			fields: fields{
				elements: []interface{}{"3", "4"},
				size:     2,
			},
			wantIs: false,
		},
		{
			name: "empty queue",
			fields: fields{
				elements: nil,
				size:     0,
			},
			wantIs: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := &ArrayQueue{
				elements: tt.fields.elements,
				size:     tt.fields.size,
			}
			if gotIs := queue.Empty(); gotIs != tt.wantIs {
				t.Errorf("ArrayQueue.Empty() = %v, want %v", gotIs, tt.wantIs)
			}
		})
	}
}

func TestArrayQueue_Size(t *testing.T) {
	type fields struct {
		elements []interface{}
		size     int
	}
	tests := []struct {
		name     string
		fields   fields
		wantSize int
	}{
		{
			name: "filled queue",
			fields: fields{
				elements: []interface{}{3, 4, 6},
				size:     3,
			},
			wantSize: 3,
		},
		{
			name: "empty queue",
			fields: fields{
				elements: nil,
				size:     0,
			},
			wantSize: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := &ArrayQueue{
				elements: tt.fields.elements,
				size:     tt.fields.size,
			}
			if gotSize := queue.Size(); gotSize != tt.wantSize {
				t.Errorf("ArrayQueue.Size() = %v, want %v", gotSize, tt.wantSize)
			}
		})
	}
}

func TestArrayQueue_Clear(t *testing.T) {
	type fields struct {
		elements []interface{}
		size     int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "filled queue",
			fields: fields{
				elements: []interface{}{3, 4, 6},
				size:     3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := &ArrayQueue{
				elements: tt.fields.elements,
				size:     tt.fields.size,
			}
			queue.Clear()
			if queue.Size() != 0 || len(queue.elements) != 0 {
				t.Errorf("ArrayQueue.Clear() err")
			}
		})
	}
}

func TestArrayQueue_Values(t *testing.T) {
	type fields struct {
		elements []interface{}
		size     int
	}
	tests := []struct {
		name       string
		fields     fields
		wantValues []interface{}
	}{
		{
			name: "filled queue",
			fields: fields{
				elements: []interface{}{"34", "324"},
				size:     2,
			},
			wantValues: []interface{}{"34", "324"},
		},
		{
			name: "empty queue",
			fields: fields{
				elements: []interface{}{},
				size:     0,
			},
			wantValues: []interface{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := &ArrayQueue{
				elements: tt.fields.elements,
				size:     tt.fields.size,
			}
			if gotValues := queue.Values(); !reflect.DeepEqual(gotValues, tt.wantValues) {
				t.Errorf("ArrayQueue.Values() = %v, want %v", gotValues, tt.wantValues)
			}
		})
	}
}

func TestArrayQueue_String(t *testing.T) {
	type fields struct {
		elements []interface{}
		size     int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "filled queue",
			fields: fields{
				elements: []interface{}{"123", "434", "123"},
				size:     3,
			},
			want: "123,434,123",
		},
		{
			name: "empty queue",
			fields: fields{
				elements: nil,
				size:     0,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := &ArrayQueue{
				elements: tt.fields.elements,
				size:     tt.fields.size,
			}
			if got := queue.String(); got != tt.want {
				t.Errorf("ArrayQueue.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
