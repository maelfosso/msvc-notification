package models

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Data is generic for all Data
type Data struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Type string             `bson:"type,omitempty" json:"type,omitempty"`
}

// ProjectInvitationData contains data for user.invited.project event
type ProjectInvitationData struct {
	*Data
	ProjectID    string `json:"projectId,omitempty" bson:"projectId,omitempty"`
	ProjectName  string `json:"projectName,omitempty" bson:"projectName,omitempty"`
	UserEmail    string `json:"email,omitempty" bson:"userEmail,omitempty"`
	InvitationID string `json:"invitationId,omitempty" bson:"invitationId,omitempty"`
}

// NewProjectInvitationData create a data
func NewProjectInvitationData() *ProjectInvitationData {
	return &ProjectInvitationData{}
}

// ParseData from JSON
func (data *ProjectInvitationData) ParseData(body []byte) error {
	err := json.Unmarshal(body, data)
	if err != nil {
		return err
	}

	return nil
}

// ParseTemplate load the corresponding HTML Template
func (data *ProjectInvitationData) ParseTemplate(templateFileName string) (string, error) {
	log.Println("Project ID", data.ProjectID)

	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}

	log.Println("Template loaded... ", buf.String())

	return buf.String(), nil
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
