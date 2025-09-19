package enumerator

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"gopkg.in/yaml.v2"
)

// Enumerator интерфейсный тип объявляющий функции для работы с нестроковыми константами
type Enumerator interface {
	Stringer
	Validator
	Unstringer
}

// Stringer  интерфейсный тип объявляющий функции для преобразования
// не строкового перечислимого типа в строку
type Stringer interface {
	fmt.Stringer
}

// Unstringer  интерфейсный тип объявляющий функции для преобразования строки
// не строкового перечислимого типа
type Unstringer interface {
	SetValue(s string) bool
}

// Validator интерфейсный тип объявляющий функции для проверки допустимости значения
// не строкового перечислимого типа
type Validator interface {
	IsValid() bool
}

// MarshalJSON маршалинг значений не строковых перечислимых типов в формат json
func MarshalJSON(c Stringer) ([]byte, error) {
	var s string

	if s = c.String(); s == "" {

		err := &json.UnsupportedValueError{
			Value: reflect.ValueOf(c),
			Str:   fmt.Sprintf("\"%b\"", c),
		}
		return nil, err
	}

	return json.Marshal(s)
}

// UnmarshalJSON анмаршалинг строкового значания поля в json-документе в значение не строкового константного типа
func UnmarshalJSON(c Unstringer, data []byte) error {
	var (
		s   string
		err error
	)
	if err = json.Unmarshal(data, &s); err != nil {
		return err
	}
	if !c.SetValue(s) {

		err = &json.UnsupportedValueError{
			Value: reflect.ValueOf(c),
			Str:   fmt.Sprintf("\"%s\" is not valid value", s),
		}
		return err
	}
	return err
}

// MarshalYAML маршалинг значений не строковых перечислимых типов в формат yaml
func MarshalYAML(c Stringer) (interface{}, error) {
	var s string

	if s = c.String(); s == "" {
		return nil, errors.New(fmt.Sprintf("value:\"%b\" is not convertable to string", c))
	}
	return s, nil
	//return yaml.Marshal(s)
}

// UnmarshalYAML анмаршалинг строки в поле yaml-документа в значение не строкового константного типа
func UnmarshalYAML(c Unstringer, unmarshal func(interface{}) error) error {
	var (
		s   string
		err error
	)
	if err = unmarshal(&s); err != nil {
		return err
	}

	if !c.SetValue(s) {

		err = &yaml.TypeError{
			Errors: []string{fmt.Sprintf("\"%s\" is not valid value", s)},
		}

		return err
	}
	return err
}
