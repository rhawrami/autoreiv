package cols

import (
	"runtime"
)

func ApplyNewCol(c *Int64Col, fn func(int) int) Int64Col {
	return Int64Col{}
}

func MkChunks(n, maxWorkers int) []Task {
	return []Task{}
}

type Task struct {
	start int
	stop  int
}

func ChunkSizeAndNWorkers(n, maxWorkers int) (int, int) {
	var nWorkers, chunkSize int
	nWorkers = runtime.NumCPU()
	// use MaxWorkers if possible
	nWorkers = min(nWorkers, maxWorkers)
	// equally distribute tasks
	remainder := 0
	if n%nWorkers != 0 {
		remainder = 1
	}
	chunkSize = (n / nWorkers) + remainder

	// if chunkSize < 1 kiB, just use 1 worker
	bytesInChunk := 64 * chunkSize
	kiB := 1 << 10
	if bytesInChunk < kiB {
		nWorkers = 1
	}
	return chunkSize, nWorkers
}

// TODO - OFF BY ONE I BELIEVE
func UnrollLoop8Vals(inSlc []int64, outSlc []int64, nullSlc []bool, t Task, fn func(int64) int64) {
	width := t.start - t.stop + 1

	// wrap fn with null short-circuit
	fnWrapped := func(x int64, i int) int64 {
		if nullSlc[i] {
			return 0
		}
		return fn(x)
	}

	// if task size < 8, regular loop
	if width < 8 {
		on := t.start
		for on <= t.stop {
			outSlc[on] = fnWrapped(inSlc[on], on)
			on += 1
		}
		return
	}

	// loop apply
	ctr := t.start
	completed := 0
	for ctr <= t.stop {
		outSlc[ctr] = fnWrapped(inSlc[ctr], ctr)
		outSlc[ctr+1] = fnWrapped(inSlc[ctr+1], ctr+1)
		outSlc[ctr+2] = fnWrapped(inSlc[ctr+2], ctr+2)
		outSlc[ctr+3] = fnWrapped(inSlc[ctr+3], ctr+3)
		outSlc[ctr+4] = fnWrapped(inSlc[ctr+4], ctr+4)
		outSlc[ctr+5] = fnWrapped(inSlc[ctr+5], ctr+5)
		outSlc[ctr+6] = fnWrapped(inSlc[ctr+6], ctr+6)
		outSlc[ctr+7] = fnWrapped(inSlc[ctr+7], ctr+7)
		ctr += 8
		completed += 8
	}

	// remainder
	toDo := width - completed
	if toDo != 0 {
		on := t.stop - toDo + 1
		for on < t.stop {
			outSlc[on] = fnWrapped(inSlc[on], on)
			on += 1
		}
	}
}
