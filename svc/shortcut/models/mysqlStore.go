package models

import (
	"database/sql"
	"log"
	"shortlink/api"

	// UUID Generator
	"github.com/rs/xid"

	// MySQL Driver
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// MySQLStore implements the ShortcutStore interface
type MySQLStore struct {
	db *sql.DB
}

// NewMySQLStore creates an instance with a given database connection
func NewMySQLStore(db *sql.DB) *MySQLStore {
	return &MySQLStore{db}
}

func (store *MySQLStore) incVisits(urlToken string) {
	updq := "UPDATE shortcuts SET visits = visits + 1 WHERE token=?"
	res, err := store.db.Exec(updq, urlToken)
	if err != nil {
		log.Printf("Could not increment vists for shortcut: %s: %v", urlToken, err)
		return
	}
	aff, err := res.RowsAffected()
	if err != nil {
		log.Printf("Could not get rows affected visits increment: %s: %v", urlToken, err)
		return
	}
	if aff != 1 {
		log.Printf("Tried to increment vists for non-existant shortcut: %s", urlToken)
		return
	}
}

// Create inserts a new Shortcut into the database and returns the generated url_token
func (store *MySQLStore) Create(targetURL string) (*api.ShortcutDetail, error) {
	insq := "INSERT INTO shortcuts (token, target_url, created_at, visits) VALUES (?, ?, NOW(), 0)"
	token := generateToken()
	insert, err := store.db.Exec(insq, token, targetURL)
	if err != nil {
		log.Printf("Encountered error on insert: %v", err)
		return nil, err
	}

	newID, err := insert.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID: %v", err)
		return nil, err
	}
	selq := "SELECT token, target_url, created_at, visits FROM shortcuts WHERE id=?"
	sel := store.db.QueryRow(selq, newID)
	detail := &api.ShortcutDetail{}
	err = sel.Scan(&detail.UrlToken, &detail.TargetUrl, &detail.CreatedAt, &detail.Visits)
	if err != nil {
		log.Printf("Error reading new shortcut: %v", err)
		return nil, err
	}
	return detail, nil
}

// Delete removes a shortcut from the database
func (store *MySQLStore) Delete(urlToken string) error {
	delq := "DELETE FROM shortcuts WHERE token = ?"
	res, err := store.db.Exec(delq, urlToken)
	if err != nil {
		log.Printf("Encountered error on delete: %v", err)
		return status.Error(codes.Internal, "Internal Error")
	}
	count, err := res.RowsAffected()
	if err != nil {
		log.Printf("Encountered error determining rows affected: %v", err)
		return status.Error(codes.Internal, "Internal Error")
	}
	if count != 1 {
		return status.Error(codes.NotFound, "Shortcut Does Not Exist")
	}
	return nil
}

// Get returns the target_url of a shortcut from the database given a urlToken
// Also increments vists in a go routine
func (store *MySQLStore) Get(urlToken string) (string, error) {
	getq := "SELECT target_url FROM shortcuts WHERE token=?"
	var result string
	err := store.db.QueryRow(getq, urlToken).Scan(&result)
	if err != nil {
		log.Printf("Error getting result: %v", err)
		return "ERROR", status.Error(codes.NotFound, "Shortcut Does Not Exist")
	}
	defer store.incVisits(urlToken)
	return result, nil
}

func generateToken() string {
	guid := xid.New()
	return guid.String()
}
