package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	"github.com/stellayazilim/stella.backend.tenants.modules.user/module"
)

func main() {

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// registering validation for nontoneof
		v.RegisterValidation("isString", func(fl validator.FieldLevel) bool {
			// split values using ` `. eg. notoneof=bob rob job
			fmt.Println(fl)
			return false
		})
	}
	r := gin.Default()
	module.MainModule(r)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
