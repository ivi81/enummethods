//unstringerWithArraw.go - содержит шаблон для генерации метода SetValue рерализующего интерфейс Unstringer
//применяется в случае если множество строк которые сопоставляются со значениями перечислимого типа
// хранятся в массиве

package methods

import "text/template"

//UnstringerWithArrayTmpl - переменная, содержит шаблон на основе которого генерируется метод рерализующий интерфейс Unstringer.
// Данный шаблон применяется в том случае когда строки сопоставляеимые значениями перечислимого типа хранятся в массиве.
//  Параметры шаблона:
//   - .TypeName — название типа данных для которого генерируются методы,
//   - .ArrayName
var UnstringerWithArrayTmpl = template.Must(template.New("").Parse(`

//SetValue - конвертация строки в значение перечислимого типа.
//Реализует интерфейс Unstringer
func (m *{{.TypeName}}) SetValue(s string) bool {
	for i, v := range {{.ArrayName}} {
		if v == s {
			*m = {{.TypeName}}(i)
			return true
		}
	}
	return false
}
`))
