package enumerator

import (
	"encoding/json"
	"fmt"
	"reflect"
)

//Enumerator интерфейсный тип объявляющий функции для работы с нестроковыми константами
type Enumerator interface {
	Stringer
	Validator
	Unstringer
}

//Stringer  интерфейсный тип объявляющий функции для преобразования
//не строкового перечислимого типа в строку
type Stringer interface {
	fmt.Stringer
}

//Unstringer  интерфейсный тип объявляющий функции для преобразования строки
//не строкового перечислимого типа
type Unstringer interface {
	SetValue(s string) bool
}

//Validator интерфейсный тип объявляющий функции для проверки допустимости значения
//не строкового перечислимого типа
type Validator interface {
	IsValid() bool
}

//MarshalConstantJSON маршалинг значений не строковых перечислимых типов в json
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

//UnmarshalConstantJSON анмаршалинг строкового значания поля в json-документе в значение не строкового константного типа
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
