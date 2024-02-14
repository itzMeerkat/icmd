package git_repo_domain

import (
	"fmt"
	"github.com/itzmeerkat/icmd/infra/config"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func init() {
	repoPath := filepath.Join(config.GetAppFolder(), config.GetRepoName())
	if _, err := os.Stat(repoPath); os.IsNotExist(err) {
		clone()
	}
}

func clone() {
	cmd := exec.Command("git", "clone", config.GetRepoSSH())
	cmd.Dir = config.GetAppFolder()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func createCommitMessage() string {
	return fmt.Sprintf("Manual update at %s", time.Now().Local().Format("2006-01-02 15:04:05 MST"))
}

func UpdateRemote(newIndexContent string, wetRun bool) {
	if wetRun == false {
		fmt.Println(newIndexContent)
		return
	}
	indexPath := filepath.Join(config.GetGitRepoFolder(), "index.html")
	indexFile, err := os.OpenFile(indexPath, os.O_TRUNC|os.O_RDWR, 0755)
	if err != nil {
		panic(err)
	}

	_, err = indexFile.WriteString(newIndexContent)
	if err != nil {
		panic(err)
	}

	//println(cmt)
	cmd := exec.Command("git", "commit", "-am", createCommitMessage(), config.GetRepoSSH())
	cmd.Dir = config.GetGitRepoFolder()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	cmd = exec.Command("git", "push", "origin", "master")
	cmd.Dir = config.GetGitRepoFolder()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
