package recommendation

import (
	"math"
	"sort"
)

// Event represents an event with tags
type Event struct {
	ID   string
	Tags []string
}

// UserPreferredTags represents the user's preferred tags
type UserPreferredTags struct {
	Tags []string
}

// RecommendationService provides event recommendations based on user preferred tags
type RecommendationService struct {
	AllEvents       []Event
	UserPreferences UserPreferredTags
}

// NewRecommendationService creates a new RecommendationService
func NewRecommendationService(allEvents []Event, userPreferences UserPreferredTags) *RecommendationService {
	return &RecommendationService{
		AllEvents:       allEvents,
		UserPreferences: userPreferences,
	}
}

// Predict recommends events based on user preferred tags
func (s *RecommendationService) Predict() []string {
	userTags := make(map[string]struct{})
	for _, tag := range s.UserPreferences.Tags {
		userTags[tag] = struct{}{}
	}

	allTags := []string{}
	for _, event := range s.AllEvents {
		allTags = append(allTags, event.Tags...)
	}
	for tag := range userTags {
		allTags = append(allTags, tag)
	}

	vectors := s.countVectorizer(s.AllEvents)

	userVector := vectors[len(vectors)-1]
	eventVectors := vectors[:len(vectors)-1]

	cosineSim := make([]float64, len(eventVectors))
	for i, vec := range eventVectors {
		cosineSim[i] = s.cosineSimilarity(userVector, vec)
	}

	sortedIndices := make([]int, len(cosineSim))
	for i := range sortedIndices {
		sortedIndices[i] = i
	}
	sort.Slice(sortedIndices, func(i, j int) bool {
		return cosineSim[sortedIndices[i]] > cosineSim[sortedIndices[j]]
	})

	recommendedEvents := []string{}
	for _, idx := range sortedIndices {
		if len(recommendedEvents) >= 3 {
			break
		}

		if idx < len(s.AllEvents) {
			eventId := s.AllEvents[idx].ID
			recommendedEvents = append(recommendedEvents, eventId)
		}
	}

	return recommendedEvents
}

func (s *RecommendationService) countVectorizer(events []Event) [][]int {
	tagSet := make(map[string]struct{})
	for _, event := range events {
		for _, tag := range event.Tags {
			tagSet[tag] = struct{}{}
		}
	}

	tagArray := make([]string, 0, len(tagSet))
	for tag := range tagSet {
		tagArray = append(tagArray, tag)
	}

	vectors := make([][]int, len(events)+1)
	for i, event := range events {
		vector := make([]int, len(tagArray))
		for _, tag := range event.Tags {
			for j, t := range tagArray {
				if t == tag {
					vector[j]++
				}
			}
		}
		vectors[i] = vector
	}

	// Add user preferences as the last vector
	userVector := make([]int, len(tagArray))
	for _, tag := range s.UserPreferences.Tags {
		for j, t := range tagArray {
			if t == tag {
				userVector[j]++
			}
		}
	}
	vectors[len(events)] = userVector

	return vectors
}

func (s *RecommendationService) cosineSimilarity(vec1, vec2 []int) float64 {
	dotProduct := 0
	for i := range vec1 {
		dotProduct += vec1[i] * vec2[i]
	}

	magnitude1 := 0.0
	for _, v := range vec1 {
		magnitude1 += float64(v * v)
	}
	magnitude1 = math.Sqrt(magnitude1)

	magnitude2 := 0.0
	for _, v := range vec2 {
		magnitude2 += float64(v * v)
	}
	magnitude2 = math.Sqrt(magnitude2)

	return float64(dotProduct) / (magnitude1 * magnitude2)
}
