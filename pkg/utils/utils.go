package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func Pprint(v any) {
	prettyString, _ := json.MarshalIndent(v, "", "  ")
	fmt.Printf("%s: %s\n", reflect.TypeOf(v).Name(), prettyString)
}

func GetTypeName(v any) string {
	return reflect.TypeOf(v).Name()
}
