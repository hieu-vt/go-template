package common

const (
	DbTypeRestaurant = 1
	DbTypeFood       = 2
	DbTypeCategory   = 3
	DbTypeUser       = 4
	DbTypeOrder      = 5
	DbTypeCart       = 6
	DbTypeFoodRating = 7
)

const (
	CurrentUser = "user"

	DBMain               = "mysql"
	PluginGrpcUserClient = "Grpc_User_Client"
	PluginGrpcAuthClient = "Grpc_auth_client"
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}

const (
	EmitUserCreateOrderSuccess = "EmitUserCreateOrderSuccess"
	EmitUserOrderFailure       = "EmitUserOrderFailure"
	EmitAuthenticated          = "EmitAuthenticated"
)

const (
	EvenAuthenticated       = "EvenAuthenticated"
	EvenUserCreateOrder     = "EvenUserCreateOrder"
	EventUserUpdateLocation = "EventUserUpdateLocation"
)

const (
	OrderTracking = "OrderTracking"
)

type TrackingType string

const (
	WaitingForShipper TrackingType = "waiting_for_shipper"
	Preparing         TrackingType = "preparing"
	Cancel            TrackingType = "cancel"
	OnTheWay          TrackingType = "on_the_way"
	Delivered         TrackingType = "delivered"
)

const (
	TRACE_SERVICE = "trace-demo"
	ENVIRONMENT   = "dev"
)
