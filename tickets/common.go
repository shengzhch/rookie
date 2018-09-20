package tickets

type GoTickets interface {
	Take()
	Return()
	Active() bool
	Total() uint32
	Remindeer() uint32
}

type myGotickets struct {
	total    uint32
	tickerCh chan struct{}
	active   bool
}

func (gt *myGotickets) init(total uint32) bool {
	if gt.active {
		return false
	}
	if total == 0 {
		return false
	}

	ch := make(chan struct{}, total)
	n := int(total)

	for i := 0; i < n; i++ {
		ch <- struct{}{}
	}
	gt.tickerCh = ch
	gt.total = total
	gt.active = true
	return true
}
