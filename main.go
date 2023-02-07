package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type DD interface {
	FF()
}
type SS struct{}

func (s SS) FF() {
	fmt.Println("jj")
}
func queryDB(db sql.DB, a string) {
	db.Exec("select * from tt where name = ?", a)
	var aa any
	if ll, ok := aa.(int); ok {
		fmt.Println(ll)
	}
	// var tt *bytes.Buffer
	// var tt io.Writer
}

func main() {
	// a := 0.1
	// b := 0.2
	// c := a + b
	// fmt.Println(c)
	// return
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		hostname, _ := os.Hostname()
		c.JSON(http.StatusOK, gin.H{
			"message": "you've hit " + hostname,
		})
	})
	r.Run()

}
