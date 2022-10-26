package numbertools

import "reflect"

func IsInt(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Int
}
