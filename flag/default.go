package flag

import (
	"fmt"
)

func DefaultFlags(prefix, nameValueSeparator string, posargs []string) defaultFlags {
	return defaultFlags{
		prefix: prefix, 
		nameValueSeparator: nameValueSeparator, 
		kwargs: kwargs{},
		posargs: posargs,
	}
}

func GetDefaultFlags() defaultFlags {
	return defaultFlags{
		prefix: "--",
		nameValueSeparator: "=",
		kwargs: kwargs{},
	}
}

type defaultFlags struct { 
	prefix string
	nameValueSeparator string
	kwargs kwargs
	posargs []string
}

func (d defaultFlags) Extract() (map[string][]struct{value string; position int}, []string) {
	return d.kwargs.Publish(), d.posargs
}

func (d1 defaultFlags) isEqualTo(d2 defaultFlags) bool {
	if d1.prefix != d2.prefix {
		return false
	}
	if d1.nameValueSeparator != d2.nameValueSeparator {
		return false
	}
	if !d1.kwargs.isEqualTo(d2.kwargs) {
		return false
	}
	return true
}

type kwargs map[string][]pair

func (k kwargs) Publish() map[string][]struct{value string; position int} {
	newMap := map[string][]struct{value string; position int}{}
	for k,v := range(k) {
		list := []struct{value string; position int}{}
		for _,v2 := range(v) {
			list = append(list, struct{value string; position int}{value: v2.value, position: v2.position})
		}
		newMap[k] = list
	}
	return newMap
}

func (k1 kwargs) isEqualTo(k2 kwargs) bool {
	for key1, value1 := range(k1) {
		value2, ok := k2[key1] 
		if !ok {
			return false
		}
		if len(value2) != len(value1) {
			return false
		}
		for i:=0;i<len(value1);i++ {
			if !value2[i].isEqualTo(value1[i]) {
				return false
			}
		}
	}
	return true
}

type pair struct{
	value string
	position int
}

func (p1 pair) isEqualTo(p2 pair) bool {
	if p1.value != p2.value {
		return false
	}
	if p1.position != p2.position {
		return false
	}
	return true
}

func (f *defaultFlags) Parse() error {
	for k, v := range f.posargs{
		prefix, body := sepToPrefixAndBody(v, (*f).prefix)
		if prefix != (*f).prefix {
			f.posargs = f.posargs[k:]// TODO: should i update it here, or higher, on every iteration ?
			return nil 
		}
		fKey, fValue, err := sepToKeyAndValue(body, f.nameValueSeparator)
		if err != nil {
			f.posargs = f.posargs[k:]// TODO: should i update it here, or higher, on every iteration ?
			return fmt.Errorf("Parsing flags, got %w", err) // TODO: k or k+1
		}
		kwarg, kwargExists := (*f).kwargs[fKey]
		flagPair := pair{
			value: fValue,
			position: k,
		}
		if kwargExists {
			(*f).kwargs[fKey] = append(kwarg, flagPair)
		} else {
			(*f).kwargs[fKey] = []pair{flagPair}} 
	}

	f.posargs = []string{}// TODO: should i update it here, or higher, on every iteration ?
	return  nil 
}

func sepToPrefixAndBody(flag, prefix string) (rPrefix, body string) {
	if len(prefix) > len(flag) {
		return "", flag
	}
	body = flag[len(prefix):]
	rPrefix = flag[:len(prefix)]
	return rPrefix, body
}


func sepToKeyAndValue(body, nameValueSeparator string) (key string, value string, err error) {
	counter := 0
	for i:=0;i<len(body);i++ {
		if body[i] == nameValueSeparator[counter] {
			counter++
		} else {
			counter = 0
		}
		if counter == len(nameValueSeparator) {
			key = body[:i-counter+1]
			if len(key) == 0 {
				return "", "", fmt.Errorf("Key must be specified") // TODO: add custom error type with separator symbol, and body fields
			}
			if len(key) < len(body) {
				value = body[i+1:]
			} else {
				value = ""
			}
			return key, value, nil
		}
	}
	return body, "", nil
}
