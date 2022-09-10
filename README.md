# go-grpc

### ⚠️ If error
message like below:

protoc-gen-go-grpc: program not found or is not executable
Please specify a program using absolute path or make sure the program is available in your PATH system variable
--go-grpc_out: protoc-gen-go-grpc: Plugin failed with status code 1.

you can enter to cmd
```shell
cd  ${GOPATH}/bin/  -- find two files protoc-gen-go  protoc-gen-go-grpc
 
cp protoc-gen-go /usr/local/bin/
cp protoc-gen-go-grpc /usr/local/bin/
```



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
evans 
https://github.com/ktr0731/evans