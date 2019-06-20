package main

import (
	"testing"
)

func Test_linkedList_PushFront(t *testing.T) {
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name string
		l    *linkedList
		want string
		args args
	}{
		{
			name: "Test PushFront 1",
			l:    CreateLinkedList(),
			want: "LinkedList: {{ 789 -> 456 -> 123 }}",
			args: args{
				v: []interface{}{123, 456, 789},
			},
		},
		{
			name: "Test PushFront 2",
			l:    CreateLinkedList(),
			want: "LinkedList: {{ hello3 -> hello2 -> hello1 }}",
			args: args{
				v: []interface{}{"hello1", "hello2", "hello3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, value := range tt.args.v {
				tt.l.PushFront(value)
			}
			if got := tt.l.String(); got != tt.want {
				t.Errorf("PushFront() not work = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_linkedList_PushBack(t *testing.T) {
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name string
		l    *linkedList
		want string
		args args
	}{
		{
			name: "Test PushBack 1",
			l:    CreateLinkedList(),
			want: "LinkedList: {{ 123 -> 456 -> 789 }}",
			args: args{
				v: []interface{}{123, 456, 789},
			},
		},
		{
			name: "Test PushBack 2",
			l:    CreateLinkedList(),
			want: "LinkedList: {{ hello1 -> hello2 -> hello3 }}",
			args: args{
				v: []interface{}{"hello1", "hello2", "hello3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, value := range tt.args.v {
				tt.l.PushBack(value)
			}
			if got := tt.l.String(); got != tt.want {
				t.Errorf("PushBack() not work = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_linkedList_Remove(t *testing.T) {
	type args struct {
		it *item
	}
	tests := []struct {
		name string
		l    *linkedList
		args args
		want bool
	}{
		{
			name: "Test Remove 1",
			l:    CreateLinkedList(),
			args: args{it: &item{value: "hello"}},
			want: true,
		},
		{
			name: "Test Remove 2",
			l:    CreateLinkedList(),
			args: args{it: &item{value: "hello"}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			switch tt.name {
			case "Test Remove 1":
				tt.l.PushBack(tt.args.it.value)
				if got := tt.l.Remove(tt.l.First()); got != tt.want {
					t.Errorf("linkedList.Remove() = %v, want %v", got, tt.want)
				}
			case "Test Remove 2":
				tt.l.PushBack("hello")
				if got := tt.l.Remove(&item{value: "hello"}); got != tt.want {
					t.Errorf("linkedList.Remove() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
