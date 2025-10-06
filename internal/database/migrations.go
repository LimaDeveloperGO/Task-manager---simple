package database

import (
	"log"
)

func (db *DB) runMigrations() error {
	migrations := []string{
		createTasksTable,
	}

	for i, migration := range migrations {
		if err := db.executeMigration(migration, i+1); err != nil {
			return err
		}
	}

	log.Println("✅ Migrações executadas com sucesso!")
	return nil
}

func (db *DB) executeMigration(query string, version int) error {
	log.Printf("🔄 Executando migração %d...", version)

	_, err := db.Exec(query)
	if err != nil {
		log.Printf("❌ Erro na migração %d: %v", version, err)
		return err
	}

	log.Printf("✅ Migração %d executada com sucesso!", version)
	return nil
}

const createTasksTable = `
CREATE TABLE IF NOT EXISTS tasks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    description TEXT DEFAULT '',
    completed BOOLEAN DEFAULT FALSE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER IF NOT EXISTS update_tasks_updated_at 
    AFTER UPDATE ON tasks
    FOR EACH ROW
BEGIN
    UPDATE tasks SET updated_at = CURRENT_TIMESTAMP WHERE id = NEW.id;
END;
`
