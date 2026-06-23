// TN — benign output rendering; fixed app-authored result, no model output.
package aisec

import "fmt"

func RenderSummary(count int, total float64) string {
	return fmt.Sprintf("Processed %d items totalling %.2f.", count, total)
}
