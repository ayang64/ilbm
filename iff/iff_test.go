package iff

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
)

func TestDecode(t *testing.T) {
	filepath.Walk("../test-images",
		func(p string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			if ext := strings.ToLower(path.Ext(p)); ext != ".iff" {
				return nil
			}

			t.Logf("trying %q", p)
			return nil
		})
}
