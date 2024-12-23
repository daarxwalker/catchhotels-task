package dice_service

import (
	"math/rand/v2"
	"time"

	"github.com/gofiber/fiber/v2"

	"catchhotels/entity/dice_roll_entity"
	"catchhotels/service/dataverse_service"
	"catchhotels/service/session_service"
)

type DiceService struct {
	dataverseService *dataverse_service.DataverseService
	sessionService   *session_service.SessionService
}

const (
	Token = "dice_service"
)

func New(
	dataverseService *dataverse_service.DataverseService, sessionService *session_service.SessionService,
) *DiceService {
	return &DiceService{
		dataverseService: dataverseService,
		sessionService:   sessionService,
	}
}

func (s *DiceService) Roll(c *fiber.Ctx, name string) (uint8, error) {
	roll := uint8(rand.Uint32N(20))
	session, getSessionErr := s.sessionService.Get(c)
	if getSessionErr != nil {
		return 0, getSessionErr
	}
	if createDiceRollErr := s.dataverseService.Create(
		dice_roll_entity.Table,
		dice_roll_entity.DiceRoll{
			RollAt:      time.Now(),
			CharacterId: session.CharacterId,
			Name:        name,
			Result:      roll,
		},
		nil,
	); createDiceRollErr != nil {
		return 0, createDiceRollErr
	}
	return roll, nil
}

func (s *DiceService) MustRoll(c *fiber.Ctx, name string) uint8 {
	roll, err := s.Roll(c, name)
	if err != nil {
		panic(err)
	}
	return roll
}
