//unstringer.go - содержит шаблон для генерации методов рерализующих интерфайс Unstringer
package methods

import "text/template"

//UnstringerTmpl - переменная содержит шаблон на основе которого генерируются методы рерализующие интерфейс Stringer.
//Данный шаблон применяется в том случае когда строки сопоставляеимые значениями перечислимого типа хранятся в виде одной строки.
//  Параметры шаблона:
//   - .CurrentPkgName — название пакета в котором располагается результат генерации,
//   - .TypeName — название типа данных для которого генерируются методы,
var UnstringerTmpl = template.Must(template.New("").Parse(`

//SetValue конвертация строки в значение типа
//Реализует интерфейс Unstringer
func (m *{{.TypeName}}) SetValue(s string) bool {
	i := strings.Index(_{{.TypeName}}_name, s)
	if i != -1 {

		for index, v := range _{{.TypeName}}_index {
			if i-int(v) == 0 {
				*m = {{.TypeName}}(index)
				return true
			}
		}
	}
	return false
}

`))
