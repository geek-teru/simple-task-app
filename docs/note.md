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

### 値渡しとポインタ渡しについて整理
値渡しの特徴
* 概要
  * 構造体の コピーを渡す 方法です。
  * メモリ上で新しいインスタンスが生成され、オリジナルのデータとは独立します。
* メリット
  * 安全性が高い
  * 呼び出し先でデータを変更しても、オリジナルに影響を与えない。
  * 意図しない副作用が起きないため、堅牢な設計が可能。
  * 小さなデータに適している
  * 構造体が小さい場合、コピーコストが低いので効率的。
  * シンプルな設計
  * ポインタ管理やnilチェックが不要。
* デメリット
  * 大きなデータの場合、コピーコストが高い
  * 構造体が大きいと、メモリ消費や処理時間に影響する。
  * 変更が必要な場合に不便
  * 呼び出し先でオリジナルを変更したい場合、値渡しだと変更できない。
* 使いどころ
  * データが小さく、変更の必要がない場合。
  * レイヤードアーキテクチャのservice層 → handler層のように、不変なデータを渡す場合。


ポインタ渡しの特徴
* 概要
  * 構造体の メモリアドレス（ポインタ）を渡す 方法です。
  * 呼び出し先はオリジナルデータへの参照を取得します。
* メリット
  * メモリ効率が良い
  * 構造体が大きい場合でもコピーせずに渡せるため、効率的。
  * 呼び出し先でデータの変更が可能
  * 呼び出し先で直接データを操作する場合に便利。
* デメリット
  * 安全性が低い
  * 呼び出し先でデータを変更するとオリジナルにも影響する。
  * 意図しない副作用が起こる可能性がある。
  * コードが複雑になる
  * nilチェックが必要。
  * ポインタの扱いに注意が必要。
* 使いどころ
  * 構造体が大きい場合。
  * 呼び出し先でデータを更新する必要がある場合。
  * パフォーマンスが特に重要な処理。

# 値渡しとポインタ渡しの使い分け

| **比較項目**         | **値渡し**                      | **ポインタ渡し**                  |
|-----------------------|--------------------------------|-----------------------------------|
| **メモリ効率**         | 小さなデータなら良い             | 大きなデータに適している             |
| **データの安全性**     | 呼び出し元のデータを守れる       | 呼び出し元のデータに影響を与える可能性 |
| **変更の必要性**       | 変更しない場合に適している        | 変更が必要な場合に適している         |
| **コードの簡潔さ**     | シンプルで扱いやすい             | `nil`チェックやポインタの管理が必要    |
| **用途**             | 不変のデータ、軽量データを渡す    | 可変データ、重量級データを渡す       |

* 結論
  * 小さな構造体や不変データ → 値渡しが推奨
    * 例: UserResponseのように小さなデータをservice層からhandler層に渡す場合。
  * 大きな構造体や変更が必要なデータ → ポインタ渡しが推奨
    * 例: 大規模なデータ処理や変更を前提とした設計。

シンプルで安全な設計を目指すなら、まずは 値渡し を基本とし、パフォーマンスや変更要件を理由に ポインタ渡し を採用するか検討するのが良い。


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
go test -v ./test/repository
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

## mockの生成
1. 準備

```
go get github.com/golang/mock/mockgen@v1.6.0
go install github.com/golang/mock/mockgen@v1.6.0
go mod tidy

```

2. mock生成
```
mockgen -source=user.go -destination=./mock/user_mock.go -package=repository
```