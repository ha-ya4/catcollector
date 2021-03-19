package catcollector

import (
	"fmt"
	"net/http"
	"sync"
)

// NodeHealth nodeのヘルス情報
type NodeHealth struct {
	Status struct {
		APINode string `json:"apiNode"`
		DB      string `json:"db"`
	} `json:"status"`
}

// NodeInfo node/infoにリクエストして得られるsymbolノードの情報
type NodeInfo struct {
	Version               int    `json:"version"`
	PublicKey             string `json:"publicKey"`
	NetworkGenerationHash string `json:"networkGenerationHash"`
	Roles                 int    `json:"roles"`
	Port                  int    `json:"port"`
	NetworkIdentifier     int    `json:"networkIdentifier"`
	Host                  string `json:"host"`
	FriendlyName          string `json:"friendlyName"`
}

// NodePeers nodeが接続しているpeerのリスト
type NodePeers []*NodeInfo

// NodeServer node/serverにリクエストして得られる情報
type NodeServer struct {
	ServerInfo struct {
		RestVersion string `json:"restVersion"`
		SdkVersion  string `json:"sdkVersion"`
	} `json:"serverInfo"`
}

type symbolReqester interface {
	GetNodeInfo(node string) *NodeInfo
	GetNodePeers(node string) NodePeers
	GetNodeHealth(node string) NodeHealth
	GetNodeServer(node string) NodeServer
}

type symbolClient struct {
	c *http.Client
}

func (s symbolClient) GetNodeInfo(node string) *NodeInfo

func (s symbolClient) GetNodePeers(node string) NodePeers

func (s symbolClient) GetNodeHealth(node string) NodeHealth

func (s symbolClient) GetNodeServer(node string) NodeServer

// NodeURL nodeのhost,protcol,portを持つ構造体
type NodeURL struct {
	Host     string
	Protocol string
	Port     string
}

// Join Protocol,Host,Portを結合してURLを作る
func (url NodeURL) Join() string {
	return fmt.Sprintf("%s://%s:%s", url.Protocol, url.Host, url.Port)
}

// NodeData nodeのいろいろな情報をもつ構造体
type NodeData struct {
	Info        *NodeInfo
	Health      NodeHealth
	Server      NodeServer
	Protocol    string
	Port        string
	ReponseTime int64
}

// NodesData NodeDataのスライス
type NodesData []*NodeData

// Selection 引数conditionsで受け取った関数に各nodeの情報を渡しtrueが返るもののみにしてNodesDataスライスを作り直す
func (n NodesData) Selection(conditions func(node *NodeData) bool) NodesData {
	nodes := NodesData{}
	return nodes
}

// TakeoutURL nodeのURLのみ取り出してスライスにいれなおして返す
func (n NodesData) TakeoutURL() []NodeURL {
	url := make([]NodeURL, 0, len(n))
	return url
}

// Collector nodeのpeer情報を元にnodeの情報を取得するための構造体
type Collector struct {
	ExcludNode []string
	Nodes      NodesData
	DebugPrint bool
	client     symbolReqester
	mu         sync.Mutex
	startNode  string
}

// New Collectorを初期化して返す
func New(client *http.Client, startNode string) Collector {
	return Collector{
		client:    symbolClient{c: client},
		startNode: startNode,
	}
}

// Collect nodeのデータを集めてスライスにして返す
func (c *Collector) Collect() NodesData {
	return c.Nodes
}
