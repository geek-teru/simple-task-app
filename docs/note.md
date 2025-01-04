## Backendの設計アプローチ

### レイヤードアーキテクチャで各層を構造体にするのはなぜ
レイヤードアーキテクチャで各層で「importしてメソッドを呼び出す」のではなく「構造体（あるいはクラス）として実装する」理由は、以下のような設計上のメリットを得るため。
1. 依存性の注入 (Dependency Injection) の容易化
  構造体を使うことで、インターフェースを注入可能になり、モジュール間の結合度を下げ、テスト時に依存するリポジトリやサービスを簡単にモックに差し替えることもできる。

2. 状態を持つ実装を可能にする
  データベース接続やキャッシュを各層で跨って利用できる

### 依存性の注入（Dependency Injection: DI）とは
- 依存性の注入とは、クラスや構造体が必要とする依存オブジェクトを自身で生成するのではなく、外部から提供（注入）する設計手法。
- これにより、モジュール間の結合度を低くし、柔軟性やテストの容易性を向上させることができる。

依存オブジェクトを自身で生成する例
```go
package service

import (
    "example.com/repository"
)

type UserService struct {
    Repository *repository.MySQLUserRepository // インターフェースではなくインスタンス
}

func NewUserService() *UserService {
    repo := &repository.MySQLUserRepository{} // 自分で依存を生成
    return &UserService{Repository: repo}
}
```
MySQLUserRepositoryに依存しているため、例えばMockRepositoryに変更したい場合にUserServiceのコードを修正する必要がある。


依存オブジェクトを外部から提供する例(DI)
```go
package service

import (
    "example.com/repository"
)

type UserService struct {
    Repository repository.UserRepository // インターフェースに依存
}

func NewUserService(repo repository.UserRepository) *UserService {
    return &UserService{Repository: repo} // 外部から依存を注入
}

```
実際のMySQLUserRepositoryだけでなく、MockRepositoryなどの別実装を注入できる。

## テストについて

### testfixtures
testfixturesライブラリを利用し、テスト関数内で初期データの投入を行う。
fixtures の投入処理では テーブルのデータが完全に入れ替えられる。（テーブルが削除された後、新規作成される）

### テストデータについて
テストデータは別ファイル(package testdata)に切り出し再利用できるようにする。

## よく使う手順メモ

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
go test -v ./test
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