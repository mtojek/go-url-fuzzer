package configuration

type errorTagMapper interface {
	mapErrorTag(tag int, values ...interface{}) error
}
