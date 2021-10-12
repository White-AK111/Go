// Package assignreflect
// Task#1
// Написать функцию, которая принимает на вход структуру in (struct или кастомную struct) и values map[string]interface{}
// (key - название поля структуры, которому нужно присвоить value этой мапы).
// Необходимо по значениям из мапы изменить входящую структуру in с помощью пакета reflect.
// Функция может возвращать только ошибку error. Написать к данной функции тесты (чем больше, тем лучше - зачтется в плюс).
package assignreflect

import (
	"errors"
	"reflect"
)

// in it's a struct with a different field types
type in struct {
	StructField struct {
		IntFieldInStruct       int
		InterfaceFieldInStruct interface{}
	}
	SliceIntField []int
	IntField      int
	FloatField    float64
	StringField   string
	BoolField     bool
}

// NewIn init new in struct
func NewIn() *in {
	return &in{
		StructField: struct {
			IntFieldInStruct       int
			InterfaceFieldInStruct interface{}
		}{IntFieldInStruct: 0, InterfaceFieldInStruct: nil},
		SliceIntField: make([]int, 8),
		IntField:      0,
		FloatField:    0.00,
		StringField:   "",
		BoolField:     false,
	}
}

// AssignmentByReflect function assigned values in struct from map by reflect pkg, return error
func AssignmentByReflect(in interface{}, values map[string]interface{}) error {
	// check that in struct is not nil
	if in == nil {
		return errors.New("nil in input interface")
	}

	// get value struct
	reflectVal := reflect.Indirect(reflect.ValueOf(in))

	// check that kind of in struct is struct
	if reflectVal.Kind() != reflect.Struct {
		return errors.New("in input interface isn't a struct")
	}

	// assign value for all field in struct
	for i := 0; i < reflectVal.Type().NumField(); i++ {
		fldStructByIndex := reflectVal.Type().Field(i)
		fldMap := values[fldStructByIndex.Name]
		fldStructByName := reflectVal.FieldByName(fldStructByIndex.Name)
		if fldStructByName.Type().AssignableTo(reflect.TypeOf(fldMap)) {
			fldStructByName.Set(reflect.ValueOf(fldMap))
		} else {
			return errors.New("map has an unassignable field")
		}
	}

	return nil
}
