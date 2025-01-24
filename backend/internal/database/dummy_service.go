package database

import (
	"backend/internal/models"
	"backend/pkg/parsers"
	"fmt"
	"log"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"gorm.io/gorm"
)

const MAX_CAPACITY = 1000

type DummyService interface {
	Cook()
}

type dummyServiceImpl struct {
	db *gorm.DB
	f  *gofakeit.Faker
}

func NewDummyService(db *gorm.DB) DummyService {
	f := gofakeit.New(uint64(gofakeit.Date().UnixNano()))
	return &dummyServiceImpl{db: db, f: f}
}

func (s *service) DummyService() DummyService {
	return s.dummyService
}

func (d *dummyServiceImpl) Cook() {
	// d.coords()
	// d.address()
	// d.location()
	// d.user1()
	// d.tags()
	// d.comments()
	// d.votes()
	// d.event1()
	// d.event2()
	// d.opinion()
	// d.followers()
	// d.user2()
	// d.generateMockSurvey()
	// d.easyLoginUser()
	// d.generateRecommendations()
}

// &models.Opinion{},
func (d *dummyServiceImpl) user2() {
	var users []*models.User
	if err := d.db.Find(&users).Error; err != nil {
		log.Println("user2: Error fetching users:", err)
		return
	}
	for _, user := range users {
		var votes []*models.VoteAnswer
		var comments []*models.Comment

		if err := d.db.Model(&models.Vote{}).Where("user_id = ?", user.ID).Find(&votes).Error; err != nil {
			log.Println("user2: Error fetching votes for user", user.ID, err)
		} else {
			user.Votes = votes
		}

		if err := d.db.Model(&models.Comment{}).Where("user_id = ?", user.ID).Find(&comments).Error; err != nil {
			log.Println("user2: Error fetching comments for user", user.ID, err)
		} else {
			user.Comments = comments
		}

		if err := d.db.Save(user).Error; err != nil {
			log.Println("user2: Error saving user", user.ID, err)
		}
	}

}

func (d *dummyServiceImpl) user1() {
	// take all locations

	var locations []*models.Location

	if err := d.db.Find(&locations).Error; err != nil {
		log.Println(err)
		return
	}

	for i := 0; i < len(locations); i++ {
		p := d.f.Person()

		password := d.f.Password(true, true, true, false, false, 40)
		date := parsers.ParseDate(d.f.Date().AddDate(-100, 0, 0).Format(time.DateOnly))
		rURL := "https://i.pravatar.cc/300"
		bio := d.f.HipsterSentence(10)

		user := &models.User{
			FirstName:  p.FirstName,
			LastName:   p.LastName,
			Email:      p.Contact.Email,
			Password:   &password,
			Birthday:   &date,
			IsVerified: d.f.Bool(),
			Bio:        &bio,
			AvatarURL:  &rURL,
			Location:   locations[i],
			LocationID: &locations[i].ID,
		}

		if err := d.db.Save(user).Error; err != nil {
			log.Println("error przy user1")
		}
	}
}

func (d *dummyServiceImpl) easyLoginUser() {
	var count int64
	if err := d.db.Model(&models.User{}).Where("email = ?", "a@a.a").Count(&count).Error; err != nil {
		log.Println("Error checking user existence:", err)
		return
	}
	if count > 0 {
		log.Println("Easy login user already exists")
		return
	}

	p := d.f.Person()
	email := "a@a.a"
	password := "Passw0rd!"
	date := parsers.ParseDate(d.f.Date().AddDate(-100, 0, 0).Format(time.DateOnly))
	rURL := "https://i.pravatar.cc/300"
	bio := d.f.HipsterSentence(10)

	var location models.Location
	if err := d.db.Order("RANDOM()").First(&location).Error; err != nil {
		log.Println(err)
		return
	}

	user := &models.User{
		FirstName:  p.FirstName,
		LastName:   p.LastName,
		Email:      email,
		Password:   &password,
		Birthday:   &date,
		IsVerified: d.f.Bool(),
		Bio:        &bio,
		AvatarURL:  &rURL,
		Location:   &location,
		LocationID: &location.ID,
	}

	if err := d.db.Save(user).Error; err != nil {
		log.Println("error at easy login user")
	}
}

