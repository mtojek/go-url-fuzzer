package configuration

import (
	"fmt"
)

type mockedErrorTagMapper struct {
}

func newMockedErrorTagMapper() *mockedErrorTagMapper {
	return new(mockedErrorTagMapper)
}

func (mockedErrorTagMapper *mockedErrorTagMapper) mapErrorTag(tag int, values ...interface{}) error {
	return fmt.Errorf("%v", tag)
}
