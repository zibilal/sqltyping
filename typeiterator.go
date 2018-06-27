package sqltyping

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
	"text/scanner"
)

func TypeIterator(input interface{}, output interface{}, customValues ...func(interface{}) (interface{}, error)) (err error) {

	if input == nil || output == nil {
		return nil
	}

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

		if checkTypes && oval.Kind() != reflect.Map && oval.Kind() != reflect.Struct && oval.Kind() != reflect.Ptr {
			err = fmt.Errorf("expecting output type of map or struct")
		} else if checkTypes {

			for _, k := range ival.MapKeys() {
				mival := ival.MapIndex(k)

				if oval.Kind() == reflect.Ptr {
					if oval.IsNil() {
						tmpOval := reflect.New(oval.Type().Elem())
						oval.Set(tmpOval)
						oval = oval.Elem()
						otyp = otyp.Elem()
					}
				}

				if oval.Kind() == reflect.Struct {
					foval := oval.FieldByName(k.String())
					if !foval.IsValid() {
						for s := 0; s < oval.NumField(); s++ {
							oftype := otyp.Field(s)
							oftag := string(oftype.Tag)
							ofsplit := strings.Split(oftag, " ")

							for _, split := range ofsplit {
								osplit := strings.Split(split, ":")
								if len(osplit) == 2 {
									oosplit := strings.Replace(strings.Split(osplit[1], ",")[0], "\"", "", -1)
									if k.String() == oosplit {
										foval = oval.Field(s)
										break
									}
								}
							}

						}
					}
					if !foval.IsValid() {
						continue
					}

					if istr, ok := mival.Interface().(string); ok && foval.Kind() == reflect.String {
						foval.Set(reflect.ValueOf(istr))
					} else if iint, ok := mival.Interface().(int); ok && foval.Kind() == reflect.Int {
						foval.Set(reflect.ValueOf(iint))
					} else if iint8, ok := mival.Interface().(int8); ok && foval.Kind() == reflect.Int8 {
						foval.Set(reflect.ValueOf(iint8))
					} else if iint16, ok := mival.Interface().(int16); ok && foval.Kind() == reflect.Int16 {
						foval.Set(reflect.ValueOf(iint16))
					} else if iint32, ok := mival.Interface().(int32); ok && foval.Kind() == reflect.Int32 {
						foval.Set(reflect.ValueOf(iint32))
					} else if iint64, ok := mival.Interface().(int64); ok && foval.Kind() == reflect.Int64 {
						foval.Set(reflect.ValueOf(iint64))
					} else if ifloat32, ok := mival.Interface().(float32); ok && foval.Kind() == reflect.Float32 {
						foval.Set(reflect.ValueOf(ifloat32))
					} else if ifloat64, ok := mival.Interface().(float64); ok && foval.Kind() == reflect.Float64 {
						foval.Set(reflect.ValueOf(ifloat64))
					} else if foval.Type().String() == mival.Type().String() {
						foval.Set(mival)
					} else if mival.Kind() == reflect.Interface {
						var isHandled bool
						for _, c := range customValues {
							result, resultError := c(mival.Interface())
							if resultError == nil && reflect.ValueOf(result).Type().String() == foval.Type().String() {
								foval.Set(reflect.ValueOf(result))
								isHandled = true
							}
						}

						if !isHandled {
							elemival := reflect.Indirect(mival.Elem())

							if elemival.Kind() == reflect.Slice && foval.Kind() == reflect.Slice {
								mSlice := reflect.MakeSlice(foval.Type(), 0, elemival.Len())
								for idx := 0; idx < elemival.Len(); idx++ {
									theOutput := reflect.New(foval.Type().Elem())

									tmpelemival := elemival.Index(idx)
									if elemival.Index(idx).Kind() == reflect.Interface {
										tmpelemival = elemival.Index(idx).Elem()
									}
									err = TypeIterator(tmpelemival.Interface(), theOutput.Interface(), customValues...)
									mSlice = reflect.Append(mSlice, theOutput.Elem())
								}
								foval.Set(mSlice)
							} else {
								pval := reflect.Indirect(mival.Elem())
								err = TypeIterator(pval.Interface(), foval.Addr().Interface(), customValues...)
								if err != nil {
									return
								}
							}
						}
					} else if foval.Type().String() == mival.Type().String() {
						foval.Set(mival)
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
								eerr := TypeIterator(mival.Interface(), vvtyp.Interface(), customValues...)

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
				ibuff.WriteString("{table_name:" + ityp.Name() + "")
			}

			for i := 0; i < ival.NumField(); i++ {

				fin := ival.Field(i)
				ftin := ityp.Field(i)

				if !checkTypes {

					ibuff := output.(*bytes.Buffer)
					itag := ftin.Tag

					fieldName := ""
					if itag != "" {
						// always find the first set of tag
						bySpace := strings.Split(string(itag), " ")
						for j := 0; j < len(bySpace); j++ {
							isplit := strings.Split(bySpace[j], ":")
							if len(isplit) == 2 && isplit[1] != "" {
								isplit1 := strings.Replace(isplit[1], "\"", "", -1)
								iisplit := strings.Split(isplit1, ",")
								if len(iisplit) == 2 && iisplit[1] == "omitempty" {
									fieldName = ""
								} else {
									fieldName = iisplit[0]
								}
							}
						}
					} else {
						fieldName = ftin.Name
					}
					if fieldName != "" {
						fieldName = strings.Replace(fieldName, "\"", "", -1)
						if fin.Type().String() == "time.Time" {
							dTime := fin.Interface().(time.Time)
							if !IsEmpty(dTime) {
								ibuff.WriteString(fmt.Sprintf(",column_name:%v", fieldName))
								str := dTime.Format("2006-01-02 03:04:05")
								err = TypeIterator(str, ibuff, customValues...)
								if err != nil {
									return
								}
							}
						} else {
							ibuff.WriteString(fmt.Sprintf(",column_name:%v", fieldName))
							err = TypeIterator(fin.Interface(), ibuff, customValues...)
							if err != nil {
								return
							}
						}
					}

				} else {
					var fout reflect.Value
					var ftout reflect.StructField

					if fout = oval.FieldByName(ftin.Name); !fout.IsValid() {
						for j := 0; j < oval.NumField(); j++ {
							ftout = otyp.Field(j)
							if itag, otag := ftin.Tag, ftout.Tag; itag != "" && otag != "" {
								var scanner1 scanner.Scanner
								scanner1.Init(strings.NewReader(string(itag)))

								for tok := scanner1.Scan(); tok != scanner.EOF; tok = scanner1.Scan() {
									switch tok {
									case scanner.String:
										text := scanner1.TokenText()
										if strings.Contains(string(otag), text) {
											fout = oval.Field(j)
											break
										}
									}
								}
							}
						}
					}

					if !fout.IsValid() {
						continue
					}

					if fout.Kind() == reflect.Interface {
						fout.Set(fin)
					} else if fin.Type().String() == "time.Time" {
						if fout.Type().String() == "time.Time" {
							fout.Set(fin)
						} else if fout.Kind() == reflect.String {
							dTime := fin.Interface().(time.Time)
							str := dTime.Format("2006-01-02 03:04:05")
							fout.Set(reflect.ValueOf(str))
						}
					} else {
						if fout.IsValid() && fin.IsValid() {
							iout := reflect.New(fout.Type())
							err = TypeIterator(fin.Interface(), iout.Interface(), customValues...)
							if err != nil {
								return
							}
							fout.Set(iout.Elem())
						}
					}
				}

			}

			if !checkTypes {
				ibuff := output.(*bytes.Buffer)
				ibuff.WriteString("}")
			}
		}

		return nil
	case reflect.Slice:
		if !checkTypes {
			obuff := output.(*bytes.Buffer)
			for i := 0; i < ival.Len(); i++ {
				iItem := ival.Index(i)
				err = TypeIterator(iItem.Interface(), obuff, customValues...)
				if err != nil {
					return
				}
			}
		} else if oval.Kind() == reflect.Interface {
			oval.Set(ival)
		} else if oval.Kind() == reflect.Slice {
			outSlice := reflect.MakeSlice(reflect.SliceOf(otyp.Elem()), 0, ival.Len())
			for i := 0; i < ival.Len(); i++ {

				oItem := reflect.New(otyp.Elem())
				iItem := ival.Index(i)
				err = TypeIterator(iItem.Interface(), oItem.Interface(), customValues...)
				if err != nil {
					return
				}
				outSlice = reflect.Append(outSlice, oItem.Elem())
			}
			oval.Set(outSlice)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.String:

		if !checkTypes {
			ibuff := output.(*bytes.Buffer)
			ibuff.WriteString(fmt.Sprintf("|%v", ival.Interface()))
		} else {
			if oval.Kind() == ival.Kind() {
				oval.Set(ival)
			} else if ival.Kind() == reflect.Float64 {
				switch oval.Kind() {
				case reflect.Int:
					if dval, ok := ival.Interface().(int); ok {
						oval.Set(reflect.ValueOf(dval))
					}
				case reflect.Int8:
					if dval, ok := ival.Interface().(int8); ok {
						oval.Set(reflect.ValueOf(dval))
					}
				case reflect.Int16:
					if dval, ok := ival.Interface().(int16); ok {
						oval.Set(reflect.ValueOf(dval))
					}
				case reflect.Int32:
					if dval, ok := ival.Interface().(int32); ok {
						oval.Set(reflect.ValueOf(dval))
					}
				case reflect.Int64:
					if dval, ok := ival.Interface().(int64); ok {
						oval.Set(reflect.ValueOf(dval))
					}
				case reflect.Uint:
					if dval, ok := ival.Interface().(uint); ok {
						oval.Set(reflect.ValueOf(dval))
					}
				case reflect.Uint8:
					if dval, ok := ival.Interface().(uint8); ok {
						oval.Set(reflect.ValueOf(dval))
					}
				case reflect.Uint16:
					if dval, ok := ival.Interface().(uint16); ok {
						oval.Set(reflect.ValueOf(dval))
					}
				case reflect.Uint32:
					if dval, ok := ival.Interface().(uint32); ok {
						oval.Set(reflect.ValueOf(dval))
					}
				case reflect.Uint64:
					if dval, ok := ival.Interface().(uint64); ok {
						oval.Set(reflect.ValueOf(dval))
					}
				case reflect.Float32:
					if dval, ok := ival.Interface().(float32); ok {
						oval.Set(reflect.ValueOf(dval))
					}
				}
			} else {
				err = fmt.Errorf("input type %T cannot by set in output of type %T", ival.Interface(), oval.Interface())
				return
			}
		}

		return nil
	case reflect.Interface:
		pival := ival.Elem()
		err = TypeIterator(pival.Interface(), output, customValues...)
		if err != nil {
			return
		}

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
