# go-url-fuzzer

[![Build Status](https://travis-ci.org/mtojek/go-url-fuzzer.svg?branch=master)](https://travis-ci.org/mtojek/go-url-fuzzer)

Status: **Done**

Discover hidden files and directories on a web server. The application tries to find url relative paths of the given website by comparing them with a given set. Go-url-fuzzer is inspired by [Indir Scanner](http://indir.uw-team.org/), which is written in Perl. Comparing to Indir Scanner, the application supports concurrent url fuzzing.

## Features

* Fuzz url set from an input file
* Concurrent relative path search
* Configurable number of fuzzing workers
* Configurable time wait periods between fuzz tests per worker
* Custom HTTP headers support
* Various HTTP methods support

## Usage

~~~
$ go-url-fuzzer --help
usage: go-url-fuzzer [<flags>] <fuzz-set-file> <base-url>

Discover hidden files and directories on a web server.

Flags:
  --help            Show help (also see --help-long and --help-man).
  -h, --header="Name: value"
                    Custom HTTP header added to every fuzz request, format: "name: value"
  -m, --method=GET  HTTP method used in tests (GET, POST, PUT, DELETE, HEAD, OPTIONS)
  -o, --output=output_file.txt
                    Output text file with found urls and statuses
  -t, --timeout=5s  Fuzzed url response timeout
  -e, --http-error-code=404
                    HTTP error code
  -n, --workers-number=4
                    Number of workers
  -w, --wait-period=0s
                    Time wait period between fuzz tests per worker
  --version         Show application version.

Args:
  <fuzz-set-file>  File containing fuzz entry set, one entry per line
  <base-url>       Number of packets to send
~~~

Example:
~~~
go-url-fuzzer -h "User-Agent: curl" -h "Cookie: token=1" -m "GET" -m "POST" resources/input-data/fuzz_02.txt http://domain.tld/any-dir/
~~~

## License

The MIT License (MIT)

Copyright (c) 2015 Marcin Tojek

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

