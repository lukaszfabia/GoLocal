package recommendation

import (
	"backend/internal/models"
	"log"
	"math"
	"sort"
)

func newFunction(s *recommendationServiceImpl, allEvents []*models.Event, userPreferences *models.Recommendation, count int) []uint {
	vectors := s.countVectorizer(allEvents, userPreferences)

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

	recommendedEvents := []uint{}
	for _, idx := range sortedIndices {
		if len(recommendedEvents) >= count {
			break
		}

		if idx < len(allEvents) {
			eventId := allEvents[idx].ID
			recommendedEvents = append(recommendedEvents, eventId)
			log.Printf("Added recommended event ID: %d", eventId)
		}
	}
	return recommendedEvents
}

func (s *recommendationServiceImpl) countVectorizer(events []*models.Event, userPreferences *models.Recommendation) [][]int {
	tagSet := make(map[string]struct{})
	for _, event := range events {
		for _, tag := range event.Tags {
			tagSet[tag.Name] = struct{}{}
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
				if t == tag.Name {
					vector[j]++
				}
			}
		}
		vectors[i] = vector
	}

	// Add user preferences as the last vector
	userVector := make([]int, len(tagArray))
	for _, tag := range userPreferences.Tags {
		for j, t := range tagArray {
			if t == tag.Name {
				userVector[j]++
			}
		}
	}
	vectors[len(events)] = userVector

	return vectors
}

func (s *recommendationServiceImpl) cosineSimilarity(vec1, vec2 []int) float64 {
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
