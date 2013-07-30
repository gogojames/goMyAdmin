package gomyadmin

type Field struct {
	Value map[string]Class
}

func (f Field) GetValue(fname string) Class {
	return f.Value[fname]
}
