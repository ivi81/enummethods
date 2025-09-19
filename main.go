package main

// TODO: процессор должен работать с go generate (go generate enummethods -type typeName1, typeName2), должен формировать типовой набор методов только для тех типов данных
//чьи имена переданы через параметр -type. Должен искать и извлекать из анализируемого кода одноименное с обрабатываемым типом данных название переменной содержащей массив
// строковых значений. Должен
import (
	"bytes"

	"path"

	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ivi81/enummethods/tmpl/methods"
	"github.com/ivi81/enummethods/version"
)

var (
	typeNames   = flag.String("type", "", "comma-separated list of type names; must be set")
	output      = flag.String("output", "", "output file name; default same file as srcdir/srcFileName.go")
	showVersion = flag.Bool("version", false, "show version information")
	//arrayName = flag.String("array", "", "string array name; default same as base type name start with lower letter")
	//trimprefix  = flag.String("trimprefix", "", "trim the `prefix` from the generated constant names")
	linecomment = flag.Bool("linecomment", false, "use line comment text as printed text when present")
	buildTags   = flag.String("tags", "", "comma-separated list of build tags to apply")
)

// Usage is a replacement usage function for the flags package.
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of enummethods:\n")
	fmt.Fprintf(os.Stderr, "\tennummethods [flags] -type T [directory]\n")
	fmt.Fprintf(os.Stderr, "\tennummethods [flags] -type T files... # Must be a single package\n")

	flag.PrintDefaults()
}

// isDirectory reports whether the named file is a directory.
func isDirectory(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
	}
	return info.IsDir()
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("enummethods: ")
	flag.Usage = Usage
	flag.Parse()

	if *showVersion {
		info := version.Get()
		fmt.Println(info)
		return
	}

	if len(*typeNames) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	types := strings.Split(*typeNames, ",")
	var tags []string
	if len(*buildTags) > 0 {
		tags = strings.Split(*buildTags, ",")
	}

	// We accept either one directory or a list of files. Which do we have?
	args := flag.Args()

	if len(args) == 0 {
		// Default: Обрабатываем пакет находящийся в текущей директории.
		args = []string{"."}
	}

	var dir string
	if len(args) == 1 && isDirectory(args[0]) {
		dir = args[0]
	} else {
		if len(tags) != 0 {
			log.Fatal("-tags option applies only to directories, not when files are specified")
		}
		dir = filepath.Dir(args[0])
	}

	g := Generator{
		lineComment: *linecomment,
	}

	g.parsePackage(args, tags)

	for typeName, info := range g.generate(types) {

		outputPath := *output
		var fileName string

		if outputPath == "" {
			fileName = fmt.Sprintf("%s_enummethods.go", strings.ToLower(typeName))
			outputPath = filepath.Join(dir, fileName)
		} else {
			fileName = path.Base(outputPath)
		}

		var buf bytes.Buffer

		params := struct {
			TypeName   string
			ConstNames ConstSlice
			ArrayName  string
			Package    string
			FileName   string
			WithArray  bool
		}{
			TypeName:   typeName,
			ConstNames: info.constNames,
			ArrayName:  convertLowerLetter(typeName),
			FileName:   fileName,
			Package:    info.pkgName,
			WithArray:  info.withArray,
		}
		methods.PkgTmpl.Execute(&buf, params)
		methods.ValidatorTmpl.Execute(&buf, params)
		methods.JsonerTmpl.Execute(&buf, params)
		methods.YamlerTmpl.Execute(&buf, params)

		if info.withArray {
			methods.StringerWithArrayTmpl.Execute(&buf, params)
			methods.UnstringerWithArrayTmpl.Execute(&buf, params)
		} else {
			methods.UnstringerTmpl.Execute(&buf, params)
		}

		b := buf.Bytes()
		os.WriteFile(outputPath, b, 0644)

	}
}
