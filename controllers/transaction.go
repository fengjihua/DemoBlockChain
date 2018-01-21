package controllers

import "DemoBlockChain/lib"

/*
Transaction :事务
*/
type Transaction struct {
	ID      int64   `bson:"id" json:"id"`
	From    string  `bson:"from" json:"from"`
	To      string  `bson:"to" json:"to"`
	Bitcoin float32 `bson:"bitcoin" json:"bitcoin"`
}

/*
Create :创建事务
*/
func (t *Transaction) Create() (bool, error) {
	t.ID, _ = lib.GetNewUID()
	lib.Log.Debug("Create Transaction:", t)
	return true, nil
}