func (d *dummyServiceImpl) location() {
	var coords []*models.Coords
	var addresses []*models.Address

	// take all cords and addressess
	if err := d.db.Find(&coords).Error; err != nil {
		log.Println(err)
	}

	if err := d.db.Find(&addresses).Error; err != nil {
		log.Println(err)
	}

	var len = min(len(coords), len(addresses))

	if len == 0 {
		log.Println("Length of min. slice is 0")
		return
	}

	for i := 0; i < len; i++ {
		a := *addresses[i]
		c := *coords[i]

		loc := &models.Location{
			City:      d.f.City(),
			Country:   d.f.Country(),
			Zip:       d.f.Zip(),
			Address:   &a,
			AddressID: a.ID,
			Coords:    &c,
			CoordsID:  c.ID,
		}

		if err := d.db.Save(loc).Error; err != nil {
			log.Println(err)
		}
	}
}

func (d *dummyServiceImpl) address() {
	for i := 0; i < MAX_CAPACITY; i++ {
		address := &models.Address{
			Street:         d.f.Street(),
			StreetNumber:   d.f.StreetNumber(),
			AdditionalInfo: d.f.SentenceSimple(),
		}

		if err := d.db.Save(address).Error; err != nil {
			log.Println(err)
		}
	}
}

func (d *dummyServiceImpl) coords() {
	for i := 0; i < MAX_CAPACITY; i++ {
		longitude := d.f.Longitude()
		latitude := d.f.Latitude()
		point := fmt.Sprintf("SRID=4326;POINT(%f %f)", longitude, latitude)

		if err := d.db.Save(&models.Coords{
			Longitude: longitude,
			Latitude:  latitude,
			Geom:      point,
		}).Error; err != nil {
			log.Println(err)
		}
	}
}

func (d *dummyServiceImpl) comments() {
	// take events and users
	var users []*models.User
	var events []*models.Event

	if err := d.db.Find(&events).Error; err != nil {
		log.Println(err)
		return
	}

	if err := d.db.Find(&users).Error; err != nil {
		log.Println(err)
		return
	}

	for i := 0; i < len(events); i++ {
		// assign about 100 comments per event
		for j := 0; j < d.f.Number(10, 100); j++ {
			comment := &models.Comment{
				Content: d.f.SentenceSimple(),
				EventID: events[i].ID,
				UserID:  users[d.f.Number(0, len(users)-1)].ID,
			}

			if err := d.db.Save(comment).Error; err != nil {
				log.Println(err)
			}
		}
	}
}

func (d *dummyServiceImpl) votes() {
	var users []*models.User
	var events []*models.Event

	if err := d.db.Find(&events).Error; err != nil {
		log.Println(err)
		return
	}

	if err := d.db.Find(&users).Error; err != nil {
		log.Println(err)
		return
	}

	for i := 0; i < len(events); i++ {
		// assign about 2 votes per event
		for j := 0; j < d.f.Number(1, 2); j++ {
			// Create a new vote
			vote := &models.Vote{
				EventID: events[i].ID,
				Text:    d.f.Question(),
				VoteType: func() models.VoteType {
					if d.f.Bool() {
						return models.CanChangeVote
					}
					return models.CannotChangeVote
				}(),
				Options: []models.VoteOption{
					{Text: d.f.Noun()},
					{Text: d.f.Noun()},
				},
			}

			if err := d.db.Create(vote).Error; err != nil {
				log.Println(err)
				continue
			}

			// Assign the vote to a random user
			user := users[d.f.Number(0, len(users)-1)]
			voteAnswer := &models.VoteAnswer{
				VoteID:       vote.ID,
				UserID:       user.ID,
				VoteOptionID: vote.Options[d.f.Number(0, len(vote.Options)-1)].ID,
			}

			if err := d.db.Create(voteAnswer).Error; err != nil {
				log.Println(err)
			}
		}
	}
}

