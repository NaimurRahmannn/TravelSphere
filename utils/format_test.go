package utils

import "testing"


func TestFormatPopulation(t *testing.T) {
	cases := []struct {
		name string
		in   int
		want string
	}{
		{"zero", 0, "0"},
		{"hundreds stays raw", 999, "999"},
		{"exactly one thousand", 1000, "1.0K"},
		{"thousands", 88400, "88.4K"},
		{"exactly one million", 1_000_000, "1.0M"},
		{"millions", 169828911, "169.8M"},
		{"exactly one billion", 1_000_000_000, "1.0B"},
		{"billions", 1_400_000_000, "1.4B"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := FormatPopulation(c.in)
			if got != c.want {
				t.Errorf("FormatPopulation(%d) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}