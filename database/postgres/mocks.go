package postgres

import (
	"cms-api/config"
	"strconv"

	"github.com/go-pg/pg/v10"
)

func (p *Postgres) SeedWithMockData(env string) error {
	if env == "local" {
		var opt, parseErr = pg.ParseURL(config.PostgresConfig)
		if parseErr != nil {
			return parseErr
		}
		db := pg.Connect(opt)
		defer db.Close()

		err := cleanupExistingData(db)
		if err != nil {
			return err
		}
		err = fillWithNewData(db)
		if err != nil {
			return err
		}
	}
	return nil
}

func cleanupExistingData(db *pg.DB) error {
	_, err := db.Query(nil, "delete from actions")
	if err != nil {
		return err
	}
	_, err = db.Query(nil, "delete from adventures")
	if err != nil {
		return err
	}
	_, err = db.Query(nil, "delete from dramas")
	if err != nil {
		return err
	}
	_, err = db.Query(nil, "delete from horrors")
	if err != nil {
		return err
	}
	_, err = db.Query(nil, "delete from contents")
	if err != nil {
		return err
	}
	_, err = db.Query(nil, "delete from authors")
	if err != nil {
		return err
	}
	_, err = db.Query(nil, "delete from author_creds")
	if err != nil {
		return err
	}
	_, err = db.Query(nil, "delete from subscribers")
	return err
}

func fillWithNewData(db *pg.DB) error {
	_, err := db.Query(nil, `insert into author_creds values (1,0000000000,'test_password')`)
	if err != nil {
		return err
	}
	_, err = db.Query(nil, `insert into author_creds values (2,1111111111,'test_password')`)
	if err != nil {
		return err
	}
	_, err = db.Query(nil, `insert into authors values (1,'Author One')`)
	if err != nil {
		return err
	}
	_, err = db.Query(nil, `insert into authors values (2,'Author Two')`)
	if err != nil {
		return err
	}
	for contentId := 0; contentId < 1000; contentId++ {
		contentIdStr := strconv.FormatInt(int64(contentId), 10)
		numOfVotes := contentId
		contentBody := "Content Body " + contentIdStr
		contentTitle := "Content Title " + contentIdStr
		contentSummary := "Content Summary " + contentIdStr
		authorId := 1
		if contentId%2 == 0 {
			authorId = 2
		}
		_, err = db.Query(nil, `insert into contents values (?,?,?, 'https://cdn.pixabay.com/photo/2021/08/25/20/42/field-6574455_1280.jpg', ?,?,?)`, contentId, authorId, contentBody, contentTitle, contentSummary, numOfVotes)
		if err != nil {
			return err
		}
		if contentId%2 == 0 {
			_, err = db.Query(nil, `insert into actions values (?)`, contentId)
			if err != nil {
				return err
			}
		}
		if contentId%3 == 0 {
			_, err = db.Query(nil, `insert into adventures values (?)`, contentId)
			if err != nil {
				return err
			}
		}
		if contentId%4 == 0 {
			_, err = db.Query(nil, `insert into dramas values (?)`, contentId)
			if err != nil {
				return err
			}
		}
		if contentId%5 == 0 {
			_, err = db.Query(nil, `insert into horrors values (?)`, contentId)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
