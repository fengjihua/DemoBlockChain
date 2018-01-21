package models

/*
Transaction :事物
*/
type Transaction struct {
	ID      int64   `bson:"id" json:"id"`
	From    string  `bson:"from" json:"from"`
	To      string  `bson:"to" json:"to"`
	Bitcoin float32 `bson:"bitcoin" json:"bitcoin"`
}
