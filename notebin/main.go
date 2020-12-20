package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"xorm.io/xorm"
)

var db *xorm.Engine

// ranking table structure
type Notes struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Uuid    string `json:"uuid"`
}

func main() {
	// init things
	db, _ = xorm.NewEngine("sqlite3", "./info.db")
	db.Sync2(&Notes{})
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})
	router.GET("/note", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, "<h1 id=\"title\"></h1>\n<div id=\"note\"></div>\n<script src=\"https://cdn.jsdelivr.net/npm/dompurify@2.2.2/dist/purify.min.js\"></script>\n<script>\n    // Debugging by uncommenting following line\n    // let _debug = {key: \"sha1:f03e8a370aa8dc80f63a6d67401a692ae72fa530\"}\n    fetch(\"/note/\" + window.location.hash.substr(1)).\n        then(response => {\n            return response.json()\n        }).\n        then(data => {\n            // no xss in production env\n            if (window._debug && window._debug.key == \"sha1:f03e8a370aa8dc80f63a6d67401a692ae72fa530\"){\n                title = data.title\n                console.log(title)\n            }\n            else\n                title = DOMPurify.sanitize(data.title);\n            document.getElementById(\"title\").innerHTML = title;\n\n\n        if (window._debug && window._debug.key == \"sha1:f03e8a370aa8dc80f63a6d67401a692ae72fa530\"){\n                content = data.content\n                console.log(content)\n            }\n            else\n                content = DOMPurify.sanitize(data.content);\n            document.getElementById(\"note\").innerHTML = content;\n        })\n</script>\n\n")
	})

	router.GET("/note/:uuid", func(c *gin.Context) {
		uuid := c.Param("uuid")
		var n Notes
		db.
			Where("uuid = ?", uuid).
			Get(&n)
		c.JSON(200, gin.H{
			"title":   n.Title,
			"content": n.Content,
		})
	})

	router.POST("/note", func(c *gin.Context) {
		content := c.DefaultPostForm("content", "NOTHING")
		title := c.DefaultPostForm("title", "NOTHING")
		u := uuid.New().String()
		newNote := Notes{
			Content: content,
			Title:   title,
			Uuid:    u,
		}
		db.Insert(&newNote)
		c.Redirect(301, "/note#"+u)
	})

	router.Run(":8080")
}
