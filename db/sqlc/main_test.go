package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/patrickchap/clipsapi/util"
)


var testQueries Store 

var testConn *pgxpool.Pool

func TestMain(m * testing.M){

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = NewStore(conn)


	user, err := testQueries.GetUser(context.Background(), 2)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(user)
	os.Exit(m.Run())
}
