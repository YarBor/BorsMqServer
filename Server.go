package MqServer

import (
	"MqServer/ConsumerGroup"
	"MqServer/MessageMem"
	"MqServer/RaftServer"
	pb "MqServer/rpc"
	"context"
	"google.golang.org/grpc"
	"math"
	"sync"
)

type Server interface {
	Serve() error
	Stop() error
}
type ServerClient struct {
	pb.RaftCallClient
	Conn *grpc.ClientConn
}

type connections struct {
	mu             sync.RWMutex
	Conn           map[string]*ServerClient
	MetadataLeader *ServerClient
}

type broker struct {
	pb.UnimplementedMqServerCallServer
	RaftServer RaftServer.RaftServer
	Url        string
	ID         string
	Key        string

	MetaDataController   MetaDataController
	PartitionsController PartitionsController
}

// 客户端和server之间的心跳

// 注册消费者
func (s *broker) RegisterConsumer(_ context.Context, req *pb.RegisterConsumerRequest) (*pb.RegisterConsumerResponse, error) {
	res := s.MetaDataController.RegisterConsumer(req)
	if res.Response.Mode == pb.Response_Success {
		res.Credential.Key = s.Key
	}
	return res, nil
}

// 注册生产者
func (s *broker) RegisterProducer(_ context.Context, req *pb.RegisterProducerRequest) (*pb.RegisterProducerResponse, error) {
	res := s.MetaDataController.RegisterProducer(req)
	if res.Response.Mode == pb.Response_Success {
		res.Credential.Key = s.Key
	}
	return res, nil
}

// 创建话题
func (s *broker) CreateTopic(_ context.Context, req *pb.CreateTopicRequest) (*pb.CreateTopicResponse, error) {
	res := s.MetaDataController.CreateTopic(req)
	if res == nil {
		panic("Ub")
	}
	return res, nil
}
func (s *broker) QueryTopic(_ context.Context, req *pb.QueryTopicRequest) (*pb.QueryTopicResponse, error) {
	res := s.MetaDataController.QueryTopic(req)
	if res == nil {
		panic("Ub")
	}
	return res, nil
}
func (s *broker) DestroyTopic(_ context.Context, req *pb.DestroyTopicRequest) (*pb.DestroyTopicResponse, error) {
	res := s.MetaDataController.DestroyTopic(req)
	if res == nil {
		panic("Ub")
	}
	return res, nil
}

// 注销
func (s *broker) UnRegisterConsumer(_ context.Context, req *pb.UnRegisterConsumerRequest) (*pb.UnRegisterConsumerResponse, error) {
	res := s.MetaDataController.UnRegisterConsumer(req)
	if res == nil {
		panic("Ub")
	}
	return res, nil
}

func (s *broker) UnRegisterProducer(_ context.Context, req *pb.UnRegisterProducerRequest) (*pb.UnRegisterProducerResponse, error) {
	res := s.MetaDataController.UnRegisterProducer(req)
	if res == nil {
		panic("Ub")
	}
	return res, nil
}

// 拉取消息
func (s *broker) PullMessage(_ context.Context, req *pb.PullMessageRequest) (*pb.PullMessageResponse, error) {
	// TODO:
	return nil, nil
}

// 推送消息
func (s *broker) PushMessage(_ context.Context, req *pb.PushMessageRequest) (*pb.PushMessageResponse, error) {
	// TODO:
	return nil, nil
}

func (s *broker) Heartbeat(_ context.Context, req *pb.Ack) (*pb.Response, error) {
	// TODO:
	return nil, nil
}

func (s *broker) JoinConsumerGroup(_ context.Context, req *pb.JoinConsumerGroupRequest) (*pb.JoinConsumerGroupResponse, error) {
	//TODO:
	return nil, nil
}

func (s *broker) LeaveConsumerGroup(_ context.Context, req *pb.LeaveConsumerGroupRequest) (*pb.LeaveConsumerGroupResponse, error) {
	//TODO:
	return nil, nil
}

func (s *broker) CheckSourceTerm(_ context.Context, req *pb.CheckSourceTermRequest) (*pb.CheckSourceTermResponse, error) {
	//TODO:
	return nil, nil
}

