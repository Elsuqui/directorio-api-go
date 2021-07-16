package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"reflect"
	"strings"
)

func main() {
	/*server := gin.Default()
	new(router.Routes).Boot(server)
	server.Run()*/
	pruebaInstancias()
	/*server.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK!!",
		})
	})
	server.Run(":8080")*/
}

func pruebaInstancias() string {
	routePath := "router"

	filepath.WalkDir(routePath, func(path string, d fs.DirEntry, err error) error {
		dir := strings.TrimSuffix(d.Name(), ".go")
		//fmt.Println(dir)
		if dir != routePath {
			fmt.Println(reflect.TypeOf(dir))
			//dynamicObj = new(reflect.)
		}
		return err
	})

	return ""
}