func (d *dummyServiceImpl) followers() {
	var users []*models.User
	if err := d.db.Find(&users).Error; err != nil {
		log.Println("Error fetching users:", err)
		return
	}

	for _, user := range users {
		numFollowing := d.f.Number(1, 10)
		numFollowers := d.f.Number(1, 10)

		user.SkipValidation = true

		for i := 0; i < numFollowers; i++ {
			follower := users[d.f.Number(0, len(users)-1)]

			if follower.ID == user.ID {
				continue
			}

			d.db.Model(&user).Association("Followers").Append(follower)
		}

		for i := 0; i < numFollowing; i++ {
			following := users[d.f.Number(0, len(users)-1)]

			if following.ID == user.ID {
				continue
			}

			d.db.Model(&user).Association("Following").Append(following)
		}
	}
}

func (d *dummyServiceImpl) opinion() {
	// TODO
}

func (d *dummyServiceImpl) tags() {
	// check if there are already 300 tags
	tags := []*models.Tag{}
	if err := d.db.Find(&tags).Error; err != nil {
		log.Println(err)
		return
	}

	if len(tags) >= 300 {
		log.Println("Tags already exist")
		return
	}

	var i = 0
	for i < 300 {
		t := &models.Tag{
			Name: d.f.Hobby(),
		}

		if err := d.db.Save(t).Error; err != nil {
			log.Println(err)
		} else {
			i++
		}
	}
}

func (d *dummyServiceImpl) event1() {
	// take all locations and people

	var locations []*models.Location

	if err := d.db.Find(&locations).Error; err != nil {
		log.Println(err)
		return
	}

	var ppl []*models.User

	if err := d.db.Find(&ppl).Error; err != nil {
		log.Println(err)
		return
	}

	for i := 0; i < len(locations); i++ {
		event := &models.Event{}

		// add max 3
		for i := 0; i < d.f.Number(0, 3); i++ {
			event.EventOrganizers = append(event.EventOrganizers, ppl[d.f.Number(0, len(ppl)-1)])
		}

		rType := d.f.RandomString(models.EventTypes)

		event.Title = generateEventTitle(models.EventType(rType))
		event.EventType = models.EventType(rType)
		event.Description = d.f.SentenceSimple()
		sDate := d.f.FutureDate()
		fDate := sDate.AddDate(0, d.f.Number(1, 12), d.f.Number(1, 25))
		event.StartDate = &sDate
		event.FinishDate = &fDate

		url := "https://picsum.photos/seed/picsum/350/200"
		event.ImageURL = &url

		event.IsAdultOnly = d.f.Bool()

		event.Location = locations[i]
		event.LocationID = locations[i].ID

		// take some rand tags
		var tags []*models.Tag
		l := d.f.Number(1, 4)
		if err := d.db.Limit(l).Order("RANDOM()").Find(&tags).Error; err != nil {
			log.Println(err)
		}

		event.Tags = tags

		if err := d.db.Save(event).Error; err != nil {
			log.Println(err)
		}
	}
}

func (d *dummyServiceImpl) event2() {
	var events []*models.Event

	if err := d.db.Find(&events).Error; err != nil {
		log.Println("Error fetching events:", err)
		return
	}

	for _, event := range events {
		var votes []*models.Vote
		var comments []*models.Comment

		if err := d.db.Model(&models.Vote{}).Where("event_id = ?", event.ID).Find(&votes).Error; err != nil {
			log.Println("Error fetching votes for event", event.ID, err)
		} else {
			event.Votes = votes
		}

		if err := d.db.Model(&models.Comment{}).Where("event_id = ?", event.ID).Find(&comments).Error; err != nil {
			log.Println("Error fetching comments for event", event.ID, err)
		} else {
			event.Comments = comments
		}

		if err := d.db.Save(event).Error; err != nil {
			log.Println("Error saving event", event.ID, err)
		}
	}
}

func generateEventTitle(eventType models.EventType) string {
	gofakeit.Seed(0)

	var title string
	switch eventType {
	case models.Workshop:
		title = fmt.Sprintf("%s %s Workshop", gofakeit.Adjective(), gofakeit.Noun())
	case models.Cultural:
		title = fmt.Sprintf("%s Cultural Event: %s", gofakeit.Adjective(), gofakeit.Noun())
	case models.Sports:
		title = fmt.Sprintf("The Big %s Sports Event", gofakeit.Noun())
	case models.Social:
		title = fmt.Sprintf("Social Gathering: %s", gofakeit.Noun())
	case models.Community:
		title = fmt.Sprintf("%s Community Meetup", gofakeit.Adjective())
	case models.Charity:
		title = fmt.Sprintf("Charity Event: %s", gofakeit.Noun())
	case models.Party:
		title = fmt.Sprintf("%s Party", gofakeit.Adjective())
	default:
		title = "Generic Event"
	}

	return title
}

