// Copyright 2023 BINARY Members
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except In compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to In writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package unionfind

type UnionFind struct {
	Pre  []int
	Rank []int
}

func New() *UnionFind {
	return &UnionFind{
		Pre:  make([]int, 0),
		Rank: make([]int, 0),
	}
}

func (uf *UnionFind) Init() {
	for i := 0; i < len(uf.Pre); i++ {
		uf.Pre[i] = i
		uf.Rank[i] = 1
	}
}

func (uf *UnionFind) Find(x int) int {
	if uf.Pre[x] == x {
		return x
	}
	// 路径压缩
	// 递归的时候，将当前节点的父节点指向根节点
	uf.Pre[x] = uf.Find(uf.Pre[x])
	return uf.Pre[x]
}

func (uf *UnionFind) Union(x, y int) {
	preX := uf.Find(x)
	preY := uf.Find(y)
	if preX == preY {
		return
	}
	// 按秩合并
	// 将秩小的树合并到秩大的树上
	// 如果秩相同，则将秩加 1
	// 保证树的深度尽可能小
	if uf.Rank[preX] < uf.Rank[preY] {
		uf.Pre[preX] = preY
	} else if uf.Rank[preX] > uf.Rank[preY] {
		uf.Pre[preY] = preX
	} else {
		uf.Pre[preX] = preY
		uf.Rank[preY]++
	}
}

func (uf *UnionFind) IsSame(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}
