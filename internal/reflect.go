package internal

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true"`
	Age  int    `required:"true"`
}

func GetReflectType() {
	sample := Sample{"test", 10}
	sampleType := reflect.TypeOf(sample)

	fmt.Println(sampleType.Field(0).Name)
}

func (s Sample) IsValid() bool {
	t := reflect.TypeOf(s)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(s).Field(i).Interface() != ""
		}
	}
	return true
}
