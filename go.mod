module github.com/huantingwei/go

replace github.com/huantingwei/go => ./

replace github.com/huantingwei/go/gweb => ./gweb

go 1.15

require (
	github.com/huantingwei/go/gweb v0.0.0-00010101000000-000000000000
	go.mongodb.org/mongo-driver v1.4.1
)
