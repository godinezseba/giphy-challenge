package adapter

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"giphy/challenge/src/domain/entities"
	"giphy/challenge/src/infrastructure/http/giphy/dto"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

func createFormData(gif *entities.GIF) (*bytes.Buffer, *multipart.Writer, error) {
	apiToken := os.Getenv("GIPHY_API_KEY")

	fileName := fmt.Sprintf("%s.gif", strings.ReplaceAll(gif.Name, " ", "_"))

	var buffer bytes.Buffer
	formWriter := multipart.NewWriter(&buffer)

	decoded, err := base64.StdEncoding.DecodeString(gif.Content)
	if err != nil {
		log.Println("[HTTP > Giphy > SaveContent] Error transforming file to binary")

		return nil, nil, err
	}

	fileFormWriter, err := formWriter.CreateFormFile("file", fileName)
	if err != nil {
		log.Println("[HTTP > Giphy > SaveContent] Error creating request on file")

		return nil, nil, err
	}

	_, err = fileFormWriter.Write(decoded)
	if err != nil {
		panic(err)
	}

	if err := formWriter.WriteField("api_key", apiToken); err != nil {
		log.Println("[HTTP > Giphy > SaveContent] Error creating request on APIKey")

		return nil, nil, err
	}

	if err := formWriter.WriteField("tags", strings.Join(gif.Tags, ",")); err != nil {
		log.Println("[HTTP > Giphy > SaveContent] Error creating request on Tags")

		return nil, nil, err
	}

	return &buffer, formWriter, nil
}

func (g *gifHTTPAdapter) SaveContent(ctx context.Context, gif *entities.GIF) (*entities.GIF, error) {
	baseURL := os.Getenv("GIPHY_UPLOAD_URL")

	formData, formWriter, err := createFormData(gif)

	if err != nil {
		return nil, err
	}

	resp, err := g.httpClient.Post(
		baseURL,
		formWriter.FormDataContentType(),
		formData,
	)

	if err != nil {
		log.Println("[HTTP > Giphy > SaveContent] Error posting GIF")

		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("[HTTP > Giphy > SaveContent] Unexpected response status code: %d", resp.StatusCode)

		return nil, errors.New("unexpected response status code")
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

	log.Println("[HTTP > Giphy > SaveContent] Success creating GIF", response)

	gif.ProviderID = response.Data.ID

	gifWithURL, err := g.findByID(gif)
	if err != nil {
		return nil, err
	}

	return gifWithURL, nil
}
