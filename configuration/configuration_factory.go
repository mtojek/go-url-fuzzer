package configuration

import (
	"errors"
	"gopkg.in/alecthomas/kingpin.v2"
)

type ConfigurationFactory struct{}

func NewConfigurationFactory() *ConfigurationFactory {
	return new(ConfigurationFactory)
}

func (this *ConfigurationFactory) FromCommandLine() *Configuration {
	configuration := this.createFlagsBoundConfiguration()
	this.parseCommandLine()
	return configuration
}

func (this *ConfigurationFactory) createFlagsBoundConfiguration() *Configuration {
	configuration := newConfiguration()
	configuration.headers = kingpin.Flag("header", "Custom HTTP header added to every fuzz request, format: \"name: value\"").Short('h').PlaceHolder("\"Name: value\"").StringMap()
	configuration.methods = kingpin.Flag("method", "HTTP method used in tests (GET, POST, PUT, DELETE, HEAD, OPTIONS)").Short('m').Default("GET").Strings()
	configuration.outputFile = kingpin.Flag("output", "Output text file with found urls and statuses").Short('o').PlaceHolder("output_file.txt").String()
	configuration.reportDirectory = kingpin.Flag("report", "Target output directory of an url fuzzing HTML report").Short('r').PlaceHolder("report_directory").String()
	configuration.urlResponseTimeout = kingpin.Flag("timeout", "Fuzzed url response timeout").Short('t').Default("5s").Duration()
	configuration.workerWaitPeriod = kingpin.Flag("wait-period", "Time wait period between fuzz tests per worker").Short('w').Default("0s").Duration()
	configuration.fuzzSetFile = kingpin.Arg("fuzz-set-file", "File containing fuzz entry set, one entry per line").Required().File()
	configuration.baseURL = kingpin.Arg("base-url", "Number of packets to send").Required().URL()

	return configuration
}

func (this *ConfigurationFactory) parseCommandLine() {
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version("0.1").Author("Marcin Tojek").Validate(this.validateFlagsArguments)
	kingpin.Parse()
}

func (this *ConfigurationFactory) validateFlagsArguments(*kingpin.Application) error {
	return errors.New("blabla")
}
