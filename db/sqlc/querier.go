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
	CreateRate(ctx context.Context, arg CreateRateParams) (Rate, error)
	DeleteBookmark(ctx context.Context, id int32) error
	DeleteCategory(ctx context.Context, id int32) error
	DeleteComment(ctx context.Context, id int32) error
	DeleteRate(ctx context.Context, id int32) error
	GetAccount(ctx context.Context, id int32) (Account, error)
	GetAccountByEmail(ctx context.Context, email string) (Account, error)
	GetBook(ctx context.Context, id int32) (BookDetail, error)
	GetBookmark(ctx context.Context, id int32) (Bookmark, error)
	GetCategory(ctx context.Context, id int32) (Category, error)
	GetComment(ctx context.Context, id int32) (Comment, error)
	GetRate(ctx context.Context, id int32) (Rate, error)
	ListBookmarkedBooksByAccountId(ctx context.Context, createdBy int32) ([]ListBookmarkedBooksByAccountIdRow, error)
	ListBooks(ctx context.Context, arg ListBooksParams) ([]ListBooksRow, error)
	ListBooksByCategoryId(ctx context.Context, arg ListBooksByCategoryIdParams) ([]Book, error)
	ListCategories(ctx context.Context) ([]Category, error)
	ListCommentsByBookId(ctx context.Context, arg ListCommentsByBookIdParams) ([]Comment, error)
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error)
	UpdateComment(ctx context.Context, arg UpdateCommentParams) (Comment, error)
	UpdateRate(ctx context.Context, arg UpdateRateParams) (Rate, error)
}

var _ Querier = (*Queries)(nil)
