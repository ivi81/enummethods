// yamler.go - содержит шаблон для генерации методов рерализующих интерфейсы json.Marshaler, json.UnMarshaler
package methods

import "text/template"

// YamlerTemplate - переменная содержит шаблон на основе которого генерируются методы рерализующие интерфейсы json.Marshaler, json.UnMarshaler
//
//	Параметры шаблона:
//	 -.TypeName — название типа данных для которого генерируются методы,
var YamlerTmpl = template.Must(template.New("").Parse(`

//MarshalYAML - реализует метод интерфейса yaml.Marshaler
func (m {{.TypeName}}) MarshalYAML() (interface{}, error) {
	return enumerator.MarshalYAML(m)
}

//UnmarshalYAML - реализует метод интерфейса yaml.UnMarshaler
func (m *{{.TypeName}}) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return enumerator.UnmarshalYAML(m, unmarshal)
}
`))
