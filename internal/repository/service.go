package repository

import (
	"test-task-go/internal/controller/request"
	"test-task-go/internal/model"
	"test-task-go/pkg"

	"gorm.io/gorm"
)

type ServiceRepository struct {
	Db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) *ServiceRepository {
	return &ServiceRepository{db}
}

func (r *ServiceRepository) Fetch(pagination *request.PaginationService) error {
	var services []model.Service

	query := r.Db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit())
	if pagination.HasSort() {
		query = query.Order(pagination.GetSort())
	}
	if pagination.HasFilter() {
		query = query.Where(pagination.GetFilterField()+" = ?", pagination.GetFilterValue())
	}

	result := query.Find(&services)

	if result.Error != nil {
		return pkg.ErrBadParamInput
	}

	pagination.SetRows(services)
	return nil
}

func (r *ServiceRepository) GetById(id uint64) (m model.Service, err error) {
	result := r.Db.First(&m, id)

	if result.Error != nil {
		return m, pkg.ErrNotFound
	}

	return m, nil
}

func (r *ServiceRepository) Store(m *model.Service) error {
	result := r.Db.Create(&m)

	if result.Error != nil {
		return pkg.ErrBadParamInput
	}

	return nil
}

func (r *ServiceRepository) DeleteById(id uint64) error {
	m := model.Service{}

	result := r.Db.First(&m, id)
	if result.Error != nil {
		return pkg.ErrNotFound
	}

	result = r.Db.Delete(&m, id)
	if result.Error != nil {
		return pkg.ErrInternalServerError
	}

	return nil
}
