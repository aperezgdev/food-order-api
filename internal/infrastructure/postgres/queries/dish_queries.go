package queries

const (
	DishSchema = `CREATE TABLE IF NOT EXISTS Dishes(
		id uuid DEFAULT uuid_generate_v1(),
		name VARCHAR NOT NUll,
		description VARCHAR NOT NULL,
		price decimal(5,2) NOT NULL,
		createdOn timestamp NOT NULL DEFAULT NOW(),
		PRIMARY KEY (id)
	);`
	DishGetAll = `SELECT * FROM Dishes`
	DishCreate = `INSERT INTO Dishes(name, description, price) VALUES (:name, :description, :price)`
	DishUpdate = `UPDATE Dishes SET name= :name, description= :description, price= :price WHERE id= :id`
	DishDelete = `DELETE FROM Dishes WHERE id = $1`
)
