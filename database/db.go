package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type Server struct {
	ID         int
	Name       string
	IP         string
	Username   string
	Password   string
	SSHKeyPath string
}

func InitDB(dbPath string) error {
	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS servers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		ip TEXT NOT NULL UNIQUE,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		ssh_key_path TEXT NOT NULL
	)
`)
	if err != nil {
		return err
	}

	return nil
}

func OpenDB(dbPath string) error {
	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	return nil
}

func AddServer(name, ip, username, password, sshKeyPath string) error {
	fmt.Printf("Adding server: Name=%s, IP=%s, Username=%s, SSHKeyPath=%s\n", name, ip, username, sshKeyPath)

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM servers WHERE ip = ?", ip).Scan(&count)
	if err != nil {
		fmt.Printf("Error checking for existing IP: %v\n", err)
		return err
	}

	if count > 0 {
		return fmt.Errorf("a server with IP %s already exists", ip)
	}

	_, err = db.Exec("INSERT INTO servers (name, ip, username, password, ssh_key_path) VALUES (?, ?, ?, ?, ?)",
		name, ip, username, password, sshKeyPath)
	if err != nil {
		fmt.Printf("Error adding server: %v\n", err)
		return err
	}

	return nil
}

func ListServers() ([]Server, error) {
	rows, err := db.Query("SELECT id, name, ip, username, password, ssh_key_path FROM servers")
	if err != nil {
		fmt.Printf("Error querying servers: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var servers []Server
	for rows.Next() {
		var s Server
		err := rows.Scan(&s.ID, &s.Name, &s.IP, &s.Username, &s.Password, &s.SSHKeyPath)
		if err != nil {
			fmt.Printf("Error scanning server row: %v\n", err)
			return nil, err
		}
		servers = append(servers, s)
	}
	fmt.Printf("Found %d servers in the database\n", len(servers))
	return servers, nil
}

func DeleteServer(id int) error {
	_, err := db.Exec("DELETE FROM servers WHERE id = ?", id)
	if err != nil {
		fmt.Printf("Error deleting server: %v\n", err)
		return err
	}
	return nil
}
