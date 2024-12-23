package character_request

type CreateCharacter struct {
	Name    string `json:"name" validate:"required"`
	RaceId  string `json:"raceId" validate:"required"`
	ClassId string `json:"classId" validate:"required"`
}
