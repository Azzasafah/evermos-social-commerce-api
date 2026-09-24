package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"evermos-backend/internal/dto"
)

type ProvCityService interface {
	GetProvinces() ([]dto.Province, error)
	GetCities(provID string) ([]dto.City, error)
	GetDetailProvince(provID string) (*dto.Province, error)
	GetDetailCity(cityID string) (*dto.City, error)
}

type provCityService struct {
	client       *http.Client
	cacheMutex   sync.RWMutex
	provinces    []dto.Province
	citiesByProv map[string][]dto.City
}

func NewProvCityService() ProvCityService {
	return &provCityService{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		citiesByProv: make(map[string][]dto.City),
	}
}

func (s *provCityService) doRequest(url string, target interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API error: status code %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

func (s *provCityService) GetProvinces() ([]dto.Province, error) {
	s.cacheMutex.RLock()
	if len(s.provinces) > 0 {
		defer s.cacheMutex.RUnlock()
		return s.provinces, nil
	}
	s.cacheMutex.RUnlock()

	var result []dto.Province
	url := "https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json"
	if err := s.doRequest(url, &result); err != nil {
		// Fallback sample provinces
		result = s.fallbackProvinces()
	}

	s.cacheMutex.Lock()
	s.provinces = result
	s.cacheMutex.Unlock()

	return result, nil
}

func (s *provCityService) GetCities(provID string) ([]dto.City, error) {
	s.cacheMutex.RLock()
	if cities, ok := s.citiesByProv[provID]; ok && len(cities) > 0 {
		defer s.cacheMutex.RUnlock()
		return cities, nil
	}
	s.cacheMutex.RUnlock()

	var result []dto.City
	url := fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/regencies/%s.json", provID)
	if err := s.doRequest(url, &result); err != nil {
		result = s.fallbackCities(provID)
	}

	s.cacheMutex.Lock()
	s.citiesByProv[provID] = result
	s.cacheMutex.Unlock()

	return result, nil
}

func (s *provCityService) GetDetailProvince(provID string) (*dto.Province, error) {
	provinces, err := s.GetProvinces()
	if err != nil {
		return nil, err
	}

	for _, p := range provinces {
		if p.ID == provID {
			return &p, nil
		}
	}

	// Fallback lookup
	return &dto.Province{
		ID:   provID,
		Name: "PROVINSI " + provID,
	}, nil
}

func (s *provCityService) GetDetailCity(cityID string) (*dto.City, error) {
	// First check cache across known provinces
	s.cacheMutex.RLock()
	for _, cities := range s.citiesByProv {
		for _, c := range cities {
			if c.ID == cityID {
				defer s.cacheMutex.RUnlock()
				return &c, nil
			}
		}
	}
	s.cacheMutex.RUnlock()

	// If cityID is 4 digits, first 2 digits are province ID (e.g. 1101 -> 11)
	if len(cityID) >= 2 {
		provID := cityID[:2]
		cities, err := s.GetCities(provID)
		if err == nil {
			for _, c := range cities {
				if c.ID == cityID {
					return &c, nil
				}
			}
		}
	}

	return &dto.City{
		ID:         cityID,
		ProvinceID: cityID[:2],
		Name:       "KOTA " + cityID,
	}, nil
}

func (s *provCityService) fallbackProvinces() []dto.Province {
	return []dto.Province{
		{ID: "11", Name: "ACEH"},
		{ID: "12", Name: "SUMATERA UTARA"},
		{ID: "13", Name: "SUMATERA BARAT"},
		{ID: "14", Name: "RIAU"},
		{ID: "15", Name: "JAMBI"},
		{ID: "16", Name: "SUMATERA SELATAN"},
		{ID: "17", Name: "BENGKULU"},
		{ID: "18", Name: "LAMPUNG"},
		{ID: "19", Name: "KEPULAUAN BANGKA BELITUNG"},
		{ID: "21", Name: "KEPULAUAN RIAU"},
		{ID: "31", Name: "DKI JAKARTA"},
		{ID: "32", Name: "JAWA BARAT"},
		{ID: "33", Name: "JAWA TENGAH"},
		{ID: "34", Name: "DI YOGYAKARTA"},
		{ID: "35", Name: "JAWA TIMUR"},
		{ID: "36", Name: "BANTEN"},
		{ID: "51", Name: "BALI"},
	}
}

func (s *provCityService) fallbackCities(provID string) []dto.City {
	if provID == "11" {
		return []dto.City{
			{ID: "1101", ProvinceID: "11", Name: "KABUPATEN SIMEULUE"},
			{ID: "1102", ProvinceID: "11", Name: "KABUPATEN ACEH SINGKIL"},
			{ID: "1103", ProvinceID: "11", Name: "KABUPATEN ACEH SELATAN"},
			{ID: "1171", ProvinceID: "11", Name: "KOTA BANDA ACEH"},
		}
	}
	return []dto.City{
		{ID: provID + "01", ProvinceID: provID, Name: "KABUPATEN PUSAT"},
		{ID: provID + "71", ProvinceID: provID, Name: "KOTA PUSAT"},
	}
}
