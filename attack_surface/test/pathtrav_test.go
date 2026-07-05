package attacksurface
import ("os"; "testing")
func TestRead(t *testing.T) {
    q := os.Getenv("Q")
    os.ReadFile("/data/" + q) // SINK — test file
}
