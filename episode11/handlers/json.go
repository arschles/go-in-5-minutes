package handlers

import (
	"fmt"
)

func jsonErr(err error, descr string) string {
	return fmt.Sprintf(`{"error":"%s", "description":"%s"}`, err.Error(), descr)
}

func jsonErrStr(descr string) string {
	return fmt.Sprintf(`{"error":"%s"}`, descr)
}

func jsonKVP(key, val string) string {
	return fmt.Sprintf(`{"%s":"%s"}`, key, val)
}
