package database

import (
	"backend/internal/models"
	"backend/pkg"
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
	d.coords()
	d.address()
	d.location()
	d.user1()
	d.tags()
	d.voteComment()
	d.event1()
	d.event2()
	d.opinion()
	d.followers()
	d.user2()
}

// &models.Opinion{},
func (d *dummyServiceImpl) user2() {
	var users []*models.User
	if err := d.db.Find(&users).Error; err != nil {
		log.Println("user2: Error fetching users:", err)
		return
	}
	for _, user := range users {
		var votes []*models.Vote
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
		date := pkg.ParseDate(d.f.Date().AddDate(-100, 0, 0).Format(time.DateOnly))
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
		if err := d.db.Save(&models.Coords{
			Longitude: d.f.Longitude(),
			Latitude:  d.f.Latitude(),
		}).Error; err != nil {
			log.Println(err)
		}
	}
}

func (d *dummyServiceImpl) voteComment() {
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

			vote := &models.Vote{
				State:   models.ParticipationStatuses[d.f.Number(0, 2)],
				EventID: events[i].ID,
				UserID:  users[d.f.Number(0, len(users)-1)].ID,
			}

			if err := d.db.Save(vote).Error; err != nil {
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
