// Hygiene — magic numbers smell.
package quality
func ShipCost(total float64) float64 {
    if total < 49.99 { return 9.99 }
    if total < 99.99 { return 4.99 }
    return 0
}
