package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/yogyrahmawan/tcgdex-go-sdk/pkg/model"
)

func decodeJSONResponse[T any](resp *http.Response, target *T) error {
	if resp.StatusCode != http.StatusOK {
		var httpErr model.TcgdexHttpError
		if err := json.NewDecoder(resp.Body).Decode(&httpErr); err != nil {
			return fmt.Errorf("decode error response: %w", err)
		}
		return errors.New(httpErr.String())
	}

	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		return fmt.Errorf("decode response body: %w", err)
	}

	return nil
}
