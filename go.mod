module zsp

go 1.13

require (
	clpars4 v0.0.0-00010101000000-000000000000 // indirect
	github.com/soniah/evaler v2.2.0+incompatible // indirect
	github.com/zzzzzzzzzzz0/zhscript-go v0.0.0-00010101000000-000000000000
	util4 v0.0.0-00010101000000-000000000000
	zsp2 v0.0.0-00010101000000-000000000000
)

replace (
	clpars4 => ./src/clpars4
	github.com/zzzzzzzzzzz0/zhscript-go => ../zhscript-go
	util4 => ./src/util4
	zsp2 => ./src/zsp2
)
