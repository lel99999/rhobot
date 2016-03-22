package config

import (
	"fmt"
	"os"
	"regexp"
)

// Config represents Rhobot's full set of configuration options
type Config struct {
	pgHost     string
	pgPort     string
	PgDatabase string
	pgUser     string
	pgPassword string

	gocdHost     string
	gocdPort     string
	gocdUser     string
	gocdPassword string
}

// NewDefaultConfig creates a new configuration object with default settings
func NewDefaultConfig() *Config {
	return &Config{
		pgHost:       "localhost",
		pgPort:       "5432",
		PgDatabase:   "postgres",
		pgUser:       "postgres",
		pgPassword:   "password",
		gocdHost:     "http://localhost",
		gocdPort:     "8153",
		gocdUser:     "admin",
		gocdPassword: "password",
	}
}

// NewConfig creates a new configuration object from environment variables
func NewConfig() (config *Config) {
	config = NewDefaultConfig()

	if os.Getenv("PGHOST") != "" {
		config.pgHost = os.Getenv("PGHOST")
	}

	if os.Getenv("PGPORT") != "" {
		config.pgPort = os.Getenv("PGPORT")
	}

	if os.Getenv("PGDATABASE") != "" {
		config.PgDatabase = os.Getenv("PGDATABASE")
	}

	if os.Getenv("PGUSER") != "" {
		config.pgUser = os.Getenv("PGUSER")
	}

	if os.Getenv("PGDATABASE") != "" {
		config.pgPassword = os.Getenv("PGPASSWORD")
	}

	if os.Getenv("GOCDHOST") != "" {
		config.gocdHost = os.Getenv("GOCDHOST")
	}

	if os.Getenv("GOCDPORT") != "" {
		config.gocdPort = os.Getenv("GOCDPORT")
	}

	if os.Getenv("GOCDUSER") != "" {
		config.gocdUser = os.Getenv("GOCDUSER")
	}

	if os.Getenv("GOCDPASSWORD") != "" {
		config.gocdPassword = os.Getenv("GOCDPASSWORD")
	}

	return
}

// SetDBURI extracts Postgres connection variables from a DB URI
func (config *Config) SetDBURI(dbURI string) {
	dbURIRegex := regexp.MustCompile(`postgres://(?P<pg_user>\w+):(?P<pg_password>\w+)@(?P<pg_host>\w+):(?P<pg_port>\w+)/(?P<pg_database>\w+).*`)

	match := dbURIRegex.FindStringSubmatch(dbURI)
	if match == nil {
		fmt.Println("Invalid DB URI!")
	} else if len(match) < 4 {
		fmt.Println("Too few regex matches with DB URI!")
	}

	config.pgUser = match[1]
	config.pgPassword = match[2]
	config.pgHost = match[3]
	config.PgDatabase = match[4]
}

// DBURI generates a DB URI from the proper configruation options
func (config *Config) DBURI() (dbURI string) {
	dbURI = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=require",
		config.pgUser,
		config.pgPassword,
		config.pgHost,
		config.pgPort,
		config.PgDatabase)

	return
}

// SetGoCDHost sets the host value of the GoCD server
func (config *Config) SetGoCDHost(host string) {
	config.gocdHost = host
}

// GoCDURL returns the host of the GoCD server
func (config *Config) GoCDURL() string {
	return fmt.Sprintf("%s:%s", config.gocdHost, config.gocdPort)
}
