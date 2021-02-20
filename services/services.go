package services

// NotificationService offered by the micro-service
type NotificationService interface {
	List() error
	MarkAsRead() error
	Create() error
}

// type notificationService struct {
// 	email string

// }

// // NewService make a new NotificationService
// func NewService() NotificationService {
// 	return notificationService{}
// }

// func (notificationService) List() error {

// }

// func (notificationService) MarkAsRead() error {

// }

// func (notificationService) Create() error {

// }
