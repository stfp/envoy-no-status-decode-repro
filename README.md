# Summary

nodejs sends %-encoded grpc message trailers, and envoy does not decode them when transcoding to json

## json to node via envoy

This shows the issue: we're hitting a nodejs grpc upstream (using https://www.npmjs.com/package/@grpc/grpc-js), with envoy doing the json transcoding.

curl -v  -H 'cluster: echojs' http://127.0.0.1:8180/echo/echo

	{"code":5,"message":"#%20node:%20here's%20an%20error%20message%20with%20spaces%20&%20some%20interesting%20characters!","details":[]}* Closing connection 0

## json to go via envoy

The same call to a go upstream looks fine (b/c the go grpc impl doesn't percent encode grpc status message headers)

curl -v  -H 'cluster: echo' http://127.0.0.1:8180/echo/echo

	{"code":5,"message":"# here's an error message with spaces & some interesting characters!","details":[]}

## pcaps show how these headers look different on the wire

Node:
`        Header: grpc-message: #%20node:%20here's%20an%20error%20message%20with%20spaces%20&%20some%20interesting%20characters!`

Go:
`        Header: grpc-message: # here's an error message with spaces & some interesting characters!`

See go.pcap, node.pcap and Decode.md

## grpcurl decodes correctly

### grpc to node upstream via envoy

grpcurl -H 'cluster: echojs' -plaintext -protoset echo/proto.pb localhost:8180 grpcecho.EchoService.Echo

	ERROR:
	  Code: NotFound
	  Message: # node: here's an error message with spaces & some interesting characters!

### grpc to go upstream via envoy

grpcurl -H 'cluster: echo' -plaintext -protoset echo/proto.pb localhost:8180 grpcecho.EchoService.Echo

	ERROR:
	  Code: NotFound
	  Message: # here's an error message with spaces & some interesting characters!


### grpc direct to node upstream

grpcurl -plaintext -protoset echo/proto.pb localhost:8998 grpcecho.EchoService.Echo

	ERROR:
	  Code: NotFound
	  Message: # node: here's an error message with spaces & some interesting characters!


### grpc direct to go upstream

grpcurl -plaintext -protoset echo/proto.pb localhost:8999 grpcecho.EchoService.Echo

	ERROR:
	  Code: NotFound
	  Message: # here's an error message with spaces & some interesting characters!





