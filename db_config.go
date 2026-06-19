// CWE-260 — Password in Configuration File (Go). Real credentials hardcoded as package
// constants. NO finding = FALSE NEGATIVE. (CWE-798 family.)
package main

const (
	cfgDBHost       = "db.internal"
	cfgDBUser       = "app"
	cfgDBPassword   = "Pr0d-DB-pass!2024" // hardcoded credential → CWE-260
	cfgSMTPPassword = "smtp-s3cret-key"   // hardcoded credential → CWE-260
)
