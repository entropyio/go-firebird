module Firebird

go 1.12

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.5.0
	github.com/gohouse/gorose/v2 v2.1.3
	github.com/gorilla/websocket v1.4.1
	github.com/mattn/go-sqlite3 v2.0.2+incompatible
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
	github.com/robfig/cron v1.2.0
	github.com/spf13/cast v1.3.1
	golang.org/x/net v0.0.0-20190620200207-3b0461eec859
)

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20200108215511-5d647ca15757
	golang.org/x/mod => github.com/golang/mod v0.1.0
	golang.org/x/net => github.com/golang/net v0.0.0-20191209160850-c0dbc17a3553
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190911185100-cd5d95a43a6e
	golang.org/x/sys => github.com/golang/sys v0.0.0-20200107162124-548cf772de50
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20200108203644-89082a384178
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20191204190536-9bdfabe68543
)
