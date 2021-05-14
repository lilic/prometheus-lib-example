module github.com/lilic/prometheus-lib-example/fancy-app

go 1.16

require (
	github.com/lilic/prometheus-lib-example/lib v0.0.0-00010101000000-000000000000
	github.com/prometheus/client_golang v1.10.0
)

replace github.com/lilic/prometheus-lib-example/lib => /Users/lili/go/src/github.com/lilic/prometheus-lib-example/lib
