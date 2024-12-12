package music_library

import (
	"embed"

	migrate "github.com/kva3umoda/sql-migrate"
)

//go:embed migrations/*
var migrationFiles embed.FS

// MigrationSource embedded migration source.
var MigrationSource = migrate.NewEmbedFileSystemMigrationSource(migrationFiles, "migrations")
