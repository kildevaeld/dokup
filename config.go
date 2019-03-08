package dokup

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/kildevaeld/dict"
)

type validator func(st reflect.Value, m dict.Map) error

type creator struct {
	v []validator
	t reflect.Type
}

type ConfigEntryCreator interface {
	Create(m map[string]interface{}) (interface{}, error)
}

func (c *creator) Create(m map[string]interface{}) (interface{}, error) {
	i := reflect.New(c.t)
	for _, v := range c.v {
		if err := v(i, m); err != nil {
			return nil, err
		}
	}
	return i.Interface(), nil
}

func to_string(fieldName string) func(st reflect.Value, m dict.Map) error {

	return func(st reflect.Value, m dict.Map) error {
		field := st.Elem().FieldByName(fieldName)

		value := m.Get(fieldName)
		if value == nil {
			return nil
		}

		if str, ok := value.(string); ok {
			field.SetString(str)
			return nil
		}

		return errors.New("not a string")
	}
}

func to_slice(fieldName string, valueR reflect.Type) func(st reflect.Value, m dict.Map) error {
	return func(st reflect.Value, m dict.Map) error {
		field := st.Elem().FieldByName(fieldName)
		val := m.Get(fieldName)
		if val == nil {
			return nil
		}

		sliceType := reflect.SliceOf(field.Type().Elem())
		var slice reflect.Value

		switch t := val.(type) {
		case string:
			if valueR.Kind() != reflect.String {
				return errors.New("cannot set string")
			}
			split := strings.Split(strings.TrimSpace(t), "\n")

			slice = reflect.MakeSlice(sliceType, len(split)-1, len(split))

			for _, str := range split {
				if str == "" {
					continue
				}
				slice = reflect.Append(slice, reflect.ValueOf(str))
			}
		case []string:
			if valueR.Kind() != reflect.String {
				return errors.New("cannot set stringslice")
			}
			slice = reflect.ValueOf(t)
		case []interface{}:
			slice = reflect.MakeSlice(sliceType, len(t)-1, len(t))
			for _, r := range t {
				destType := reflect.TypeOf(r)

				if valueR.Kind() != destType.Kind() {
					fmt.Printf("not same type")
				}
				slice = reflect.Append(slice, reflect.ValueOf(r))
			}

		}
		if !slice.IsValid() {
			return nil
		}
		field.Set(slice)
		return nil
	}
}

func to_struct(fieldName string) validator {
	return nil
}

func to_dict(fieldName string) validator {
	return nil
}

func to_bool(fieldName string) validator {
	return func(v reflect.Value, m dict.Map) error {
		if val := m.Get(fieldName); val != nil {
			b, ok := val.(bool)
			if !ok {
				return errors.New("not a boolean")
			}
			v.Elem().FieldByName(fieldName).SetBool(b)
		}

		return nil
	}
}

func field_creator(val reflect.Value, typ reflect.StructField) validator {
	switch val.Kind() {
	case reflect.String:
		return to_string(typ.Name)
	case reflect.Struct:
		return to_struct(typ.Name)
	case reflect.Slice:
		return to_slice(typ.Name, typ.Type.Elem())
	case reflect.Bool:
		return to_bool(typ.Name)
	}
	return nil
}

func RegisterConfigEntry(name string, e interface{}) error {

	if entry, ok := e.(ConfigEntryCreator); ok {
		configEntries[name] = entry
		return nil
	}

	v := reflect.ValueOf(e)
	t := reflect.TypeOf(e)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}

	if v.Kind() != reflect.Struct {
		return errors.New("not a struct")
	}

	c := &creator{
		t: t,
	}
	for i := 0; i < v.NumField(); i++ {
		valueField := v.Field(i)
		typeField := v.Type().Field(i)

		val := field_creator(valueField, typeField)
		if val == nil {
			return fmt.Errorf("no creator for %s", typeField.Type.Name())
		}

		c.v = append(c.v, val)
	}

	configEntries[name] = c
	return nil
}

func HasConfigEntry(name string) bool {
	_, ok := configEntries[name]
	return ok
}

func CreateConfigEntry(name string, e map[string]interface{}) (interface{}, error) {

	creator, ok := configEntries[name]
	if !ok {
		return nil, errors.New("no entry named")
	}

	return creator.Create(e)
}
