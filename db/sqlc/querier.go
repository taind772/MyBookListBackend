// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"context"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateBook(ctx context.Context, arg CreateBookParams) (Book, error)
	CreateBookCategory(ctx context.Context, arg CreateBookCategoryParams) (BookCategory, error)
	CreateBookmark(ctx context.Context, arg CreateBookmarkParams) (Bookmark, error)
	CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error)
	CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error)
	CreateFeedback(ctx context.Context, arg CreateFeedbackParams) (Feedback, error)
	CreateRate(ctx context.Context, arg CreateRateParams) (Rate, error)
	DeleteBook(ctx context.Context, id int32) error
	DeleteBookCategory(ctx context.Context, arg DeleteBookCategoryParams) error
	DeleteBookmark(ctx context.Context, id int32) error
	DeleteCategory(ctx context.Context, id int32) error
	DeleteComment(ctx context.Context, id int32) error
	DeleteFeedback(ctx context.Context, id int32) error
	DeleteRate(ctx context.Context, id int32) error
	GetAccount(ctx context.Context, id int32) (Account, error)
	GetBook(ctx context.Context, id int32) (Book, error)
	GetBookmark(ctx context.Context, id int32) (Bookmark, error)
	GetCategory(ctx context.Context, id int32) (Category, error)
	GetComment(ctx context.Context, id int32) (Comment, error)
	GetFeedback(ctx context.Context, id int32) (Feedback, error)
	GetRate(ctx context.Context, id int32) (Rate, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	ListBookmarksByAccountId(ctx context.Context, arg ListBookmarksByAccountIdParams) ([]Bookmark, error)
	ListBooks(ctx context.Context, arg ListBooksParams) ([]Book, error)
	ListBooksByCategoryId(ctx context.Context, arg ListBooksByCategoryIdParams) ([]Book, error)
	ListCategories(ctx context.Context) ([]Category, error)
	ListComments(ctx context.Context, arg ListCommentsParams) ([]Comment, error)
	ListCommentsByBookId(ctx context.Context, arg ListCommentsByBookIdParams) ([]Comment, error)
	ListFeedbacks(ctx context.Context, arg ListFeedbacksParams) ([]Feedback, error)
	ListRatesByAccountId(ctx context.Context, arg ListRatesByAccountIdParams) ([]Rate, error)
	ListRatesByBookId(ctx context.Context, arg ListRatesByBookIdParams) ([]Rate, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
	UpdateBook(ctx context.Context, arg UpdateBookParams) (Book, error)
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error)
	UpdateComment(ctx context.Context, arg UpdateCommentParams) (Comment, error)
	UpdateFeedback(ctx context.Context, arg UpdateFeedbackParams) (Feedback, error)
	UpdateRate(ctx context.Context, arg UpdateRateParams) (Rate, error)
}

var _ Querier = (*Queries)(nil)