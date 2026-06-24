package main

import "crypto/md5"

func HashMD5(data []byte) [16]byte { return md5.Sum(data) } // SINK CWE-328
