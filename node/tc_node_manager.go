package node

import (
	"context"
	"log"
	"net"
	"strings"

	"github.com/nminhquan/mas/transaction"

	"github.com/nminhquan/mas/config"

	"github.com/nminhquan/mas/client"
	"github.com/nminhquan/mas/dao"
	"github.com/nminhquan/mas/db"

	"github.com/nminhquan/mas/model"

	"github.com/nminhquan/mas/credentials"
	pb "github.com/nminhquan/mas/proto"
	"google.golang.org/grpc"
	grpc_creds "google.golang.org/grpc/credentials"
)

type set map[string]bool

var (
	cacheService *db.CacheService
	txnDao       dao.TxnCoordinatorDAO
)

type TCNodeManager struct {
	port string
}

func init() {
	log.Println("TC node mgr init")
	cacheService = db.NewCacheService(config.RedisHost, "")
	txnDao = dao.NewTxnCoordinatorDAO(cacheService, nil)
}

func CreateTCNodeManagerServer(port string) *TCNodeManager {
	return &TCNodeManager{port}
}

func (nm *TCNodeManager) Start() {
	lis, err := net.Listen("tcp", nm.port)
	if err != nil {
		log.Fatalf("[TCNodeManager] failed to listen: %v port : %v", err, nm.port)
	}
	// register JWTServerInterceptor for authentication
	creds, err := grpc_creds.NewServerTLSFromFile(credentials.SSL_SERVER_CERT, credentials.SSL_SERVER_PRIVATE_KEY)
	if err != nil {
		log.Fatalf("[TCNodeManager] Cannot get credentials: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(credentials.JWTServerInterceptor), grpc.Creds(creds))
	pb.RegisterNodeServiceServer(s, nm)
	log.Printf("[TCNodeManager] RPC server started at port: %s", nm.port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("[TCNodeManager] failed to serve: %v", err)
	}
}

func (nm *TCNodeManager) AddNode(ctx context.Context, in *pb.NodeRequest) (*pb.NodeReturn, error) {
	peerMap := txnDao.GetPeersList()

	currentPeers := strings.Split(peerMap[in.ClusterId], ",")
	if in.JoinRaft {
		for _, raftMember := range currentPeers {
			log.Printf("[TCNodeManager] %v JoinRaft, sending request to: %v", in.RaftNodeAddr, raftMember)
			client := client.CreateNodeClient("", raftMember)
			ok := client.JoinExistingRaft(in.ClusterId, in.NodeId, in.RaftNodeAddr)
			if ok {
				log.Printf("[TCNodeManager] %v JoinRaft successfully", in.RaftNodeAddr)
				break
			} else {
				log.Printf("[TCNodeManager] %v JoinRaft retrying...", in.RaftNodeAddr)
			}
		}

	}

	peerSet := make(set)
	for _, peer := range currentPeers {
		peerSet[peer] = true
	}
	peerSet[in.RmNodeAddr] = true
	log.Println("[TCNodeManager] AddNode with peerString = ", makePeerString(peerSet))
	_, err := txnDao.InsertPeers(in.ClusterId, makePeerString(peerSet))
	if err != nil {
		return &pb.NodeReturn{Message: model.RPC_MESSAGE_FAIL}, err
	}
	transaction.RefreshPeerList()
	return &pb.NodeReturn{Message: model.RPC_MESSAGE_OK}, nil
}

func makePeerString(peerSet set) string {
	var keys string
	for key := range peerSet {
		keys += key + ","
	}
	return strings.Trim(keys, ",")
}
