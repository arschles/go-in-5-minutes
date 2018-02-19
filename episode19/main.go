package main

import (
	"encoding/json"
	"fmt"
	"strings"

	gson "github.com/bitly/go-simplejson"
)

func main() {
	js, err := gson.NewJson([]byte(aDict))
	if err != nil {
		logger.Fatalf("error decoding (%s)", err)
	}
	str, err := jsonToString(js)
	if err != nil {
		logger.Fatalf("error converting js to string (%s)", err)
	}
	logger.Println(str)

	js, err = gson.NewJson([]byte(aList))
	if err != nil {
		logger.Fatalf("error decoding (%s)", err)
	}
	str, err = jsonToString(js)
	if err != nil {
		logger.Fatalf("error converting js to string (%s)", err)
	}
	logger.Println(str)

	js, err = gson.NewJson([]byte(anInt))
	if err != nil {
		logger.Fatalf("error decoding (%s)", err)
	}
	str, err = jsonToString(js)
	if err != nil {
		logger.Fatalf("error converting to string (%s)", err)
	}
	logger.Println(str)

	js, err = gson.NewJson([]byte(aString))
	if err != nil {
		logger.Fatalf("error decoding (%s)", err)
	}
	str, err = jsonToString(js)
	if err != nil {
		logger.Fatalf("error converting js to string (%s)", err)
	}
	logger.Println(str)
}

func jsonToString(js *gson.Json) (string, error) {
	if dict, err := js.Map(); err == nil {
		return ifaceToString(dict), nil
	} else if arr, err := js.Array(); err == nil {
		return ifaceToString(arr), nil
	} else if str, err := js.String(); err == nil {
		return ifaceToString(str), nil
	} else if num, err := js.Int(); err == nil {
		return ifaceToString(num), nil
	} else if num, err := js.Int64(); err == nil {
		return ifaceToString(num), nil
	}

	return "", fmt.Errorf("unsupported js object %#v (%T)", js, js)
}

func ifaceToString(iface interface{}) string {
	switch t := iface.(type) {
	case map[string]interface{}:
		strs := make([]string, len(t))
		i := 0
		for key, val := range t {
			strs[i] = fmt.Sprintf("%s: %s", key, ifaceToString(val))
			i++
		}
		return "{" + strings.Join(strs, ", ") + "}"
	case []interface{}:
		strs := make([]string, len(t))
		i := 0
		for _, val := range t {
			strs[i] = ifaceToString(val)
			i++
		}
		return "[" + strings.Join(strs, ", ") + "]"
	case int:
		return fmt.Sprintf("%d", t)
	case json.Number:
		return fmt.Sprintf("%s", t)
	case string:
		return fmt.Sprintf(`"%s"`, t)
	}
	return fmt.Sprintf("unsupported value %#v (%T)", iface, iface)
}