func (s *broker) GetCorrespondPartition(_ context.Context, req *pb.GetCorrespondPartitionRequest) (*pb.GetCorrespondPartitionResponse, error) {
	//TODO:
	return nil, nil
}

func (s *broker) RegisterConsumerGroup(_ context.Context, req *pb.RegisterConsumerGroupRequest) (*pb.RegisterConsumerGroupResponse, error) {
	//TODO:
	return nil, nil
}

func (s *broker) UnRegisterConsumerGroup(_ context.Context, req *pb.UnRegisterConsumerGroupRequest) (*pb.UnRegisterConsumerGroupResponse, error) {
	//TODO:
	return nil, nil
}

func ResponseFailure() *pb.Response {
	return &pb.Response{Mode: pb.Response_Failure}
}
func ResponseErrTimeout() *pb.Response {
	return &pb.Response{Mode: pb.Response_ErrTimeout}
}
func ResponseErrNotLeader() *pb.Response {
	return &pb.Response{Mode: pb.Response_ErrNotLeader}
}
func ResponseErrSourceNotExist() *pb.Response {
	return &pb.Response{Mode: pb.Response_ErrSourceNotExist}
}
func ResponseErrSourceAlreadyExist() *pb.Response {
	return &pb.Response{Mode: pb.Response_ErrSourceAlreadyExist}
}
func ResponseErrPartitionChanged() *pb.Response {
	return &pb.Response{Mode: pb.Response_ErrPartitionChanged}
}
func ResponseErrRequestIllegal() *pb.Response {
	return &pb.Response{Mode: pb.Response_ErrRequestIllegal}
}
func ResponseErrSourceNotEnough() *pb.Response {
	return &pb.Response{Mode: pb.Response_ErrSourceNotEnough}
}

func ResponseSuccess() *pb.Response {
	return &pb.Response{Mode: pb.Response_Success}
}
func ResponseNotServer() *pb.Response {
	return &pb.Response{Mode: pb.Response_NotServe}
}

type Partition struct {
	T            string
	P            string
	Consumers    *ConsumerGroup.GroupManager
	MessageEntry *MessageMem.MessageEntry
}

func newPartition(t, p string, MaxEntries, MaxSize uint64, handleTimeout ConsumerGroup.SessionLogoutNotifier) *Partition {
	return &Partition{
		T:            t,
		P:            p,
		Consumers:    ConsumerGroup.NewConsumerHeartBeatManager(handleTimeout),
		MessageEntry: MessageMem.NewMessageEntry(MaxEntries, MaxSize),
	}
}

func (p *Partition) registerConsumer(c *ConsumerGroup.Consumer) error {
	return p.Consumers.RegisterConsumer(c)
}

var (
	defaultMaxEntries = uint64(math.MaxUint64)
	defaultMaxSize    = uint64(math.MaxUint64)
)

type PartitionsController struct {
	mu            sync.RWMutex
	P             map[string]*Partition // key: "Topic/Partition"
	handleTimeout ConsumerGroup.SessionLogoutNotifier
}

func NewPartitionsController(handleTimeout ConsumerGroup.SessionLogoutNotifier) *PartitionsController {
	return &PartitionsController{
		P:             make(map[string]*Partition),
		handleTimeout: handleTimeout,
	}
}

func (pc *PartitionsController) getPartition(t, p string) *Partition {
	pc.mu.RLock()
	part, ok := pc.P[t+"/"+p]
	pc.mu.RUnlock()
	if ok {
		return part
	}
	return nil
}

func (ptc *PartitionsController) RegisterPart(t, p string, MaxEntries, MaxSize uint64) {
	ptc.mu.Lock()
	defer ptc.mu.Unlock()
	if MaxSize == -1 {
		MaxSize = defaultMaxSize
	}
	if MaxEntries == -1 {
		MaxEntries = defaultMaxEntries
	}
	ptc.P[t+"/"+p] = newPartition(t, p, MaxEntries, MaxSize, ptc.handleTimeout)
}
