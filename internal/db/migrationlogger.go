package db

import "log"

// MigrationLogger is used to log the activity in the migration process
type MigrationLogger struct {
	verbose bool
}

// Printf function
func (ml *MigrationLogger) Printf(format string, v ...interface{}) {
	log.Printf(format, v)
}

// Verbose will enable verbose logging
func (ml *MigrationLogger) Verbose() bool {
	return ml.verbose
}
