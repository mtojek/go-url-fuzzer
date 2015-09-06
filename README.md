# go-url-fuzzer

Status: **Work in progress**, not ready yet

Discover hidden files and directories on a web server. The application tries to find url relative paths of the given website by comparing them with a given set. Go-url-fuzzer is inspired by [Indir Scanner](http://indir.uw-team.org/), which is written in Perl. Comparing to Indir Scanner, the application supports concurrent url fuzzing.

## Features

* Fuzz url set from an input file
* Concurrent relative path search
* Configurable number of fuzzing workers
* Configurable time wait periods between fuzz tests per worker
* Custom HTTP headers support
* Different HTTP method support
* HTML url fuzzing reports
