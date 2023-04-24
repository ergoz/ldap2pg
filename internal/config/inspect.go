package config

// Either an SQL string or a predefined list of YAML rows.
type RowsOrSQL interface{}

// Like pgx.RowToFunc, but from YAML
type YamlToFunc[T any] func(row interface{}) (T, error)

func IsPredefined(q RowsOrSQL) bool {
	switch q.(type) {
	case string:
		return false
	default:
		return true
	}
}
