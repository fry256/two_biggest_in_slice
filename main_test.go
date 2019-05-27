package main

import "testing"

func Test_twoBiggestEls(t *testing.T) {
	type args struct {
		numbers []float32
	}
	tests := []struct {
		name    string
		args    args
		want    float32
		want1   float32
		wantErr bool
	}{
		{
			name: "Empty slice",
			args: args{
				numbers: []float32{},
			},
			want:    0,
			want1:   0,
			wantErr: true,
		},
		{
			name: "Short slice",
			args: args{
				numbers: []float32{1.0},
			},
			want:    0,
			want1:   0,
			wantErr: true,
		},
		{
			name: "Two of the same item",
			args: args{
				numbers: []float32{3.0, 3.0},
			},
			want:    0,
			want1:   0,
			wantErr: true,
		},
		{
			name: "Haven't different elements",
			args: args{
				numbers: []float32{2.0, 2.0, 2.0, 2.0, 2.0, 2.0},
			},
			want:    0,
			want1:   0,
			wantErr: true,
		},
		{
			name: "Correct answer 1",
			args: args{
				numbers: []float32{12.0, 4.0, 8.1, 4.4, 6.767},
			},
			want:    8.1,
			want1:   12.0,
			wantErr: false,
		},
		{
			name: "Correct answer 2",
			args: args{
				numbers: []float32{4.4, 54.5, 12.0, 43.3, 63.0, 12.2},
			},
			want:    54.5,
			want1:   63.0,
			wantErr: false,
		},
		{
			name: "Correct answer 3",
			args: args{
				numbers: []float32{34.5, 21.2, 1.0, 43.4},
			},
			want:    34.5,
			want1:   43.4,
			wantErr: false,
		},
		{
			name: "Correct answer 4",
			args: args{
				numbers: []float32{5, 5, 5, 5, 4, 4},
			},
			want:    4,
			want1:   5,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := twoBiggestEls(tt.args.numbers)
			if (err != nil) != tt.wantErr {
				t.Errorf("twoBiggestEls() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("twoBiggestEls() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("twoBiggestEls() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
