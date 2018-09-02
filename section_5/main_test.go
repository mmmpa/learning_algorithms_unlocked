package main

import (
	"testing"
	"math/rand"
	"github.com/k0kubun/pp"
)

func shuffle(data []*Node) []*Node {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}

	return data
}

func TestCompute(t *testing.T) {
	names := []string{
		"",
		"undershorts",
		"socks",
		"compression shorts",
		"hose",
		"cup",
		"pants",
		"skates",
		"leg pads",
		"T-shirt",
		"chest pad",
		"sweater",
		"mask",
		"catch glove",
		"blocker",
	}
	nodes := make([]*Node, len(names))

	for i, n := range names {
		nodes[i] = NewNode(n)
	}

	nodes[1].AddNear(nodes[3])
	nodes[2].AddNear(nodes[4])
	nodes[3].AddNear(nodes[4])
	nodes[3].AddNear(nodes[5])
	nodes[4].AddNear(nodes[6])
	nodes[5].AddNear(nodes[6])
	nodes[6].AddNear(nodes[7])
	nodes[7].AddNear(nodes[8])
	nodes[8].AddNear(nodes[13])
	nodes[9].AddNear(nodes[10])
	nodes[10].AddNear(nodes[11])
	nodes[11].AddNear(nodes[12])
	nodes[12].AddNear(nodes[13])
	nodes[13].AddNear(nodes[14])

	shuffle(nodes[1:])
	namePrinter(topologicalSort(nodes[1:]))
}

type N struct {
	Name   string
	Weight int
}

func TestCompute2(t *testing.T) {
	names := []struct {
		Name   string
		Weight int
		Nears  []string
	}{
		{"マリネードを合わせる", 2, []string{"チキンをマリネードにつける"}},
		{"にんにくを刻む", 4, []string{"にんにくと生姜を加える"}},
		{"生姜を刻む", 3, []string{"にんにくと生姜を加える"}},
		{"人参を刻む", 4, []string{"人参セロリピーナッツに火を入れる"}},
		{"セロリを刻む", 3, []string{"人参セロリピーナッツに火を入れる"}},
		{"ピーナッツを洗う", 2, []string{"人参セロリピーナッツに火を入れる"}},
		{"クッキングソースを合わせる", 3, []string{"クッキングソースをかける"}},
		{"チキンを焼く", 6, []string{"チキンをマリネードにつける"}},
		{"チキンをマリネードにつける", 15, []string{"チキンに少し火を入れる"}},
		{"チキンに少し火を入れる", 4, []string{"にんにくと生姜を加える"}},
		{"にんにくと生姜を加える", 1, []string{"チキンを仕上げる"}},
		{"チキンを仕上げる", 2, []string{"チキンを外す"}},
		{"チキンを外す", 1, []string{"人参セロリピーナッツに火を入れる"}},
		{"人参セロリピーナッツに火を入れる", 4, []string{"チキンを戻す"}},
		{"チキンを戻す", 1, []string{"クッキングソースをかける"}},
		{"クッキングソースをかける", 1, []string{"ソースが濃くなるまで火を入れる"}},
		{"ソースが濃くなるまで火を入れる", 3, []string{"料理を火から外す"}},
		{"料理を火から外す", 1, []string{}},
	}
	nodes := make([]*Node, len(names))

	for i, n := range names {
		nodes[i] = NewNodeWithWeight(n.Name, n.Weight)
	}

	for i, n := range names {
		for _, nn := range n.Nears {
			for _, node := range nodes {
				if node.Name == nn {
					nodes[i].AddNear(node)
				}
			}
		}
	}

	shuffle(nodes)

	re := criticalPath(nodes)
	pp.Println(re.PathList())

	if re.Weight != 39 {
		t.Fail()
	}
}
