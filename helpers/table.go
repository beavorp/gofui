package helpers

import (
	"syscall/js"
)

type Column[T any] struct {
	Key      string
	Header   *js.Value
	Ceil     func(data T) *js.Value
	FilterFn func(data T) bool
}

type TableConfig[T any] struct {
	Columns []Column[T]
}

type Table[T any] struct {
	data          []T
	displayedData []T

	config   TableConfig[T]
	onChange func()
}

func NewTable[T any](config TableConfig[T]) *Table[T] {
	return &Table[T]{
		data:     make([]T, 0),
		onChange: func() {},
		config:   config,
	}
}

func (t *Table[T]) SetColumns(cols []Column[T]) {
	t.config.Columns = cols
	t.computeDisplayedData()
	t.onChange()
}

func (t *Table[T]) SetData(data []T) {
	t.data = data
	t.computeDisplayedData()
	t.onChange()
}

func (t *Table[T]) computeDisplayedData() {
	// Start by counting the number of filter funcs
	funcs := make([]func(data T) bool, 0, len(t.config.Columns))
	for _, col := range t.config.Columns {
		if col.FilterFn != nil {
			funcs = append(funcs, col.FilterFn)
		}
	}

	t.displayedData = t.data
	newData := make([]T, 0, len(t.data))

	for _, fn := range funcs {
		for _, data := range t.displayedData {
			// If the filter func returns false, we skip the data
			if !fn(data) {
				continue
			}

			newData = append(newData, data)
		}

		t.displayedData = newData
	}

	// By the end of the loop, t.displayedData will contain only the data that passed all the filter funcs
}

func (t *Table[T]) OnChange(fn func()) {
	t.onChange = fn
}

func (t *Table[T]) GetHeaders() []*js.Value {
	headers := make([]*js.Value, 0, len(t.config.Columns))
	for _, col := range t.config.Columns {
		headers = append(headers, col.Header)
	}

	return headers
}

func (t *Table[T]) GetRows() []T {
	rows := make([]T, 0, len(t.displayedData))
	for _, data := range t.displayedData {
		rows = append(rows, data)
	}

	return rows
}

func (t *Table[T]) GetCeils(data T) []*js.Value {
	ceils := make([]*js.Value, 0, len(t.config.Columns))
	for _, col := range t.config.Columns {
		if col.Ceil == nil {
			continue
		}
		ceils = append(ceils, col.Ceil(data))
	}

	return ceils
}
