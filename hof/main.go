package main

import (
	"context"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/shomali11/slacker"
	"log"
	"os"
	"xorm.io/xorm"
)

var db *xorm.Engine

// ranking table structure
type Rankings struct {
	Id       int64  `json:"id"`
	Rank     int    `json:"rank"`
	TeamName string `json:"team_name"`
}

// list
func listAllHandler(_ slacker.BotContext, _ slacker.Request, response slacker.ResponseWriter) {
	var rankings []Rankings
	// query database
	_ = db.Table("rankings").Find(&rankings)
	// convert array of ranking objects to readable string
	var respText string
	for _, ranking := range rankings {
		respText += fmt.Sprintf("%d - %s\n", ranking.Rank, ranking.TeamName)
	}
	// send response
	response.Reply(respText)
}

// rank <team_name>
func rankHandler(_ slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
	teamName := request.Param("team_name")
	var r Rankings
	// query database
	_, _ = db.
		Where(fmt.Sprintf("team_name = '%s'", teamName)).
		Get(&r)
	// send response
	response.Reply(fmt.Sprintf("%d - %s", r.Rank, r.TeamName))
}

func main() {
	// init things
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"))
	db, _ = xorm.NewEngine("sqlite3", "./info.db")
	// list definition
	bot.Command("list", &slacker.CommandDefinition{
		Description: "Show top 20 teams' rankings in last WeCTF.",
		Handler:     listAllHandler,
	})
	// rank <team_name> definition
	bot.Command("rank <team_name>", &slacker.CommandDefinition{
		Description: "Show the ranking of a team in last WeCTF.",
		Example:     "rank by7ch",
		Handler:     rankHandler,
	})
	// blablabla copied from docs
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
