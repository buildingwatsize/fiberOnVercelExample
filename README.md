# Fiber on Vercel example

A demonstration of the deployment of [GoFiber](https://gofiber.io/) to [Vercel](https://vercel.com/) using the integrated [https://github.com/gofiber/adaptor](https://github.com/gofiber/adaptor) package; for more details, see [/api/index.go](/api/index.go)

> Bonus: I added the function `/api/http.go` for `net/http` package.

## How to test?

### Fiber Version

Go to <https://fiber-on-vercel-example.vercel.app/api>

Expected: Return string response like below

```text
Hello World 👋!
```

### net/http Version

Go to <https://fiber-on-vercel-example.vercel.app/api/http>

Expected: Return JSON like below

```json
{"msg":"Hello World!"}
```

> made with ❤️ [watsize](https://github.com/buildingwatsize)
