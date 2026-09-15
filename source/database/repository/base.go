package repository

import (
	"github.com/abdulrhman-elghnam/golang/source/database"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository() *Repository {
	return &Repository{
		database.DB,
	}
}

func (r *Repository) Create(data any) error {
	return r.db.Create(data).Error
}

func (r *Repository) CreateMany(data any) error {
	return r.db.Create(data).Error
}

func (r *Repository) FindByID(id any, result any) error {
	return r.db.First(result, id).Error
}

func (r *Repository) FindOne(
	filter map[string]any,
	result any,
) error {
	return r.db.
		Where(filter).
		First(result).
		Error
}

func (r *Repository) Find(
	filter map[string]any,
	result any,
) error {
	query := r.db

	if filter != nil {
		query = query.Where(filter)
	}

	return query.Find(result).Error
}

func (r *Repository) FindWithSelect(
	filter map[string]any,
	selectFields []string,
	result any,
) error {
	query := r.db

	if filter != nil {
		query = query.Where(filter)
	}

	return query.
		Select(selectFields).
		Find(result).
		Error
}

func (r *Repository) UpdateOne(
	filter map[string]any,
	update map[string]any,
	model any,
) error {
	return r.db.
		Model(model).
		Where(filter).
		Updates(update).
		Error
}

func (r *Repository) UpdateByID(
	id any,
	update map[string]any,
	model any,
) error {
	return r.db.
		Model(model).
		Where("id = ?", id).
		Updates(update).
		Error
}

func (r *Repository) FindOneAndUpdate(
	filter map[string]any,
	update map[string]any,
	result any,
) error {
	return r.db.
		Model(result).
		Where(filter).
		Updates(update).
		Error
}

func (r *Repository) DeleteOne(
	filter map[string]any,
	model any,
) error {
	return r.db.
		Where(filter).
		Delete(model).
		Error
}

func (r *Repository) DeleteByID(
	id any,
	model any,
) error {
	return r.db.
	Where("id = ?", id).
	Delete(model).
	Error
}

func (r *Repository) DeleteMany(
	filter map[string]any,
	model any,
) error {
	return r.db.
	Where(filter).
	Delete(model).
	Error
}

func (r *Repository) Count(
	filter map[string]any,
	model any,
) (int64, error) {
	var count int64

	err := r.db.
		Model(model).
		Where(filter).
		Count(&count).
		Error

	return count, err
}

func (r *Repository) Preload(
	filter map[string]any,
	relation string,
	result any,
) error {
	query := r.db

	if filter != nil {
		query = query.Where(filter)
	}

	return query.
		Preload(relation).
		Find(result).
		Error
}

func (r *Repository) PreloadOne(
	filter map[string]any,
	relation string,
	result any,
) error {
	query := r.db

	if filter != nil {
		query = query.Where(filter)
	}

	return query.
		Preload(relation).
		First(result).
		Error
}

func (r *Repository) Paginate(
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

	err := r.db.
		Model(model).
		Where(filter).
		Count(&count).
		Error

	if err != nil {
		return 0, 0, err
	}

	pages := int((count + int64(size) - 1) / int64(size))

	offset := (page - 1) * size

	err = r.db.
		Where(filter).
		Limit(size).
		Offset(offset).
		Find(result).
		Error

	return count, pages, err
}

func (r *Repository) Order(
	filter map[string]any,
	order string,
	result any,
) error {
	return r.db.
		Where(filter).
		Order(order).
		Find(result).
		Error
}

func (r *Repository) Limit(
	filter map[string]any,
	limit int,
	result any,
) error {
	return r.db.
		Where(filter).
		Limit(limit).
		Find(result).
		Error
}

func (r *Repository) Offset(
	filter map[string]any,
	offset int,
	result any,
) error {
	return r.db.
		Where(filter).
		Offset(offset).
		Find(result).
		Error
}

func (r *Repository) Select(
	fields []string,
	result any,
) error {
	return r.db.
		Select(fields).
		Find(result).
		Error
}

func (r *Repository) Joins(
	join string,
	filter map[string]any,
	result any,
) error {
	query := r.db.Joins(join)

	if filter != nil {
		query = query.Where(filter)
	}

	return query.Find(result).Error
}