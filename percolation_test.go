package percolation

import(
  "github.com/stretchr/testify/assert"
  "github.com/iloire/percolation"
  "testing"
)

func TestCellCoordenatesMapper(t *testing.T) { 
  p:= new(percolation.Percolation)
  p.Initialize(5)
  assert.Equal(t, 0, p.CellCoordenatesToInt(0,0))
  assert.Equal(t, 4, p.CellCoordenatesToInt(0,4))
  assert.Equal(t, 5, p.CellCoordenatesToInt(1,0))
  assert.Equal(t, 10, p.CellCoordenatesToInt(2,0))
}

func TestOnceCellShouldPercolate(t *testing.T) { 
  p:= new(percolation.Percolation)
  p.Initialize(1)
  assert.False(t, p.Percolates(), "1 single closed cell matrix should not percolate") 
}

func TestTwoDimMatrixShouldNotPercolate(t *testing.T) { 
  p:= new(percolation.Percolation)
  p.Initialize(2)
  assert.False(t, p.Percolates(), "2 single closed cell matrix should not percolate") 
}

func TestPercolationPercolatesWithFullColumnOpen(t *testing.T) { 
  p:= new(percolation.Percolation)
  p.Initialize(5)
  p.Open(percolation.Cell{0,0})
  p.Open(percolation.Cell{1,0})
  p.Open(percolation.Cell{2,0})
  p.Open(percolation.Cell{3,0})
  p.Open(percolation.Cell{4,0})

  assert.True(t, p.Percolates()) 
}

func TestPercolationUnionCellOpenToTheRight(t *testing.T) { 
  p:= new(percolation.Percolation)
  p.Initialize(5)
  c1:=percolation.Cell{0,1}
  c2:=percolation.Cell{0,2}
  p.Open(c1)
  p.Open(c2)

  assert.True(t, p.IsConnected(c1, c2)) 
}

func TestPercolationUnionCellOpenToTheRightBorder(t *testing.T) { 
  p:= new(percolation.Percolation)
  p.Initialize(5)
  c1:=percolation.Cell{0,3}
  c2:=percolation.Cell{0,4}
  p.Open(c1)
  p.Open(c2)

  assert.True(t, p.IsConnected(c1, c2)) 
}

func TestPercolationUnionCellOpenToTheLeft(t *testing.T) { 
  p:= new(percolation.Percolation)
  p.Initialize(5)
  c1:=percolation.Cell{0,2}
  c2:=percolation.Cell{0,1}
  p.Open(c1)
  p.Open(c2)

  assert.True(t, p.IsConnected(c1, c2)) 
}

func TestPercolationUnionCellOpenToTheLeftBorder(t *testing.T) { 
  p:= new(percolation.Percolation)
  p.Initialize(5)
  c1:=percolation.Cell{0,1}
  c2:=percolation.Cell{0,0}
  p.Open(c1)
  p.Open(c2)

  assert.True(t, p.IsConnected(c1, c2)) 
}

func TestPercolationUnionCellOpenTop(t *testing.T) { 
  p:= new(percolation.Percolation)
  p.Initialize(5)
  c1:=percolation.Cell{2,0}
  c2:=percolation.Cell{1,0}
  p.Open(c1)
  p.Open(c2)

  assert.True(t, p.IsConnected(c1, c2)) 
}

func TestPercolationUnionCellOpenTopBorder(t *testing.T) { 
  p:= new(percolation.Percolation)
  p.Initialize(5)
  c1:=percolation.Cell{1,0}
  c2:=percolation.Cell{0,0}
  p.Open(c1)
  p.Open(c2)

  assert.True(t, p.IsConnected(c1, c2)) 
}

func TestPercolationUnionCellOpenBottom(t *testing.T) { 
  p:= new(percolation.Percolation)
  p.Initialize(5)
  c1:=percolation.Cell{0,0}
  c2:=percolation.Cell{1,0}
  p.Open(c1)
  p.Open(c2)

  assert.True(t, p.IsConnected(c1, c2)) 
}

func TestPercolationUnionCellOpenBottomBorder(t *testing.T) { 
  p:= new(percolation.Percolation)
  p.Initialize(5)
  c1:=percolation.Cell{3,0}
  c2:=percolation.Cell{4,0}
  p.Open(c1)
  p.Open(c2)

  assert.True(t, p.IsConnected(c1, c2)) 
}
