package mongo



type UserDetails{
	Name string	`json: username`
	Password string `json: password`
	UserId int `json: userID`
	Email string `json: email`
}

type MesssageDetails{

	message string `json: message`
	SenderID string `json: senderid`
	ReceiverID string `json: receiverid`
	messageID string `json: messageid`


}







