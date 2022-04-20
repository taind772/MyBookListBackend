// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: feedback.sql

package db

import (
	"context"
)

const createFeedback = `-- name: CreateFeedback :one
INSERT INTO feedbacks (
    content,
    created_by
) VALUES (
    $1, $2
) RETURNING id, content, created_by, is_viewed, is_processing, is_resolved, message, modified_at, created_at
`

type CreateFeedbackParams struct {
	Content   string `json:"content"`
	CreatedBy int32  `json:"created_by"`
}

func (q *Queries) CreateFeedback(ctx context.Context, arg CreateFeedbackParams) (Feedback, error) {
	row := q.db.QueryRowContext(ctx, createFeedback, arg.Content, arg.CreatedBy)
	var i Feedback
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.CreatedBy,
		&i.IsViewed,
		&i.IsProcessing,
		&i.IsResolved,
		&i.Message,
		&i.ModifiedAt,
		&i.CreatedAt,
	)
	return i, err
}

const deleteFeedback = `-- name: DeleteFeedback :exec
DELETE FROM feedbacks
WHERE id = $1
`

func (q *Queries) DeleteFeedback(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteFeedback, id)
	return err
}

const getFeedback = `-- name: GetFeedback :one
SELECT id, content, created_by, is_viewed, is_processing, is_resolved, message, modified_at, created_at
FROM feedbacks
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetFeedback(ctx context.Context, id int32) (Feedback, error) {
	row := q.db.QueryRowContext(ctx, getFeedback, id)
	var i Feedback
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.CreatedBy,
		&i.IsViewed,
		&i.IsProcessing,
		&i.IsResolved,
		&i.Message,
		&i.ModifiedAt,
		&i.CreatedAt,
	)
	return i, err
}

const listFeedbacks = `-- name: ListFeedbacks :many
SELECT id, content, created_by, is_viewed, is_processing, is_resolved, message, modified_at, created_at
FROM feedbacks
LIMIT $1
OFFSET $2
`

type ListFeedbacksParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListFeedbacks(ctx context.Context, arg ListFeedbacksParams) ([]Feedback, error) {
	rows, err := q.db.QueryContext(ctx, listFeedbacks, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Feedback{}
	for rows.Next() {
		var i Feedback
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.CreatedBy,
			&i.IsViewed,
			&i.IsProcessing,
			&i.IsResolved,
			&i.Message,
			&i.ModifiedAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateFeedback = `-- name: UpdateFeedback :one
UPDATE feedbacks
SET content = $2
WHERE id = $1
RETURNING id, content, created_by, is_viewed, is_processing, is_resolved, message, modified_at, created_at
`

type UpdateFeedbackParams struct {
	ID      int32  `json:"id"`
	Content string `json:"content"`
}

func (q *Queries) UpdateFeedback(ctx context.Context, arg UpdateFeedbackParams) (Feedback, error) {
	row := q.db.QueryRowContext(ctx, updateFeedback, arg.ID, arg.Content)
	var i Feedback
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.CreatedBy,
		&i.IsViewed,
		&i.IsProcessing,
		&i.IsResolved,
		&i.Message,
		&i.ModifiedAt,
		&i.CreatedAt,
	)
	return i, err
}
