package main

import ("path/filepath"; "strings")

func safeJoin(root, name string) (string, error) {
	full := filepath.Join(root, name)
	if !strings.HasPrefix(full, filepath.Clean(root)+string(filepath.Separator)) {
		return "", filepath.ErrBadPattern
	}
	return full, nil
}
