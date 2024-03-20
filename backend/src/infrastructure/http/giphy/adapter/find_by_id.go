package adapter

import (
	"encoding/json"
	"giphy/challenge/src/domain/entities"
	"giphy/challenge/src/infrastructure/http/giphy/dto"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func (g *gifHTTPAdapter) findByID(gif *entities.GIF) (*entities.GIF, error) {
	baseURL := os.Getenv("GIPHY_FIND_URL")
	apiToken := os.Getenv("GIPHY_API_KEY")

	queryParams := url.Values{}
	queryParams.Set("api_key", apiToken)

	reqURL := baseURL + "/" + gif.ProviderID + "?" + queryParams.Encode()

	resp, err := g.httpClient.Get(reqURL)
	if err != nil {
		log.Println("[HTTP > Giphy > FindByID] Error finding GIF")

		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("[HTTP > Giphy > FindByID] Unexpected response status code: %d", resp.StatusCode)

		return nil, err
	}

	responseBuffer, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("[HTTP > Giphy > FindByID] Error reading response")

		return nil, err
	}

	var response *dto.FindByIDResponse

	if err = json.Unmarshal(responseBuffer, &response); err != nil {
		log.Println("[HTTP > Giphy > FindByID] Error converting response")

		return nil, err
	}

	gif.URL = response.Data.Images.Original.URL

	return gif, nil
}
