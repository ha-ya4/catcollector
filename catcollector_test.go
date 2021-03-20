package catcollector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeURLJoin(t *testing.T) {
	url := NodeURL{
		Host:     "symbol.dev",
		Protocol: "http",
		Port:     "3000",
	}
	expected := "http://symbol.dev:3000"
	assert.Equal(t, expected, url.Join())
}

func TestNodesDataSelection(t *testing.T) {
	data := [][]string{
		{"3000", "http"},
		{"3000", "https"},
		{"3300", "https"},
		{"4000", "http"},
		{"3000", "http"},
	}
	nodes := NodesData{}
	for _, d := range data {
		nodes = append(nodes, &NodeData{APIPort: d[0], Protocol: d[1]})
	}

	conditions := func(node *NodeData) bool {
		return node.APIPort == "3000"
	}
	assert.Len(t, nodes.Selection(conditions), 3)

	conditions = func(node *NodeData) bool {
		return node.APIPort == "3000" && node.Protocol == "http"
	}
	assert.Len(t, nodes.Selection(conditions), 2)
}

func TestNodesDataTakeoutURL(t *testing.T) {
	url := []NodeURL{
		{Host: "nem", Port: "3000", Protocol: "http"},
		{Host: "cat", Port: "4000", Protocol: "http"},
		{Host: "symbol", Port: "8080", Protocol: "https"},
		{Host: "xym", Port: "3000", Protocol: "http"},
		{Host: "xem", Port: "3030", Protocol: "https"},
	}

	nodes := NodesData{}
	for _, u := range url {
		nodes = append(nodes, &NodeData{
			Info:     &NodeInfo{Host: u.Host},
			APIPort:  u.Port,
			Protocol: u.Protocol,
		})
	}

	assert.ElementsMatch(t, url, nodes.TakeoutURL())
}

func TestCollectorOnryNodeInfo(t *testing.T) {
	c := New(nil, "")
	assert.True(t, c.IncludeHealth)
	assert.True(t, c.IncludeServer)

	c.OnryNodeInfo()
	assert.False(t, c.IncludeHealth)
	assert.False(t, c.IncludeServer)
}

func TestCollectorIncludeAllInfo(t *testing.T) {
	c := New(nil, "")
	c.OnryNodeInfo()
	assert.False(t, c.IncludeHealth)
	assert.False(t, c.IncludeServer)

	c.IncluedAllInfo()
	assert.True(t, c.IncludeHealth)
	assert.True(t, c.IncludeServer)
}
