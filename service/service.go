package service

import (
	"fmt"
	"github.com/c12s/blackhole/helper"
	"github.com/c12s/blackhole/model"
	"github.com/c12s/blackhole/queue"
	storage "github.com/c12s/blackhole/storage"
	aPb "github.com/c12s/scheme/apollo"
	pb "github.com/c12s/scheme/blackhole"
	sg "github.com/c12s/stellar-go"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

type Server struct {
	Queue      *queue.BlackHole
	Apollo     string
	instrument map[string]string
}

func (s *Server) getTK(ctx context.Context, req *pb.PutReq) (*queue.TaskQueue, error) {
	span, _ := sg.FromGRPCContext(ctx, "blackhole.getTK")
	fmt.Println("SPAN: ", span)

	defer span.Finish()
	fmt.Println(span)

	if req.Mtdata.ForceNamespaceQueue {
		fmt.Println("FORCED")
		tk, err := s.Queue.GetTK(sg.NewTracedGRPCContext(ctx, span), req.Mtdata.Namespace)
		if err != nil {
			span.AddLog(&sg.KV{"Queue GetTK ForceNamespace error", err.Error()})
			return nil, err
		}
		return tk, nil
	}
	tk, err := s.Queue.GetTK(sg.NewTracedGRPCContext(ctx, span), req.Mtdata.Queue)
	fmt.Println()
	fmt.Println("TO QUEUE")
	fmt.Println(tk.Namespace, tk.Queue)
	fmt.Println(tk)
	fmt.Println()
	if err != nil {
		span.AddLog(&sg.KV{"Queue GetTK error", err.Error()})
		return nil, err
	}
	return tk, nil
}

func (s *Server) extractRBACData(req *pb.PutReq, token string) map[string]string {
	data := map[string]string{
		"intent":         "mutate",
		"kind":           strings.ToLower(pb.TaskKind_name[int32(req.Kind)]),
		"service":        "blackhole",
		"userid":         req.UserId,
		"namespace":      req.Mtdata.Namespace,
		"forceNamespace": strconv.FormatBool(req.Mtdata.ForceNamespaceQueue),
		"queue":          req.Mtdata.Queue,
		"token":          token,
	}

	temp := map[string][]string{}
	for _, t := range req.Tasks {
		temp[t.RegionId] = append(temp[t.RegionId], t.ClusterId)
	}

	for k, v := range temp {
		data[k] = strings.Join(v, ",")
	}

	return data
}

func (s *Server) Put(ctx context.Context, req *pb.PutReq) (*pb.Resp, error) {
	span, _ := sg.FromGRPCContext(ctx, "blackhole.put")
	defer span.Finish()
	fmt.Println(span)
	fmt.Println("SERIALIZE ", span.Serialize())

	fmt.Println("STIGLO: ", req)

	token, terr := helper.ExtractToken(ctx)
	if terr != nil {
		span.AddLog(&sg.KV{"token error", terr.Error()})
	}

	client := NewApolloClient(s.Apollo)
	resp, err := client.Auth(
		helper.AppendToken(
			sg.NewTracedGRPCContext(ctx, span),
			token,
		),
		&aPb.AuthOpt{
			Data: map[string]string{"intent": "auth"},
		},
	)
	if err != nil {
		span.AddLog(&sg.KV{"apollo resp error", err.Error()})
		return &pb.Resp{Msg: err.Error()}, nil
	}

	if !resp.Value {
		span.AddLog(
			&sg.KV{"apollo.auth value", strconv.FormatBool(resp.Value)},
			&sg.KV{"apollo.auth message", resp.Data["message"]},
		)
		return &pb.Resp{Msg: resp.Data["message"]}, nil
	}

	tk, err := s.getTK(sg.NewTracedGRPCContext(ctx, span), req)
	if err != nil {
		span.AddLog(&sg.KV{"blackhole.Put getTK error", err.Error()})
		return nil, err
	}

	pResp, err := tk.PutTasks(
		helper.AppendToken(
			sg.NewTracedContext(ctx, span),
			token,
		),
		req,
	)
	if err != nil {
		span.AddLog(&sg.KV{"blackhole.Put PutTasks error", err.Error()})
		log.Println(err)
	}

	span.AddLog(&sg.KV{"blackhole.Put ok", pResp.Msg})

	// return to client that task is accepted!
	return &pb.Resp{Msg: pResp.Msg}, nil
}

func Run(db storage.DB, conf *model.BlackHoleConfig) {
	trace := sg.Init("blackhole")
	defer trace.Finish()

	span := trace.Span("run")
	defer span.Finish()
	fmt.Println(span)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	lis, err := net.Listen("tcp", conf.Address)
	if err != nil {
		span.AddLog(&sg.KV{"blackhole.run error", err.Error()})
		log.Fatalf("failed to initializa TCP listen: %v", err)
	}
	defer lis.Close()

	server := grpc.NewServer()
	blackholeServer := &Server{
		Queue:      queue.New(sg.NewTracedContext(ctx, span), db, conf.Opts, conf.Celestial, conf.Apollo),
		Apollo:     conf.Apollo,
		instrument: conf.InstrumentConf,
	}

	n, err := sg.NewCollector(blackholeServer.instrument["address"], blackholeServer.instrument["stopic"])
	if err != nil {
		fmt.Println(err)
		return
	}
	c, err := sg.InitCollector(blackholeServer.instrument["location"], n)
	if err != nil {
		fmt.Println(err)
		return
	}
	go c.Start(ctx, 15*time.Second)

	fmt.Println("BlackHoleService RPC Started")
	pb.RegisterBlackHoleServiceServer(server, blackholeServer)
	server.Serve(lis)
}
