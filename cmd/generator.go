package main

import (
	"fmt"
	"go/ast"
	"log"
	"strings"

	"golang.org/x/tools/go/packages"
)

// Generator содержит состояние анализа кода.
//В основном используется для буферизации выходных данных для format.Source.
//  Поля:
//   pkg - сканируемый пакет
type Generator struct {
	pkg *Package

	trimPrefix  string
	lineComment bool

	logf func(format string, args ...interface{}) // test logging hook; nil when not testing
}

// parsePackage анализирует пакет.
// parsePackage exits if there is an error.
func (g *Generator) parsePackage(patterns []string, tags []string) {
	cfg := &packages.Config{
		Mode:       packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedSyntax,
		Tests:      false,
		BuildFlags: []string{fmt.Sprintf("-tags=%s", strings.Join(tags, " "))},
		Logf:       g.logf,
	}
	pkgs, err := packages.Load(cfg, patterns...)
	if err != nil {
		log.Fatal(err)
	}
	if len(pkgs) != 1 {
		log.Fatalf("error: %d packages matching %v", len(pkgs), strings.Join(patterns, " "))
	}
	fmt.Printf("PKG PATH:%s", patterns)
	fmt.Printf("PKG NAME:%s", pkgs[0].Name)
	g.addPackage(pkgs[0])
}

// addPackage - добавляет информацию о пакете и информацию о синтаксических файлах в Generator.
func (g *Generator) addPackage(pkg *packages.Package) {
	g.pkg = &Package{
		name:  pkg.Name,
		defs:  pkg.TypesInfo.Defs,
		files: make([]*File, len(pkg.Syntax)),
	}

	for i, file := range pkg.Syntax {
		g.pkg.files[i] = &File{
			file:        file,
			pkg:         g.pkg,
			trimPrefix:  g.trimPrefix,
			lineComment: g.lineComment,
		}
	}
}

// generate производит генерацию информации для списка заданных наименований типов.
func (g *Generator) generate(typeNames []string) map[string]TypeInfo {
	typeInfos := make(map[string]TypeInfo, len(typeNames))

	for _, file := range g.pkg.files {

		if file.file != nil {
			file.typeNames = typeNames
			file.typeInfos = make(map[string]TypeInfo, len(typeNames))
			ast.Inspect(file.file, file.genDecl)

			for k, v := range file.typeInfos {
				info := typeInfos[k]

				if len(v.constNames) != 0 {
					info.constNames = append(info.constNames, v.constNames...)
				}

				info.withArray = v.withArray
				info.pkgName=g.pkg.name
				typeInfos[k] = info
			}
		}
	}

	return typeInfos
}
