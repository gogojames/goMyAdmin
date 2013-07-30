package gomyadmin

type Row struct {
	Fields map[string]Field
}

func (r Row) GetFields(fname string) Field {
	return r.Fields[fname]
}

func (r Row) SetField() {

}
