package config

import (
	"log"
	"os"
	"path/filepath"
)

type PomoCliConfig interface {
	GetDataStoreFileCompletePath() string
	GetDataStoreFileFolderPath() string
}

type pomoCliConfig struct {
	homeDir    string
	dbFileName string
}

var pomoClientConfig PomoCliConfig

func GetConfig() PomoCliConfig {
	return pomoClientConfig
}

func InitConfig() PomoCliConfig {

	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal("Unable to get home directory:", err)
	}

	pomoClientConfig = &pomoCliConfig{
		dbFileName: "pomocli.db",
		homeDir:    homeDir,
	}

	return pomoClientConfig
}

func (pc *pomoCliConfig) GetDataStoreFileCompletePath() string {
	return filepath.Join(pc.GetDataStoreFileFolderPath(), pc.dbFileName)
}

func (pc *pomoCliConfig) GetDataStoreFileFolderPath() string {
	return filepath.Join(pc.homeDir, ".local", "share", "pomocli")
}
