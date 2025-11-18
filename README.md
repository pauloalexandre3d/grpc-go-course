### Generate code manually

```bash
protoc -I greet/proto --go_out=. --go_opt=module=github.com/pauloalexandre3d/grpc-go-course --go-grpc_out=. --go-grpc_opt=module=github.com/pauloalexandre3d/grpc-go-course greet/proto/dummy.proto
```