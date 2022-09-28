package main

import (
	"fmt"
	"math"
)

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

type Edge struct {
	to, cap, flow, rev int
}

type Graph struct {
	V     int
	E     [][]Edge
	level []int
}

func newGraph(V int) *Graph {
	return &Graph{V: V, E: make([][]Edge, V), level: make([]int, V)}
}

func (g *Graph) addEdge(from, to, cap int) {
	g.E[from] = append(g.E[from], Edge{to, cap, 0, len(g.E[to])})
	g.E[to] = append(g.E[to], Edge{from, 0, 0, len(g.E[from]) - 1})
}

func (g *Graph) bfs(s int, t int) bool {
	for i := range g.level {
		g.level[i] = -1
	}

	var queue []int = []int{s}
	g.level[s] = 0

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for _, e := range g.E[u] {
			if e.flow < e.cap && g.level[e.to] < 0 {
				g.level[e.to] = g.level[u] + 1
				queue = append(queue, e.to)
			}
		}
	}
	return g.level[t] >= 0
}

func (g *Graph) sendFlow(u int, t int, flow int, next []int) int {
	if u == t {
		return flow
	}

	for next[u] < len(g.E[u]) {
		e := &g.E[u][next[u]]

		if g.level[e.to] == g.level[u]+1 && e.flow < e.cap {
			tmp_flow := g.sendFlow(e.to, t, Min(flow, e.cap-e.flow), next)
			if tmp_flow > 0 {
				e.flow += tmp_flow
				g.E[e.to][e.rev].flow -= tmp_flow
				return tmp_flow
			}
		}
		next[u] += 1

	}

	return 0
}

func (g *Graph) DinicMaxFlow(s, t int) int {
	if s == t {
		return -1
	}

	res := 0
	for g.bfs(s, t) {
		next := make([]int, g.V)
		for flow := g.sendFlow(s, t, math.MaxInt, next); flow > 0; {
			res += flow
			flow = g.sendFlow(s, t, math.MaxInt, next)
		}
	}
	return res
}

func main() {
	var V, E, start, finish int
	fmt.Scanf("%d %d %d %d", &V, &E, &start, &finish)

	g := newGraph(V)
	for i := 0; i < E; i++ {
		var u, v, cap int
		fmt.Scanf("%d %d %d", &u, &v, &cap)
		g.addEdge(u, v, cap)
	}

	maxFlow := g.DinicMaxFlow(start, finish)
	fmt.Println(maxFlow)
}
