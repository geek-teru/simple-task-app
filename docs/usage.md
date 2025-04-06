# 利用方法

## DB関連
  ### 起動
  ```
  docker-compose up -d

  # build
  docker compose up -d --build backend

  # delete
  docker-compose down --rmi all --volumes --remove-orphans
  ```

  ### 設定確認
  ```
  # 文字コード
  docker exec -it postgres.local psql -U admin -d sampledb -c "\l sampledb"

  # タイムゾーン
  docker exec -it postgres.local psql -U admin -d sampledb -c "SHOW timezone;"

  # ログ関連
  docker exec -it postgres.local psql -U admin -d sampledb -c "SHOW logging_collector;"
  docker exec -it postgres.local psql -U admin -d sampledb -c "SHOW log_directory;"
  docker exec -it postgres.local psql -U admin -d sampledb -c "SHOW log_filename;"
  docker exec -it postgres.local psql -U admin -d sampledb -c "SHOW log_statement;"
  docker exec -it postgres.local psql -U admin -d sampledb -c "SHOW log_min_duration_statement;"

  # パフォーマンス関連
  docker exec -it postgres.local psql -U admin -d sampledb -c "SHOW shared_buffers;"
  docker exec -it postgres.local psql -U admin -d sampledb -c "SHOW work_mem;"
  docker exec -it postgres.local psql -U admin -d sampledb -c "SHOW maintenance_work_mem;"
  docker exec -it postgres.local psql -U admin -d sampledb -c "SHOW max_connections;"
  ```


  ### テーブル確認
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
  ```

## ORM(ent)関連
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

  ### マイグレーション
  ```
  docker compose run --rm migrate bash -c'go run cmd/main.go migrate'
  ```


## テスト関連
  ### テスト実行
  ```
  go test -v ./...
  go test -v ./test/repository
  ```

## mock関連
  ### mockの生成
  1. 準備

  ```
  go get github.com/golang/mock/mockgen@v1.6.0
  go install github.com/golang/mock/mockgen@v1.6.0
  go mod tidy

  ```

  2. mock生成
  ```
  mockgen -source=user.go -destination=./mock/user_mock.go -package=repository
  mockgen -source=task.go -destination=./mock/task_mock.go -package=repository
  ```

## E2Eテスト
  ```
  # サインアップ
  curl -X POST  http://localhost:8080/signup -H "Content-Type: application/json" -d '{"name": "teru", "email":"teru@example.com", "password": "terupassword"}'

  # サインイン
  token=$(curl -X POST  http://localhost:8080/signin -H "Content-Type: application/json" -d '{"name": "teru", "email":"teru@example.com", "password": "terupassword"}'| tr -d '"')

  # task登録
  curl -X POST http://localhost:8080/task \
  -H "Authorization: Bearer $token" \
  -H "Content-Type: application/json" \
  -d '{"title": "task01", "description": "task01description", "status": "TODO", "due_date": "2024-01-01T00:00:00Z"}'

  # task一覧
  curl -X GET http://localhost:8080/task?p=1 \
  -H "Authorization: Bearer $token" \
  -H "Content-Type: application/json"

  # task参照
  curl -X GET http://localhost:8080/task/10001 \
  -H "Authorization: Bearer $token"
  ```