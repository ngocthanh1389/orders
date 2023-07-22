package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // nolint sql driver name: "postgres"
	"github.com/urfave/cli"
)

const (
	PostgresHostFlag    = "postgres-host"
	DefaultPostgresHost = "127.0.0.1"

	PostgresPortFlag    = "postgres-port"
	DefaultPostgresPort = 5432

	PostgresUserFlag    = "postgres-user"
	DefaultPostgresUser = "postgres"

	PostgresPasswordFlag    = "postgres-password"
	DefaultPostgresPassword = "postgres"

	PostgresDatabaseFlag    = "postgres-database"
	DefaultPostgresDatabase = "postgres"

	PostgresMigrationPath = "migration-path"
	DefaultMigrationPath  = "./internal/migrations"
)

// PostgresSQLFlags creates new cli flags for PostgreSQL client.
func PostgresSQLFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   PostgresHostFlag,
			Usage:  "PostgresSQL host to connect",
			EnvVar: "POSTGRES_HOST",
			Value:  DefaultPostgresHost,
		},
		cli.IntFlag{
			Name:   PostgresPortFlag,
			Usage:  "PostgresSQL port to connect",
			EnvVar: "POSTGRES_PORT",
			Value:  DefaultPostgresPort,
		},
		cli.StringFlag{
			Name:   PostgresUserFlag,
			Usage:  "PostgresSQL user to connect",
			EnvVar: "POSTGRES_USER",
			Value:  DefaultPostgresUser,
		},
		cli.StringFlag{
			Name:   PostgresPasswordFlag,
			Usage:  "PostgresSQL password to connect",
			EnvVar: "POSTGRES_PASSWORD",
			Value:  DefaultPostgresPassword,
		},
		cli.StringFlag{
			Name:   PostgresDatabaseFlag,
			Usage:  "Postgres database to connect",
			EnvVar: "POSTGRES_DATABASE",
			Value:  DefaultPostgresDatabase,
		},
		cli.StringFlag{
			Name:   PostgresMigrationPath,
			Value:  DefaultMigrationPath,
			EnvVar: "MIGRATION_PATH",
		},
	}
}

// NewDBFromContext creates a DB instance from cli flags configuration.
func NewDBFromContext(c *cli.Context) (*sqlx.DB, error) {
	const driverName = "postgres"

	connStr := FormatDSN(map[string]string{
		"host":     c.String(PostgresHostFlag),
		"port":     c.String(PostgresPortFlag),
		"user":     c.String(PostgresUserFlag),
		"password": c.String(PostgresPasswordFlag),
		"dbname":   c.String(PostgresDatabaseFlag),
		"sslmode":  "disable",
	})

	return sqlx.Connect(driverName, connStr)
}
