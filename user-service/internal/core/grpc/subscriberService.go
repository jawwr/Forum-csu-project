package grpc

import (
	"user-service/internal/core/interface/service"
	"user-service/internal/core/mapper"
	serviceManager "user-service/internal/core/service"
	pb "user-service/proto/generated/subscriberService"
)

type _subscriberServiceImpl struct {
	pb.UnimplementedSubscriberServiceServer
	subscriberService service.SubscriberService
}

func NewSubscriberService(manager serviceManager.Manager) pb.SubscriberServiceServer {
	return _subscriberServiceImpl{subscriberService: manager.SubscriberService}
}

func (s _subscriberServiceImpl) GetAllUserSubscribers(user *pb.UserRequest, stream pb.SubscriberService_GetAllUserSubscribersServer) error {
	ctx := stream.Context()
	savedSubscribers, err := s.subscriberService.GetAllSubscribers(ctx, int(user.Id))
	if err != nil {
		return err
	}
	for _, user := range savedSubscribers {
		subscriber := mapper.ToPbSubscriber(&user)
		if err := stream.Send(subscriber); err != nil {
			return err
		}
	}
	return nil
}
