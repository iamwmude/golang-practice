package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	router.GET("/test", tableView)

	router.Run(":8080")
}

func tableView(c *gin.Context) {
	c.HTML(http.StatusOK, "user/test.tmpl", getTableInfo())
}

type playerInfo struct {
	Account    string
	Position   int
	Is_banker  bool
	Bet        int
	Profit     int
	Tax        int
	Bet_status int
}

func getTableInfo() gin.H {
	res := gin.H{}
	t := template.Must(template.New("").Parse(`<table>{{range .}}<tr><td>{{.Account}}</td><td>{{.Position}}</td><td>{{.Is_banker}}</td><td>{{.Bet}}</td><td>{{.Profit}}</td><td>{{.Tax}}</td><td>{{.Bet_status}}</td></tr>{{end}}</table>`))

	data := []playerInfo{
		playerInfo{
			Account:    "a1",
			Position:   0,
			Is_banker:  false,
			Bet:        9999999,
			Profit:     -100,
			Tax:        0,
			Bet_status: 1,
		},
		playerInfo{
			Account:    "a2",
			Position:   0,
			Is_banker:  false,
			Bet:        100,
			Profit:     190,
			Tax:        10,
			Bet_status: 2,
		},
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, &data); err != nil {
		log.Fatal(err)
	}
	res["body"] = template.HTML(tpl.String())
	return res
}
