//file.go - содержит тип данных
package main

import (
	"go/ast"
	"go/token"
	"go/types"
	"log"
	"strings"
)

//File - тип данных описывающий данные извлеченные из обработанного файла с кодом.
// Поля структуры:
//   pkg       - содержит пакет которому принадлежит файл
//   file      - содержит синтаксическое дерево кода обработанного файла
//   typeNames - список имен обрабатываемых типов данных
//   values    - аккамулирует информацию необходимую для генерации
type File struct {
	pkg  *Package
	file *ast.File
	// These fields are reset for each type being generated.
	typeNames []string
	typeInfos map[string]TypeInfo

	trimPrefix  string
	lineComment bool
}

// genDecl processes one declaration clause.
func (f *File) genDecl(node ast.Node) bool {
	switch decl := node.(type) {
	case *ast.GenDecl:
		switch decl.Tok {
		case token.CONST:
			f.getConsts(decl.Specs)
			break
		case token.VAR:
			f.checkForArray(decl.Specs)
			break
		case token.TYPE:
			break
		default:
			return true
		}
	default:
		return true
	}
	return false
}

//isValidTypeNames - проверка названия обрабатываемого типа
func (f *File) isValidTypeNames(name string) bool {
	for _, typeName := range f.typeNames {
		if typeName == typeName {
			return true
		}
	}
	return false
}

//isSimilarToTypeName - проверка условия того что название константного типа и имя переменной похожи
func (f *File) isSimilarToTypeName(name string) (string, bool) {
	for _, typeName := range f.typeNames {
		if name == convertLowerLetter(typeName) {
			return typeName, true
		}
	}
	return "", false
}

//getConst - получаем список констант
func (f *File) getConsts(specs []ast.Spec) {
	typeName := ""
	for _, spec := range specs {
		vspec := spec.(*ast.ValueSpec)
		if typ, ok := f.isValidConst(vspec); ok {
			typeName = typ
		}
		if typeName == "" {
			continue
		}
		f.getConstName(vspec, typeName)
	}
}

//isValidConst - устанваливает валидность типа данных для констант
func (f *File) isValidConst(vspec *ast.ValueSpec) (string, bool) {
	//проверяем случай когда при названии константы не указан тип
	if vspec.Type == nil && len(vspec.Values) > 0 {
		if ce, ok := vspec.Values[0].(*ast.CallExpr); ok {
			if id, ok := ce.Fun.(*ast.Ident); ok {
				if f.isValidTypeNames(id.Name) {
					return id.Name, true
				}
			}
		}
		return "", false
	}
	//Проверяем силучай когда название типа указанно в поле Type (как правило это первый элемент
	//в списке констант)
	if vspec.Type != nil {
		if ident, ok := vspec.Type.(*ast.Ident); ok && f.isValidTypeNames(ident.Name) {
			return ident.Name, true
		}
	}
	return "", false
}

//getConstName - извлечние имени константы из синтаксического обекта
func (f *File) getConstName(vspec *ast.ValueSpec, typeName string) {

	//извлекаем имена констант
	typeInfo := f.typeInfos[typeName]
	for _, name := range vspec.Names {
		if name.Name == "_" {
			continue
		}

		obj, ok := f.pkg.defs[name]
		if !ok {
			log.Fatalf("no value for constant %s", name)
		}
		info := obj.Type().Underlying().(*types.Basic).Info()
		if info&types.IsInteger == 0 {
			log.Fatalf("can't handle non-integer constant type %s", typeName)
		}
		typeInfo.constNames = append(typeInfo.constNames, name.Name)
	}
	f.typeInfos[typeName] = typeInfo
}

//checkForArray - получаем название переменной массива строк
func (f *File) checkForArray(specs []ast.Spec) {
	for _, spec := range specs {
		vspec := spec.(*ast.ValueSpec)
		varName := vspec.Names[0].Name

		if typeName, withArrayOk := f.isSimilarToTypeName(varName); withArrayOk {
			//Проверяем тип элементов массива, должен быть string
			cl := vspec.Values[0].(*ast.CompositeLit)
			at := cl.Type.(*ast.ArrayType)
			if at.Elt.(*ast.Ident).Name != "string" {
				log.Fatalf("no []string type for %s", varName)
			}
			//сохраняем в общий список
			info := f.typeInfos[typeName]
			info.withArray = withArrayOk
			f.typeInfos[typeName] = info
		}
	}
}

//convertLowerLetter - конвертировать первую букву слова
//из заглавной в строчную
func convertLowerLetter(word string) string {
	lower := strings.ToLower(word[0:1])
	return (lower + word[1:])
}
