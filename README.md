# snail_girl
田螺姑娘: 帮你实现自动撸码的愿望

### 自动生成code
``
go run ./cmd/domain/main.go by_file your_file/domain/xxx.go
``

``or``

``
go build ./cmd/domain/main.go && ./main by_file your_file/domain/xxx.go
``
### 修改模板后执行
``make gen``