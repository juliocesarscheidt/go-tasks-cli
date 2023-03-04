package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
)

type Task struct {
	Name      string         `json:"name"`
	Done      bool           `json:"done"`
	CreatedAt sql.NullString `json:"created_at"`
	UpdatedAt sql.NullString `json:"updated_at"`
	DeletedAt sql.NullString `json:"deleted_at"`
}

func listTasks(db *sql.DB) ([]*Task, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	query := "SELECT name, done, created_at, updated_at, deleted_at FROM tasks order by created_at DESC"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.Name, &task.Done, &task.CreatedAt, &task.UpdatedAt, &task.DeletedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func getTask(db *sql.DB, name string) (*Task, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	var task Task
	query := "SELECT name, done, created_at, updated_at, deleted_at FROM tasks WHERE name = ? LIMIT 1"
	row := db.QueryRowContext(ctx, query, name)
	if row == nil {
		return nil, nil
	}
	err := row.Scan(&task.Name, &task.Done, &task.CreatedAt, &task.UpdatedAt, &task.DeletedAt)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}
	return &task, nil
}

func createTask(db *sql.DB, name string, done bool) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	now := time.Now()
	query := "INSERT INTO tasks (name, done, created_at, updated_at, deleted_at) VALUES (?, ?, ?, ?, ?)"
	_, err := db.ExecContext(ctx, query, name, done, now, now, nil)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	db, err := sql.Open("sqlite3", "file:/tmp/go_tasks.db")
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(1)
	db.Exec(`CREATE TABLE IF NOT EXISTS tasks (name VARCHAR(255), done TINYINT default 0, created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL, updated_at DATETIME DEFAULT NULL, deleted_at DATETIME DEFAULT NULL)`)

	rootCmd := &cobra.Command{
		Use:   "go-tasks",
		Short: "Running go tasks",
		Example: `
Displays help menu for go-tasks CLI.`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	rootCmd.Version = "v1.0.0"
	rootCmd.SetVersionTemplate("go-tasks version: {{.Version}}\n")

	var taskname string
	var taskdone bool

	taskCmd := &cobra.Command{
		Use:   "tasks",
		Short: "Commands for tasks.",
		Long:  "Commands for tasks.        Tasks are chores.",
		Args:  cobra.MinimumNArgs(1),
		Example: `
  Interact with tasks
  $ go-tasks tasks list
  $ go-tasks tasks get
  $ go-tasks tasks create`,
	}

	// sub commands
	taskLsCmd := &cobra.Command{
		Use:   "list",
		Short: "ls.",
		Long:  "list.        List tasks.",
		Example: `
  List tasks.
  $ go-tasks tasks list`,
		RunE: func(cmd *cobra.Command, args []string) error {
			tasks, err := listTasks(db)
			if err != nil {
				return err
			}
			if len(tasks) <= 0 {
				return nil
			}

			tbl := table.NewWriter()
			tbl.AppendHeader(table.Row{"Name", "Done", "Created At", "Updated At", "Deleted At"})
			for _, task := range tasks {
				tbl.AppendRow([]interface{}{
					task.Name, task.Done, task.CreatedAt.String, task.UpdatedAt.String, task.DeletedAt.String,
				})
			}
			tbl.SetIndexColumn(1)
			fmt.Println(tbl.Render())

			return nil
		},
	}

	taskGetCmd := &cobra.Command{
		Use:   "get",
		Short: "g.",
		Long:  "get.        Get task by name.",
		Example: `
  Get task by name.
  $ go-tasks tasks get --name [taskname]`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(taskname) <= 0 {
				fmt.Println("Missing Task Name")
				os.Exit(1)
			}

			task, err := getTask(db, taskname)
			if err != nil {
				return err
			}
			if task == nil {
				return nil
			}

			tbl := table.NewWriter()
			tbl.AppendHeader(table.Row{"Name", "Done", "Created At", "Updated At", "Deleted At"})
			tbl.AppendRow([]interface{}{
				task.Name, task.Done, task.CreatedAt.String, task.UpdatedAt.String, task.DeletedAt.String,
			})
			tbl.SetIndexColumn(1)
			fmt.Println(tbl.Render())

			return nil
		},
	}
	taskGetCmd.Flags().StringVarP(&taskname, "name", "n", "", "Task Name")
	taskGetCmd.MarkFlagRequired("name")

	taskCreateCmd := &cobra.Command{
		Use:   "create",
		Short: "c.",
		Long:  "create.        Create a task.",
		Example: `
  Create a task.
  $ go-tasks tasks create --name [taskname] --done [true|false]`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(taskname) <= 0 {
				fmt.Println("Missing Task Name")
				os.Exit(1)
			}
			return createTask(db, taskname, taskdone)
		},
	}
	taskCreateCmd.Flags().StringVarP(&taskname, "name", "n", "", "Task Name")
	taskCreateCmd.MarkFlagRequired("name")
	taskCreateCmd.Flags().BoolVarP(&taskdone, "done", "d", false, "Task Done")
	// taskCreateCmd.MarkFlagRequired("done")

	taskCmd.AddCommand(taskLsCmd)
	taskCmd.AddCommand(taskGetCmd)
	taskCmd.AddCommand(taskCreateCmd)

	rootCmd.AddCommand(taskCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error %v", err)
		os.Exit(1)
	}
}
