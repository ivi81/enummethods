## Overview ##

This tool is used to automatically complement the data types used to describe enumerations with standard methods that satisfy the interface Enumerator and json.Marshaler, json.Unmarshaler

Enumerator combine the next interface:

* Stringer (correspond fmt.Stringer interface)
* Unstringer - used to set the value of an enumerated type from a string
* Validator - used to check the validity of an enumerated type value

json.Marshaler, json.Unmarshaler - used to marshaling, unmarshaling json string to correspond enumeration value.

## Install ##
go get github.com/ivi81/enummethods