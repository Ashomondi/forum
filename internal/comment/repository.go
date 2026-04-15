package comment

import "database/sql"

type Repository interface {
	Create(comment *Comment) error
	GetTopLevelByPost(postID, limit, offset int) ([]Comment, error)
	GetRepliesByParentID(parentID int) ([]Comment, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

// TODO: Return custom repository errors as opposed to database errorrs

func (r *repository) Create(comment *Comment) error {
	query := `INSERT INTO comments (user_id, post_id, parent_id, content) VALUES (?, ?, ?, ?)`

	result, err := r.db.Exec(
		query,
		comment.UserID,
		comment.PostID,
		comment.ParentID, // can be nil. nil becomes NULL
		comment.Content,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	comment.ID = int(id)
	return nil
}

// NOTE: Using OFFSET-based pagination doesn't scale well should you have thousands of records.
// Because it reads all records up to your offset, then discards them.
// However, for the scale of this application, that's not likely to be a problem.

func (r *repository) GetTopLevelByPost(postID, limit, offset int) ([]Comment, error) {
	query := `
		SELECT id, user_id, post_id, parent_id, content, created_at
		FROM comments
		WHERE post_id = ?
		AND parent_id IS NULL
		ORDER BY created_at DESC
		LIMIT ? OFFSET ?
	`

	rows, err := r.db.Query(query, postID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []Comment

	for rows.Next() {
		var c Comment
		var parentID sql.NullInt64

		err := rows.Scan(
			&c.ID,
			&c.UserID,
			&c.PostID,
			&parentID,
			&c.Content,
			&c.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		if parentID.Valid {
			id := int(parentID.Int64)
			c.ParentID = &id
		}

		comments = append(comments, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *repository) GetRepliesByParentID(parentID int) ([]Comment, error) {
	// NOTE: The 100 comment reply LIMIT here potentially hides replies (if you have more than 100).
	// But for the scale of this application, that is not a problem likely to occur.

		query := `
		SELECT id, user_id, post_id, parent_id, content, created_at
		FROM comments
		WHERE parent_id = ?
		ORDER BY created_at ASC
		LIMIT 100
	`

	rows, err := r.db.Query(query, parentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []Comment

	for rows.Next() {
		var c Comment
		var parentID sql.NullInt64

		err := rows.Scan(
			&c.ID,
			&c.UserID,
			&c.PostID,
			&parentID,
			&c.Content,
			&c.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		if parentID.Valid {
			id := int(parentID.Int64)
			c.ParentID = &id
		}

		comments = append(comments, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}
