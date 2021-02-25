

# json to node via envoy

curl -v  -H 'cluster: echojs' http://127.0.0.1:8180/echo/echo

	{"code":5,"message":"#%20node:%20here's%20an%20error%20message%20with%20spaces%20&%20some%20interesting%20characters!","details":[]}* Closing connection 0


# json to go via envoy

curl -v  -H 'cluster: echo' http://127.0.0.1:8180/echo/echo

	{"code":5,"message":"# here's an error message with spaces & some interesting characters!","details":[]}


# grpc to node upstream via envoy

grpcurl -H 'cluster: echojs' -plaintext -protoset echo/proto.pb localhost:8180 grpcecho.EchoService.Echo

	ERROR:
	  Code: NotFound
	  Message: # node: here's an error message with spaces & some interesting characters!

# grpc to go upstream via envoy

grpcurl -plaintext -protoset echo/proto.pb localhost:8998 grpcecho.EchoService.Echo

	ERROR:
	  Code: NotFound
	  Message: # node: here's an error message with spaces & some interesting characters!


# grpc direct to node upstream

grpcurl -plaintext -protoset echo/proto.pb localhost:8998 grpcecho.EchoService.Echo

	ERROR:
	  Code: NotFound
	  Message: # node: here's an error message with spaces & some interesting characters!



# grpc direct to go upstream

grpcurl -plaintext -protoset echo/proto.pb localhost:8999 grpcecho.EchoService.Echo

	ERROR:
	  Code: NotFound
	  Message: # here's an error message with spaces & some interesting characters!





