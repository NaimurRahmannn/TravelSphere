package utils

import "fmt"

// FormatPopulation turns a raw count into a compact human label:
// 169828911 -> "169.8M", 88400 -> "88.4K". Keeps the cards readable.
func FormatPopulation(n int) string {
	switch {
	case n >= 1_000_000_000:
		return fmt.Sprintf("%.1fB", float64(n)/1_000_000_000)
	case n >= 1_000_000:
		return fmt.Sprintf("%.1fM", float64(n)/1_000_000)
	case n >= 1_000:
		return fmt.Sprintf("%.1fK", float64(n)/1_000)
	default:
		return fmt.Sprintf("%d", n)
	}
}