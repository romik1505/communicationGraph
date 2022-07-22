package mapper

import (
	"github.com/romik1505/communicationGraph/internal/app/model"
)

type Graph struct {
	Nodes  []Node `json:"nodes"`
	Links  []Link `json:"links"`
	Detail Detail `json:"details"`
}

type Node struct {
	UserID   string `json:"user_id"`
	UserName string `json:"username"`
}

type Link struct {
	SourceID string `json:"source_id"`
	TargetID string `json:"target_id"`
	Count    int    `json:"count"`
}

type Detail struct {
	Total int     `json:"total"`
	Min   int     `json:"min"`
	Avg   float32 `json:"avg"`
	Max   int     `json:"max"`
}

func ConverterLink(c model.Communication) Link {
	return Link{
		SourceID: c.SourceID.String,
		TargetID: c.TargetID.String,
		Count:    int(c.Count.Int32),
	}
}

// ConverterLinks Возращает дуги графа
func ConverterLinks(c []model.Communication) []Link {
	res := make([]Link, 0, len(c))
	for _, v := range c {
		res = append(res, ConverterLink(v))
	}
	return res
}

// ConvertNodes Возвращает вершины графа
func ConvertNodes(cs []model.Communication) []Node {
	set := make(map[string]struct{}, 0)
	for _, v := range cs {
		set[v.SourceID.String] = struct{}{}
		set[v.TargetID.String] = struct{}{}
	}
	nodes := make([]Node, 0)
	for key, _ := range set {
		nodes = append(nodes, Node{
			UserID: key,
		})
	}
	return nodes
}
