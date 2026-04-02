package dialects

import (
	"fmt"
	"time"

	"github.com/pressly/goose/v3/database/dialect"
)

// NewManticoreMysql returns a new [dialect.Querier] for NewManticoreMysql dialect.
func NewManticoreMysql() dialect.QuerierExtender {
	return &manticoreMysql{}
}

type manticoreMysql struct{}

var _ dialect.QuerierExtender = (*manticoreMysql)(nil)

func (m *manticoreMysql) CreateTable(tableName string) string {
	q := `CREATE TABLE IF NOT EXISTS %s (
       version_id bigint,
       is_applied int,
       tstamp timestamp
    )`
	return fmt.Sprintf(q, tableName)
}

func (m *manticoreMysql) InsertVersion(tableName string) string {
	now := time.Now().Unix()
	q := `INSERT INTO %s (version_id, is_applied, tstamp) VALUES (?, ?, %d)`
	return fmt.Sprintf(q, tableName, now)
}

func (m *manticoreMysql) DeleteVersion(tableName string) string {
	q := `DELETE FROM %s WHERE version_id = ?`
	return fmt.Sprintf(q, tableName)
}

func (m *manticoreMysql) GetMigrationByVersion(tableName string) string {
	return fmt.Sprintf("SELECT tstamp, is_applied FROM %s WHERE version_id = ? LIMIT 1", tableName)
}

func (m *manticoreMysql) ListMigrations(tableName string) string {
	q := `SELECT version_id, is_applied FROM %s ORDER BY id DESC`
	return fmt.Sprintf(q, tableName)
}

func (m *manticoreMysql) GetLatestVersion(tableName string) string {
	return fmt.Sprintf("SELECT IFNULL(MAX(version_id), 0) FROM %s", tableName)
}

func (m *manticoreMysql) TableExists(tableName string) string {
	return fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
}
