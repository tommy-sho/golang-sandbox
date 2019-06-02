package main

import "testing"

func TestCalcualteDigit(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name        string
		args        args
		want        int
		wantErr     bool
		wantErrType Error
	}{
		{
			name: "Success",
			args: args{
				a: 10,
				b: 5,
			},
			want:    10,
			wantErr: false,
		},
		{
			name: "Success",
			args: args{
				a: 10,
				b: 5,
			},
			want:    10,
			wantErr: false,
		},
		{
			name: "Success",
			args: args{
				a: 10,
				b: 5,
			},
			want:    10,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalcualteDigit(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalcualteDigit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CalcualteDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
