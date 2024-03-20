package adapter

import (
	"context"
	"encoding/json"
	"giphy/challenge/src/domain/entities"
	"giphy/challenge/src/infrastructure/http/giphy/dto"
	"giphy/challenge/src/infrastructure/http/giphy/dto/mappers"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func (g *gifHTTPAdapter) FindBySearch(ctx context.Context, search string, maxNumber int) ([]*entities.GIF, error) {
	baseURL := os.Getenv("GIPHY_SEARCH_URL")
	apiToken := os.Getenv("GIPHY_API_KEY")

	queryParams := url.Values{}
	queryParams.Set("api_key", apiToken)
	queryParams.Set("q", search)
	queryParams.Set("limit", strconv.Itoa(maxNumber))

	reqURL := baseURL + "?" + queryParams.Encode()

	resp, err := g.httpClient.Get(reqURL)
	if err != nil {
		log.Println("[HTTP > Giphy > FindBySearch] Error finding GIFs")

		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("[HTTP > Giphy > FindBySearch] Unexpected response status code: %d", resp.StatusCode)

		return nil, err
	}

	responseBuffer, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("[HTTP > Giphy > FindBySearch] Error reading response")

		return nil, err
	}

	var response *dto.FindBySearchResponse

	if err = json.Unmarshal(responseBuffer, &response); err != nil {
		log.Println("[HTTP > Giphy > FindBySearch] Error converting response")

		return nil, err
	}

	log.Print(response.Data[0])

	return mappers.GIFDTOsToEntities(response.Data), nil
}
