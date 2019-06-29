package greeting

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	println("Set message")

	message = "Hey!"

	os.Exit(m.Run())
}

func TestHey(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Success",
			args: args{
				name: "Tom",
			},
			want: "HOGEHOGE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hey(tt.args.name); got != tt.want {
				t.Errorf("Hey() = %v, want %v", got, tt.want)
			}
		})
	}
}
