# server

## 简易server

```go
package main

import `net/http`

func main() {
	http.ListenAndServe(":8080", nil)
}
```

我们如果跟一下源码，就会发现本质是使用http.Server的方法，因此可以等价于如下
```go
package main

import `net/http`

func main() {
	s := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	s.ListenAndServe()
}
```

## handler

一个接口，需要我们实现这个接口就能实现路由处理
```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

// 类型转化
type HandlerFunc func(ResponseWriter, *Request)
```

简单的路由处理
```go
package main

import `net/http`

func rootHander(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("root"))
}

func main() {
	s := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	http.HandleFunc("/", http.HandlerFunc(rootHander))
	s.ListenAndServe()
}
```

### 内置Handler

```go
func AllowQuerySemicolons(h Handler) Handler
func FileServer(root FileSystem) Handler
func MaxBytesHandler(h Handler, n int64) Handler
// 404
func NotFoundHandler() Handler
// 跳转
func RedirectHandler(url string, code int) Handler
// 移除url前缀后在调用handler
func StripPrefix(prefix string, h Handler) Handler
func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler
```

### Get Request

gquery字段：查询是在url中，因此可以通过URL获得
```go
func loginHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")
	fmt.Printf("ID: %v\n", id)
	w.Write([]byte(fmt.Sprintf("Hello %v", id)))
}
```

返回值是 `type Values map[string][]string` 
- 可以是 `?id=1&id=2` 这种类型
- Get方法只取第一个值。
- 可以使用 `map[key]` 来取值

### Post Request

#### Form

通过form表单来处理post请求, enctype
- 简单文本，使用 `application/x-www-form-urlencoded`，此种方法只会简单的url编码
- 大量数据，使用 `multipart/form-data`，此种数据会为每一个数据生成MIME

```html
<form method="post" action="http://localhost:8080/register" enctype="application/x-www-form-urlencoded">
    <input type="text" name="name" />
    <input type="text" name="password" />
    <input type="submit" />
</form>
```

然后在go处理
```go
// Form url.Values => type Values map[string][]string
func registerHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.Write([]byte(fmt.Sprintf("Register %v", r.Form.Get("name"))))
}
```

但是Form会处理Get里的字段，可以使用`PostForm`只处理Post


但是改成下面的，就会无法请求
```html
<form method="post" action="http://localhost:8080/register" enctype="multipart/form-data">
    <input type="text" name="name" />
    <input type="text" name="password" />
    <input type="submit" />
</form>
```

我们需要保证后端解析方法和发送数据的方法一致
- `ParseMultipartForm` 可以解析，并且使用 `r.MultipartForm`来获得数据

或者直接使用 `FormValue` 函数来获得值，post数据使用`PostFormValue`
```go
func (r *Request) FormValue(key string) string {
	if r.Form == nil {
		r.ParseMultipartForm(defaultMaxMemory)
	}
	if vs := r.Form[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}
```

#### 上传文件

定义form表单
```html
<form method="post" action="http://localhost:8080/upload" enctype="multipart/form-data">
    <input type="file" name="upload" />
    <input type="submit" />
</form>
```

后端处理
```go
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("upload")
	if err != nil {
		fmt.Fprint(w, "Error uploading")
		return
	}
	data, _ := io.ReadAll(file)
	fmt.Fprintf(w, string(data))
}
```

### Response

设置状态码
```go
w.WriteHeader(302) // 重定向302
```

设置header
```go
w.Header().Set("Content-Type", "application/json")
```



