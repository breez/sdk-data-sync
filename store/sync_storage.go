package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/markbates/pkger"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/pkger"
	_ "github.com/mattn/go-sqlite3"
)

var _ = pkger.Include("/store/migrations")
var ErrSetConflict = errors.New("set conflict")

type StoredRecord struct {
	Id       string
	Data     []byte
	Revision int64
}
type SyncStorage interface {
	SetRecord(ctx context.Context, id string, existingRevision int64, data []byte) error
	ListChanges(ctx context.Context, id string, sinceRevision int64) ([]StoredRecord, error)
}

type SQLiteSyncStorage struct {
	db *sql.DB
}

func Connect(file string) (*SQLiteSyncStorage, error) {
	if err := os.MkdirAll(path.Dir(file), 0700); err != nil {
		return nil, fmt.Errorf("failed to create databases directory %w", err)
	}
	needMigration := false
	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		needMigration = true
	}
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite3 database %w", err)
	}

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to create migration driver %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"pkger:///store/migrations",
		file, driver)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate migrations %w", err)
	}
	if needMigration {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			return nil, fmt.Errorf("failed to run migrations %w", err)
		}
	}
	return &SQLiteSyncStorage{db: db}, nil
}

func (s *SQLiteSyncStorage) SetRecord(ctx context.Context, id string, data []byte, existingRevision int64) (int64, error) {
	var revision int64
	err := s.db.QueryRow("SELECT revision FROM records WHERE id = ?", id).Scan(&revision)
	if err != sql.ErrNoRows {
		if err != nil {
			return 0, fmt.Errorf("failed to get record's latest revision: %w", err)
		}
		if existingRevision != revision {
			return 0, ErrSetConflict
		}
	}

	res, err := s.db.Exec("INSERT OR REPLACE INTO records (id, data) VALUES (?, ?)", id, data)
	if err != nil {
		return 0, fmt.Errorf("failed to insert record: %w", err)

	}
	return res.LastInsertId()
}

func (s *SQLiteSyncStorage) ListChanges(ctx context.Context, sinceRevision int64) ([]StoredRecord, error) {
	rows, err := s.db.Query("SELECT id, data, revision FROM records WHERE revision > ?", sinceRevision)
	if err != nil {
		return nil, fmt.Errorf("failed to query records: %w", err)
	}
	defer rows.Close()

	records := make([]StoredRecord, 0)
	for rows.Next() {
		record := StoredRecord{}
		err = rows.Scan(&record.Id, &record.Data, &record.Revision)
		if err != nil {
			return nil, fmt.Errorf("failed to scan record: %w", err)
		}
		records = append(records, record)
	}
	return records, nil
}
