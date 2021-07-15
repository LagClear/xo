package sqlserver

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// AuthPermission represents a row from 'django.auth_permission'.
type AuthPermission struct {
	ID            int    `json:"id"`              // id
	Name          string `json:"name"`            // name
	ContentTypeID int    `json:"content_type_id"` // content_type_id
	Codename      string `json:"codename"`        // codename
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the AuthPermission exists in the database.
func (ap *AuthPermission) Exists() bool {
	return ap._exists
}

// Deleted returns true when the AuthPermission has been marked for deletion from
// the database.
func (ap *AuthPermission) Deleted() bool {
	return ap._deleted
}

// Insert inserts the AuthPermission to the database.
func (ap *AuthPermission) Insert(ctx context.Context, db DB) error {
	switch {
	case ap._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case ap._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO django.auth_permission (` +
		`name, content_type_id, codename` +
		`) VALUES (` +
		`@p1, @p2, @p3` +
		`); SELECT ID = CONVERT(BIGINT, SCOPE_IDENTITY())`
	// run
	logf(sqlstr, ap.Name, ap.ContentTypeID, ap.Codename)
	rows, err := db.QueryContext(ctx, sqlstr, ap.Name, ap.ContentTypeID, ap.Codename)
	if err != nil {
		return logerror(err)
	}
	defer rows.Close()
	// retrieve id
	var id int64
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return logerror(err)
		}
	}
	if err := rows.Err(); err != nil {
		return logerror(err)
	} // set primary key
	ap.ID = int(id)
	// set exists
	ap._exists = true
	return nil
}

// Update updates a AuthPermission in the database.
func (ap *AuthPermission) Update(ctx context.Context, db DB) error {
	switch {
	case !ap._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case ap._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE django.auth_permission SET ` +
		`name = @p1, content_type_id = @p2, codename = @p3 ` +
		`WHERE id = @p4`
	// run
	logf(sqlstr, ap.Name, ap.ContentTypeID, ap.Codename, ap.ID)
	if _, err := db.ExecContext(ctx, sqlstr, ap.Name, ap.ContentTypeID, ap.Codename, ap.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the AuthPermission to the database.
func (ap *AuthPermission) Save(ctx context.Context, db DB) error {
	if ap.Exists() {
		return ap.Update(ctx, db)
	}
	return ap.Insert(ctx, db)
}

// Upsert performs an upsert for AuthPermission.
func (ap *AuthPermission) Upsert(ctx context.Context, db DB) error {
	switch {
	case ap._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `MERGE django.auth_permission AS t ` +
		`USING (` +
		`SELECT @p1 id, @p2 name, @p3 content_type_id, @p4 codename ` +
		`) AS s ` +
		`ON s.id = t.id ` +
		`WHEN MATCHED THEN ` +
		`UPDATE SET ` +
		`t.name = s.name, t.content_type_id = s.content_type_id, t.codename = s.codename ` +
		`WHEN NOT MATCHED THEN ` +
		`INSERT (` +
		`name, content_type_id, codename` +
		`) VALUES (` +
		`s.name, s.content_type_id, s.codename` +
		`);`
	// run
	logf(sqlstr, ap.ID, ap.Name, ap.ContentTypeID, ap.Codename)
	if _, err := db.ExecContext(ctx, sqlstr, ap.ID, ap.Name, ap.ContentTypeID, ap.Codename); err != nil {
		return err
	}
	// set exists
	ap._exists = true
	return nil
}

// Delete deletes the AuthPermission from the database.
func (ap *AuthPermission) Delete(ctx context.Context, db DB) error {
	switch {
	case !ap._exists: // doesn't exist
		return nil
	case ap._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM django.auth_permission ` +
		`WHERE id = @p1`
	// run
	logf(sqlstr, ap.ID)
	if _, err := db.ExecContext(ctx, sqlstr, ap.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	ap._deleted = true
	return nil
}

// AuthPermissionByID retrieves a row from 'django.auth_permission' as a AuthPermission.
//
// Generated from index 'PK__auth_per__3213E83FA54D36B9'.
func AuthPermissionByID(ctx context.Context, db DB, id int) (*AuthPermission, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name, content_type_id, codename ` +
		`FROM django.auth_permission ` +
		`WHERE id = @p1`
	// run
	logf(sqlstr, id)
	ap := AuthPermission{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&ap.ID, &ap.Name, &ap.ContentTypeID, &ap.Codename); err != nil {
		return nil, logerror(err)
	}
	return &ap, nil
}

// AuthPermissionByContentTypeID retrieves a row from 'django.auth_permission' as a AuthPermission.
//
// Generated from index 'auth_permission_content_type_id_2f476e4b'.
func AuthPermissionByContentTypeID(ctx context.Context, db DB, contentTypeID int) ([]*AuthPermission, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name, content_type_id, codename ` +
		`FROM django.auth_permission ` +
		`WHERE content_type_id = @p1`
	// run
	logf(sqlstr, contentTypeID)
	rows, err := db.QueryContext(ctx, sqlstr, contentTypeID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*AuthPermission
	for rows.Next() {
		ap := AuthPermission{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&ap.ID, &ap.Name, &ap.ContentTypeID, &ap.Codename); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &ap)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// AuthPermissionByContentTypeIDCodename retrieves a row from 'django.auth_permission' as a AuthPermission.
//
// Generated from index 'auth_permission_content_type_id_codename_01ab375a_uniq'.
func AuthPermissionByContentTypeIDCodename(ctx context.Context, db DB, contentTypeID int, codename string) (*AuthPermission, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name, content_type_id, codename ` +
		`FROM django.auth_permission ` +
		`WHERE content_type_id = @p1 AND codename = @p2`
	// run
	logf(sqlstr, contentTypeID, codename)
	ap := AuthPermission{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, contentTypeID, codename).Scan(&ap.ID, &ap.Name, &ap.ContentTypeID, &ap.Codename); err != nil {
		return nil, logerror(err)
	}
	return &ap, nil
}

// DjangoContentType returns the DjangoContentType associated with the AuthPermission's (ContentTypeID).
//
// Generated from foreign key 'auth_permission_content_type_id_2f476e4b_fk_django_content_type_id'.
func (ap *AuthPermission) DjangoContentType(ctx context.Context, db DB) (*DjangoContentType, error) {
	return DjangoContentTypeByID(ctx, db, ap.ContentTypeID)
}