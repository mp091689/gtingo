package country

import (
	"testing"
)

func TestCountry_GetCode(t *testing.T) {
	type fields struct {
		code string
		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{"simple get code", fields{"482", "Ukraine"}, []string{"482"}},
		{"parse range code", fields{"600-603", "South Africa"}, []string{"600", "601", "602", "603"}},
		{"parse short code", fields{"11", "Wakanda"}, []string{"011"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Country{
				code: tt.fields.code,
				name: tt.fields.name,
			}
			got := c.GetCode()

			wantFunc := func() bool {
				for _, want := range tt.want {
					if want == got {
						return true
					}
				}

				return false
			}

			if !wantFunc() {
				t.Errorf("Country.GetCode() = %v, want range of %v", got, tt.want)
			}
		})
	}
}

func TestCountry_GetName(t *testing.T) {
	type fields struct {
		code string
		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"check", fields{"482", "Ukraine"}, "Ukraine"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Country{
				code: tt.fields.code,
				name: tt.fields.name,
			}
			if got := c.GetName(); got != tt.want {
				t.Errorf("Country.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCountry(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"get random country"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCountry()
			if len(got.GetName()) == 0 || len(got.GetCode()) == 0 {
				t.Errorf("NewCountry() return empty name or code = %+v", got)
			}
		})
	}
}
