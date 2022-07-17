package ui

import (
	_ "image/png"
	"testing"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
)

//func TestNewGame(t *testing.T) {
//	tests := []struct {
//		name string
//		want bool
//	}{
//		{
//			name: "case",
//			want: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewGame(); got != tt.want {
//				t.Errorf("NewGame() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestGame_Update(t *testing.T) {
	type fields struct {
		sqSelected     int
		mvLast         int
		bFlipped       bool
		bGameOver      bool
		showValue      string
		images         map[int]*ebiten.Image
		audios         map[int]*audio.Player
		audioContext   *audio.Context
		singlePosition *PositionStruct
	}
	type args struct {
		screen *ebiten.Image
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "case",
			fields: fields{
				sqSelected:     0,
				mvLast:         0,
				bFlipped:       false,
				bGameOver:      false,
				showValue:      "1",
				images:         nil,
				audios:         nil,
				audioContext:   nil,
				singlePosition: NewPositionStruct(),
			},
			args:    args{screen: nil},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				sqSelected:     tt.fields.sqSelected,
				mvLast:         tt.fields.mvLast,
				bFlipped:       tt.fields.bFlipped,
				bGameOver:      tt.fields.bGameOver,
				showValue:      tt.fields.showValue,
				images:         tt.fields.images,
				audios:         tt.fields.audios,
				audioContext:   tt.fields.audioContext,
				singlePosition: tt.fields.singlePosition,
			}
			if err := g.Update(tt.args.screen); (err != nil) != tt.wantErr {
				t.Errorf("Game.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGame_Layout(t *testing.T) {
	type fields struct {
		sqSelected     int
		mvLast         int
		bFlipped       bool
		bGameOver      bool
		showValue      string
		images         map[int]*ebiten.Image
		audios         map[int]*audio.Player
		audioContext   *audio.Context
		singlePosition *PositionStruct
	}
	type args struct {
		outsideWidth  int
		outsideHeight int
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantScreenWidth  int
		wantScreenHeight int
	}{
		{
			name: "case",
			fields: fields{
				sqSelected:     0,
				mvLast:         0,
				bFlipped:       false,
				bGameOver:      false,
				showValue:      "",
				images:         nil,
				audios:         nil,
				audioContext:   nil,
				singlePosition: NewPositionStruct(),
			},
			args: args{
				outsideWidth:  0,
				outsideHeight: 0,
			},
			wantScreenWidth:  520,
			wantScreenHeight: 576,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				sqSelected:     tt.fields.sqSelected,
				mvLast:         tt.fields.mvLast,
				bFlipped:       tt.fields.bFlipped,
				bGameOver:      tt.fields.bGameOver,
				showValue:      tt.fields.showValue,
				images:         tt.fields.images,
				audios:         tt.fields.audios,
				audioContext:   tt.fields.audioContext,
				singlePosition: tt.fields.singlePosition,
			}
			gotScreenWidth, gotScreenHeight := g.Layout(tt.args.outsideWidth, tt.args.outsideHeight)
			if gotScreenWidth != tt.wantScreenWidth {
				t.Errorf("Game.Layout() gotScreenWidth = %v, want %v", gotScreenWidth, tt.wantScreenWidth)
			}
			if gotScreenHeight != tt.wantScreenHeight {
				t.Errorf("Game.Layout() gotScreenHeight = %v, want %v", gotScreenHeight, tt.wantScreenHeight)
			}
		})
	}
}

