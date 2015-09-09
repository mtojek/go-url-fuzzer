package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	headers            = kingpin.Flag("header", "Custom HTTP header added to every fuzz request, format: \"name: value\"").Short('h').PlaceHolder("\"Name: value\"").StringMap()
	methods            = kingpin.Flag("method", "HTTP method used in tests (GET, POST, PUT, DELETE, HEAD, OPTIONS)").Short('m').Default("GET").Strings()
	outputFile         = kingpin.Flag("output", "Output text file with found urls and statuses").Short('o').PlaceHolder("output_file.txt").String()
	reportDirectory    = kingpin.Flag("report", "Target output directory of an url fuzzing HTML report").Short('r').PlaceHolder("report_directory").File()
	urlResponseTimeout = kingpin.Flag("timeout", "Fuzzed url response timeout, in milliseconds").Short('t').Default("5000").Uint64()
	workerWaitPeriod   = kingpin.Flag("wait-period", "Time wait period between fuzz tests per worker, in milliseconds").Short('w').Default("0").Uint64()
	fuzzSetFile        = kingpin.Arg("fuzz-set-file", "File containing fuzz entry set, one entry per line").Required().ExistingFile()
	baseURL            = kingpin.Arg("base-url", "Number of packets to send").Required().URL()
)

func main() {
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate).Version("0.1").Author("Marcin Tojek")
	kingpin.Parse()
}
