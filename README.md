# go-grpc

if has error 
message like :

protoc-gen-go-grpc: program not found or is not executable
Please specify a program using absolute path or make sure the program is available in your PATH system variable
--go-grpc_out: protoc-gen-go-grpc: Plugin failed with status code 1.

you can enter to cmd
cd to go path find protoc-gen-go and protoc-gen-go 
cp protoc-gen-go /usr/local/bin/
cp protoc-gen-go-grpc /usr/local/bin/


```shell
make  xxx
-- run server

-- enter server
evans --host localhost --port 50051 --reflection repl

show package
show message
show service
-- use service 
service  CalculatorService
call xxx
```
