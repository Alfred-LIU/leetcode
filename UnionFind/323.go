package UnionFind

func countComponents(n int, edges [][]int) int {
	uf := createUF(n)

	for _, edge := range edges {
		uf.union(edge[0], edge[1])
	}

	//clean
	for i := 0; i < n; i++ {
		uf.nodes[i] = uf.find(uf.nodes[i])
	}

	set := make(map[int]struct{})
	for _, node := range uf.nodes {
		set[node] = struct{}{}
	}

	return len(set)
}

type UF struct {
	nodes []int
	rank  []int
}

func createUF(n int) UF {
	nodes := make([]int, 0)
	rank := make([]int, 0)
	for i := 0; i < n; i++ {
		nodes = append(nodes, i)
		rank = append(rank, 1)
	}
	return UF{
		nodes: nodes,
		rank:  rank,
	}
}

func (uf *UF) union(i, j int) {
	fI := uf.find(i)
	fJ := uf.find(j)

	if uf.rank[fI] > uf.rank[fJ] {
		uf.nodes[fJ] = fI
		uf.rank[fI] += uf.rank[fJ]
	} else {
		uf.nodes[fI] = fJ
		uf.rank[fJ] += uf.rank[fI]
	}
}

func (uf *UF) find(i int) int {
	for uf.nodes[i] != i {
		i = uf.nodes[i]
	}
	return i
}
