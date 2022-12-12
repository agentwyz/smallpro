gin框架的底层就是基于http/net库进行开发的

```go
func (engine *Engine) Run(addr ...string) (err error) {
	defer func() { debugPrintError(err) }()

	if engine.isUnsafeTrustedProxies() {
		debugPrint("[WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.\n" +
			"Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.")
	}

	address := resolveAddress(addr)
	debugPrint("Listening and serving HTTP on %s\n", address)
	err = http.ListenAndServe(address, engine.Handler())
	return
}
```
r这个变量可以调用`Run`这个方法, 这个方法属于Egine这个类的, 所以r一定是`Engine`类的对象

然后来解析一下`Engine`这个类

结构体使用了哪些接口
```go


```
