package flag

import (
	"fmt"
)

func GetDefaultFlags() defaultFlags {
	return defaultFlags{
		prefix: "--",
		nameValueSeparator: "=",
		flags: map[string]string{},
	}
}

type defaultFlags struct { 
	prefix string
	nameValueSeparator string
	flags map[string]string
}

func (f *defaultFlags) parse(cmd []string) ([]string, error) {
	for k, v := range cmd {
		prefix, body := sepToPrefixAndBody(v, (*f).prefix)
		if prefix != (*f).prefix {
			return cmd[k:], nil // TODO: k or k+1
		}
		fKey, fValue, err := sepToKeyAndValue(body, f.nameValueSeparator)
		if err != nil {
			return cmd[k:], fmt.Errorf("Parsing flags, got %w", err) // TODO: k or k+1
		}
		(*f).flags[fKey] = fValue
	}
	return []string{}, nil // TODO: k or k+1
}

func sepToPrefixAndBody(flag, prefix string) (rPrefix, body string) {
	body = flag[len(prefix):]
	rPrefix = flag[:len(prefix)]
	return rPrefix, body
}


func sepToKeyAndValue(body, nameValueSeparator string) (key string, value string, err error) {
	counter := 0
	for i:=0;i<len(body);i++ {
		if body[i] == nameValueSeparator[counter] {
		} else {
			counter = 0
		}
		if counter == len(nameValueSeparator) {
			key = body[:i-counter]
			if len(key) == 0 {
				return "", "", fmt.Errorf("Key must be specified") // TODO: add custom error type with separator symbol, and body fields
			}
			value = body[i-counter:]
			return key, value, nil
		}
	}
	return body, "", nil
}
