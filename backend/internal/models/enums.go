package models

type EventType string

const (
	Workshop  EventType = "WORKSHOP"
	Cultural  EventType = "CULTURAL"
	Sports    EventType = "SPORTS"
	Social    EventType = "SOCIAL"
	Community EventType = "COMMUNITY"
	Charity   EventType = "CHARITY"
	Party     EventType = "PARTY"
)

type ParticipationStatus string

const (
	Interested      ParticipationStatus = "INTERESTED"
	WillParticipate ParticipationStatus = "WILL_PARTICIPATE"
	NotInterested   ParticipationStatus = "NOT_INTERESTED"
)
