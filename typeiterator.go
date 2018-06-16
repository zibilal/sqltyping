package sqltyping

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func TypeIterator(input interface{}, output interface{}) (err error) {

	ival := reflect.Indirect(reflect.ValueOf(input))
	ityp := ival.Type()
	oval := reflect.Indirect(reflect.ValueOf(output))
	otyp := oval.Type()

	checkTypes := !(oval.Type().String() == "bytes.Buffer")

	switch ival.Kind() {
	case reflect.Map:

		if !checkTypes {
			obuff := output.(*bytes.Buffer)
			obuff.WriteString(ityp.Name() + "\n")
		}

		if checkTypes && oval.Kind() != reflect.Map && oval.Kind() != reflect.Struct {
			err = fmt.Errorf("expecting output type of map or struct")
		} else {

			for _, k := range ival.MapKeys() {
				mival := ival.MapIndex(k)

				if oval.Kind() == reflect.Struct {

					foval := oval.FieldByName(k.String())

					if !foval.IsValid() {
						for s := 0; s < oval.NumField(); s++ {
							tmptyp := otyp.Field(s)

							tmptag := tmptyp.Tag
							if tmptag != "" {
								osplit := strings.Split(string(tmptag), "|")
								oosplit := strings.Split(osplit[1], ",")

								if k.String() == oosplit[0] {
									foval = oval.Field(s)
								}
							}
						}
					}

					if checkTypes && ( foval.Type().String() == "interface {}" || foval.Type().String() == mival.Type().String() ){
						foval.Set(mival)
					} else if !checkTypes {
						obuff := output.(*bytes.Buffer)
						fieldName := k.String()
						obuff.WriteString(fmt.Sprintf(" %s", fieldName))

						TypeIterator(mival.Interface(), obuff)
					}

				} else { // assumes output of type Map
					if ityp.Elem().String() == otyp.Elem().String() {
						oval.SetMapIndex(k, mival)
					} else {
						switch otyp.Elem().String() {
						case "int":
							if itmp, ok := mival.Interface().(int); ok {
								oval.SetMapIndex(k, reflect.ValueOf(itmp))
							} else if mival.Kind() == reflect.String {
								var itmp64 int64
								itmp64, err = strconv.ParseInt(mival.Interface().(string), 10, 64)
								if err != nil {
									return
								}
								oval.SetMapIndex(k, reflect.ValueOf(int(itmp64)))
							}
						case "int8":
							if itmp, ok := mival.Interface().(int8); ok {
								oval.SetMapIndex(k, reflect.ValueOf(itmp))
							} else if mival.Kind() == reflect.String {
								var itmp64 int64
								itmp64, err = strconv.ParseInt(mival.Interface().(string), 10, 64)
								if err != nil {
									return
								}
								oval.SetMapIndex(k, reflect.ValueOf(int8(itmp64)))
							}
						case "int16":
							if itmp, ok := mival.Interface().(int16); ok {
								oval.SetMapIndex(k, reflect.ValueOf(itmp))
							} else if mival.Kind() == reflect.String {
								var itmp64 int64
								itmp64, err = strconv.ParseInt(mival.Interface().(string), 10, 64)
								if err != nil {
									return
								}
								oval.SetMapIndex(k, reflect.ValueOf(int16(itmp64)))
							}
						case "int32":
							if itmp, ok := mival.Interface().(int32); ok {
								oval.SetMapIndex(k, reflect.ValueOf(itmp))
							} else if mival.Kind() == reflect.String {
								var itmp64 int64
								itmp64, err = strconv.ParseInt(mival.Interface().(string), 10, 64)
								if err != nil {
									return
								}
								oval.SetMapIndex(k, reflect.ValueOf(int32(itmp64)))
							}
						case "int64":
							if itmp, ok := mival.Interface().(int64); ok {
								oval.SetMapIndex(k, reflect.ValueOf(itmp))
							} else if mival.Kind() == reflect.String {
								var itmp64 int64
								itmp64, err = strconv.ParseInt(mival.Interface().(string), 10, 64)
								if err != nil {
									return
								}
								oval.SetMapIndex(k, reflect.ValueOf(itmp64))
							}
						case "uint":
							if itmp, ok := mival.Interface().(uint); ok {
								oval.SetMapIndex(k, reflect.ValueOf(itmp))
							} else if mival.Kind() == reflect.String {
								var itmp64 uint64
								itmp64, err = strconv.ParseUint(mival.Interface().(string), 10, 64)
								if err != nil {
									return
								}
								oval.SetMapIndex(k, reflect.ValueOf(uint(itmp64)))
							}
						case "uint8":
							if itmp, ok := mival.Interface().(uint8); ok {
								oval.SetMapIndex(k, reflect.ValueOf(itmp))
							} else if mival.Kind() == reflect.String {
								var itmp64 uint64
								itmp64, err = strconv.ParseUint(mival.Interface().(string), 10, 64)
								if err != nil {
									return
								}
								oval.SetMapIndex(k, reflect.ValueOf(uint8(itmp64)))
							}
						case "uint16":
							if itmp, ok := mival.Interface().(uint16); ok {
								oval.SetMapIndex(k, reflect.ValueOf(itmp))
							} else if mival.Kind() == reflect.String {
								var itmp64 uint64
								itmp64, err = strconv.ParseUint(mival.Interface().(string), 10, 64)
								if err != nil {
									return
								}
								oval.SetMapIndex(k, reflect.ValueOf(uint16(itmp64)))
							}
						case "uint32":
							if itmp, ok := mival.Interface().(uint32); ok {
								oval.SetMapIndex(k, reflect.ValueOf(itmp))
							} else if mival.Kind() == reflect.String {
								var itmp64 uint64
								itmp64, err = strconv.ParseUint(mival.Interface().(string), 10, 64)
								if err != nil {
									return
								}
								oval.SetMapIndex(k, reflect.ValueOf(uint32(itmp64)))
							}
						case "uint64":
							if itmp, ok := mival.Interface().(uint64); ok {
								oval.SetMapIndex(k, reflect.ValueOf(itmp))
							} else if mival.Kind() == reflect.String {
								var itmp64 uint64
								itmp64, err = strconv.ParseUint(mival.Interface().(string), 10, 64)
								if err != nil {
									return
								}
								oval.SetMapIndex(k, reflect.ValueOf(itmp64))
							}
						case "float32":
							if itmp, ok := mival.Interface().(float32); ok {
								oval.SetMapIndex(k, reflect.ValueOf(itmp))
							} else if mival.Kind() == reflect.String {
								var ftmp64 float64
								ftmp64, err = strconv.ParseFloat(mival.Interface().(string), 64)
								if err != nil {
									return
								}
								oval.SetMapIndex(k, reflect.ValueOf(float32(ftmp64)))
							}
						case "float64":
							if itmp, ok := mival.Interface().(float32); ok {
								oval.SetMapIndex(k, reflect.ValueOf(itmp))
							} else if mival.Kind() == reflect.String {
								var ftmp64 float64
								ftmp64, err = strconv.ParseFloat(mival.Interface().(string), 64)
								if err != nil {
									return
								}
								oval.SetMapIndex(k, reflect.ValueOf(ftmp64))
							}
						case "interface {}":
							oval.SetMapIndex(k, mival)
						default:
							if otyp.Elem().Kind() == reflect.Struct {
								vvtyp := reflect.New(otyp.Elem())
								eerr := TypeIterator(mival.Interface(), vvtyp.Interface())

								if eerr != nil {
									return eerr
								}

								oval.SetMapIndex(k, vvtyp.Elem())
							} else {
								err = errors.New("unsupported type pairs")
							}
						}
					}
				}
			}
		}

		if !checkTypes {
			obuff := output.(*bytes.Buffer)
			obuff.WriteString("\n")
		}

		return
	case reflect.Struct:

		if checkTypes && oval.Kind() != reflect.Struct {
			err = fmt.Errorf("expecting output type of struct")
		} else {
			if !checkTypes {
				ibuff := output.(*bytes.Buffer)
				ibuff.WriteString( " { \"table_name\": \"" + ityp.Name() + "\"")
			}

			for i := 0; i < ival.NumField(); i++ {

				fin := ival.Field(i)
				ftin := ityp.Field(i)

				if !checkTypes {

					ibuff := output.(*bytes.Buffer)
					itag := ftin.Tag

					fieldName := ""
					if itag != "" {
						isplit := strings.Split(string(itag), ":")
						if len(isplit) == 2 && isplit[1] != "" {
							isplit1 := strings.Replace(isplit[1], "\"", "", -1)
							iisplit := strings.Split(isplit1, ",")
							if len(iisplit) == 2 && iisplit[1] == "omitempty" {
								fieldName = ""
							} else {
								fieldName = iisplit[0]
							}
						}
					} else {
						fieldName = ftin.Name
					}
					if fieldName != "" {
						fieldName = strings.Replace(fieldName, "\"", "", -1)
						ibuff.WriteString(fmt.Sprintf(", \"column_name\": \"%v\"", fieldName))

						if fin.Type().String() == "time.Time" {
							dTime := fin.Interface().(time.Time)
							str := dTime.Format("2006-01-02 03:04:05")
							TypeIterator(str, ibuff)
						} else {
							TypeIterator(fin.Interface(), ibuff)
						}
					}


				} else {
					var fout reflect.Value
					var ftout reflect.StructField

					if fout = oval.FieldByName(ftin.Name); !fout.IsValid() {
						for j := 0; j < oval.NumField(); j++ {
							ftout = otyp.Field(j)
							if itag, otag := ftin.Tag, ftout.Tag; itag != "" && otag != "" {
								isplit := strings.Split(string(itag), "|")
								osplit := strings.Split(string(otag), "|")
								iisplit := strings.Split(isplit[1], ",")
								oosplit := strings.Split(osplit[1], ",")

								if len(isplit) == 2 && len(osplit) == 2 && iisplit[0] == oosplit[0] {
									fout = oval.Field(i)
								}
							}
						}
					}

					if fout.Kind() == reflect.Interface {
						fout.Set(fin)
					} else {
						if fout.IsValid() && fin.IsValid() {
							iout := reflect.New(fout.Type())
							TypeIterator(fin.Interface(), iout.Interface())
							fout.Set(iout.Elem())
						}
					}
				}

			}

			if !checkTypes {
				ibuff := output.(*bytes.Buffer)
				ibuff.WriteString(" }")
			}
		}

		return nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.String:

		if !checkTypes {
			ibuff := output.(*bytes.Buffer)
			ibuff.WriteString(fmt.Sprintf("|%v ", ival.Interface()))
		} else {
			oval.Set(ival)
		}

		return nil
	default:
		err = fmt.Errorf("unsupported type %T", input)
		return
	}

	return
}

// IsEmpty is an helper function to decide whether a value is empty or not
// This function is mean to be used to decide whether a struct variable is empty or not
func IsEmpty(t interface{}) bool {
	return reflect.DeepEqual(t, reflect.Zero(reflect.TypeOf(t)).Interface())
}
