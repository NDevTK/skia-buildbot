// Package sqlsubscriptionstore implements subscription.Store using an SQL database.

package sqlfavoritestore

import (
	"context"
	"time"

	"go.skia.org/infra/go/skerr"
	"go.skia.org/infra/go/sql/pool"
	"go.skia.org/infra/perf/go/favorites"
)

// statement is an SQL statement identifier.
type statement int

const (
	// The identifiers for all the SQL statements used.
	getFavorite statement = iota
	insertFavorite
	updateFavorite
	deleteFavorite
	listFavorites
)

// statements holds all the raw SQL statemens.
var statements = map[statement]string{
	getFavorite: `
		SELECT
			*
		FROM
			Favorites
		WHERE
			id=$1
	`,
	insertFavorite: `
		INSERT INTO
			Favorites (user_id, name, url, description, last_modified)
		VALUES
			($1, $2, $3, $4, $5)
	`,

	updateFavorite: `
		UPDATE
			Favorites
		SET
			(name, url, description, last_modified) = ($1, $2, $3, $4)
		WHERE
			id=$5
	`,
	deleteFavorite: `
		DELETE
		FROM
			Favorites
		WHERE
			id=$1
	`,
	listFavorites: `
		SELECT
			id,
			name,
			url,
			description
		FROM
			Favorites
		WHERE
			user_id=$1
	`,
}

// FavoriteStore implements the favorite.Store interface using an SQL
// database.
type FavoriteStore struct {
	db pool.Pool
}

// New returns a new *FavoriteStore.
func New(db pool.Pool) *FavoriteStore {
	return &FavoriteStore{
		db: db,
	}
}

// Get implements the favorites.Store interface.
func (s *FavoriteStore) Get(ctx context.Context, id int64) (*favorites.Favorite, error) {
	fav := &favorites.Favorite{}
	if err := s.db.QueryRow(ctx, statements[getFavorite], id).Scan(
		&fav.ID,
		&fav.UserId,
		&fav.Name,
		&fav.Url,
		&fav.Description,
		&fav.LastModified,
	); err != nil {
		return nil, skerr.Wrapf(err, "Failed to load favorite.")
	}
	return fav, nil
}

// Create implements the favorites.Store interface.
func (s *FavoriteStore) Create(ctx context.Context, req *favorites.SaveRequest) error {
	now := time.Now().Unix()
	if _, err := s.db.Exec(ctx, statements[insertFavorite], req.UserId, req.Name, req.Url, req.Description, now); err != nil {
		return skerr.Wrapf(err, "Failed to insert favorite")
	}
	return nil
}

// Create implements the favorites.Store interface.
func (s *FavoriteStore) Update(ctx context.Context, req *favorites.SaveRequest, id int64) error {
	now := time.Now().Unix()
	if _, err := s.db.Exec(ctx, statements[updateFavorite], req.Name, req.Url, req.Description, now, id); err != nil {
		return skerr.Wrapf(err, "Failed to update favorite with id=%d", id)
	}
	return nil
}

// Delete implements the favorites.Store interface.
func (s *FavoriteStore) Delete(ctx context.Context, id int64) error {
	if _, err := s.db.Exec(ctx, statements[deleteFavorite], id); err != nil {
		return skerr.Wrapf(err, "Failed to delete favorite with id=%d", id)
	}
	return nil
}

// List implements the favorites.Store interface.
func (s *FavoriteStore) List(ctx context.Context, userId string) ([]*favorites.Favorite, error) {
	rows, err := s.db.Query(ctx, statements[listFavorites], userId)
	if err != nil {
		return nil, err
	}

	ret := []*favorites.Favorite{}
	for rows.Next() {
		f := &favorites.Favorite{}
		if err := rows.Scan(&f.ID, &f.Name, &f.Url, &f.Description); err != nil {
			return nil, err
		}

		ret = append(ret, f)
	}
	return ret, nil
}
