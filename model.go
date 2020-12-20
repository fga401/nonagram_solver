package main

type Clue []int

type Puzzle struct {
	RowSize  int
	ColSize  int
	RowClues []Clue
	ColClues []Clue
}

type GridType = int8

var GridTypeEnum = struct {
	Blank   GridType
	Colored GridType
	Unknown GridType
}{-1, 1, 0}

type Solution struct {
	RowSize int
	ColSize int
	grids   [][]GridType
}

type Line []GridType

func NewSolution(rowSize, colSize int) Solution {
	ret := Solution{
		RowSize: rowSize,
		ColSize: colSize,
		grids:   make([][]GridType, rowSize),
	}
	for r := 0; r < rowSize; r++ {
		row := make([]GridType, colSize, colSize)
		ret.grids[r] = row
	}
	return ret
}

func (s *Solution) Copy() Solution {
	ret := Solution{
		RowSize: s.RowSize,
		ColSize: s.ColSize,
		grids:   make([][]GridType, s.RowSize),
	}
	for r := 0; r < s.RowSize; r++ {
		row := make([]GridType, s.ColSize, s.ColSize)
		copy(row, s.grids[r])
		ret.grids[r] = row
	}
	return ret
}

func (s *Solution) GetGrid(rIdx, cIdx int) GridType {
	return s.grids[rIdx][cIdx]
}

func (s *Solution) SetGrid(rIdx, cIdx int, grid GridType) {
	s.grids[rIdx][cIdx] = grid
}

func (s *Solution) GetRow(idx int) Line {
	ret := make(Line, s.ColSize, s.ColSize)
	for c := 0; c < s.ColSize; c++ {
		ret[c] = s.grids[idx][c]
	}
	return ret
}

func (s *Solution) GetCol(idx int) Line {
	ret := make(Line, s.RowSize, s.RowSize)
	for r := 0; r < s.RowSize; r++ {
		ret[r] = s.grids[r][idx]
	}
	return ret
}

func (s *Solution) SetRow(idx int, row Line) {
	for c := 0; c < s.ColSize; c++ {
		s.grids[idx][c] = row[c]
	}
}

func (s *Solution) SetCol(idx int, col Line) {
	for r := 0; r < s.RowSize; r++ {
		s.grids[r][idx] = col[r]
	}
}

func (l *Line) Union(other Line)  {
	for i := 0; i < len(*l); i++ {
		if (*l)[i] != other[i] {
			(*l)[i] = GridTypeEnum.Unknown
		}
	}
}

func (l *Line) Copy() Line {
	ret := make(Line, len(*l), len(*l))
	copy(ret, *l)
	return ret
}

func (l *Line) EqualTo(other Line) bool {
	for i := 0; i < len(*l); i++ {
		if (*l)[i] != other[i] {
			return false
		}
	}
	return true
}

