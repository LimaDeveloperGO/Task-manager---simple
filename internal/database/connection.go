package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

type DB struct {
	*sql.DB
}

func NewConnection() (*DB, error) {
	dbDir := "database"
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		return nil, fmt.Errorf("erro ao criar diretório database: %v", err)
	}

	dbPath := filepath.Join(dbDir, "tasks.db")

	sqlDB, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar com SQLite: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("erro ao testar conexão: %v", err)
	}

	db := &DB{sqlDB}

	if err := db.runMigrations(); err != nil {
		return nil, fmt.Errorf("erro ao executar migrações: %v", err)
	}

	log.Println("✅ Conectado ao SQLite com sucesso!")
	return db, nil
}

func (db *DB) Close() error {
	return db.DB.Close()
}
