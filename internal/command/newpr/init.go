package newpr

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func cloneProject(version string, projectName string) (string, error) {
	cmd := exec.Command("git", "clone", "--branch", version, "--depth", "1", "https://github.com/sajad-dev/gingo", projectName)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	dir = fmt.Sprintf("%s/%s", dir, projectName)

	err = os.RemoveAll(fmt.Sprintf("%s/.git", dir))
	if err != nil {
		return "", err
	}

	return dir, nil
}

func changeRepo(path string, repo string) error {
	pattern := `github\.com\/sajad-dev\/gingo\/`
	replacement := fmt.Sprintf("%s/", repo)

	patternM := `module github.com/sajad-dev/gingo`
	replacementM := fmt.Sprintf("module %s", repo)

	re, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	reM, err := regexp.Compile(patternM)
	if err != nil {
		return err
	}

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

		var updated []byte
		
		if filepath.Ext(path) == ".mod" {
			updated = reM.ReplaceAll(data, []byte(replacementM))
		} else {
			updated = re.ReplaceAll(data, []byte(replacement))
		}

		if string(data) != string(updated) {
			err = os.WriteFile(path, updated, 0644)
			if err != nil {
				return err
			}
		}

		return nil
	})
	return nil
}

func Init(cmd *cobra.Command, args []string) error {
	version, err := cmd.Flags().GetString("version")
	if err != nil {
		color.Red(err.Error())
		return err
	}

	projectName, err := cmd.Flags().GetString("name")
	if err != nil {
		color.Red(err.Error())
		return err
	}

	repo, err := cmd.Flags().GetString("repo")
	if err != nil {
		color.Red(err.Error())
		return err
	}

	path, err := cloneProject(version, projectName)
	if err != nil {
		color.Red(err.Error())
		return err
	}

	err = changeRepo(path, repo)
	if err != nil {
		color.Red(err.Error())
		return err
	}

	color.Green(fmt.Sprintf("Create project on path : %s", path))

	return nil
}
