package pathxf

import "os"

const base = "/srv/app/data/"

func Run(name string) ([]byte, error) { return os.ReadFile(base + name) } // SINK CWE-22
