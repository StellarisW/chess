package ui

import "testing"

func Test_inBoard(t *testing.T) {
	type args struct {
		sq int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case",
			args: args{sq: 5},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inBoard(tt.args.sq); got != tt.want {
				t.Errorf("inBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inFort(t *testing.T) {
	type args struct {
		sq int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case",
			args: args{sq: 5},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inFort(tt.args.sq); got != tt.want {
				t.Errorf("inFort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getY(t *testing.T) {
	type args struct {
		sq int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case",
			args: args{sq: 6},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getY(tt.args.sq); got != tt.want {
				t.Errorf("getY() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getX(t *testing.T) {
	type args struct {
		sq int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case",
			args: args{sq: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getX(tt.args.sq); got != tt.want {
				t.Errorf("getX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_squareXY(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case",
			args: args{
				x: 5,
				y: 3,
			},
			want: 53,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := squareXY(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("squareXY() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_squareFlip(t *testing.T) {
	type args struct {
		sq int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case",
			args: args{sq: 3},
			want: 251,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := squareFlip(tt.args.sq); got != tt.want {
				t.Errorf("squareFlip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_xFlip(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case",
			args: args{x: 3},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xFlip(tt.args.x); got != tt.want {
				t.Errorf("xFlip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yFlip(t *testing.T) {
	type args struct {
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case",
			args: args{y: 3},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := yFlip(tt.args.y); got != tt.want {
				t.Errorf("yFlip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mirrorSquare(t *testing.T) {
	type args struct {
		sq int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case",
			args: args{sq: 4},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mirrorSquare(tt.args.sq); got != tt.want {
				t.Errorf("mirrorSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_squareForward(t *testing.T) {
	type args struct {
		sq int
		sd int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case",
			args: args{
				sq: 3,
				sd: 6,
			},
			want: 179,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := squareForward(tt.args.sq, tt.args.sd); got != tt.want {
				t.Errorf("squareForward() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jiangSpan(t *testing.T) {
	type args struct {
		sqSrc int
		sqDst int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case",
			args: args{
				sqSrc: 3,
				sqDst: 8,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jiangSpan(tt.args.sqSrc, tt.args.sqDst); got != tt.want {
				t.Errorf("jiangSpan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shiSpan(t *testing.T) {
	type args struct {
		sqSrc int
		sqDst int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case",
			args: args{
				sqSrc: 5,
				sqDst: 7,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shiSpan(tt.args.sqSrc, tt.args.sqDst); got != tt.want {
				t.Errorf("shiSpan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_xiangSpan(t *testing.T) {
	type args struct {
		sqSrc int
		sqDst int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case",
			args: args{
				sqSrc: 3,
				sqDst: 8,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xiangSpan(tt.args.sqSrc, tt.args.sqDst); got != tt.want {
				t.Errorf("xiangSpan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_xiangPin(t *testing.T) {
	type args struct {
		sqSrc int
		sqDst int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case",
			args: args{
				sqSrc: 5,
				sqDst: 9,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xiangPin(tt.args.sqSrc, tt.args.sqDst); got != tt.want {
				t.Errorf("xiangPin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maPin(t *testing.T) {
	type args struct {
		sqSrc int
		sqDst int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case",
			args: args{
				sqSrc: 3,
				sqDst: 9,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maPin(tt.args.sqSrc, tt.args.sqDst); got != tt.want {
				t.Errorf("maPin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_noRiver(t *testing.T) {
	type args struct {
		sq int
		sd int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case",
			args: args{
				sq: 3,
				sd: 9,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := noRiver(tt.args.sq, tt.args.sd); got != tt.want {
				t.Errorf("noRiver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasRiver(t *testing.T) {
	type args struct {
		sq int
		sd int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case",
			args: args{
				sq: 3,
				sd: 8,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasRiver(tt.args.sq, tt.args.sd); got != tt.want {
				t.Errorf("hasRiver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sameRiver(t *testing.T) {
	type args struct {
		sqSrc int
		sqDst int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case",
			args: args{
				sqSrc: 3,
				sqDst: 8,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sameRiver(tt.args.sqSrc, tt.args.sqDst); got != tt.want {
				t.Errorf("sameRiver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sameX(t *testing.T) {
	type args struct {
		sqSrc int
		sqDst int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case",
			args: args{
				sqSrc: 9,
				sqDst: 3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sameX(tt.args.sqSrc, tt.args.sqDst); got != tt.want {
				t.Errorf("sameX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sameY(t *testing.T) {
	type args struct {
		sqSrc int
		sqDst int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case",
			args: args{
				sqSrc: 9,
				sqDst: 5,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sameY(tt.args.sqSrc, tt.args.sqDst); got != tt.want {
				t.Errorf("sameY() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sideTag(t *testing.T) {
	type args struct {
		sd int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case",
			args: args{sd: 4},
			want: 40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sideTag(tt.args.sd); got != tt.want {
				t.Errorf("sideTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_oppSideTag(t *testing.T) {
	type args struct {
		sd int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case",
			args: args{sd: 4},
			want: -16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oppSideTag(tt.args.sd); got != tt.want {
				t.Errorf("oppSideTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_src(t *testing.T) {
	type args struct {
		mv int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case",
			args: args{mv: 3},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := src(tt.args.mv); got != tt.want {
				t.Errorf("src() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dst(t *testing.T) {
	type args struct {
		mv int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case",
			args: args{mv: 2},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dst(tt.args.mv); got != tt.want {
				t.Errorf("dst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_move(t *testing.T) {
	type args struct {
		sqSrc int
		sqDst int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case",
			args: args{
				sqSrc: 3,
				sqDst: 3,
			},
			want: 771,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := move(tt.args.sqSrc, tt.args.sqDst); got != tt.want {
				t.Errorf("move() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mirrorMove(t *testing.T) {
	type args struct {
		mv int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case",
			args: args{mv: 6},
			want: 3592,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mirrorMove(tt.args.mv); got != tt.want {
				t.Errorf("mirrorMove() = %v, want %v", got, tt.want)
			}
		})
	}
}
