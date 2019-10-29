module github.com/crusttech/crust-subscription

go 1.12

replace gopkg.in/Masterminds/squirrel.v1 => github.com/Masterminds/squirrel v1.1.0

require (
	github.com/crusttech/crust-server v0.0.0-20191028201848-9d88d3346b1f
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/ianlancetaylor/demangle v0.0.0-20181102032728-5e5cf60278f6 // indirect
	github.com/spf13/pflag v1.0.5
	golang.org/x/arch v0.0.0-20190927153633-4e8777c89be4 // indirect
)
