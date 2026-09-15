package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (repo *Repository) Create(data any) error {
	return repo.db.Create(data).Error
}

func (repo *Repository) CreateMany(data any) error {
	return repo.db.Create(data).Error
}

func (repo *Repository) FindByID(id any, result any) error {
	return repo.db.First(result, id).Error
}

func (repo *Repository) FindOne(
	filter map[string]any,
	result any,
) error {
	return repo.db.
		Where(filter).
		First(result).
		Error
}

func (repo *Repository) Find(
	filter map[string]any,
	result any,
) error {
	query := repo.db

	if filter != nil {
		query = query.Where(filter)
	}

	return query.Find(result).Error
}

func (repo *Repository) FindWithSelect(
	filter map[string]any,
	selectFields []string,
	result any,
) error {
	query := repo.db

	if filter != nil {
		query = query.Where(filter)
	}

	return query.
		Select(selectFields).
		Find(result).
		Error
}

func (repo *Repository) UpdateOne(
	filter map[string]any,
	update map[string]any,
	model any,
) error {
	return repo.db.
		Model(model).
		Where(filter).
		Updates(update).
		Error
}

func (repo *Repository) UpdateByID(
	id any,
	update map[string]any,
	model any,
) error {
	return repo.db.
		Model(model).
		Where("id = ?", id).
		Updates(update).
		Error
}

func (repo *Repository) FindOneAndUpdate(
	filter map[string]any,
	update map[string]any,
	result any,
) error {
	return repo.db.
		Model(result).
		Where(filter).
		Updates(update).
		Error
}

func (repo *Repository) DeleteOne(
	filter map[string]any,
	model any,
) error {
	return repo.db.
		Where(filter).
		Delete(model).
		Error
}

func (repo *Repository) DeleteByID(
	id any,
	model any,
) error {
	return repo.db.
	Where("id = ?", id).
	Delete(model).
	Error
}

func (repo *Repository) DeleteMany(
	filter map[string]any,
	model any,
) error {
	return repo.db.
	Where(filter).
	Delete(model).
	Error
}

func (repo *Repository) Count(
	filter map[string]any,
	model any,
) (int64, error) {
	var count int64

	err := repo.db.
		Model(model).
		Where(filter).
		Count(&count).
		Error

	return count, err
}

func (repo *Repository) Preload(
	filter map[string]any,
	relation string,
	result any,
) error {
	query := repo.db

	if filter != nil {
		query = query.Where(filter)
	}

	return query.
		Preload(relation).
		Find(result).
		Error
}

func (repo *Repository) PreloadOne(
	filter map[string]any,
	relation string,
	result any,
) error {
	query := repo.db

	if filter != nil {
		query = query.Where(filter)
	}

	return query.
		Preload(relation).
		First(result).
		Error
}

func (repo *Repository) Paginate(
	filter map[string]any,
	page int,
	size int,
	result any,
	model any,
) (int64, int, error) {

	if page < 1 {
		page = 1
	}

	if size < 1 {
		size = 5
	}

	var count int64

	err := repo.db.
		Model(model).
		Where(filter).
		Count(&count).
		Error

	if err != nil {
		return 0, 0, err
	}

	pages := int((count + int64(size) - 1) / int64(size))

	offset := (page - 1) * size

	err = repo.db.
		Where(filter).
		Limit(size).
		Offset(offset).
		Find(result).
		Error

	return count, pages, err
}

func (repo *Repository) Order(
	filter map[string]any,
	order string,
	result any,
) error {
	return repo.db.
		Where(filter).
		Order(order).
		Find(result).
		Error
}

func (repo *Repository) Limit(
	filter map[string]any,
	limit int,
	result any,
) error {
	return repo.db.
		Where(filter).
		Limit(limit).
		Find(result).
		Error
}

func (repo *Repository) Offset(
	filter map[string]any,
	offset int,
	result any,
) error {
	return repo.db.
		Where(filter).
		Offset(offset).
		Find(result).
		Error
}

func (repo *Repository) Select(
	fields []string,
	result any,
) error {
	return repo.db.
		Select(fields).
		Find(result).
		Error
}

func (repo *Repository) Joins(
	join string,
	filter map[string]any,
	result any,
) error {
	query := repo.db.Joins(join)

	if filter != nil {
		query = query.Where(filter)
	}

	return query.Find(result).Error
}