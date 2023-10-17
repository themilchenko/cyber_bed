package convert

import "github.com/cyber_bed/internal/models"

func InputRecognitionResultsToModels(results models.RecResponse, count int) []models.Plant {
	plants := make([]models.Plant, 0)
	counter := 0

	for _, result := range results.Results {
		if counter+1 == count {
			break
		}

		plants = append(plants, models.Plant{
			CommonName: result.Species.Name,
		})
		counter++
	}

	return plants
}

func InputSearchResultsToModels(results models.SearchResponse, count int) []models.Plant {
	plants := make([]models.Plant, 0)
	counter := 0

	for _, result := range results.Data {
		if counter+1 == count {
			break
		}

		plants = append(plants, models.Plant{
			ExternalID: uint64(result.ID),
			CommonName: result.ScName,
			ImageUrl:   result.ImageURL,
		})
		counter++
	}

	return plants
}
