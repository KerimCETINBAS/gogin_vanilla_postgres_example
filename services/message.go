package services

import (
	"github.com/kerimcetinbas/goginpostgrestut/repositories"
	"github.com/kerimcetinbas/goginpostgrestut/types"
)

type messageService struct {
	messageRepository repositories.IMessageRepository
}
type IMessageService interface {
	GetMessages() (*[]types.Message, error)
	CreateMessage(data *types.MessageCreateDto) error
}

func MessageService() repositories.IMessageRepository {
	return &messageService{
		messageRepository: repositories.MessageRepository(),
	}
}

func (s *messageService) GetMessages() (*[]types.Message, error) {
	return s.messageRepository.GetMessages()
}

func (s *messageService) CreateMessage(data *types.MessageCreateDto) error {

	return s.messageRepository.CreateMessage(data)
}
