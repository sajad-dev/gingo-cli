package newpr

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestNewpr_CloneProject(t *testing.T) {
// 	path, err := cloneProject("v0.1.0", "x")
// 	assert.NoError(t, err)
// 	_, err = os.Stat(fmt.Sprintf("%s/go.mod", path))
// 	assert.NoError(t, err)
// 	err = os.RemoveAll(path)
// 	assert.NoError(t,err)
// }

func TestNewpr_ChangeRepo(t *testing.T) {
	path, err := cloneProject("v0.1.0", "x")
	pattern := `github\.com\/sajad-dev\/gingo\/`

	patternM := `module github.com/sajad-dev/gingo`

	filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		match, err := regexp.MatchString(pattern, string(data))
		if err != nil {
			return err
		}

		matchm, err := regexp.MatchString(patternM, string(data))
		if err != nil {
			return err
		}

		if match || matchm {
			return fmt.Errorf("Found match with pattern")
		}

		return nil
	})
	assert.NoError(t, err)

	os.RemoveAll(path)
}
