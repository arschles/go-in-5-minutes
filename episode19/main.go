package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	if js, err := unmarshal(aDict); err != nil {
		logger.Fatalf("error decoding (%s)", err)
	} else {
		if str, err := ifaceToString(js); err != nil {
			logger.Fatalf("error converting js (%s)", err)
		} else {
			logger.Println(str)
		}
	}

	if js, err := unmarshal(aList); err != nil {
		logger.Fatalf("error decoding (%s)", err)
	} else {
		if str, err := ifaceToString(js); err != nil {
			logger.Fatalf("error converting js (%s)", err)
		} else {
			logger.Println(str)
		}
	}

	if js, err := unmarshal(anInt); err != nil {
		logger.Fatalf("error decoding (%s)", err)
	} else {
		if str, err := ifaceToString(js); err != nil {
			logger.Fatalf("error converting js (%s)", err)
		} else {
			logger.Println(str)
		}
	}

	if js, err := unmarshal(aString); err != nil {
		logger.Fatalf("error decoding (%s)", err)
	} else {
		if str, err := ifaceToString(js); err != nil {
			logger.Fatalf("error converting js (%s)", err)
		} else {
			logger.Println(str)
		}
	}
}

func unmarshal(str string) (interface{}, error) {
	var iface interface{}
	decoder := json.NewDecoder(strings.NewReader(str))
	decoder.UseNumber()
	if err := decoder.Decode(&iface); err != nil {
		return nil, err
	}
	return iface, nil
}

func ifaceToString(iface interface{}) (string, error) {
	switch t := iface.(type) {
	case map[string]interface{}:
		strs := make([]string, len(t))
		i := 0
		for key, val := range t {
			str, err := ifaceToString(val)
			if err != nil {
				return "", err
			}
			strs[i] = fmt.Sprintf("%s: %s", key, str)
			i++
		}
		return "{" + strings.Join(strs, ", ") + "}", nil
	case []interface{}:
		strs := make([]string, len(t))
		i := 0
		for _, val := range t {
			str, err := ifaceToString(val)
			if err != nil {
				return "", err
			}
			strs[i] = str
			i++
		}
		return "[" + strings.Join(strs, ", ") + "]", nil
	case int:
		return fmt.Sprintf("%d", t), nil
	case json.Number:
		return fmt.Sprintf("%s", t), nil
	case string:
		return fmt.Sprintf(`"%s"`, t), nil
	}
	return "", fmt.Errorf("unsupported value %#v (%T)", iface, iface)
}
