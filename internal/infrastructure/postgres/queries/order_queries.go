package queries

const (
	OrderSchema = `CREATE TABLE IF NOT EXISTS Orders(
		id uuid DEFAULT uuid_generate_v1(),
		status VARCHAR NOT NUll,
		description VARCHAR NOT NULL,
		price decimal(5,2) NOT NULL,
		createdOn timestamp NOT NULL DEFAULT NOW(),
		PRIMARY KEY (id)
	);`
)
