module github.com/martinkunc/gowasmer

go 1.13

require (
	github.com/martinkunc/wagather v0.0.0-20200727110959-e905a78a3d7d
	github.com/wasmerio/go-ext-wasm v0.3.1
)

replace (
	golang.org/x/crypto => golang.org/x/crypto v0.0.0-20190820162420-60c769a6c586
	golang.org/x/net => golang.org/x/net v0.0.0-20191004110552-13f9640d40b9
	golang.org/x/sys => golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a // pinned to release-branch.go1.13
	golang.org/x/tools => golang.org/x/tools v0.0.0-20190821162956-65e3620a7ae7 // pinned to release-branch.go1.13
)
