//jsoner.go - содержит шаблон для генерации методов рерализующих интерфейсы json.Marshaler, json.UnMarshaler
package methods

import "text/template"

//JsonerTemplate - переменная содержит шаблон на основе которого генерируются методы рерализующие интерфейсы json.Marshaler, json.UnMarshaler
//  Параметры шаблона:
//   -.TypeName — название типа данных для которого генерируются методы,
var JsonerTmpl = template.Must(template.New("").Parse(`

//MarshalJSON - реализует метод интерфейса json.Marshaler
func (m {{.TypeName}}) MarshalJSON() ([]byte, error) {
	return enumerator.MarshalJSON(m)
}

//UnmarshalJSON - реализует метод интерфейса json.UnMarshaler
func (m *{{.TypeName}}) UnmarshalJSON(data []byte) error {
	return enumerator.UnmarshalJSON(m, data)
}
`))
