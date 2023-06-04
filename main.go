package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stellayazilim/stella.backend.tenants/modules"
)

func main() {
	r := gin.Default()
	modules.MainModule(r)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
