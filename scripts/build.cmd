go build -o build/inspect.exe inspect.go 
go build -o build/randfile.exe randfile.go
xcopy build\inspect.exe F:\D\Programs\bin /Y
xcopy build\randfile.exe F:\D\Programs\bin /Y
