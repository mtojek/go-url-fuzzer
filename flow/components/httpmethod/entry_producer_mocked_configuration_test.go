package httpmethod

type entryProducerMockedConfiguration struct {
	methods []string
}

func newEntryProducerMockedConfiguration(methods []string) *entryProducerMockedConfiguration {
	return &entryProducerMockedConfiguration{methods: methods}
}

func (e *entryProducerMockedConfiguration) Methods() []string {
	return e.methods
}
