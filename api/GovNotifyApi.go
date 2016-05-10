package api

// IGovNotifyAPI - Client to handle communication with Gov Notify API.
type IGovNotifyAPI interface {

	// SendEmail - Emailrequest object representing request to notify by email
	//SendEmailReq(EmailRequest) NotificationCreatedResponse

	// SendEmail -  Valid email address, TemplateID of a valid GovNotify Template
	//SendEmail(string, string)

	// SendSms - SMSrequest object representing request to notify by email
	//SendSmsReq(SMSRequest) NotificationCreatedResponse

	// SendSms -  Valid Phonenumber, TemplateID of a valid GovNotify Template
	//SendSms(string, string)

	// checkStatus - Status request object representing request to check the status of notification
	//checkStatus(StatusRequest) StatusResponse

	// checkStatus - notificationId id of the notification
	//checkStatus(string) StatusResponse
}
