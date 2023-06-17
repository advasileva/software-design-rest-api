package students

type student struct {
	Id         int64
	Name       string `pg:"name,notnull"`
	Age        int64  `pg:"age"`
	Profession string `pg:"profession,"`
}

type grade struct {
	Id        int64
	StudentId int64  `pg:"student_id"`
	Subject   string `pg:"subject"`
	Grade     int64  `pg:"grade"`
}
