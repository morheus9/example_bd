package main

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	migrate "github.com/kva3umoda/sql-migrate"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/Azaliya1995/music_library"
	"github.com/Azaliya1995/music_library/internal/config"
	"github.com/Azaliya1995/music_library/pkg/log"
)

var migrateCMD = &cobra.Command{
	Use:   "migrate",
	Short: "Run migrations",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Init()
		if err != nil {
			log.Error("Init config failed",
				zap.Error(err),
			)
			return err
		}

		err = log.Init(&cfg.LogConfig)
		if err != nil {
			log.Error("Init logger failed",
				zap.Error(err),
			)

			return errors.Wrap(err, "failed to init logger")
		}

		defer log.Sync()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		go waitSignalExit(cancel)

		migrate.SetLogger(newMigrateLogger())
		migrate.SetCreateTable(true)
		migrate.SetSchema(cfg.DatabaseConfig.Schema)
		migrate.SetTable("migrations")
		migrate.SetIgnoreUnknown(false)

		dialect, err := migrate.GetDialect(migrate.Postgres)
		if err != nil {
			log.Error("get migrate dialect failed",
				zap.Error(err),
			)

			return errors.Wrap(err, "get dialect")
		}

		db, err := connDB(cfg.DatabaseConfig)
		if err != nil {
			log.With(zap.Error(err)).Fatal("failed create database connection")

			return errors.Wrap(err, "failed create database connection")
		}

		n, err := migrate.ExecContext(
			ctx,
			db,
			dialect,
			music_library.MigrationSource,
			migrate.Up,
		)
		if err != nil {
			log.With(zap.Error(err)).Fatal("migrate failed")

			return errors.Wrap(err, "migration failed")
		}
		log.Sugar().Infof("Applied %d migrations!", n)

		return nil
	},
}

func connDB(conf config.DatabaseConfig) (*sql.DB, error) {
	config, err := pgx.ParseConfig(conf.GetDSN())
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse pgx config")
	}

	config.DefaultQueryExecMode = pgx.QueryExecModeExec

	db := stdlib.OpenDB(*config)

	db.SetConnMaxLifetime(conf.ConnMaxLifetime)
	db.SetConnMaxIdleTime(conf.ConnMaxIdleTime)
	db.SetMaxIdleConns(conf.MaxIdleConns)
	db.SetMaxOpenConns(conf.MaxOpenConns)

	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "failed to ping db")
	}

	return db, nil
}

var _ migrate.Logger = (*migrateLogger)(nil)

type migrateLogger struct {
}

func newMigrateLogger() *migrateLogger {
	return &migrateLogger{}
}

func (m *migrateLogger) Tracef(format string, v ...any) {
	log.Debugf(format, v...)
}

func (m *migrateLogger) Infof(format string, v ...any) {
	log.Infof(format, v...)
}

func (m *migrateLogger) Errorf(format string, v ...any) {
	log.Errorf(format, v...)
}
