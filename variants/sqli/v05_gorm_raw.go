package main

import (
    "net/http"
    "gorm.io/gorm"
)

func gormRaw(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
    id := r.FormValue("id")
    db.Raw("SELECT * FROM users WHERE id=" + id).Scan(&[]map[string]any{}) // SINK CWE-89
}
