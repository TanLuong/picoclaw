package main

import (
	"encoding/json"
	"fmt"
)

type ModelConfig struct {
	ModelName string   `json:"model_name"`
	Model     string   `json:"model"`
	apiKeys   []string
}

func (c *ModelConfig) UnmarshalJSON(data []byte) error {
	type Alias ModelConfig
	aux := &struct {
		APIKey  string   `json:"api_key"`
		APIKeys []string `json:"api_keys"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if aux.APIKey != "" {
		c.apiKeys = append(c.apiKeys, aux.APIKey)
	}
	if len(aux.APIKeys) > 0 {
		c.apiKeys = append(c.apiKeys, aux.APIKeys...)
	}
	return nil
}

func (c ModelConfig) MarshalJSON() ([]byte, error) {
	type Alias ModelConfig
	aux := &struct {
		APIKey  string   `json:"api_key,omitempty"`
		APIKeys []string `json:"api_keys,omitempty"`
		Alias
	}{
		Alias: (Alias)(c),
	}
	if len(c.apiKeys) == 1 {
		aux.APIKey = c.apiKeys[0]
	} else if len(c.apiKeys) > 1 {
		aux.APIKeys = c.apiKeys
	}
	return json.Marshal(aux)
}

func main() {
	j := `{"model_name": "test", "model": "test", "api_key": "test_key", "api_keys": ["test_key2"]}`
	var cfg ModelConfig
	json.Unmarshal([]byte(j), &cfg)
	fmt.Printf("%+v\n", cfg)

	bytes, _ := json.Marshal(cfg)
	fmt.Println(string(bytes))
}
