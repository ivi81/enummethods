// interface_test.go - тесты корректности работы автогенерируемой реализации функционала интерфейсных типов
package enumerator_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/ivi81/enummethods/enumerator/testcons"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}

// TestEnumerator тестирование функциональности интерфейса Enumerator
func TestEnumerator(t *testing.T) {

	t.Run("TEST1 : Test Stringer", func(t *testing.T) {

		enumValue := testcons.First
		assert.EqualValues(t, fmt.Sprint(enumValue), "First")
	})

	t.Run("TEST2 : Test Unstringer", func(t *testing.T) {

		var enumValue testcons.TestEnum
		if assert.True(t, enumValue.SetValue("Second")) {
			assert.False(t, enumValue.SetValue("NotSecond"))
		}
	})
	t.Run("TEST3 : Test Validator", func(t *testing.T) {

		var enumValue testcons.TestEnum
		if assert.True(t, enumValue.SetValue("Second")) {
			assert.True(t, enumValue.IsValid())
		}
	})
}

// TestYamler  - тестирование реализации yaml.Unmarshaler yaml.Unmarshaler
func TestYamler(t *testing.T) {

	t.Run("Test : yaml.Unmarshaler as simple value", func(t *testing.T) {

		var enumValue testcons.TestEnum
		b := `"Third"`

		err := yaml.Unmarshal([]byte(b), &enumValue)
		assert.NoError(t, err)
	})

	t.Run("Test : yaml.Unmarshaler as struct field", func(t *testing.T) {

		var temp struct {
			Value testcons.TestEnum `yaml:"value"`
		}

		b := `{"value":"Third"}`
		err := yaml.Unmarshal([]byte(b), &temp)

		assert.NoError(t, err)
		assert.Equal(t, testcons.Third, temp.Value)
	})

	t.Run("Test : yaml.Marshaler as simple value", func(t *testing.T) {

		var enumValue testcons.TestEnum
		enumValue = testcons.Third

		exp := "Third\n"
		b, err := yaml.Marshal(&enumValue)
		assert.NoError(t, err)
		assert.Equal(t, []byte(exp), b)
	})

	t.Run("Test : yaml.Marshaler as struct field", func(t *testing.T) {

		var temp struct {
			Value testcons.TestEnum `yaml:"value"`
		}
		temp.Value = testcons.Third

		exp := "value: Third\n"

		b, err := yaml.Marshal(&temp)
		assert.NoError(t, err)
		assert.Equal(t, []byte(exp), b)
	})

}

// TestJsoner - тестирование реализации json.Unmarshaler json.Marshaler
func TestJsoner(t *testing.T) {
	t.Run("Test : json.Unmarshaler as simple value", func(t *testing.T) {

		var enumValue testcons.TestEnum
		b := `"Third"`
		err := json.Unmarshal([]byte(b), &enumValue)

		assert.NoError(t, err)
		assert.Equal(t, testcons.Third, enumValue)
	})

	t.Run("Test : json.Unmarshaler as struct field", func(t *testing.T) {

		var temp struct {
			Value testcons.TestEnum `json:"value"`
		}

		b := `{"value":"Third"}`
		err := json.Unmarshal([]byte(b), &temp)

		assert.NoError(t, err)
		assert.Equal(t, testcons.Third, temp.Value)
	})

	t.Run("Test : json.Marshaler as simple value", func(t *testing.T) {

		var enumValue testcons.TestEnum
		enumValue = testcons.Third

		exp := `"Third"`
		b, err := json.Marshal(&enumValue)
		assert.NoError(t, err)
		assert.Equal(t, []byte(exp), b)
	})

	t.Run("Test : json.Marshaler as struct field", func(t *testing.T) {

		var temp struct {
			Value testcons.TestEnum `json:"value"`
		}
		temp.Value = testcons.Third

		exp := `{"value":"Third"}`

		b, err := json.Marshal(&temp)
		assert.NoError(t, err)
		assert.Equal(t, []byte(exp), b)
	})
}
