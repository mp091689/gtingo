package gtingo

import (
	"testing"
)

func TestGtin_Generate(t *testing.T) {
	type args struct {
		format int
	}
	tests := []struct {
		name       string
		args       args
		wantLength int
		wantErr    bool
	}{
		{"generate gtin-8", args{Gtin8}, Gtin8, false},
		{"generate gtin-12", args{Gtin12}, Gtin12, false},
		{"generate gtin-13", args{Gtin13}, Gtin13, false},
		{"generate gtin-14", args{Gtin14}, Gtin14, false},
		{"generate wrang format", args{3}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGtin()
			got, err := g.Generate(tt.args.format)
			if (err != nil) != tt.wantErr {
				t.Errorf("Gtin.Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.wantLength {
				t.Errorf("Gtin.Generate() = %v, want %v", got, tt.wantLength)
			}
		})
	}
}

func TestGtin_Validate(t *testing.T) {
	type args struct {
		number string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"valid", args{"950831448628"}, true},
		{"invalid", args{"950831448623"}, false},
		{"wrong format", args{"111"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGtin()
			got := g.Validate(tt.args.number)
			if got != tt.want {
				t.Errorf("Gtin.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGtin_Calculate(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"calculate gtin-8", args{"7704353"}, "77043535", false},
		{"calculate gtin-12", args{"88026537250"}, "880265372506", false},
		{"calculate gtin-13", args{"622265431361"}, "6222654313617", false},
		{"calculate gtin-14", args{"4629681347651"}, "46296813476510", false},
		{"calculate wrong short gtin", args{"123"}, "", true},
		{"calculate wrong long gtin", args{"123468901234567890"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGtin()
			got, err := g.Calculate(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Gtin.Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Gtin.Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
