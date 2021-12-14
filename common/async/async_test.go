package async

import (
	"fmt"
	"testing"
	"time"
)

func TestAsync_Add(t *testing.T) {
	type args struct {
		f func()
	}
	tests := []struct {
		name string
		as   *Async
		args args
	}{
		// TODO: Add test cases.
		{
			name: "1",
			as:   New(),
			args: args{
				f: func() { fmt.Println("test1") },
			},
		},
		// TODO: Add test cases.
		{
			name: "2",
			as:   New(),
			args: args{
				f: func() { fmt.Println("test2") },
			},
		},
		// TODO: Add test cases.
		{
			name: "3",
			as:   New(),
			args: args{
				f: func() { fmt.Println("test3") },
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.as.Add(tt.args.f)
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.as.RunAndWait()
		})
	}

}

func TestAsync_RunAndWait(t *testing.T) {
	tests := []struct {
		name string
		as   *Async
	}{
		{
			name: "1",
			as:   New(),
		},
	}
	tests[0].as.fs = append(tests[0].as.fs, func() { fmt.Println("test1") })
	tests[0].as.fs = append(tests[0].as.fs, func() { fmt.Println("test2") })
	tests[0].as.fs = append(tests[0].as.fs, func() { fmt.Println("test3") })
	time.Sleep(10 * time.Millisecond)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.as.RunAndWait()
			fmt.Println(tt.as.cost)
			fmt.Printf("time cost = %v", tt.as.cost)
		})
	}
}
