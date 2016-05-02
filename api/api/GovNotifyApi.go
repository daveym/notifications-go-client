package api

//  GovNotifyApi - Client to handle communication with Gov Notify API.
type GovNotifyApi interface {
	sendEmail(request emailRequest) NotificationCreatedResponse
}
