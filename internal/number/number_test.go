package number

import "testing"

func TestNumber_Stringify(t *testing.T) {
	tests := []struct {
		name string
		n    Number
		want string
	}{
		{"ok gtin-13", []int{8, 9, 6, 6, 2, 4, 5, 4, 7, 3, 4, 6, 2}, "8966245473462"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Stringify(); got != tt.want {
				t.Errorf("Number.Stringify() = %v, want %v", got, tt.want)
			}
		})
	}
}