func TestGame_loadResource(t *testing.T) {
	type fields struct {
		sqSelected     int
		mvLast         int
		bFlipped       bool
		bGameOver      bool
		showValue      string
		images         map[int]*ebiten.Image
		audios         map[int]*audio.Player
		audioContext   *audio.Context
		singlePosition *PositionStruct
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "case",
			fields: fields{
				sqSelected:     0,
				mvLast:         0,
				bFlipped:       false,
				bGameOver:      false,
				showValue:      "",
				images:         make(map[int]*ebiten.Image),
				audios:         make(map[int]*audio.Player),
				audioContext:   new(audio.Context),
				singlePosition: NewPositionStruct(),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				sqSelected:     tt.fields.sqSelected,
				mvLast:         tt.fields.mvLast,
				bFlipped:       tt.fields.bFlipped,
				bGameOver:      tt.fields.bGameOver,
				showValue:      tt.fields.showValue,
				images:         tt.fields.images,
				audios:         tt.fields.audios,
				audioContext:   tt.fields.audioContext,
				singlePosition: tt.fields.singlePosition,
			}
			if got := g.loadResource(); got != tt.want {
				t.Errorf("Game.loadResource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGame_drawBoard(t *testing.T) {
	type fields struct {
		sqSelected     int
		mvLast         int
		bFlipped       bool
		bGameOver      bool
		showValue      string
		images         map[int]*ebiten.Image
		audios         map[int]*audio.Player
		audioContext   *audio.Context
		singlePosition *PositionStruct
	}
	type args struct {
		screen *ebiten.Image
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "case",
			fields: fields{
				sqSelected:     0,
				mvLast:         0,
				bFlipped:       false,
				bGameOver:      false,
				showValue:      "",
				images:         make(map[int]*ebiten.Image),
				audios:         make(map[int]*audio.Player),
				audioContext:   new(audio.Context),
				singlePosition: NewPositionStruct(),
			},
			args: args{screen: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				sqSelected:     tt.fields.sqSelected,
				mvLast:         tt.fields.mvLast,
				bFlipped:       tt.fields.bFlipped,
				bGameOver:      tt.fields.bGameOver,
				showValue:      tt.fields.showValue,
				images:         tt.fields.images,
				audios:         tt.fields.audios,
				audioContext:   tt.fields.audioContext,
				singlePosition: tt.fields.singlePosition,
			}
			g.drawBoard(tt.args.screen)
		})
	}
}

func TestGame_drawChess(t *testing.T) {
	type fields struct {
		sqSelected     int
		mvLast         int
		bFlipped       bool
		bGameOver      bool
		showValue      string
		images         map[int]*ebiten.Image
		audios         map[int]*audio.Player
		audioContext   *audio.Context
		singlePosition *PositionStruct
	}
	type args struct {
		x      int
		y      int
		screen *ebiten.Image
		img    *ebiten.Image
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "case",
			fields: fields{
				sqSelected:     0,
				mvLast:         0,
				bFlipped:       false,
				bGameOver:      false,
				showValue:      "",
				images:         make(map[int]*ebiten.Image),
				audios:         make(map[int]*audio.Player),
				audioContext:   new(audio.Context),
				singlePosition: NewPositionStruct(),
			},
			args: args{
				x:      0,
				y:      0,
				screen: nil,
				img:    nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				sqSelected:     tt.fields.sqSelected,
				mvLast:         tt.fields.mvLast,
				bFlipped:       tt.fields.bFlipped,
				bGameOver:      tt.fields.bGameOver,
				showValue:      tt.fields.showValue,
				images:         tt.fields.images,
				audios:         tt.fields.audios,
				audioContext:   tt.fields.audioContext,
				singlePosition: tt.fields.singlePosition,
			}
			g.drawChess(tt.args.x, tt.args.y, tt.args.screen, tt.args.img)
		})
	}
}

func TestGame_clickSquare(t *testing.T) {
	type fields struct {
		sqSelected     int
		mvLast         int
		bFlipped       bool
		bGameOver      bool
		showValue      string
		images         map[int]*ebiten.Image
		audios         map[int]*audio.Player
		audioContext   *audio.Context
		singlePosition *PositionStruct
	}
	type args struct {
		screen *ebiten.Image
		sq     int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "case",
			fields: fields{
				sqSelected:     0,
				mvLast:         0,
				bFlipped:       false,
				bGameOver:      false,
				showValue:      "",
				images:         make(map[int]*ebiten.Image),
				audios:         make(map[int]*audio.Player),
				audioContext:   new(audio.Context),
				singlePosition: NewPositionStruct(),
			},
			args: args{
				screen: nil,
				sq:     0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				sqSelected:     tt.fields.sqSelected,
				mvLast:         tt.fields.mvLast,
				bFlipped:       tt.fields.bFlipped,
				bGameOver:      tt.fields.bGameOver,
				showValue:      tt.fields.showValue,
				images:         tt.fields.images,
				audios:         tt.fields.audios,
				audioContext:   tt.fields.audioContext,
				singlePosition: tt.fields.singlePosition,
			}
			g.clickSquare(tt.args.screen, tt.args.sq)
		})
	}
}

func TestGame_playAudio(t *testing.T) {
	type fields struct {
		sqSelected     int
		mvLast         int
		bFlipped       bool
		bGameOver      bool
		showValue      string
		images         map[int]*ebiten.Image
		audios         map[int]*audio.Player
		audioContext   *audio.Context
		singlePosition *PositionStruct
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "case",
			fields: fields{
				sqSelected:     0,
				mvLast:         0,
				bFlipped:       false,
				bGameOver:      false,
				showValue:      "",
				images:         make(map[int]*ebiten.Image),
				audios:         make(map[int]*audio.Player),
				audioContext:   new(audio.Context),
				singlePosition: NewPositionStruct(),
			},
			args: args{value: BanValue},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				sqSelected:     tt.fields.sqSelected,
				mvLast:         tt.fields.mvLast,
				bFlipped:       tt.fields.bFlipped,
				bGameOver:      tt.fields.bGameOver,
				showValue:      tt.fields.showValue,
				images:         tt.fields.images,
				audios:         tt.fields.audios,
				audioContext:   tt.fields.audioContext,
				singlePosition: tt.fields.singlePosition,
			}
			g.playAudio(tt.args.value)
		})
	}
}
