package main

import (
	"testing"
	"time"
)

func TestAccounting_TotalAmount(t *testing.T) {
	t.Run("oneMonth", TotalAmount_OneMonth)
	t.Run("oneDay", TotalAmount_OneDay)
	t.Run("crossMonth", TotalAmount_CrossMonth)
	t.Run("crossFourMonth", TotalAmount_CrossFourMonth)
	t.Run("crossYear", TotalAmount_CrossYear)
}

func TotalAmount_OneMonth(t *testing.T) {
	db := &DB{
		Budget: []Budget{
			{
				Date:   "202101",
				Amount: 3100,
			},
		},
	}
	acc := &Accounting{}
	acc.DB = db
	start, _ := time.Parse("2006-01-02", "2021-01-01")
	end, _ := time.Parse("2006-01-02", "2021-01-30")
	amount := acc.TotalAmount(start, end)
	if amount != 3000 {
		t.Errorf("expect: %v, res: %v", 3000, amount)
	} else {
		t.Logf("Success")
	}
}

func TotalAmount_OneDay(t *testing.T) {
	db := &DB{
		Budget: []Budget{
			{
				Date:   "202101",
				Amount: 3100,
			},
		},
	}
	acc := &Accounting{}
	acc.DB = db
	start, _ := time.Parse("2006-01-02", "2021-01-01")
	end, _ := time.Parse("2006-01-02", "2021-01-01")
	amount := acc.TotalAmount(start, end)
	if amount != 100 {
		t.Errorf("expect: %v, res: %v", 100, amount)
	} else {
		t.Logf("Success")
	}
}

func TotalAmount_CrossMonth(t *testing.T) {
	db := &DB{
		Budget: []Budget{
			{
				Date:   "202101",
				Amount: 3100,
			},
			{
				Date:   "202102",
				Amount: 2800,
			},
		},
	}
	acc := &Accounting{}
	acc.DB = db
	start, _ := time.Parse("2006-01-02", "2021-01-31")
	end, _ := time.Parse("2006-01-02", "2021-02-02")
	amount := acc.TotalAmount(start, end)
	if amount != 300 {
		t.Errorf("expect: %v, res: %v", 300, amount)
	} else {
		t.Logf("Success")
	}
}

func TotalAmount_CrossFourMonth(t *testing.T) {
	db := &DB{
		Budget: []Budget{
			{
				Date:   "202101",
				Amount: 31,
			},
			{
				Date:   "202102",
				Amount: 280,
			},
			{
				Date:   "202103",
				Amount: 3100,
			},
			{
				Date:   "202104",
				Amount: 30000,
			},
		},
	}
	acc := &Accounting{}
	acc.DB = db
	start, _ := time.Parse("2006-01-02", "2021-01-01")
	end, _ := time.Parse("2006-01-02", "2021-04-30")
	amount := acc.TotalAmount(start, end)
	if amount != 33411 {
		t.Errorf("expect: %v, res: %v", 33411, amount)
	} else {
		t.Logf("Success")
	}
}

func TotalAmount_CrossYear(t *testing.T) {
	db := &DB{
		Budget: []Budget{
			{
				Date:   "202012",
				Amount: 31,
			},
			{
				Date:   "202101",
				Amount: 310,
			},
			{
				Date:   "202102",
				Amount: 2800,
			},
			{
				Date:   "202103",
				Amount: 31000,
			},
		},
	}
	acc := &Accounting{}
	acc.DB = db
	start, _ := time.Parse("2006-01-02", "2020-09-01")
	end, _ := time.Parse("2006-01-02", "2021-05-01")
	amount := acc.TotalAmount(start, end)
	if amount != 34141 {
		t.Errorf("expect: %v, res: %v", 34141, amount)
	} else {
		t.Logf("Success")
	}
}
