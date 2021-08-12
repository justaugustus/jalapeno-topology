module github.com/jalapeno/topology

go 1.15

replace (
	github.com/sbezverk/gobmp => /Users/brucemcdougall/go/jalapeno-gobmp/gobmp
)

require (
	github.com/Shopify/sarama v1.27.0
	github.com/arangodb/go-driver v0.0.0-20200831144536-17278d36b7e8
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/sbezverk/gobmp v0.0.1-beta.0.20210708163419-62d7fc1982bd
	go.uber.org/atomic v1.7.0
)
