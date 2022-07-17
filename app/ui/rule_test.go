package ui

import (
	"reflect"
	"testing"
)

func TestNewPositionStruct(t *testing.T) {
	tests := []struct {
		name string
		want *PositionStruct
	}{
		{
			name: "case1",
			want: &PositionStruct{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPositionStruct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPositionStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPositionStruct_clearBoard(t *testing.T) {
	type fields struct {
		sdPlayer    int
		ucpcSquares [256]int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "case",
			fields: fields{
				sdPlayer:    0,
				ucpcSquares: [256]int{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PositionStruct{
				sdPlayer:    tt.fields.sdPlayer,
				ucpcSquares: tt.fields.ucpcSquares,
			}
			p.clearBoard()
		})
	}
}

func TestPositionStruct_startup(t *testing.T) {
	type fields struct {
		sdPlayer    int
		ucpcSquares [256]int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "case",
			fields: fields{
				sdPlayer:    0,
				ucpcSquares: [256]int{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PositionStruct{
				sdPlayer:    tt.fields.sdPlayer,
				ucpcSquares: tt.fields.ucpcSquares,
			}
			p.startup()
		})
	}
}

func TestPositionStruct_changeSide(t *testing.T) {
	type fields struct {
		sdPlayer    int
		ucpcSquares [256]int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "case",
			fields: fields{
				sdPlayer:    0,
				ucpcSquares: [256]int{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PositionStruct{
				sdPlayer:    tt.fields.sdPlayer,
				ucpcSquares: tt.fields.ucpcSquares,
			}
			p.changeSide()
		})
	}
}

func TestPositionStruct_addPiece(t *testing.T) {
	type fields struct {
		sdPlayer    int
		ucpcSquares [256]int
	}
	type args struct {
		sq int
		pc int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "case",
			fields: fields{
				sdPlayer:    0,
				ucpcSquares: [256]int{},
			},
			args: args{
				sq: 0,
				pc: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PositionStruct{
				sdPlayer:    tt.fields.sdPlayer,
				ucpcSquares: tt.fields.ucpcSquares,
			}
			p.addPiece(tt.args.sq, tt.args.pc)
		})
	}
}

func TestPositionStruct_delPiece(t *testing.T) {
	type fields struct {
		sdPlayer    int
		ucpcSquares [256]int
	}
	type args struct {
		sq int
		pc int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "case",
			fields: fields{
				sdPlayer:    0,
				ucpcSquares: [256]int{},
			},
			args: args{
				sq: 0,
				pc: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PositionStruct{
				sdPlayer:    tt.fields.sdPlayer,
				ucpcSquares: tt.fields.ucpcSquares,
			}
			p.delPiece(tt.args.sq, tt.args.pc)
		})
	}
}

func TestPositionStruct_movePiece(t *testing.T) {
	type fields struct {
		sdPlayer    int
		ucpcSquares [256]int
	}
	type args struct {
		mv int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "case",
			fields: fields{
				sdPlayer:    0,
				ucpcSquares: [256]int{},
			},
			args: args{mv: 1},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PositionStruct{
				sdPlayer:    tt.fields.sdPlayer,
				ucpcSquares: tt.fields.ucpcSquares,
			}
			if got := p.movePiece(tt.args.mv); got != tt.want {
				t.Errorf("PositionStruct.movePiece() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPositionStruct_undoMovePiece(t *testing.T) {
	type fields struct {
		sdPlayer    int
		ucpcSquares [256]int
	}
	type args struct {
		mv         int
		pcCaptured int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "case",
			fields: fields{
				sdPlayer:    0,
				ucpcSquares: [256]int{},
			},
			args: args{
				mv:         0,
				pcCaptured: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PositionStruct{
				sdPlayer:    tt.fields.sdPlayer,
				ucpcSquares: tt.fields.ucpcSquares,
			}
			p.undoMovePiece(tt.args.mv, tt.args.pcCaptured)
		})
	}
}

func TestPositionStruct_makeMove(t *testing.T) {
	type fields struct {
		sdPlayer    int
		ucpcSquares [256]int
	}
	type args struct {
		mv int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "",
			fields: fields{
				sdPlayer:    0,
				ucpcSquares: [256]int{},
			},
			args: args{mv: 1},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PositionStruct{
				sdPlayer:    tt.fields.sdPlayer,
				ucpcSquares: tt.fields.ucpcSquares,
			}
			if got := p.makeMove(tt.args.mv); got != tt.want {
				t.Errorf("PositionStruct.makeMove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPositionStruct_undoMakeMove(t *testing.T) {
	type fields struct {
		sdPlayer    int
		ucpcSquares [256]int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "case",
			fields: fields{
				sdPlayer:    0,
				ucpcSquares: [256]int{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PositionStruct{
				sdPlayer:    tt.fields.sdPlayer,
				ucpcSquares: tt.fields.ucpcSquares,
			}
			p.undoMakeMove()
		})
	}
}

func TestPositionStruct_nullMove(t *testing.T) {
	type fields struct {
		sdPlayer    int
		ucpcSquares [256]int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "case",
			fields: fields{
				sdPlayer:    0,
				ucpcSquares: [256]int{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PositionStruct{
				sdPlayer:    tt.fields.sdPlayer,
				ucpcSquares: tt.fields.ucpcSquares,
			}
			p.nullMove()
		})
	}
}

func TestPositionStruct_undoNullMove(t *testing.T) {
	type fields struct {
		sdPlayer    int
		ucpcSquares [256]int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "case",
			fields: fields{
				sdPlayer:    0,
				ucpcSquares: [256]int{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PositionStruct{
				sdPlayer:    tt.fields.sdPlayer,
				ucpcSquares: tt.fields.ucpcSquares,
			}
			p.undoNullMove()
		})
	}
}

func TestPositionStruct_generateMoves(t *testing.T) {
	type fields struct {
		sdPlayer    int
		ucpcSquares [256]int
	}
	type args struct {
		mvs      []int
		bCapture bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "case",
			fields: fields{
				sdPlayer:    0,
				ucpcSquares: [256]int{},
			},
			args: args{
				mvs:      nil,
				bCapture: false,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PositionStruct{
				sdPlayer:    tt.fields.sdPlayer,
				ucpcSquares: tt.fields.ucpcSquares,
			}
			if got := p.generateMoves(tt.args.mvs, tt.args.bCapture); got != tt.want {
				t.Errorf("PositionStruct.generateMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPositionStruct_legalMove(t *testing.T) {
	type fields struct {
		sdPlayer    int
		ucpcSquares [256]int
	}
	type args struct {
		mv int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "case",
			fields: fields{
				sdPlayer:    0,
				ucpcSquares: [256]int{},
			},
			args: args{mv: 1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PositionStruct{
				sdPlayer:    tt.fields.sdPlayer,
				ucpcSquares: tt.fields.ucpcSquares,
			}
			if got := p.legalMove(tt.args.mv); got != tt.want {
				t.Errorf("PositionStruct.legalMove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPositionStruct_checked(t *testing.T) {
	type fields struct {
		sdPlayer    int
		ucpcSquares [256]int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "case",
			fields: fields{
				sdPlayer:    0,
				ucpcSquares: [256]int{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PositionStruct{
				sdPlayer:    tt.fields.sdPlayer,
				ucpcSquares: tt.fields.ucpcSquares,
			}
			if got := p.checked(); got != tt.want {
				t.Errorf("PositionStruct.checked() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPositionStruct_isMate(t *testing.T) {
	type fields struct {
		sdPlayer    int
		ucpcSquares [256]int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "case",
			fields: fields{
				sdPlayer:    0,
				ucpcSquares: [256]int{},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PositionStruct{
				sdPlayer:    tt.fields.sdPlayer,
				ucpcSquares: tt.fields.ucpcSquares,
			}
			if got := p.isMate(); got != tt.want {
				t.Errorf("PositionStruct.isMate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPositionStruct_drawValue(t *testing.T) {
	type fields struct {
		sdPlayer    int
		ucpcSquares [256]int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "case",
			fields: fields{
				sdPlayer:    0,
				ucpcSquares: [256]int{},
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PositionStruct{
				sdPlayer:    tt.fields.sdPlayer,
				ucpcSquares: tt.fields.ucpcSquares,
			}
			if got := p.drawValue(); got != tt.want {
				t.Errorf("PositionStruct.drawValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPositionStruct_mirror(t *testing.T) {
	type fields struct {
		sdPlayer    int
		ucpcSquares [256]int
	}
	type args struct {
		posMirror *PositionStruct
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "case",
			fields: fields{
				sdPlayer:    0,
				ucpcSquares: [256]int{},
			},
			args: args{posMirror: NewPositionStruct()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PositionStruct{
				sdPlayer:    tt.fields.sdPlayer,
				ucpcSquares: tt.fields.ucpcSquares,
			}
			p.mirror(tt.args.posMirror)
		})
	}
}

func TestPositionStruct_printBoard(t *testing.T) {
	type fields struct {
		sdPlayer    int
		ucpcSquares [256]int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "case",
			fields: fields{
				sdPlayer:    0,
				ucpcSquares: [256]int{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PositionStruct{
				sdPlayer:    tt.fields.sdPlayer,
				ucpcSquares: tt.fields.ucpcSquares,
			}
			p.printBoard()
		})
	}
}
