package main

import "fmt"

// Solve solves the puzzle
// todo: assertion?
func Solve(puzzle Puzzle) Solution {
	sol := NewSolution(puzzle.RowSize, puzzle.ColSize)
	for iter := 1; ; iter++ {
		fmt.Printf("The %dth Row iteration:\n", iter)
		refresh := false
		for r := 0; r < puzzle.RowSize; r++ {
			row := sol.GetRow(r)
			filledRow := fillLine(puzzle.RowClues[r], row)
			if !row.EqualTo(filledRow) {
				refresh = true
				sol.SetRow(r, filledRow)
			}
		}
		PrintSolution(sol)
		fmt.Printf("The %dth Col iteration:\n", iter)
		for c := 0; c < puzzle.ColSize; c++ {
			col := sol.GetCol(c)
			filledCol := fillLine(puzzle.ColClues[c], col)
			if !col.EqualTo(filledCol) {
				refresh = true
				sol.SetCol(c, filledCol)
			}
		}
		PrintSolution(sol)
		if !refresh {
			break
		}
	}
	return sol
}

// fillLine returns a new Line that all definite grids are filled
// line would not be changed after calling.
func fillLine(clue Clue, line Line) Line {
	var ret Line
	feasible := enumerate(clue, line)
	if len(feasible) > 0 {
		ret = feasible[0]
		for i := 1; i < len(feasible); i++ {
			ret.Union(feasible[i])
		}
	}
	return ret
}

// enumerate returns all feasible solutions given clue and current line
// line would not be changed after calling.
func enumerate(clue Clue, line Line) []Line {
	ret := make([]Line, 0)
	try(clue, line, 0, 0, &ret)
	return ret
}

// try fills segment recursively and append a feasible solution to ret
// line would not be changed after calling.
func try(clue Clue, line Line, clueIdx int, lineIdx int, ret *[]Line) {
	if lineIdx >= len(line) && clueIdx < len(clue) {
		return
	}
	if clueIdx == len(clue) {
		filled := fillBlank(line, lineIdx, len(line))
		if isValid(clue, line) {
			*ret = append(*ret, line.Copy())
		}
		resetLine(line, filled)
		return
	}
	blanks := make([]int, 0)
	requiredLength := requiredLength(clue, clueIdx)
	for i := lineIdx; i < len(line); i++ {
		if requiredLength > len(line)-i {
			break
		}
		if isFillable(line, i, clue[clueIdx]) {
			segment := clue[clueIdx]
			filled := fillSegment(line, segment, i)
			try(clue, line, clueIdx+1, i+segment+1, ret)
			resetLine(line, filled)
		}
		blanks = append(blanks, fillBlank(line, i, i+1)...)
	}
	resetLine(line, blanks)
}

// isValid returns true if line is a feasible solution against clue
func isValid(clue Clue, line Line) bool {
	clueIdx := 0
	for i := 0; i < len(line); {
		switch line[i] {
		case GridTypeEnum.Colored:
			j := i
			for ; j < len(line); j++ {
				if line[j] != GridTypeEnum.Colored {
					break
				}
			}
			if clueIdx < len(clue) && j-i == clue[clueIdx] {
				clueIdx += 1
			} else {
				return false
			}
			i = j
		case GridTypeEnum.Blank:
			i++
		case GridTypeEnum.Unknown:
			return false
		}
	}
	if clueIdx != len(clue) {
		return false
	}
	return true
}

// requiredLength return the minimum length to fill all segments that in clue[clueIdx:]
func requiredLength(clue Clue, clueIdx int) int {
	ret := 0
	for i := clueIdx; i < len(clue); i++ {
		ret += clue[i]
	}
	ret += len(clue) - clueIdx - 1
	return ret
}

// isFillable returns true if line can be filled segment from lineIdx
func isFillable(line Line, lineIdx int, segment int) bool {
	if lineIdx+segment > len(line) {
		return false
	}
	for i := lineIdx; i < lineIdx+segment; i++ {
		if line[i] == GridTypeEnum.Blank {
			return false
		}
	}
	if lineIdx+segment < len(line) && line[lineIdx+segment] == GridTypeEnum.Colored {
		return false
	}
	return true
}

// fillSegment fills a segment (marks as Colored) to line from beginIdx.
// If it's not the end of line, it will append a Blank
func fillSegment(line Line, segment int, lineIdx int) []int {
	filled := make([]int, 0)
	for i := lineIdx; i < lineIdx+segment; i++ {
		if line[i] == GridTypeEnum.Unknown {
			filled = append(filled, i)
			line[i] = GridTypeEnum.Colored
		}
	}
	if lineIdx+segment < len(line) && line[lineIdx+segment] == GridTypeEnum.Unknown {
		filled = append(filled, lineIdx+segment)
		line[lineIdx+segment] = GridTypeEnum.Blank
	}
	return filled
}

// fillBlank fills grids of line to Blank from beginIdx to endIdx.
func fillBlank(line Line, beginIdx int, endIdx int) []int {
	filled := make([]int, 0)
	for i := beginIdx; i < MinInt(len(line), endIdx); i++ {
		if line[i] == GridTypeEnum.Unknown {
			filled = append(filled, i)
			line[i] = GridTypeEnum.Blank
		}
	}
	return filled
}

// resetLine resets all grids whose index is in filled to Unknown
func resetLine(line Line, filled []int) {
	for _, idx := range filled {
		line[idx] = GridTypeEnum.Unknown
	}
}
