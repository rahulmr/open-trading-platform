module github.com/ettec/open-trading-platform/go/execution-venues/fix-sim-execution-venue

go 1.13

require (
	github.com/ettec/open-trading-platform/go/common v0.0.0
	github.com/ettec/open-trading-platform/go/execution-venues/common v0.0.0
	github.com/ettec/open-trading-platform/go/model v0.0.0
	github.com/quickfixgo/quickfix v0.6.0
	github.com/shopspring/decimal v0.0.0-20191009025716-f1972eb1d1f5
	google.golang.org/grpc v1.25.1
)

replace github.com/ettec/open-trading-platform/go/common v0.0.0 => ../../common

replace github.com/ettec/open-trading-platform/go/execution-venues/common v0.0.0 => ../common

replace github.com/ettec/open-trading-platform/go/model v0.0.0 => ../../model
