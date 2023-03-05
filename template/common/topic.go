package common

const (
	// restaurant
	//TopicUserLikeRestaurant                   = "TopicUserLikeRestaurant"
	//TopicUserDislikeRestaurant                = "TopicUserDislikeRestaurant"
	TopicCreateRestaurantFoodsAfterCreateFood = "TopicCreateRestaurantFoodsAfterCreateFood"

	// socket
	TopicHandleOrderWhenUserOrderFood       = "TopicHandleOrderWhenUserOrderFood"
	TopicEmitEvenWhenUserCreateOrderSuccess = "TopicEmitEvenWhenUserCreateOrderSuccess"

	// order
	TopicCreateOrderTrackingAfterCreateOrderDetail = "TopicCreateOrderTrackingAfterCreateOrderDetail"

	// Food
	TopicUserLikeFood    = "TopicUserLikeFood"
	TopicUserDislikeFood = "TopicUserDislikeFood"
)
