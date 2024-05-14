package actions

import (
	"strconv"
	"strings"
	"time"
)

const (
	IDIncomeArival  = 1
	IDIncomeSited   = 2
	IDIncomeWaiting = 3
	IDIncomeLeft    = 4

	IDOutcomeLeft   = 11
	IDOutcomeSitted = 12
	IDOutcomeError  = 13
)

type Action interface {
	GetID() int
	String() string
}

type ActionIncomeArival struct {
	Time time.Time
	Name string
}

func (a ActionIncomeArival) GetID() int {
	return IDIncomeArival
}

func (a ActionIncomeArival) String() string {
	return strings.Join([]string{a.Time.Format("15:04"), strconv.FormatInt(IDIncomeArival, 10), a.Name}, " ")
}

func NewActionIncomeArival(time time.Time, name string) ActionIncomeArival {
	var a ActionIncomeArival
	a.Time = time
	a.Name = name
	return a
}

type ActionIncomeSited struct {
	Time  time.Time
	Name  string
	Table int
}

func (a ActionIncomeSited) GetID() int {
	return IDIncomeSited
}

func (a ActionIncomeSited) String() string {
	return strings.Join([]string{a.Time.Format("15:04"),
		strconv.FormatInt(IDIncomeSited, 10), a.Name,
		strconv.FormatInt(int64(a.Table+1), 10)}, " ")
}

func NewActionIncomeSited(time time.Time, name string, table int) ActionIncomeSited {
	var a ActionIncomeSited
	a.Time = time
	a.Name = name
	a.Table = table
	return a
}

type ActionIncomeWaiting struct {
	Time time.Time
	Name string
}

func (a ActionIncomeWaiting) String() string {
	return strings.Join([]string{a.Time.Format("15:04"), strconv.FormatInt(IDIncomeWaiting, 10), a.Name}, " ")
}

func (a ActionIncomeWaiting) GetID() int {
	return IDIncomeWaiting
}

func NewActionIncomeWaiting(time time.Time, name string) ActionIncomeWaiting {
	var a ActionIncomeWaiting
	a.Time = time
	a.Name = name
	return a
}

type ActionIncomeLeft struct {
	Time time.Time
	Name string
}

func (a ActionIncomeLeft) GetID() int {
	return IDIncomeLeft
}

func (a ActionIncomeLeft) String() string {
	return strings.Join([]string{a.Time.Format("15:04"), strconv.FormatInt(IDIncomeLeft, 10), a.Name}, " ")
}

func NewActionIncomeLeft(time time.Time, name string) ActionIncomeLeft {
	var a ActionIncomeLeft
	a.Time = time
	a.Name = name
	return a
}

type ActionOutcomeLeft struct {
	Time time.Time
	Name string
}

func (a ActionOutcomeLeft) GetID() int {
	return IDOutcomeLeft
}

func (a ActionOutcomeLeft) String() string {
	return strings.Join([]string{a.Time.Format("15:04"), strconv.FormatInt(IDOutcomeLeft, 10), a.Name}, " ")
}

func NewActionOutcomeLeft(time time.Time, name string) ActionOutcomeLeft {
	var a ActionOutcomeLeft
	a.Time = time
	a.Name = name
	return a
}

type ActionOutcomeSitted struct {
	Time  time.Time
	Name  string
	Table int
}

func (a ActionOutcomeSitted) GetID() int {
	return IDOutcomeSitted
}

func (a ActionOutcomeSitted) String() string {
	return strings.Join([]string{a.Time.Format("15:04"),
		strconv.FormatInt(IDOutcomeSitted, 10), a.Name,
		strconv.FormatInt(int64(a.Table+1), 10)}, " ")
}

func NewActionOutcomeSitted(time time.Time, name string) ActionOutcomeSitted {
	var a ActionOutcomeSitted
	a.Time = time
	a.Name = name
	return a
}

type ActionOutcomeError struct {
	Time time.Time
	Err  string
}

func (a ActionOutcomeError) GetID() int {
	return IDOutcomeError
}

func (a ActionOutcomeError) String() string {
	return strings.Join([]string{a.Time.Format("15:04"), strconv.FormatInt(IDOutcomeError, 10), a.Err}, " ")
}

func NewActionOutcomeError(time time.Time, err string) ActionOutcomeError {
	var a ActionOutcomeError
	a.Time = time
	a.Err = err
	return a
}
