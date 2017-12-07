# Install
`cd $GOPATH/src`
`go install -u github.com/dvwallin/perinc`

# Usage
`$GOPATH/bin/perinc <base_price> <percentage> <years>`

For example:
`$GOPATH/bin/perinc 725 4 12` which will increase the base price 725 with 4% per year for 12 years