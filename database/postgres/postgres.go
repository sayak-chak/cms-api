package postgres

import (
	"cms-api/config"
	"cms-api/custom_errors"
	databaseModels "cms-api/models/database"
	models "cms-api/models/requests"
	responseModels "cms-api/models/responses"
	"cms-api/utils"

	"github.com/go-pg/pg/v10"
)

type Postgres struct{}

func (r *Postgres) GetAuthors() []string {
	var opt, parseErr = pg.ParseURL(config.PostgresConfig)
	if parseErr != nil {
		panic(parseErr)
	}
	db := pg.Connect(opt)
	defer db.Close()

	var authorNames []string
	db.Model(&databaseModels.Author{}).Column("name").Select(&authorNames)

	return authorNames
}

func (r *Postgres) AddSubscriber(subscriberEmail string) error {
	var opt, parseErr = pg.ParseURL(config.PostgresConfig)
	if parseErr != nil {
		return parseErr
	}
	db := pg.Connect(opt)
	defer db.Close()

	_, err := db.Model(&databaseModels.Subscriber{
		Email: subscriberEmail,
	}).OnConflict("DO NOTHING").Insert()

	return err
}

func (r *Postgres) DeleteSubscriber(subscriberEmail string) error {
	var opt, parseErr = pg.ParseURL(config.PostgresConfig)
	if parseErr != nil {
		return parseErr
	}
	db := pg.Connect(opt)
	defer db.Close()

	_, err := db.Model(&databaseModels.Subscriber{}).Where("email =?", subscriberEmail).Delete() //No need to wrap this, does not throw exception
	return err
}

func (r *Postgres) AddContent(addContentRequest *models.AddContentRequest) (int, error) {
	var opt, parseErr = pg.ParseURL(config.PostgresConfig)
	if parseErr != nil {
		return -1, parseErr
	}
	db := pg.Connect(opt)
	defer db.Close()
	utils.LogDbQueries(db)

	content := &databaseModels.Content{
		ImageSrc:      addContentRequest.ImageSrc,
		Title:         addContentRequest.Title,
		Summary:       addContentRequest.Summary,
		AuthorCredsId: addContentRequest.AuthorId,
		Content:       addContentRequest.Body,
		Votes:         1,
	}
	_, err := db.Model(content).Returning("id").Insert() //content.Id has the PQ, this translates to Insert into content..values..return id
	if err != nil {
		return -1, custom_errors.CouldntCompleteYourOperation()
	}

	return content.Id, nil
}

func (r *Postgres) AddContentToTag(tagTableModel interface{}) error {
	var opt, parseErr = pg.ParseURL(config.PostgresConfig)
	if parseErr != nil {
		return parseErr
	}
	db := pg.Connect(opt)
	defer db.Close()

	utils.LogDbQueries(db)

	_, err := db.Model(tagTableModel).Insert()

	if err != nil {
		return custom_errors.CouldntCompleteYourOperation()
	}

	return nil
}

func (r *Postgres) TopContents() (*[]responseModels.TopContentsResponse, error) {
	var opt, parseErr = pg.ParseURL(config.PostgresConfig)
	if parseErr != nil {
		return nil, parseErr
	}
	db := pg.Connect(opt)
	defer db.Close()
	var topContents []responseModels.TopContentsResponse
	utils.LogDbQueries(db)

	err := db.Model(&databaseModels.Author{}).
		Column(`author.name`, `a.id`, `a.author_creds_id`, `a.content`, `a.image_src`, `a.title`, `a.summary`, `a.votes`).
		Join(`INNER JOIN 
			(SELECT * FROM contents ORDER BY votes DESC LIMIT ?) 
			AS a 
			ON author.author_creds_id=a.author_creds_id`,
			config.NumberOfTopEntriesToConsider).
		Select(&topContents)
	if err != nil {
		return nil, err
	}

	return &topContents, nil
}

func (r *Postgres) TopContentsByTag(tagTable string) (*[]responseModels.TopContentsResponse, error) {
	var opt, parseErr = pg.ParseURL(config.PostgresConfig)
	if parseErr != nil {
		return nil, parseErr
	}
	db := pg.Connect(opt)
	defer db.Close()
	var topContents []responseModels.TopContentsResponse
	utils.LogDbQueries(db)

	//SQL injection is not possible here because tagTable isn't user input
	err := db.Model(&databaseModels.Author{}).Column(`author.name`, `a.id`, `a.author_creds_id`, `a.content`, `a.image_src`, `a.title`, `a.summary`, `a.votes`).
		Join(`INNER JOIN
			(SELECT * FROM contents INNER JOIN `+tagTable+` ON contents.id = `+tagTable+`.content_id ORDER BY contents.votes LIMIT ?) 
			AS a 
			ON author.author_creds_id=a.author_creds_id`,
			config.NumberOfLastEntriesToConsiderWhenSearchingByTag).
		Select(&topContents)
	if err != nil {
		return nil, err
	}
	return &topContents, nil
}

