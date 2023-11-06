// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"juice/app/public/ent/videocomment"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// VideoComment is the model entity for the VideoComment schema.
type VideoComment struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// CommentID holds the value of the "comment_id" field.
	CommentID int `json:"comment_id,omitempty"`
	// PcommentID holds the value of the "pcomment_id" field.
	PcommentID int `json:"pcomment_id,omitempty"`
	// VideoID holds the value of the "video_id" field.
	VideoID int `json:"video_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// CommentText holds the value of the "comment_text" field.
	CommentText string `json:"comment_text,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime   time.Time `json:"update_time,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*VideoComment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case videocomment.FieldID, videocomment.FieldCommentID, videocomment.FieldPcommentID, videocomment.FieldVideoID, videocomment.FieldUserID:
			values[i] = new(sql.NullInt64)
		case videocomment.FieldCommentText:
			values[i] = new(sql.NullString)
		case videocomment.FieldCreateTime, videocomment.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the VideoComment fields.
func (vc *VideoComment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case videocomment.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			vc.ID = uint64(value.Int64)
		case videocomment.FieldCommentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field comment_id", values[i])
			} else if value.Valid {
				vc.CommentID = int(value.Int64)
			}
		case videocomment.FieldPcommentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field pcomment_id", values[i])
			} else if value.Valid {
				vc.PcommentID = int(value.Int64)
			}
		case videocomment.FieldVideoID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field video_id", values[i])
			} else if value.Valid {
				vc.VideoID = int(value.Int64)
			}
		case videocomment.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				vc.UserID = int(value.Int64)
			}
		case videocomment.FieldCommentText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comment_text", values[i])
			} else if value.Valid {
				vc.CommentText = value.String
			}
		case videocomment.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				vc.CreateTime = value.Time
			}
		case videocomment.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				vc.UpdateTime = value.Time
			}
		default:
			vc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the VideoComment.
// This includes values selected through modifiers, order, etc.
func (vc *VideoComment) Value(name string) (ent.Value, error) {
	return vc.selectValues.Get(name)
}

// Update returns a builder for updating this VideoComment.
// Note that you need to call VideoComment.Unwrap() before calling this method if this VideoComment
// was returned from a transaction, and the transaction was committed or rolled back.
func (vc *VideoComment) Update() *VideoCommentUpdateOne {
	return NewVideoCommentClient(vc.config).UpdateOne(vc)
}

// Unwrap unwraps the VideoComment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (vc *VideoComment) Unwrap() *VideoComment {
	_tx, ok := vc.config.driver.(*txDriver)
	if !ok {
		panic("ent: VideoComment is not a transactional entity")
	}
	vc.config.driver = _tx.drv
	return vc
}

// String implements the fmt.Stringer.
func (vc *VideoComment) String() string {
	var builder strings.Builder
	builder.WriteString("VideoComment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", vc.ID))
	builder.WriteString("comment_id=")
	builder.WriteString(fmt.Sprintf("%v", vc.CommentID))
	builder.WriteString(", ")
	builder.WriteString("pcomment_id=")
	builder.WriteString(fmt.Sprintf("%v", vc.PcommentID))
	builder.WriteString(", ")
	builder.WriteString("video_id=")
	builder.WriteString(fmt.Sprintf("%v", vc.VideoID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", vc.UserID))
	builder.WriteString(", ")
	builder.WriteString("comment_text=")
	builder.WriteString(vc.CommentText)
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(vc.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(vc.UpdateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// VideoComments is a parsable slice of VideoComment.
type VideoComments []*VideoComment
