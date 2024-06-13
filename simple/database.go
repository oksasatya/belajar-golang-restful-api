package simple

type Database struct {
	Name string
}

type DatabasePostgreSQL Database
type DatabaseMongoDB Database

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return (*DatabaseMongoDB)(&Database{Name: "MongoDB "})
}

func NewDatabasePostgreSQL() *DatabasePostgreSQL {
	return (*DatabasePostgreSQL)(&Database{Name: "PostgreSQL"})
}

type DatabaseRepository struct {
	DatabasePostgreSSQL *DatabasePostgreSQL
	DatabaseMongoDB     *DatabaseMongoDB
}

func NewDatabaseRepository(postgreSQL *DatabasePostgreSQL, mongodb *DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{DatabasePostgreSSQL: postgreSQL, DatabaseMongoDB: mongodb}
}
