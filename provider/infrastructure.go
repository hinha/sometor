package provider

import (
	"bytes"
	"context"
	"errors"
)

// ErrCacheMiss returned when value from cache is not found
//var ErrCacheMiss = errors.New("cache miss")

// ErrDBNotFound returned when there is no data found in the database
var ErrDBNotFound = errors.New("data not found")

// S3Management is bucket management transaction
type S3Management interface {
	PutObject(pathString string, fileReader *bytes.Reader) error
	DownloadObject(pathObject string) (string, error)
}

// DB is database interface wrapper for *sql.DB
type DB interface {
	Transaction(ctx context.Context, transactionKey string, f func(tx TX) error) error
	ExecContext(ctx context.Context, queryKey, query string, args ...interface{}) (Result, error)
	QueryContext(ctx context.Context, queryKey, query string, args ...interface{}) (Rows, error)
	QueryRowContext(ctx context.Context, queryKey, query string, args ...interface{}) Row
}

// TX is database transaction
type TX interface {
	ExecContext(ctx context.Context, queryKey, query string, args ...interface{}) (Result, error)
	QueryContext(ctx context.Context, queryKey, query string, args ...interface{}) (Rows, error)
	QueryRowContext(ctx context.Context, queryKey, query string, args ...interface{}) Row
}

// A Result summarizes an executed SQL command.
type Result interface {
	// LastInsertId returns the integer generated by the database
	// in response to a command. Typically this will be from an
	// "auto increment" column when inserting a new row. Not all
	// databases support this feature, and the syntax of such
	// statements varies.
	LastInsertId() (int64, error)

	// RowsAffected returns the number of rows affected by an
	// update, insert, or delete. Not every database or database
	// driver may support this.
	RowsAffected() (int64, error)
}

// Row single result of database query
type Row interface {
	Scan(dest ...interface{}) error
}

// Rows multiple result of database query
type Rows interface {
	Close() error
	Columns() ([]string, error)
	Err() error
	Next() bool
	NextResultSet() bool
	Scan(dest ...interface{}) error
}
