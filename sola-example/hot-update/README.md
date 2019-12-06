# Hot Update

执行 `sola-hot` 进行热更新，运行过程中将生成临时可执行文件 `sola-dev`。

linux 系统使用 `sola-linux` 监听端口可实现平滑切换（不中断请求）。 

```go
import (
	linux "github.com/ddosakura/sola/v2/extension/sola-linux"
)

linux.Listen("127.0.0.1:3000", app)
linux.Keep()
```
