# vts
Go implemented CLI tool for display current system timestamp number, or convert someone timestamp into formated string which human read.

## Installation

	go get -u github.com/tfenng/vts
	
make sure the `$GOPATH/bin` is added into `$PATH`

## Usage

Hello World:

	$ vts
	$ vts 1563270875
	$ vts      70875
	$ vts | xargs vts

Synopsis:

	vts [flags] [TIMESTAMP]
	
See also `vts --help`.	
