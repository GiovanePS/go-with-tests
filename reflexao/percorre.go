package reflexao

import "reflect"

func percorre(x interface{}, fn func(input string)) {
	value := getValue(x)

	numberOfValues := 0
	var getField func(int) reflect.Value

	switch value.Kind() {
	case reflect.String:
		fn(value.String())

	case reflect.Struct:
		numberOfValues = value.NumField()
		getField = value.Field

	case reflect.Slice:
		numberOfValues = value.Len()
		getField = value.Index

	case reflect.Map:
		iter := value.MapRange()
		for iter.Next() {
			percorre(iter.Value().Interface(), fn)
		}
	}

	for i := 0; i < numberOfValues; i++ {
		percorre(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	value := reflect.ValueOf(x)

	switch value.Kind() {
	case reflect.Pointer:
		return value.Elem()
	}

	return value
}
