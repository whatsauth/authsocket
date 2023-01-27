# websocket
websocket backend for whatsauth

in main before fiber declaration:
go wasocket.RunHub()


use in controller :
a:=wasocket.RunSocket(c)

## Publish
GOPROXY=proxy.golang.org

```sh
git tag v0.0.1
git push origin --tags
go list -m github.com/whatsauth/wasocket@v0.0.1
```
