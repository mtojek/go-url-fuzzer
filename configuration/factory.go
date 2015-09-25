package configuration

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

// Factory allows for creating a struct representing configuration.
type Factory struct{}

// NewFactory returns a new instance of configuration f
func NewFactory() *Factory {
	return new(Factory)
}

// FromCommandLine returns a configuration created from the command line parameters.
func (f *Factory) FromCommandLine() *Configuration {
	configuration := f.createFlagsBoundConfiguration()
	f.parseFlagsArguments(configuration)
	return configuration
}

func (f *Factory) createFlagsBoundConfiguration() *Configuration {
	configuration := newConfiguration()
	configuration.headers = kingpin.Flag("header", "Custom HTTP header added to every fuzz request, format: \"name: value\"").Short('h').PlaceHolder("\"Name: value\"").StringMap()
	configuration.methods = kingpin.Flag("method", "HTTP method used in tests (GET, POST, PUT, DELETE, HEAD, OPTIONS)").Short('m').Default("GET").Enums("GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS")
	configuration.outputFile = kingpin.Flag("output", "Output text file with found urls and statuses").Short('o').PlaceHolder("output_file.txt").String()
	configuration.urlResponseTimeout = kingpin.Flag("timeout", "Fuzzed url response timeout").Short('t').Default("5s").Duration()
	configuration.httpErrorCode = kingpin.Flag("http-error-code", "HTTP error code").Short('e').Default("404").Uint64()
	configuration.workersNumber = kingpin.Flag("workers-number", "Number of workers").Short('n').Default("4").Uint64()
	configuration.workerWaitPeriod = kingpin.Flag("wait-period", "Time wait period between fuzz tests per worker").Short('w').Default("0s").Duration()
	configuration.fuzzSetFile = kingpin.Arg("fuzz-set-file", "File containing fuzz entry set, one entry per line").Required().File()
	configuration.baseURL = kingpin.Arg("base-url", "Number of packets to send").Required().URL()

	return configuration
}

func (f *Factory) parseFlagsArguments(flagsBoundConfiguration *Configuration) {
	configurationValidator := newValidator(flagsBoundConfiguration)
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version("1.0").Author("Marcin Tojek").Validate(configurationValidator.validate)
	kingpin.CommandLine.Name = "go-url-fuzzer"
	kingpin.CommandLine.Help = "Discover hidden files and directories on a web server."
	kingpin.Parse()
}