func (r *Postgres) GetPhoneFor(otp string) (int, error) {
	var opt, parseErr = pg.ParseURL(config.PostgresConfig)
	if parseErr != nil {
		return -1, parseErr
	}
	db := pg.Connect(opt)
	defer db.Close()
	var mobileNumber int
	utils.LogDbQueries(db)

	_, err := db.Model(&databaseModels.AuthorTempCreds{}).Where("temp_password = ?", otp).Returning("mobile").Delete(&mobileNumber)
	if err != nil {
		return -1, err
	}
	return mobileNumber, nil
}

func (r *Postgres) CleanUpOldOtps(currTime int64, otpExpiryPeriod int) error {
	var opt, parseErr = pg.ParseURL(config.PostgresConfig)
	if parseErr != nil {
		return parseErr
	}
	db := pg.Connect(opt)
	defer db.Close()
	utils.LogDbQueries(db)

	_, err := db.Model(&databaseModels.AuthorTempCreds{}).Where("? - creation_time > ?", currTime, otpExpiryPeriod).Delete()
	return err
}

func (r *Postgres) AddAuthorCreds(newAuthorAccCreationReq *models.NewAuthorAccCreationRequest, phone int, hashedPassword string) (int, error) {
	var opt, parseErr = pg.ParseURL(config.PostgresConfig)
	if parseErr != nil {
		return -1, parseErr
	}
	db := pg.Connect(opt)
	defer db.Close()
	utils.LogDbQueries(db)
	newAuthorAccCreds := &databaseModels.AuthorCreds{
		Mobile:         phone,
		HashedPassword: hashedPassword,
	}
	_, err := db.Model(newAuthorAccCreds).Insert()
	if err != nil {
		return -1, err
	}
	return newAuthorAccCreds.Id, nil
}

func (r *Postgres) AddAuthorDetails(newAuthorAccCreationReq *models.NewAuthorAccCreationRequest, authorCredId int) error {
	var opt, parseErr = pg.ParseURL(config.PostgresConfig)
	if parseErr != nil {
		return parseErr
	}
	db := pg.Connect(opt)
	defer db.Close()

	utils.LogDbQueries(db)

	_, err := db.Model(&databaseModels.Author{
		AuthorCredsId: authorCredId,
		Name:          newAuthorAccCreationReq.Name,
	}).Insert()
	return err
}

func (r *Postgres) GetPasswordAndIdFor(mobileNumber int) (string, int, error) {
	var opt, parseErr = pg.ParseURL(config.PostgresConfig)
	if parseErr != nil {
		return "", -1, parseErr
	}
	db := pg.Connect(opt)
	defer db.Close()

	var author databaseModels.AuthorCreds
	utils.LogDbQueries(db)

	err := db.Model(&author).Where("mobile = ?", mobileNumber).Select(&author)
	if err != nil {
		return "", -1, err
	}

	return author.HashedPassword, author.Id, nil
}

func (r *Postgres) GetContent(contentId int) (*responseModels.ReadContentResponse, error) {
	var opt, parseErr = pg.ParseURL(config.PostgresConfig)
	if parseErr != nil {
		return nil, parseErr
	}
	db := pg.Connect(opt)
	defer db.Close()
	var content responseModels.ReadContentResponse
	utils.LogDbQueries(db)

	err := db.Model(&databaseModels.Author{}).Column(`a.id`, `a.content`, `a.title`, `a.summary`, `a.votes`, `author.name`, `author.author_creds_id`).
		Join(`INNER JOIN (SELECT * FROM contents WHERE id=?) AS a ON a.author_creds_id = author.author_creds_id`, contentId).Select(&content)
	if err != nil {
		return nil, err
	}
	return &content, nil
}

func (r *Postgres) Upvote(aricleId int) error {
	var opt, parseErr = pg.ParseURL(config.PostgresConfig)
	if parseErr != nil {
		return parseErr
	}
	db := pg.Connect(opt)
	defer db.Close()

	utils.LogDbQueries(db)

	_, err := db.Model(&databaseModels.Content{}).Set(`votes = votes + 1`).Where(`id=?`, aricleId).Update()
	return err
}
