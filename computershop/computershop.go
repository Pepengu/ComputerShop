package computershop

import (
	"bufio"
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/Pepengu/ComputerShop/computershop/actions"
	"github.com/Pepengu/ComputerShop/computershop/shop"
)

type ComputerShop struct {
	shop.Shop
}

func (cs *ComputerShop) parseInitialInfo(scanner *bufio.Scanner) error {
	scanner.Scan()
	tables_amount, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return errors.New("Tables amount is invalid")
	}

	scanner.Scan()
	times := strings.Split(scanner.Text(), " ")
	if len(times) != 2 {
		return errors.New("Times line is invalid")
	}
	start, err := time.Parse("15:04", times[0])
	if err != nil {
		return errors.New("Opening time is invalid")
	}
	end, err := time.Parse("15:04", times[1])
	if err != nil {
		return errors.New("Closing time is invalid")
	}

	scanner.Scan()
	cost, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return errors.New("Closing time is invalid")
	}

	cs.Shop = shop.New(tables_amount, start, end, cost)
	return nil
}

func (cs *ComputerShop) Init(scanner *bufio.Scanner) error {
	return cs.parseInitialInfo(scanner)
}

var validator *regexp.Regexp = regexp.MustCompile("^([01][0-9]|2[0-3]):[0-5][0-9] [0-9]+ [a-z0-9_-]+( [0-9]*)?$")

func parseAction(line string) (actions.Action, error) {
	if !validator.MatchString(line) {
		return nil, errors.New("Action is invalid")
	}

	cmd := strings.Split(line, " ")
	time, _ := time.Parse("15:04", cmd[0])

	var result actions.Action
	result = actions.NewActionOutcomeError(time.UTC(), "UnknownError")
	cmdType, _ := strconv.Atoi(cmd[1])
	switch cmdType {
	case actions.IDIncomeArival:
		result = actions.NewActionIncomeArival(time, cmd[2])

	case actions.IDIncomeLeft:
		result = actions.NewActionIncomeLeft(time, cmd[2])

	case actions.IDIncomeWaiting:
		result = actions.NewActionIncomeWaiting(time, cmd[2])

	case actions.IDIncomeSited:
		if len(cmd) != 4 {
			break
		}
		table, _ := strconv.Atoi(cmd[3])
		result = actions.NewActionIncomeSited(time, cmd[2], table-1)
	}

	return result, nil
}

func (cs *ComputerShop) applyAction(action actions.Action) actions.Action {
	switch action.GetID() {
	case actions.IDIncomeArival:
		arivial := action.(actions.ActionIncomeArival)
		if arivial.Time.Before(cs.Schedule.Open) {
			return actions.NewActionOutcomeError(arivial.Time, "NotOpenYet")
		}

		if _, ok := cs.Clients[arivial.Name]; ok {
			return actions.NewActionOutcomeError(arivial.Time, "YouShallNotPass")
		}

		cs.Clients[arivial.Name] = -1

	case actions.IDIncomeSited:
		sited := action.(actions.ActionIncomeSited)
		if cs.Tables[sited.Table].Is_busy {
			return actions.NewActionOutcomeError(sited.Time, "PlaceIsBusy")
		}

		if _, ok := cs.Clients[sited.Name]; !ok {
			return actions.NewActionOutcomeError(sited.Time, "ClientUnknown")
		}

		cs.Unsit(sited.Time, sited.Name)
		cs.Sit(sited.Time, sited.Name, sited.Table)

		cur := cs.Queue.Front()
		for ; cur != nil && cur.Value != sited.Name; cur = cur.Next() {
		}
		if cur != nil {
			cs.Queue.Remove(cur)
		}

	case actions.IDIncomeWaiting:
		waiting := action.(actions.ActionIncomeWaiting)

		has_free := false
		for _, v := range cs.Tables {
			if !v.Is_busy {
				has_free = true
				break
			}
		}
		if has_free {
			return actions.NewActionOutcomeError(waiting.Time, "ICanWaitNoLonger!")
		}

		if cs.Queue.Len() >= cs.Tables_amount {
			return actions.NewActionOutcomeLeft(waiting.Time, waiting.Name)
		}

		cs.Queue.PushBack(waiting.Name)

	case actions.IDIncomeLeft:
		left := action.(actions.ActionIncomeLeft)

		table, err := cs.Unsit(left.Time, left.Name)
		if err != nil {
			return actions.NewActionOutcomeError(left.Time, err.Error())
		}

		delete(cs.Clients, left.Name)

		if cs.Queue.Front() != nil {
			name := cs.Queue.Front().Value.(string)
			cs.Queue.Remove(cs.Queue.Front())
			cs.Sit(left.Time, name, table)
			return actions.NewActionOutcomeSitted(left.Time, name)
		}
	}

	return nil
}

func (cs *ComputerShop) RunAction(line string) error {
	action, err := parseAction(line)
	if err != nil {
		return err
	}
	fmt.Println(action.String())

	action = cs.applyAction(action)
	if action != nil {
		fmt.Println(action.String())
	}

	return nil
}

func (cs *ComputerShop) Close() {
	clients := make([]string, len(cs.Clients))
	idx := 0
	for k := range cs.Clients {
		clients[idx] = k
		idx++
	}
	slices.Sort(clients)
	for _, client := range clients {
		cs.Unsit(cs.Schedule.Close, client)
		fmt.Println(actions.NewActionOutcomeLeft(cs.Schedule.Close, client).String())
	}

	for idx, table := range cs.Tables {
		fmt.Println(idx+1, table.String())
	}
}
