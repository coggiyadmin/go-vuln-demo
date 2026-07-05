package attacksurface
import ("os"; "testing")
func TestRender(t *testing.T) {
    q := os.Getenv("Q")
    _ = "<p>" + q + "</p>" // SINK — test file
}
