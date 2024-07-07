package entity

type Link struct {
	ID       string `bson:"_id"`
	Original string `bson:"original"`
	Code     string `bson:"code"`
}
