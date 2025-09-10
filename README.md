## Overview ##

This tool is used to automatically complement the data types used to describe enumerations with standard methods that satisfy the interface Enumerator and json.Marshaler, json.Unmarshaler

Enumerator combine the next interface:

* Stringer (correspond fmt.Stringer interface)
* Unstringer - used to set the value of an enumerated type from a string
* Validator - used to check the validity of an enumerated type value

json.Marshaler, json.Unmarshaler - used to marshaling, unmarshaling json string to correspond enumeration value.

## Install ##
go get github.com/ivi81/enummethods@version

## Example of code design for generation ##
```
...
//go:generate enummethods -type=FilterConst
type FilterConst int

const (
	_lt = FilterConst(iota)
	_lte
	_gt
	_gte
	_ne
	_not
	_and
	_or
	_between
	_like
	_contains
	_field
	_value
	_from
	_to
)

const (
	LT       = _lt
	LTE      = _lte
	GT       = _gt
	GTE      = _gte
	NE       = _ne
	NOT      = _not
	AND      = _and
	OR       = _or
	BETWEEN  = _between
	LIKE     = _like
	CONTAINS = _contains
	FIELD    = _field
	VALUE    = _value
	FROM     = _from
	TO       = _to
)
```
...

```
...
go generate
```
...