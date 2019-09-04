module episode29

go 1.12

require (
	cloud.google.com/go v0.36.0 // indirect
	github.com/cockroachdb/apd v1.1.0 // indirect
	github.com/cockroachdb/cockroach-go v0.0.0-20181001143604-e0a95dfd547c // indirect
	github.com/gobuffalo/buffalo v0.14.8
	github.com/gobuffalo/buffalo-goth v1.0.3 // indirect
	github.com/gobuffalo/buffalo-plugins v1.14.1 // indirect
	github.com/gobuffalo/buffalo-pop v1.16.0
	github.com/gobuffalo/envy v1.7.0
	github.com/gobuffalo/mw-csrf v0.0.0-20190129204204-25460a055517
	github.com/gobuffalo/mw-forcessl v0.0.0-20190224202501-6d1ef7ffb276
	github.com/gobuffalo/mw-i18n v0.0.0-20190224203426-337de00e4c33
	github.com/gobuffalo/mw-paramlogger v0.0.0-20190224201358-0d45762ab655
	github.com/gobuffalo/packr v1.30.1
	github.com/gobuffalo/packr/v2 v2.5.2
	github.com/gobuffalo/pop v4.11.2+incompatible
	github.com/gobuffalo/suite v2.8.1+incompatible
	github.com/jackc/fake v0.0.0-20150926172116-812a484cc733 // indirect
	github.com/markbates/going v1.0.3 // indirect
	github.com/markbates/goth v1.55.0
	github.com/markbates/grift v1.1.0
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/shopspring/decimal v0.0.0-20180709203117-cd690d0c9e24 // indirect
	github.com/unrolled/secure v1.0.1
)

replace (
	github.com/golang/lint => golang.org/x/lint v0.0.0-20190409202823-959b441ac422
	sourcegraph.com/sourcegraph/go-diff => github.com/sourcegraph/go-diff v0.5.1
)
