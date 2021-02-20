package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Data is generic for all Data
type Data struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Type string             `bson:"type,omitempty"`
}

// ProjectInvitationData contains data for user.invited.project event
type ProjectInvitationData struct {
	*Data
	ProjectID    string `bson:"projectId,omitempty"`
	ProjectName  string `bson:"projectName,omitempty"`
	UserEmail    string `bson:"userEmail,omitempty"`
	InvitationID string `bson:"invitationId,omitempty"`
}

// Email sent to the user
type Email struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Template    string             `bson:"template,omitempty"`
	Recipient   string             `bson:"recipient,omitempty"`
	SendingTime time.Time          `bson:"sendingTime,omitempty"`
	ReadingTime time.Time          `bson:"readingTime,omitempty"`
	DataID      primitive.ObjectID `bson:"dataId,omitempty"`
}

// Notification model handled
type Notification struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Description string             `bson:"description,omitempty"`
	CreatedAt   time.Time          `bson:"createdAt,omitempty"`
	DataID      string             `bson:"dataId,omitempty"`
}
