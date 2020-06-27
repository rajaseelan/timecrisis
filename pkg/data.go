package pkg

// TimeStruct how output is marshalled into JSON
type TimeStruct struct {
	String string `json:"timestring"`
	Unix   int64  `json:"unix"`
	Host   string `json:"host"`
}
