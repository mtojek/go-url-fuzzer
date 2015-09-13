package configuration

import (
	"errors"
	"fmt"
)

type mockedErrorTagMapper struct {
}

func newMockedErrorTagMapper() *mockedErrorTagMapper {
	return new(mockedErrorTagMapper)
}

func (this *mockedErrorTagMapper) mapErrorTag(tag int, values ...interface{}) error {
	return errors.New(fmt.Sprintf("%v", tag))
}
