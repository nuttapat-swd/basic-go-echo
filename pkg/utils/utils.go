package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func GetTypeName(v any) string {
	if v == nil {
		return ""
	}
	return reflect.TypeOf(v).Name()
}

func Pprint(v any) {
	prettyString, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Printf("error pretty-printing %s: %v\n", reflect.TypeOf(v).Name(), err)
		return
	}
	fmt.Printf("%s: %s\n", GetTypeName(v), prettyString)
}
