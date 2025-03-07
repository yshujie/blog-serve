package mysql

import (
	"github.com/jinzhu/gorm"
	m "github.com/yshujie/blog-serve/internal/model"
)

type ArticleRepository struct {
	db *gorm.DB
}

// new article repository
func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

// create article
func (r *ArticleRepository) Create(article *m.Article) error {
	return r.db.Create(article).Error
}

// update article
func (r *ArticleRepository) Update(article *m.Article) error {
	return r.db.Save(article).Error
}

// delete article
func (r *ArticleRepository) Delete(article *m.Article) error {
	return r.db.Delete(article).Error
}

// get article by id
func (r *ArticleRepository) GetById(id int) (*m.Article, error) {
	var article m.Article
	return &article, r.db.Where("id = ?", id).First(&article).Error
}

// get article by title
func (r *ArticleRepository) GetByTitle(title string) (*m.Article, error) {
	var article m.Article
	return &article, r.db.Where("title = ?", title).First(&article).Error
}

// get all articles
func (r *ArticleRepository) GetAll() ([]m.Article, error) {
	var articles []m.Article
	return articles, r.db.Find(&articles).Error
}
