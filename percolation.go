package percolation

import (
  "fmt"
  "github.com/iloire/quickunionweighted"
)

// Percolation matrix (i,j)

// | [0,0] [0,1] [0,2] [0,3] |
// | [1,0] [1,1] [1,2] [1,3] |
// | [2,0] [2,1] [2,2] [2,3] |
// | [3,0] [3,1] [3,2] [3,3] |
// 

// Quick Union Set:
// 0 1 2 3  
// 4 5 6 7
// 8 9 10 11
// 12 13 14 15
//  
// 16 // top virtual cell (N*N)
// 17 // bottom virtual cell (N*N+1)
// 

type Percolation struct {
  arr [][] bool
  unionSet quickunionweighted.QuickUnion
  sizeN uint  
}

type Cell struct {
  I, J uint
}

func(p *Percolation) CellCoordenatesToInt (i,j uint) uint { // convert Cell to unique int
  return j + i * p.sizeN
}

func(p *Percolation) Initialize(N uint) {
  p.arr = make([][]bool, N)
  for j := range p.arr {
    p.arr[j] = make([]bool, N)
  } 
  p.unionSet.Initialise(N * N + 2) // two extras for top and bottom
  p.sizeN = N
}

func(p *Percolation) Open(c Cell) {
  
  if p.arr[c.I][c.J] {
    return // already open
  }
  
  p.arr[c.I][c.J] = true

  if (c.I == 0) {
    // join with top virtual cell
    p.unionSet.Union(p.CellCoordenatesToInt(c.I, c.J), (p.sizeN * p.sizeN))
  }

  if (c.I == p.sizeN - 1) {
    // join with bottom virtual cell
    p.unionSet.Union(p.CellCoordenatesToInt(c.I, c.J), (p.sizeN * p.sizeN + 1))
  }

  // unite consecutive cells
  if (c.J < p.sizeN - 1) && p.arr[c.I][c.J + 1] { // right cell
    p.unionSet.Union(p.CellCoordenatesToInt(c.I, c.J), p.CellCoordenatesToInt(c.I, c.J + 1))
  }
  if (c.J > 0) && p.arr[c.I][c.J - 1] { // left cell
    p.unionSet.Union(p.CellCoordenatesToInt(c.I, c.J), p.CellCoordenatesToInt(c.I, c.J - 1))
  }

  if (c.I > 0) && p.arr[c.I - 1][c.J] { // upper cell
    p.unionSet.Union(p.CellCoordenatesToInt(c.I, c.J), p.CellCoordenatesToInt(c.I - 1, c.J))
  }

  if (c.I < p.sizeN - 1)  && p.arr[c.I + 1][c.J] { // lower cell
    p.unionSet.Union(p.CellCoordenatesToInt(c.I, c.J), p.CellCoordenatesToInt(c.I + 1 , c.J))
  }
}

func(p *Percolation) IsOpen(c Cell) bool {
  return p.arr[c.I][c.J]
}

func(p *Percolation) IsConnected(c1, c2 Cell) bool {
  return p.unionSet.IsConnected(p.CellCoordenatesToInt(c1.I, c1.J), p.CellCoordenatesToInt(c2.I, c2.J))
}

func(p *Percolation) Percolates()  bool {
  return p.unionSet.IsConnected(p.sizeN * p.sizeN, p.sizeN * p.sizeN + 1)
}

func(p *Percolation) Print() {
  for i := range p.arr {
    fmt.Printf("\n%v", p.arr[i])
  } 
  fmt.Printf("\n%v", p.unionSet)
}
