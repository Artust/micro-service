package streaming

import (
	"avatar/services/streaming/domain/broker"
	pb "avatar/services/streaming/protos"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

func ListenListPos(input *pb.ListenListPosRequest, stream pb.Streaming_ListenListPosServer, broker broker.Broker) error {
	ctx := context.Background()
	response := &pb.ListenListPosResponse{}
	listPosResponse := pb.ListPosResponse{
		PosId:            100,
		TalkSessionId:    100,
		Status:           2,
		MainPos:          false,
		CameraId:         []int64{},
		DefaultCameraId:  0,
		ServerUri:        "",
		Name:             "",
		Address:          "",
		StartTimeDeteted: "",
	}
	response.ListPosResponse = append(response.ListPosResponse, &listPosResponse)
	err := stream.Send(response)
	if err != nil {
		return err
	}
	dataChannel, errChannel, closeConsumer := broker.Consume(ctx, fmt.Sprintf("ListenListPos-%s", input.GroupId), true)
	for {
		select {
		case <-ctx.Done():
			err := closeConsumer()
			if err != nil {
				fmt.Println(err)
			}
			err = broker.DeleteTopic(fmt.Sprintf("ListenListPos-%s", input.GroupId))
			if err != nil {
				fmt.Println(err)
			}
			return ctx.Err()
		case data := <-dataChannel:
			err := json.Unmarshal(data, response)
			if err != nil {
				log.Println("error when unmarshal events:", err)
				return err
			}
			err = stream.Send(response)
			if err != nil {
				return err
			}
		case err := <-errChannel:
			return err
		}
	}
}
