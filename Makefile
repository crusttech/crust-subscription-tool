TS := $(shell date +'%Y%m%d%H%M%S')

build:
	go build -o make-crust-sub cmd/make-crust-sub.go

keys:
	openssl ecparam -genkey -name secp521r1 -noout -out $(TS)-prv.pem
	openssl ec -in $(TS)-prv.pem -pubout -out $(TS)-pub.pem
