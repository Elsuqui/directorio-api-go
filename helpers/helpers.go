package helpers

import "os"

// Get environment parameter using a key, if not exists it returns the default value provide it
func GetEnvParamDefault(key string, defaultValue string) string {
	if defaultValue == "" {
		panic("Default value must not be an empty value")
	}
	value, exists := os.LookupEnv(key)
	if (exists && value == "") || (!exists) {
		value = defaultValue
	}
	return value
}

// Get environment parameter, if you need set a default value consider use GetEnvParamDefault
func GetEnvParam(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic("Env variable does not exist")
	}
	return value
}

func HandleError(err error, message string) {
	if err != nil {
		panic(message)
	}
}
