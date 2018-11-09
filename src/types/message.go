package model

type Message struct {
    Timestamp time.Time
    Author    string
    Message   string
}

func (this Message) MarshallJSON(){
	
}
