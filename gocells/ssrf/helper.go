package ssrfxf

import "net/http"

func Run(url string) (*http.Response, error) { return http.Get(url) } // SINK CWE-918
