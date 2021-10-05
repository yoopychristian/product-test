del go.mod
del go.sum
@RD /S /Q "vendor"
go mod init product-test
go mod tidy 
