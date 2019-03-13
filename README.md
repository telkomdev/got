### Got

a very simple Go's struct template parser

#### Usage

- Install
```shell
$ go get github.com/Bhinneka/got
```

- Parse Template to `stdio`

you can create a custom function inside your template:
```go
	funcMaps := template.FuncMap{
		"toUpper": func(v string) string {
			return strings.ToUpper(v)
		},
	}
```

example:
```go
// Person struct
type Person struct {
	ID   string
	Name string
}

func main() {

	data := Person{
		ID:   "1",
		Name: "Wuriyanto",
	}

	tmpl := `
	Hello
	Your ID : {{ .ID }}
	Your Name : {{ toUpper .Name }}
	`

	funcMaps := template.FuncMap{
		"toUpper": func(v string) string {
			return strings.ToUpper(v)
		},
	}

	err := got.Parse(tmpl, data, os.Stdout, funcMaps, got.ParseTemplateText)

	if err != nil {
		fmt.Println(err)
    }
    
    //output :

    //Hello
    //Your ID : 1
    //Your Name : WURIYANTO

}
```

- Parse template file
`my_template.tmpl`
```html
<html>
    <head>
        <title>Example</title>
    </head>
    <body>
        <h1>{{ .OrderID }}</h1>
        <ul>
            {{- range $item := .Items}}
                <li>{{ $item }}</li>
            {{- end }}
        </ul>
    </body>
</html>
```

`Order.go struct`
```go
// Order struct
type Order struct {
	OrderID string
	Items   []string
}
```

Parse this go struct to `my_template.tmpl` file
```go
func main() {
	myOrders := Order{
		OrderID: "00881119977",
		Items:   []string{"Coffee mix", "Tea", "Mendoan", "Gorengan", "Onde-onde"},
	}

	var target strings.Builder

	err := got.Parse("my_template.tmpl", myOrders, &target, nil, got.ParseTemplateFile)

	if err != nil {
		fmt.Println(err)
	}

    fmt.Println(target.String())
    
    //output:

    // <html>
    //     <head>
    //         <title>Example</title>
    //     </head>
    //     <body>
    //         <h1>00881119977</h1>
    //         <ul>
    //                 <li>Coffee mix</li>
    //                 <li>Tea</li>
    //                 <li>Mendoan</li>
    //                 <li>Gorengan</li>
    //                 <li>Onde-onde</li>
    //         </ul>
    //     </body>
    // </html>
}
```

#

#### Copyright (c) 2019 bhinneka.com