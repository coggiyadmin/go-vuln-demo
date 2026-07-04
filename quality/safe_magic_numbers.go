// SAFE — named constants.
package quality
const (
    freeShipThreshold = 99.99
    reducedShipThreshold = 49.99
    reducedShipCost = 4.99
    standardShipCost = 9.99
)
func ShipCost(total float64) float64 {
    if total < reducedShipThreshold { return standardShipCost }
    if total < freeShipThreshold { return reducedShipCost }
    return 0
}
