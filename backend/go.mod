module github.com/arry/WB_project

go 1.24.4

replace github.com/arry/WB_project/internal/model/dispayed => ../home/arry/WB_project/code/backend/internal/model/displayed

require github.com/lib/pq v1.10.9

require github.com/julienschmidt/httprouter v1.3.0 // indirect
