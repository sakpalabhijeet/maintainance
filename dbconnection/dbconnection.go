package dbconnection

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DBConnection(ctx context.Context)(*pgxpool.Pool, error) {
	dsn:= "maintenance://postgress:12345@localhost:5433/postgress?sslmode=disable"
	config,err:= pgxpool.ParseConfig(dsn)
	if err!= nil{
		log.Fatalf("Unable to parse config %v \n", err)
		return  nil, err
	}
	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = 30* time.Minute

	db, err:= pgxpool.NewWithConfig(ctx, config)
	if err!= nil{
		log.Fatalf("Unable to Pool new config %v\n",err)
		return  nil, err
	}
	if err:= db.Ping(ctx);err!=nil{
		log.Fatalf("Unable to ping the database %v\n", err)
		return nil, err
	}
	fmt.Println("Connection to Database established successfully!")
	return  db, nil
}