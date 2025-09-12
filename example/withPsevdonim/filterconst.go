package withpsevdonim

//go:generate enummethods -type=FilterConst
//go:generate stringer -type=FilterConst
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
