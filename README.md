Content improvement

    → Существует страница с размещенным на ней контентом
    
    → Пользователи имеют к ней полный доступ
    
    → Они могут выгружать его себе затем изменять и добавлять на страницу ресурса
    
    → Таким образом формируется дерево с развитием какого-либо медиа-контента
    
    → Необходимо реализовать оценку контента для получения какого-то результата на основе наиболее высоко-оцененных доработок
    
Тэги: ресурс совместного развития идей; реклама контента; продвижение




-- CREATE MIGRATION
```sh
    migrate create -ext sql -dir migrations -seq users
```
-- EXECUTE MIGRATIONS
```sh
    migrate -path ./migrations -database "postgres://username:password@localhost:5432/dbname?sslmode=disable" up
```

--gRPC
```sh
    protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        helloworld/helloworld.proto
```
--


```json
    https://helllolworld.atlassian.net/jira/software/projects/MB/boards/35/backlog
```