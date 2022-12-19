module github.com/myl7/rapidchain

go 1.19

require (
	github.com/jinzhu/copier v0.3.5
	github.com/klauspost/reedsolomon v1.11.3
	github.com/wealdtech/go-merkletree v1.0.0
)

require (
	github.com/klauspost/cpuid/v2 v2.1.1 // indirect
	github.com/stretchr/testify v1.8.1 // indirect
	golang.org/x/crypto v0.4.0 // indirect
	golang.org/x/sys v0.3.0 // indirect
)

replace github.com/wealdtech/go-merkletree => github.com/kimborgen/go-merkletree v1.0.1-0.20200621183708-d72b0ecf9ebd
