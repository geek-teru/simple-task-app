## Backend
### モジュールの初期化
```
repo=<your_repository> # github.com/<user_name>/<repository_name>の形式
go mod init ${repo}
```

### モジュールのダウンロード
```
go mod tidy

```

### entの設定
1. Userエンティティの作成
  * `ent/schema/user.go` が作成される
```
go run -mod=mod entgo.io/ent/cmd/ent new User
```

2. Userエンティティのフィールドを定義
  * `ent/schema/user.go` の `func (User) Fields()` 内を編集
  * https://entgo.io/ja/docs/schema-fields

3. アセットの作成
  * `ent/user` にアセットが作成される
```
go generate ./ent
```

## DB起動
```
docker-compose up -d

# delete
docker-compose down --rmi all --volumes --remove-orphans
```

## マイグレーション
```
go mod tidy
go run cmd/migrate.go
```

## テーブル確認
1. テーブル確認
```
docker exec -it postgres.local psql -U admin -d sampledb -c "\dt"
       List of relations        
 Schema | Name  | Type  | Owner 
--------+-------+-------+-------
 public | users | table | admin 
(1 row)
```

2. SELECT
```
docker exec -it postgres.local psql -U admin -d sampledb -c "select * from users;"
 id | name | email | password
----+------+-------+----------
(0 rows)

```

## テストの実行
```
go test -v ./...
```

## サーバー起動
```
go run main.go

```

## 動作確認用リクエスト
```
curl -X GET  http://localhost:8080/healthcheck
curl -X POST  http://localhost:8080/user -H "Content-Type: application/json" -d '{"Name": "user_x", "Email": "user_x@example.com", "Password": "password"}'
curl -X GET  http://localhost:8080/user/1
```