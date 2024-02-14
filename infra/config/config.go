package config

import (
	"github.com/spf13/viper"
	"os"
	"os/user"
	"path/filepath"
)

var folderPath string

func init() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	folderPath = filepath.Join(u.HomeDir, ".icmd")
	err = os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(folderPath)
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}

func GetAppFolder() string {
	return folderPath
}

func GetRepoName() string {
	return viper.GetString("repo_name")
}

func GetRepoSSH() string {
	return viper.GetString("repo_ssh")
}

func GetGitRepoFolder() string {
	return filepath.Join(GetAppFolder(), GetRepoName())
}
