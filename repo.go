package main

import (
	"os"
	"path"
)

const gitDir = ".git"

type Repo struct {
	WordDir string
	GitDir  string
}

func OpenRepo() (Repo, error) {
	wd, err := os.Getwd()
	if err != nil {
		return Repo{}, err
	}

	_, err = os.Stat(path.Join(wd, gitDir))
	if err != nil {
		return Repo{}, err
	}

	return Repo{
		WordDir: wd,
		GitDir:  path.Join(wd, gitDir),
	}, nil
}
