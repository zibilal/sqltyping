package sqltyping

import (
	"reflect"
	"errors"
)

// CopyValue is used to copy partial value of input struct type to the output struct type, this function is used to make
// value copy between two different but related struct convenience
// input is the original struct value
// output is the destination struct value, output should be on type of pointer of the struct
func CopyValue(input interface{}, output interface{}) error {
	ivalue := reflect.Indirect(reflect.ValueOf(input))
	ovalue := reflect.Indirect(reflect.ValueOf(output))

	if ivalue.Kind() != reflect.Struct || ovalue.Kind() != reflect.Struct {
		return errors.New("expected input is of type struct, and output is of type addressable struct")
	}

	otype := ovalue.Type()
	for i:=0; i<ovalue.NumField(); i++ {
		ft := otype.Field(i)
		result := searchValue(input, ft.Name, "", nil)
		if result != nil {
			rvalue := reflect.ValueOf(result)
			if rvalue.Type() == ft.Type {
				fv := ovalue.Field(i)
				fv.Set(rvalue)
			}
		}
	}

	return nil
}

func searchValue(v interface{}, name1, name2 string, result interface{}) interface{} {

	switch reflect.ValueOf(v).Kind() {
	case reflect.Struct:
		ival := reflect.ValueOf(v)
		ityp := ival.Type()
		for i := 0; i < ival.NumField(); i++ {
			fv := ival.Field(i)
			ft := ityp.Field(i)
			result = searchValue(fv.Interface(), name1, ft.Name, result)
		}
	default:
		if name1 == name2 {
			result = v
		}
	}

	return result
}
