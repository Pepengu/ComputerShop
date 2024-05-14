package computershop

import (
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/Pepengu/ComputerShop/computershop/actions"
)

func TestValidator(t *testing.T) {
	tests := []struct {
		n string
		s string
		b bool
	}{
		{"1", "08:48 1 client1", true},
		{"2", "38:48 1 client1", false},
		{"3", "23:68 1 client1", false},
		{"4", "23:48 a client1", false},
		{"5", "23:48 1 cli#nt1", false},
		{"6", "00:00 14 client1", true},
		{"7", "08:32 1 vasiliy", true},
	}
	for _, v := range tests {
		t.Run(v.n, func(t *testing.T) {
			s, b := v.s, v.b
			t.Parallel()
			res := validator.MatchString(s)
			if res != b {
				t.Errorf("String \"%s\" was matched incorectly. Expected: %t, got: %t", s, b, res)
			}
		})
	}
}

func TestParseActionIncomeArivial(t *testing.T) {
	t.Run("Correct", func(t *testing.T) {
		t.Parallel()
		validTime := "09:32"
		validID := strconv.Itoa(actions.IDIncomeArival)
		validName := "vasiliy"
		validString := strings.Join([]string{validTime, validID, validName}, " ")
		res, err := parseAction(validString)
		if err != nil {
			t.Errorf("Error \"%s\" occured when passing string \"%s\"", err.Error(), validString)
		}

		if res.GetID() != actions.IDIncomeArival {
			t.Errorf("Resieved wrong ID. Expected: %d, Actual: %d", actions.IDIncomeArival, res.GetID())
		}

		action := res.(actions.ActionIncomeArival)
		if action.Name != validName {
			t.Errorf("Resieved wrong name. Expected: %s, Actual: %s", action.Name, validName)
		}

		ti, _ := time.Parse("15:04", validTime)
		if action.Time.Compare(ti) != 0 {
			t.Errorf("Resieved wrong time. Expected: %s, Actual: %s", action.Time.Format("15:04"), validTime)
		}
	})
}

func TestParseActionIncomeLeft(t *testing.T) {
	t.Run("Correct", func(t *testing.T) {
		t.Parallel()
		validTime := "09:32"
		validID := strconv.Itoa(actions.IDIncomeLeft)
		validName := "vasiliy"
		validString := strings.Join([]string{validTime, validID, validName}, " ")
		res, err := parseAction(validString)
		if err != nil {
			t.Errorf("Error \"%s\" occured when passing string \"%s\"", err.Error(), validString)
		}

		if res.GetID() != actions.IDIncomeLeft {
			t.Errorf("Resieved wrong ID. Expected: %d, Actual: %d", actions.IDIncomeArival, res.GetID())
		}

		action := res.(actions.ActionIncomeLeft)
		if action.Name != validName {
			t.Errorf("Resieved wrong name. Expected: %s, Actual: %s", action.Name, validName)
		}

		ti, _ := time.Parse("15:04", validTime)
		if action.Time.Compare(ti) != 0 {
			t.Errorf("Resieved wrong time. Expected: %s, Actual: %s", action.Time.Format("15:04"), validTime)
		}
	})
}

func TestParseActionIncomeWaiting(t *testing.T) {
	t.Run("Correct", func(t *testing.T) {
		t.Parallel()
		validTime := "09:32"
		validID := strconv.Itoa(actions.IDIncomeWaiting)
		validName := "vasiliy"
		validString := strings.Join([]string{validTime, validID, validName}, " ")
		res, err := parseAction(validString)
		if err != nil {
			t.Errorf("Error \"%s\" occured when passing string \"%s\"", err.Error(), validString)
		}

		if res.GetID() != actions.IDIncomeWaiting {
			t.Errorf("Resieved wrong ID. Expected: %d, Actual: %d", actions.IDIncomeArival, res.GetID())
		}

		action := res.(actions.ActionIncomeWaiting)
		if action.Name != validName {
			t.Errorf("Resieved wrong name. Expected: %s, Actual: %s", action.Name, validName)
		}

		ti, _ := time.Parse("15:04", validTime)
		if action.Time.Compare(ti) != 0 {
			t.Errorf("Resieved wrong time. Expected: %s, Actual: %s", action.Time.Format("15:04"), validTime)
		}
	})
}

func TestParseActionIncomeSited(t *testing.T) {
	t.Run("Correct", func(t *testing.T) {
		t.Parallel()
		validTime := "09:32"
		validID := strconv.Itoa(actions.IDIncomeSited)
		validName := "vasiliy"
		validString := strings.Join([]string{validTime, validID, validName}, " ")
		res, err := parseAction(validString)
		if err != nil {
			t.Errorf("Error \"%s\" occured when passing string \"%s\"", err.Error(), validString)
		}

		if res.GetID() != actions.IDIncomeSited {
			t.Errorf("Resieved wrong ID. Expected: %d, Actual: %d", actions.IDIncomeArival, res.GetID())
		}

		action := res.(actions.ActionIncomeSited)
		if action.Name != validName {
			t.Errorf("Resieved wrong name. Expected: %s, Actual: %s", action.Name, validName)
		}

		ti, _ := time.Parse("15:04", validTime)
		if action.Time.Compare(ti) != 0 {
			t.Errorf("Resieved wrong time. Expected: %s, Actual: %s", action.Time.Format("15:04"), validTime)
		}
	})
}
