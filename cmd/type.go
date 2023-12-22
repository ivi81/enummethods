package main

// TypeInfo агрегирует информацию о типе данных необходимую для генерации методов данного типа
// Поля:
//   constNames - список названий констант
//   withArray - используется массив или нет
//   varName - название переменной содержащей строки
type TypeInfo struct {
	constNames ConstSlice
	withArray  bool
	pkgName string
}

type ConstSlice []string
