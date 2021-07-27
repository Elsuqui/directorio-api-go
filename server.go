package main

import (
	"apigolang/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load env file to set security or credentials information
	err := godotenv.Load()
	if err != nil {
		panic("Env file is required")
	}
	server := gin.Default()
	new(router.Routes).Boot(server)
	server.Run()
}

/*func pruebaInstancias() string {
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
}*/
