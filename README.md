[![license](https://img.shields.io/:license-mit-blue.svg)](https://github.com/ozgur-yalcin/verimor.go/blob/master/LICENSE.md)
[![documentation](https://pkg.go.dev/badge/github.com/ozgur-yalcin/verimor.go)](https://pkg.go.dev/github.com/ozgur-yalcin/verimor.go/src)

# Verimor.go
An easy-to-use verimor.com.tr API with golang

# Installation
```bash
go get github.com/ozgur-yalcin/verimor.go
```

# Usage
```go
package main

import (
	"fmt"

	verimor "github.com/ozgur-yalcin/verimor.go/src"
)

func main() {
	api := new(verimor.API)
	request := verimor.Request{}
	request.Username = "Kullanıcı adı"
	request.Password = "Şifre"
	request.Source = "BASLIK"

	message := new(verimor.Message)
	message.Msg = "test"
	message.No = "905555555555"
	request.Messages = append(request.Messages, message)

	send := api.Sms(request)
	if send {
		fmt.Println("mesaj iletildi")
	} else {
		fmt.Println("hata oluştu")
	}
}
```
