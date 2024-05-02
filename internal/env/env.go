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
	defaultSkipUrlList = []string{"login"}
	defaultLOgLevel = logger.INFO

)

// GetSkipURLList return a copy of the skipURLList
func GetSkipURLList() []string {
	var tmp []string
	copy(tmp, skipURLList)
	return tmp
}

func SetLogLevel()string{

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
		tmp = defaultLOgLevel
		return
	}

	// err := json.Unmarshal([]byte(tmp), &skipURLList)
	// logger.Logf(FORMAT, SKIP_URL_LIST, err)
}

func LoadEnvVariable() {
	SetkipURLList()
}

// func getORDefaultValue(key ,value string){

// }