func (d *dummyServiceImpl) generateMockSurvey() {
	// don't generate if survey already exists

	var count int64
	if err := d.db.Model(&models.PreferenceSurvey{}).Count(&count).Error; err != nil {
		log.Println("Error counting surveys:", err)
		return
	}

	if count > 0 {
		log.Println("Mock survey already exists")
		return
	}

	mockSurvey := models.PreferenceSurvey{
		Title:       "Preference survey",
		Description: "Thanks to this quiz, we will be able to personalize our recommendations of events just for You",
	}

	if err := d.db.Save(&mockSurvey).Error; err != nil {
		log.Println("Error saving mock survey:", err)
		return
	}

	questions := []models.PreferenceSurveyQuestion{
		{
			Text:     "Are you interested in adult-only activities?",
			Type:     models.Toggle,
			SurveyID: mockSurvey.ID,
		},
		{
			Text:     "Do you prefer to relax, or spend time actively?",
			Type:     models.SingleChoice,
			SurveyID: mockSurvey.ID,
		},
		{
			Text:     "What are your age/family constraints for events and activities?",
			Type:     models.SingleChoice,
			SurveyID: mockSurvey.ID,
		},
		{
			Text:     "Do you prefer indoors or outdoors events and activities?",
			Type:     models.SingleChoice,
			SurveyID: mockSurvey.ID,
		},
		{
			Text:     "What more are you interested in?",
			Type:     models.MultipleChoice,
			SurveyID: mockSurvey.ID,
		},
	}

	for _, question := range questions {
		if err := d.db.Save(&question).Error; err != nil {
			log.Println("Error saving question:", err)
			return
		}

		var options []models.PreferenceSurveyOption
		switch question.Text {
		case "Are you interested in adult-only activities?":
			options = []models.PreferenceSurveyOption{
				{Text: "Yes", QuestionID: question.ID},
				{Text: "No", QuestionID: question.ID},
			}
		case "Do you prefer to relax, or spend time actively?":
			options = []models.PreferenceSurveyOption{
				{Text: "High-energy", QuestionID: question.ID},
				{Text: "Relaxation", QuestionID: question.ID},
			}
		case "What are your age/family constraints for events and activities?":
			options = []models.PreferenceSurveyOption{
				{Text: "Family-friendly", QuestionID: question.ID},
				{Text: "Couple-friendly", QuestionID: question.ID},
				{Text: "Adult-only", QuestionID: question.ID},
			}
		case "Do you prefer indoors or outdoors events and activities?":
			options = []models.PreferenceSurveyOption{
				{Text: "Indoors", QuestionID: question.ID},
				{Text: "Outdoors", QuestionID: question.ID},
			}
		case "What more are you interested in?":
			options = []models.PreferenceSurveyOption{
				{Text: "Learning", QuestionID: question.ID},
				{Text: "Music", QuestionID: question.ID},
				{Text: "Sports", QuestionID: question.ID},
			}
		}

		for _, option := range options {
			if err := d.db.Save(&option).Error; err != nil {
				log.Println("Error saving option:", err)
				return
			}
		}
	}

	log.Println("Mock survey saved successfully")
}

func (d *dummyServiceImpl) generateRecommendations() error {
	var users []*models.User
	if err := d.db.Find(&users).Error; err != nil {
		return fmt.Errorf("error fetching users: %w", err)
	}

	var tags []*models.Tag
	if err := d.db.Find(&tags).Error; err != nil {
		return fmt.Errorf("error fetching tags: %w", err)
	}

	for _, user := range users {
		recommendation := &models.Recommendation{
			UserID: user.ID,
			Tags:   []models.Tag{},
		}

		numTags := d.f.Number(3, 8)
		for i := 0; i < numTags; i++ {
			tag := tags[d.f.Number(0, len(tags)-1)]
			recommendation.Tags = append(recommendation.Tags, *tag)
		}

		if err := d.db.Save(recommendation).Error; err != nil {
			log.Println("error saving recommendation for user", user.ID, err)
		}
	}

	return nil
}
