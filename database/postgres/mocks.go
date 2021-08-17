package postgres

import (
	// "cms-api/config"
	"cms-api/config"
	models "cms-api/models/database"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func (p *Postgres) SeedWithMockData() error {
	var opt, parseErr = pg.ParseURL(config.PostgresConfig)
	if parseErr != nil {
		return parseErr
	}
	db := pg.Connect(opt)
	defer db.Close()

	modelList := []interface{}{
		(*models.AuthorCreds)(nil),
		(*models.Author)(nil),
		(*models.AuthorTempCreds)(nil),
		(*models.Subscriber)(nil),
		(*models.Content)(nil),
		(*models.Action)(nil),
		(*models.Adventure)(nil),
		(*models.Drama)(nil),
		(*models.Horror)(nil),
	}

	for _, tableModel := range modelList {
		err := db.Model(tableModel).CreateTable(&orm.CreateTableOptions{
			Temp:          false,
			IfNotExists:   true,
			FKConstraints: true,
		})
		if err != nil {
			return err
		}
	}

	err := createIndexes(db)
	if err != nil {
		return err
	}
	return nil
}

func createIndexes(db *pg.DB) error {
	_, err := db.Exec(`CREATE INDEX votes ON contents (votes);`)
	if err != nil {
		return err
	}
	_, err = db.Exec(`CREATE INDEX mobile ON author_creds USING hash (mobile)`)
	return err
}

// func fillTagData(db *pg.DB) error {
// 	for _, genre := range config.TagList {
// 		_, err := db.Model(&models.Tag{
// 			Type: genre,
// 		}).Insert()
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
