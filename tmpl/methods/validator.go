package methods

import (
	"reflect"
	"text/template"
)

//ValidatorTmpl - переменная, содержит шаблон на основе которого генерируется метод рерализующий интерфейс Validator.
//  Параметры шаблона:
//   - .TypeName — название типа данных для которого генерируются методы,

var funcMap = template.FuncMap{
	"last": func(x int, a interface{}) bool {
		return x == int(reflect.ValueOf(a).Int()-1)
	},
}

var ValidatorTmpl = template.Must(template.New("").Funcs(funcMap).Parse(`

//IsValid проверка корректности значения
//Реализует интерфейс Validator
func (m {{.TypeName}}) IsValid() bool {

	switch m {
	case
		{{$l:=len .ConstNames -}}
		{{range $i,$v:=.ConstNames}}{{ $v }}{{if not (last $i $l) }},{{else}}:{{end}}
		{{end}}return true
	}
	return false
}
`))
