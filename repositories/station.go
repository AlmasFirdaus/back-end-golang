package repositories

import (
	"back-end-golang/models"

	"gorm.io/gorm"
)

type StationRepository interface {
	GetAllStations(page, limit int) ([]models.Station, int, error)
	GetStationByID(id uint) (models.Station, error)
	CreateStation(station models.Station) (models.Station, error)
	UpdateStation(station models.Station) (models.Station, error)
	DeleteStation(station models.Station) error
}

type stationRepository struct {
	db *gorm.DB
}

func NewStationRepository(db *gorm.DB) StationRepository {
	return &stationRepository{db}
}

// Implementasi fungsi-fungsi dari interface ItemRepository

func (r *stationRepository) GetAllStations(page, limit int) ([]models.Station, int, error) {
	var (
		stations []models.Station
		count    int64
	)
	err := r.db.Find(&stations).Count(&count).Error
	if err != nil {
		return stations, int(count), err
	}

	offset := (page - 1) * limit

	err = r.db.Limit(limit).Offset(offset).Find(&stations).Error

	return stations, int(count), err
}

func (r *stationRepository) GetStationByID(id uint) (models.Station, error) {
	var station models.Station
	err := r.db.Where("id = ?", id).First(&station).Error
	return station, err
}

func (r *stationRepository) CreateStation(station models.Station) (models.Station, error) {
	err := r.db.Create(&station).Error
	return station, err
}

func (r *stationRepository) UpdateStation(station models.Station) (models.Station, error) {
	err := r.db.Save(&station).Error
	return station, err
}

func (r *stationRepository) DeleteStation(station models.Station) error {
	err := r.db.Delete(&station).Error
	return err
}