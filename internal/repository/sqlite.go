package repository

import (
	"database/sql"

	"github.com/mayankanup/go-agent-gateway/internal/models"
)

type SQLiteRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(
	db *sql.DB,
) *SQLiteRepository {

	return &SQLiteRepository{
		db: db,
	}
}

func (r *SQLiteRepository) AppendMessage(
	conversationID string,
	message models.Message,
) error {

	query := `
	INSERT INTO messages
	(
		conversation_id,
		role,
		content
	)
	VALUES
	(
		?,
		?,
		?
	)
	`

	_, err :=
		r.db.Exec(
			query,
			conversationID,
			message.Role,
			message.Content,
		)

	return err
}

func (r *SQLiteRepository) GetConversation(
	conversationID string,
) ([]models.Message, error) {

	query := `
	SELECT role, content
	FROM messages
	WHERE conversation_id = ?
	ORDER BY id ASC
	`

	rows, err :=
		r.db.Query(
			query,
			conversationID,
		)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var messages []models.Message

	for rows.Next() {

		var m models.Message

		err :=
			rows.Scan(
				&m.Role,
				&m.Content,
			)

		if err != nil {
			return nil, err
		}

		messages =
			append(messages, m)
	}

	return messages, nil
}
