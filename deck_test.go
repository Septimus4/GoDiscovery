package main

import (
	"reflect"
	"testing"
)

func Test_deck_deal(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name  string
		d     deck
		args  args
		want  deck
		want1 deck
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.d.deal(tt.args.size)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deal() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("deal() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_deck_print(t *testing.T) {
	tests := []struct {
		name string
		d    deck
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.print()
		})
	}
}

func Test_deck_saveToFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		d       deck
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.saveToFile(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("saveToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_deck_shuffle(t *testing.T) {
	tests := []struct {
		name string
		d    deck
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.shuffle()
		})
	}
}

func Test_deck_toString(t *testing.T) {
	tests := []struct {
		name string
		d    deck
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.toString(); got != tt.want {
				t.Errorf("toString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadFromFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want deck
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := loadFromFile(tt.args.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newDeck(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{name: "newDeck", want: 52},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}
