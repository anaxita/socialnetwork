package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gocraft/dbr"
	_ "github.com/lib/pq"
	"log"
	"synergycommunity/internal/bootstrap"
	"synergycommunity/internal/domain/entity"
	"synergycommunity/internal/infrastructure/repository"
	"synergycommunity/pkg/faker"
)

func init() {
	flag.Parse()
}

var (
	posts  = flag.Int("posts", 0, "count posts for generate")
	users  = flag.Int("users", 0, "count users for generate")
	tags   = flag.Int("tags", 0, "count tags for generate")
	groups = flag.Int("groups", 0, "count groups for generate")
	subs   = flag.Int("subs", 0, "subscriptions groups for generate")
)

func main() {
	flag.Parse()

	ctx := context.Background()

	c, err := bootstrap.NewConfig()
	if err != nil {
		log.Println("Config load:", err)

		return
	}

	dbPool, err := bootstrap.NewDBConn(c.DBScheme, c.DBUsername, c.DBPassword, c.DBName, c.DBHost,
		c.DBPort)
	if err != nil {
		log.Println("DB connect:", err)

		return
	}

	defer func(dbPool *dbr.Connection) {
		err := dbPool.Close()
		if err != nil {
			log.Println("close db pool: ", err)
		}
	}(dbPool)

	repo := repository.NewRepository(dbPool)

	f := faker.New()

	if *users > 0 {
		fmt.Println("Generate users...")

		u := f.GenerateUsers(*users)
		for _, v := range u {
			_, err := repo.InsertUser(ctx, v)
			if err != nil {
				log.Println("insert users ", err)
			}
		}

		fmt.Println("Generate users success!")
	}

	if *tags > 0 {
		fmt.Println("Generate tags...")

		t := f.GenerateTags(*tags)
		for _, v := range t {
			_, err := repo.InsertTag(ctx, v)
			if err != nil {
				log.Println("insert tags ", err)
			}
		}

		fmt.Println("Generate tags success!")
	}

	opts := entity.Options{
		Page:  1,
		Limit: 999,
	}

	curUsers, _, err := repo.Users(ctx, opts)
	if err != nil {
		log.Println("[FATAL] get users: ", err)

		return
	}

	curTags, _, err := repo.Tags(ctx, opts)
	if err != nil {
		log.Println("[FATAL] get tags: ", err)

		return
	}

	if *groups > 0 {
		fmt.Println("Generate groups...")

		g := f.GenerateGroups(*groups, curUsers, curTags)
		for _, v := range g {
			_, err = repo.InsertGroup(ctx, v)
			if err != nil {
				log.Println("insert groups ", err)
			}
		}

		fmt.Println("Generate groups success!")
	}

	curGroups, _, err := repo.Groups(ctx, opts)
	if err != nil {
		log.Println("[FATAL] get groups: ", err)

		return
	}

	if *posts > 0 {
		fmt.Println("Generate posts...")

		p := f.GeneratePosts(*posts, curUsers, curGroups, curTags)
		for _, v := range p {
			_, err := repo.InsertPost(ctx, v)
			if err != nil {
				log.Println("insert posts ", err)
			}
		}

		fmt.Println("Generate posts success!")
	}

	if *subs > 0 {
		fmt.Println("Generate subs...")

		p := f.GenerateSubscriptions(*subs, curUsers, curGroups, curTags)
		for _, v := range p {
			err = repo.InsertSubscription(ctx, v)
			if err != nil {
				log.Println("insert subs ", v.Model, v.ModelID, err)
			}
		}

		fmt.Println("Generate subs success!")
	}
}
