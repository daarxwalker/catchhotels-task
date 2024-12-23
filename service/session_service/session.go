package session_service

type Session struct {
	Id          string `json:"id"`
	CharacterId string `json:"characterId"`
	Email       string `json:"email"`
	IP          string `json:"ip"`
	UserAgent   string `json:"userAgent"`
}

func (s Session) Exists() bool {
	return len(s.Id) > 0
}

func (s Session) CompareIP(ip string) bool {
	return s.IP == ip
}

func (s Session) CompareUserAgent(userAgent string) bool {
	return s.UserAgent == userAgent
}
