package chlog

import "path"

//Config for the change log
type Config struct {
	DataDirPath    string
	LogFileDirName string
}

//DefaultConfig for Writer
func DefaultConfig() Config {
	return Config{
		DataDirPath:    "data",
		LogFileDirName: "change_logs",
	}
}

func (c Config) logFilesPath() string {
	return path.Join(c.DataDirPath, c.LogFileDirName)
}

func (c Config) logFilePath(fileName string) string {
	return path.Join(c.logFilesPath(), fileName)
}
