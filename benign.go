// BENIGN-BASELINE TRUE-NEGATIVE FIXTURE.
//
// Ordinary business logic with NO security surface: no HTTP, no DB, no file
// I/O, no exec, no crypto, no secrets. The scanner MUST produce ZERO security
// findings here. Measures specificity / the noise floor.
package main

import (
	"math"
	"sort"
	"strings"
)

// Item is a product line with a unit price and quantity.
type Item struct {
	SKU       string
	UnitPrice float64
	Quantity  int
}

func toCents(n float64) float64 { return math.Round(n*100) / 100 }

func (it Item) extended() float64 { return it.UnitPrice * float64(it.Quantity) }

var tierRates = map[string]float64{"standard": 0, "silver": 0.05, "gold": 0.10}

// Subtotal sums all line extended prices, rounded to cents.
func Subtotal(items []Item) float64 {
	var sum float64
	for _, it := range items {
		sum += it.extended()
	}
	return toCents(sum)
}

// TierFor picks a discount tier from a subtotal.
func TierFor(amount float64) string {
	switch {
	case amount >= 1000:
		return "gold"
	case amount >= 250:
		return "silver"
	default:
		return "standard"
	}
}

// Total returns the final total after the tier discount.
func Total(items []Item) float64 {
	sub := Subtotal(items)
	return toCents(sub - sub*tierRates[TierFor(sub)])
}

// ByCategory groups items by the category code before the dash in the SKU.
func ByCategory(items []Item) map[string][]Item {
	groups := map[string][]Item{}
	for _, it := range items {
		code := "misc"
		if i := strings.Index(it.SKU, "-"); i > 0 {
			code = it.SKU[:i]
		}
		groups[code] = append(groups[code], it)
	}
	return groups
}

// TopSKUs returns SKUs sorted by descending extended price.
func TopSKUs(items []Item, limit int) []string {
	ordered := make([]Item, len(items))
	copy(ordered, items)
	sort.Slice(ordered, func(a, b int) bool { return ordered[a].extended() > ordered[b].extended() })
	out := []string{}
	for i := 0; i < limit && i < len(ordered); i++ {
		out = append(out, ordered[i].SKU)
	}
	return out
}
