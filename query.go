package bboy

import "encoding/json"

type Query interface {
	Bucket(string) Query
	Get(string) interface{}
	SetDB(*BBoy) Query
}

type Simple struct {
	db     *BBoy
	bucket string
	key    string
}

func NewSimple(db *BBoy) *Simple {
	return &Simple{
		db: db,
	}
}

func (q *Simple) SetDB(b *BBoy) Query {
	q.db = b
	return q
}

func (q *Simple) Bucket(b string) Query {
	return q
}

func (q *Simple) Get(k string) interface{} {
	return ""
}

// EntityGen use Simple Query for generating entities from the query's result
type Json struct {
	simple *Simple
}

func NewJson(db *BBoy) *Json {
	return &Json{
		simple: &Simple{
			db: db,
		},
	}
}

func (q *Json) SetDB(b *BBoy) Query {
	q.simple.db = b
	return q
}

func (q *Json) Bucket(b string) Query {
	q.simple.bucket = b
	return q
}

func (q *Json) Get(k string) interface{} {
	s := q.Get(k).(string)

	err := json.Unmarshal()
	return
}
