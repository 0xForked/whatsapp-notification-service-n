package console

import (
	"fmt"
	"github.com/aasumitro/gowa/configs"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	"strconv"

	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

var migrateCommand = &cobra.Command{
	Use:  "migrate",
	Long: `migrate cmd is used for database migration: migrate < up | down >`,
}

var migrateUpCommand = &cobra.Command{
	Use:  "up",
	Long: `Command to upgrade database migration`,
	Run: func(cmd *cobra.Command, args []string) {
		migration, err := initGoMigrate()
		if err != nil {
			fmt.Printf("migrate down error: %v \n", err)
			return
		}

		if err := migration.Up(); err != nil {
			fmt.Printf("migrate up error: %v \n", err)
			return
		}

		fmt.Println("Migrate up done with success")
	},
}

var migrateDownCommand = &cobra.Command{
	Use:  "down",
	Long: `Command to downgrade database`,
	Run: func(cmd *cobra.Command, args []string) {
		migration, err := initGoMigrate()
		if err != nil {
			fmt.Printf("migrate down error: %v \n", err)
			return
		}

		if err := migration.Down(); err != nil {
			fmt.Printf("migrate down error: %v \n", err)
			return
		}

		fmt.Println("Migrate down done with success")
	},
}

var migrateVersionCommand = &cobra.Command{
	Use:  "version",
	Long: `Command to see database migration version`,
	Run: func(cmd *cobra.Command, args []string) {
		migration, err := initGoMigrate()
		if err != nil {
			fmt.Printf("migrate down error: %v \n", err)
			return
		}

		version, dirty, err := migration.Version()
		if err != nil {
			fmt.Printf("migrate up error: %v \n", err)
			return
		}

		fmt.Printf("Database Version %d is dirty: %s",
			version, strconv.FormatBool(dirty))
	},
}

func initGoMigrate() (instance *migrate.Migrate, err error) {
	fileSource, err := (&file.File{}).Open("file://db/migrations")
	if err != nil {
		return nil, err
	}

	driver, err := sqlite3.WithInstance(
		configs.DbPool, &sqlite3.Config{})
	if err != nil {
		return nil, err
	}

	instance, err = migrate.NewWithInstance(
		"file", fileSource, "pokewar", driver)
	if err != nil {
		return nil, err
	}

	return instance, nil
}

func init() {
	Commands.AddCommand(migrateCommand)
	migrateCommand.AddCommand(migrateUpCommand)
	migrateCommand.AddCommand(migrateDownCommand)
	migrateCommand.AddCommand(migrateVersionCommand)
}
