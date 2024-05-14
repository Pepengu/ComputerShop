package shop

import (
	"container/list"
	"errors"
	"math"
	"strconv"
	"time"
)

type Table struct {
	Is_busy    bool
	busy_since time.Time
	Used_for   time.Time
	Earnings   int
}

func (t *Table) String() string {
	return strconv.Itoa(t.Earnings) + " " + t.Used_for.Format("15:04")
}

type Shop struct {
	Tables_amount int
	Schedule      struct {
		Open  time.Time
		Close time.Time
	}
	Cost    int
	Tables  []Table
	Clients map[string]int
	Queue   *list.List
}

func New(ta int, o time.Time, cl time.Time, co int) Shop {
	return Shop{
		ta,
		struct {
			Open  time.Time
			Close time.Time
		}{
			o,
			cl,
		},
		co,
		make([]Table, ta),
		make(map[string]int),
		list.New(),
	}
}

func (s *Shop) Sit(t time.Time, client string, table int) {
	s.Tables[table].busy_since = t
	s.Tables[table].Is_busy = true
	s.Clients[client] = table
}

func (s *Shop) Unsit(t time.Time, client string) (int, error) {
	table, ok := s.Clients[client]
	if !ok || table == -1 {
		return -1, errors.New("ClientUnknown")
	}
	s.Tables[table].Is_busy = false
	busy_for := t.Sub(s.Tables[table].busy_since)
	s.Tables[table].Used_for = s.Tables[table].Used_for.Add(busy_for)
	s.Tables[table].Earnings += int(math.Ceil(busy_for.Hours())) * s.Cost
	s.Clients[client] = -1

	return table, nil
}
