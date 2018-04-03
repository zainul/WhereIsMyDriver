package databases

import "WhereIsMyDriver/models"

// RunMigration ...
func RunMigration() {
	MigrateDB(models.User{})
	MigrateDB(models.HistoryPosition{})
}
