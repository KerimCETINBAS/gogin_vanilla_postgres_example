package repositories

import (
	"database/sql"
	"fmt"

	"github.com/kerimcetinbas/goginpostgrestut/database"
	"github.com/kerimcetinbas/goginpostgrestut/types"
	"golang.org/x/exp/slices"
)

type messageRepository struct {
	db *sql.DB
}
type IMessageRepository interface {
	GetMessages() (*[]types.Message, error)
	CreateMessage(*types.MessageCreateDto) error
}

func MessageRepository() IMessageRepository {
	return &messageRepository{
		db: database.DB,
	}
}

func (r *messageRepository) CreateMessage(data *types.MessageCreateDto) error {

	_, err := r.db.Exec(`
		INSERT INTO messages 
		(message, "user_id") 
		VALUES($1, $2)
	`, data.Message, data.UserId)

	if err != nil {
		return err
	}

	return nil
}
func (r *messageRepository) GetMessages() (*[]types.Message, error) {
	messages := make([]types.Message, 0)
	var (
		rows *sql.Rows
		err  error
	)
	rows, err = r.db.Query(`
	SELECT m.id AS id, m.message AS message, u.name AS user
	FROM messages m
	INNER JOIN users u on u.id = m.user_id`)

	if err != nil {
		return &messages, err
	}

	for rows.Next() {
		message := types.Message{}

		rows.Scan(&message.ID, &message.Message, &message.User)
		messages = append(messages, message)
	}

	slices.Reverse(messages)
	fmt.Println(messages)
	return &messages, nil
}
