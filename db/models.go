package db

type Pointer struct {
	PointerX float64 `json:"x"`
	PointerY float64 `json:"y"`
}

type Post struct {
	Room  int     `json:"room"`
	UUID  string  `json:"name"`
	Pos   Pointer `json:"pos"`
	Score int     `json:"score"`
}

type SendData struct {
	Room     int                `json:"room"`
	Question []int              `json:"question"`
	Score    map[string]int     `json:"score"`
	Pos      map[string]Pointer `json:"pos"`
}

func InitSendData() *SendData {
	return &SendData{
		Question: make([]int, 10),
		Score:    make(map[string]int),
		Pos:      make(map[string]Pointer),
	}
}
