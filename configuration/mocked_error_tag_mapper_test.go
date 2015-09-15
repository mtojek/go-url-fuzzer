package configuration

import (
	"fmt"
)

type mockedErrorTagMapper struct {
}

func newMockedErrorTagMapper() *mockedErrorTagMapper {
	return new(mockedErrorTagMapper)
}

func (m *mockedErrorTagMapper) mapErrorTag(tag int, values ...interface{}) error {
	return fmt.Errorf("%v", tag)
}
