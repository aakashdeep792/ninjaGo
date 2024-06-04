package env

import (
	"encoding/json"
	"ninjaGo/internal/logger"
	"os"
)

var skipURLList []string

const (
	SKIP_URL_LIST = "SKIP_URL_LIST"
	LOG_LEVEL     = "LOG_LEVEL"
	FORMAT        = `ERROR while loading environment variable %v:%v\n`
)

var (
	defaultSkipUrlList = []string{"health-check", "login", "sign-up"}
	defaultLogLevel    = logger.DEBUG
)

// GetSkipURLList return a copy of the skipURLList
func GetSkipURLList() []string {
	tmp := make([]string, len(skipURLList))
	copy(tmp, skipURLList)

	return tmp
}

func SetkipURLList() {
	tmp := os.Getenv(SKIP_URL_LIST)
	if tmp == "" {
		skipURLList = defaultSkipUrlList
		return
	}

	err := json.Unmarshal([]byte(tmp), &skipURLList)
	logger.Logf(FORMAT, SKIP_URL_LIST, err)
}

func SetLogLevel() {
	tmp := os.Getenv(LOG_LEVEL)
	if tmp == "" {
		tmp = defaultLogLevel

	}

	logger.UpdateLogLevel(tmp)
	logger.Logf("Setting LogLevel as %v\n", tmp)
}

func LoadEnvVariable() {
	SetkipURLList()
	SetLogLevel()
}

// func getORDefaultValue(key ,value string){

// }
