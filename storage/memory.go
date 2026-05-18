package storage

import "errors"

// MapStorage encapsula nosso mapa de URLs
type MapStorage struct {
	urls map[string]string
}

// NewMapStorage inicializa e retorna uma nova instância
func NewMapStorage() *MapStorage {
	return &MapStorage{
		urls: make(map[string]string),
	}
}

// Save guarda a URL longa associada ao código encurtado
func (s *MapStorage) Save(code string, originalURL string) {
	s.urls[code] = originalURL
}

// Get busca a URL original. Retorna erro se não encontrar.
func (s *MapStorage) Get(code string) (string, error) {
	url, exists := s.urls[code]
	if !exists {
		return "", errors.New("URL não encontrada")
	}
	return url, nil
}
