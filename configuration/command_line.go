package configuration

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	headers            = kingpin.Flag("header", "Custom HTTP header added to every fuzz request, format: \"name: value\"").Short('h').PlaceHolder("\"Name: value\"").StringMap()
	methods            = kingpin.Flag("method", "HTTP method used in tests (GET, POST, PUT, DELETE, HEAD, OPTIONS)").Short('m').Default("GET").Strings()
	outputFile         = kingpin.Flag("output", "Output text file with found urls and statuses").Short('o').PlaceHolder("output_file.txt").String()
	reportDirectory    = kingpin.Flag("report", "Target output directory of an url fuzzing HTML report").Short('r').PlaceHolder("report_directory").String()
	urlResponseTimeout = kingpin.Flag("timeout", "Fuzzed url response timeout").Short('t').Default("5s").Duration()
	workerWaitPeriod   = kingpin.Flag("wait-period", "Time wait period between fuzz tests per worker").Short('w').Default("0s").Duration()
	fuzzSetFile        = kingpin.Arg("fuzz-set-file", "File containing fuzz entry set, one entry per line").Required().File()
	baseURL            = kingpin.Arg("base-url", "Number of packets to send").Required().URL()
)

func FromCommandLine() *Configuration {
	parseCommandLine()
	return createConfiguration()
}

func parseCommandLine() {
	// TODO validator
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version("0.1").Author("Marcin Tojek")
	kingpin.Parse()
}

func createConfiguration() *Configuration {
	aConfiguration := new(Configuration)
	if headers != nil {
		aConfiguration.Headers = *headers
	}

	if methods != nil {
		aConfiguration.Methods = *methods
	}

	// TODO != nil

	return aConfiguration
}