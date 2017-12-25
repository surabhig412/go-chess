package main

import (
	"errors"
	"fmt"
	"unsafe"
)

type SPvEntry struct {
	PosKey uint64
	Move   int
}

type SPvTable struct {
	PTable     []*SPvEntry
	NumEntries int
}

func (pvt *SPvTable) Clear() {
	for i := 0; i < pvt.NumEntries; i++ {
		pvEntry := new(SPvEntry)
		pvEntry.PosKey = 0
		pvEntry.Move = NoMove
		pvt.PTable = append(pvt.PTable, pvEntry)
	}
}

func (pvt *SPvTable) Init() {
	var pvEntry SPvEntry
	pvt.NumEntries = PvSize/int(unsafe.Sizeof(pvEntry)) - 2
	pvt.PTable = nil
	pvt.Clear()
}

func GetPvLine(depth int, pos *SBoard) (int, error) {
	if depth >= MaxDepth {
		return 0, errors.New(fmt.Sprintf("Depth can't be more than %d", MaxDepth))
	}
	move, err := ProbePvTable(pos)
	if err != nil {
		return 0, err
	}
	count := 0
	for move != NoMove && count < depth {
		if count >= MaxDepth {
			return 0, errors.New(fmt.Sprintf("Depth can't be more than %d", MaxDepth))
		}
		if MoveExists(pos, move) {
			MakeMove(move, pos)
			pos.PvArray[count] = move
			count++
		} else {
			break
		}
		move, err = ProbePvTable(pos)
		if err != nil {
			return 0, err
		}

	}
	for pos.Ply > 0 {
		TakeMove(pos)
	}
	return count, nil
}

// StorePvMove stores the move in the PV table of the board
func StorePvMove(pos *SBoard, move int) error {
	index := pos.PosKey % uint64(pos.PvTable.NumEntries)
	if index < 0 || index > uint64(pos.PvTable.NumEntries-1) {
		return errors.New("Index where PV move is to be stored is invalid")
	}
	pos.PvTable.PTable[index].PosKey = pos.PosKey
	pos.PvTable.PTable[index].Move = move
	return nil
}

// ProbePvTable searches for the most significant move on the board based on its pv
func ProbePvTable(pos *SBoard) (int, error) {
	index := pos.PosKey % uint64(pos.PvTable.NumEntries)
	if index < 0 || index > uint64(pos.PvTable.NumEntries-1) {
		return NoMove, errors.New("Index where PV move is to be probed from is invalid")
	}
	if pos.PvTable.PTable[index].PosKey == pos.PosKey {
		return pos.PvTable.PTable[index].Move, nil
	}
	return NoMove, nil
}
