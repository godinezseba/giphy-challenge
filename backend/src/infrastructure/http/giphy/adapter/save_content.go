package adapter

import (
	"bytes"
	"context"
	"encoding/json"
	"giphy/challenge/src/domain/entities"
	"giphy/challenge/src/infrastructure/http/giphy/dto"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func (g *gifHTTPAdapter) SaveContent(ctx context.Context, gif *entities.GIF) (*entities.GIF, error) {
	baseURL := os.Getenv("GIPHY_UPLOAD_URL")
	apiToken := os.Getenv("GIPHY_API_KEY")

	request := dto.UploadRequest{
		ApiKey: apiToken,
		File:   gif.Content,
		Tags:   strings.Join(gif.Tags, ","),
	}

	requestBuffer, err := json.Marshal(request)
	if err != nil {
		log.Println("[HTTP > Giphy > SaveContent] Error creating request")

		return nil, err
	}

	resp, err := g.httpClient.Post(baseURL, "application/json", bytes.NewBuffer(requestBuffer))
	if err != nil {
		log.Println("[HTTP > Giphy > SaveContent] Error posting GIF")

		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("[HTTP > Giphy > SaveContent] Unexpected response status code: %d", resp.StatusCode)

		return nil, err
	}

	responseBuffer, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("[HTTP > Giphy > SaveContent] Error reading response")

		return nil, err
	}

	var response *dto.UploadResponse

	if err = json.Unmarshal(responseBuffer, &response); err != nil {
		log.Println("[HTTP > Giphy > SaveContent] Error converting response")

		return nil, err
	}

	gif.ProviderID = response.Data.ID

	gifWithURL, err := g.findByID(gif)
	if err != nil {
		return nil, err
	}

	return gifWithURL, nil
}
