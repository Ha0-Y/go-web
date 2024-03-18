# 模板

模板引擎
```txt
text/template
html/template
```

html/template包实现了数据驱动的模板，用于生成可对抗代码注入的安全HTML输出。它提供了和text/template包相同的接口，Go语言中输出HTML的场景都应使用text/template包。

## 解析模板

ParseFiles：可以解析多个文件
```go
//func ParseFiles(filenames ...string) (*Template, error) {
//    return parseFiles(nil, readFileOS, filenames...)
//}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("server/index.html")
	t.Execute(w, "hello")
}
```

ParseGlob：正则匹配文件
```go
func ParseGlob(pattern string) (*Template, error) {
	return parseGlob(nil, pattern)
}
```

## 执行模板

Execute: 只支持解析一个模板
```go
func (t *Template) Execute(wr io.Writer, data any) error {
	return t.execute(wr, data)
}
```

ExecuteTemplate: 指定模板名 `name`
```go
func (t *Template) ExecuteTemplate(wr io.Writer, name string, data any) error {
	tmpl := t.Lookup(name)
	if tmpl == nil {
		return fmt.Errorf("template: no template %q associated with template %q", name, t.name)
	}
	return tmpl.Execute(wr, data)
}
```


## action

`{{.}}` 就是一个action
```html
<!DOCTYPE html>
<html lang="en" xmlns="http://www.w3.org/1999/html">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
    {{.}}
</body>
</html>
```

if 表达
```html
{{if pipeline}} T1 {{else if pipeline}} T0 {{end}}
```

range
```html
{{range pipeline}} T1 {{end}}
```

定义变量
```html
{{ $variable := pipeline }}
```