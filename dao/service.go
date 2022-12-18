package dao

type ServiceDetail struct {
	Info          *ServiceInfo   `json:"info" description:"基本信息"`
	HTTPRule      *HttpRule      `json:"http" description:"http_rule"`
	TCPRule       *TcpRule       `json:"tcp" description:"tcp_rule"`
	GRPCRule      *GrpcRule      `json:"grpc" description:"grpc_rule"`
	LoadBalance   *LoadBalance   `json:"load_balance" description:"grpc_rule"`
	AccessControl *AccessControl `json:"access_control" description:"grpc_rule"`
}
