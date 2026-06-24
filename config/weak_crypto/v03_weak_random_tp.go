package main

import "math/rand"

func Token() int { return rand.Intn(1_000_000) } // SINK CWE-330
