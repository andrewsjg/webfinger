package mastodon

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type MastodonClient struct {
    baseURL    string
    httpClient *http.Client
}

func NewMastodonClient(baseURL string) *MastodonClient {
    return &MastodonClient{
        baseURL:    baseURL,
        httpClient: &http.Client{},
    }
}

func (c *MastodonClient) GetProfile(username string) (*Profile, error) {
    url := fmt.Sprintf("%s/api/v1/accounts/lookup?acct=%s", c.baseURL, username)
    resp, err := c.httpClient.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("failed to get profile: %s", resp.Status)
    }

    var profile Profile
    if err := json.NewDecoder(resp.Body).Decode(&profile); err != nil {
        return nil, err
    }

    return &profile, nil
}