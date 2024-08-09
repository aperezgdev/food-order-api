package queries

const (
	UserSchema = `CREATE TABLE IF NOT EXISTS Users(
		id uuid DEFAULT uuid_generate_v1(),
		name VARCHAR NOT NUll,
		email VARCHAR NOT NULL,
		PRIMARY KEY (id)
	);`
	UserCreate = `INSERT INTO Users(name, email) VALUES (:name, :email)`
)
