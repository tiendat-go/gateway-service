package grpc

import (
	"net/http"

	frt "github.com/tiendat-go/common-service/utils/format"
	jsn "github.com/tiendat-go/common-service/utils/json"
	"github.com/tiendat-go/gateway-service/internal/client/grpc"
	pbNotification "github.com/tiendat-go/proto-service/gen/notification/v1"
)

type NotificationController struct {
	pbNotification.UnimplementedNotificationServiceServer
	notificationClient *grpc.NotificationClient
}

func NewNotificationController(notificationClient *grpc.NotificationClient) *NotificationController {
	return &NotificationController{
		notificationClient: notificationClient,
	}
}

func (n *NotificationController) GetNotifications(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	userId := frt.GetString(query.Get("userId"), "1")
	res, err := n.notificationClient.GetNotifications(&pbNotification.GetNotificationsRequest{UserId: userId})
	if err != nil {
		jsn.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	jsn.WriteJSON(w, http.StatusOK, res)
}
