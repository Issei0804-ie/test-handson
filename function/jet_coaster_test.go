package function

import (
	"testing"
)

func Test同値分割分析によるCanRideJetCoaster(t *testing.T) {
	test := []struct {
		age    int
		height int
		want   bool
	}{
		{age: 16, height: 160, want: true},
		{age: 2, height: 90, want: false},
	}
	for _, tt := range test {
		got := CanRideJetCoaster(tt.age, tt.height)
		if got != tt.want {
			t.Errorf("CanRideJetCoaster(%d, %d) = %t; want %t", tt.age, tt.height, got, tt.want)
		}
	}
}

func Test境界値分析によるCanRideJetCoaster(t *testing.T) {
	test := []struct {
		age    int
		height int
		want   bool
	}{
		// 正常値
		{age: 12, height: 130, want: true},
		{age: 12, height: 179, want: true},
		{age: 69, height: 179, want: true},
		{age: 69, height: 130, want: true},
		// 異常値
		// heightの最大値±1
		{age: 12, height: 129, want: false},
		{age: 12, height: 180, want: false},
		{age: 69, height: 129, want: false},
		{age: 69, height: 180, want: false},
		// ageの最大値±1
		{age: 11, height: 130, want: false},
		{age: 11, height: 179, want: false},
		{age: 70, height: 130, want: false},
		{age: 70, height: 179, want: false},
	}
	for _, tt := range test {
		got := CanRideJetCoaster(tt.age, tt.height)
		if got != tt.want {
			t.Errorf("CanRideJetCoaster(%d, %d) = %t; want %t", tt.age, tt.height, got, tt.want)
		}
	}
}
