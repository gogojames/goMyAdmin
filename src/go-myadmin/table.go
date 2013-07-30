package gomyadmin

type Table struct {
	Rows map[int]Row
}

func (t Table) GetRows(i int) Row {
	return t.Rows[i]
}
