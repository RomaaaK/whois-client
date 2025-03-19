package repository

import (
	"errors"
	"slices"
	"time"
	"whois-client/internal/models"
)

type CommentRepository struct {
	comments []models.Comment
}

func NewCommentRepository() *CommentRepository {
	return &CommentRepository{}
}

func (c *CommentRepository) GetByUserID(id uint64) ([]models.Comment, error) {
	var result []models.Comment

	for _, comment := range c.comments {
		if comment.UserID == id {
			result = append(result, comment)
		}
	}

	if len(result) == 0 {
		return nil, errors.New("comment not found")
	}

	return result, nil
}

func (c *CommentRepository) Create(userID uint64, text string) error {
	var id uint64 = 1
	if len(c.comments) > 0 {
		id = c.comments[len(c.comments)-1].UserID + 1
	}

	c.comments = append(c.comments, models.Comment{ID: 1, UserID: id, Text: text, CreatedAt: time.Now()})
	return nil
}

func (c *CommentRepository) Delete(id uint64) error {
	for i, comment := range c.comments {
		if comment.ID == id {
			c.comments = slices.Delete(c.comments, i, i+1)
			return nil
		}
	}

	return errors.New("comment not found")
}
