# simple-task-app
  シンプルなタスク管理アプリケーション

## 環境
  ### Backend API
    * Language: Go 1.23.4
    * Web FW: Echo
    * ORM: ent
  ### Database
    * RDBMS: PostgreSQL
  ### infra
    * LB: AWS ALB
    * Computing: AWS ECS
    * DB: AWS RDS Aurora
    * IaC: Terraform
  ### Other
    * CI/CD: Github Actions
    * Container: Docker

## 実装のポイント
  * レイヤードアーキテクチャ
  * シンプルなCRUDを実装
  * 認証機能を実装
  * バリデーション機能を実装
  * テストを実装
  * Terraformでインフラのコード化
  * GithubActionsでビルド・デプロイを自動化

## DB設計
  ### サーバー
  * database名 `sampledb`
  * 文字コード、タイムゾーン、ログ設定、パフォーマンス設定はデフォルト

  ### 制約
  * 基本、すべてのカラムにNOT NULL制約をつけてデフォルト値を設定する。deleted_atのみNULLを許容する
  * 文字列はvarchar(100), varchar(255), varchar(512)
  * CHECK制約でidの負の値を禁止する
  * created_at, updateded_at, deleted_atは固定
  * created_at, updateded_atのデフォルトは `CURRENT_TIMESTAMP`とする
  * deleted_atのデフォルトは`NULL`とする

  ### ER図
  ```mermaid
  erDiagram
      User {
          int id PK "ユーザーID"
          varchar(100) name "名前"
          varchar(100) email "メールアドレス"
          varchar(100) password "パスワード"
          timestamp created_at
          timestamp updateded_at
          timestamp deleted_at
      }
      Task {
          int id PK "タスクID"
          varchar(100) title "タイトル"
          varchar(255) description "詳細"
          date due_date "期限日"
          int status "ステータス (e.g., TODO, IN_PROGRESS, DONE)"
          int user_id FK "ユーザーID"
          timestamp created_at
          timestamp updateded_at
          timestamp deleted_at
      }
      
      
      User ||--o{ Task : "1対多"
  ```

  ### インデックス
  * idとcreated_atだけインデックスを設定しておく

## ディレクトリ構成
  ディレクトリ構成は以下の通り
  ```
  backend/
  ├── cmd/
  │   └── main.go            # メインファイル                
  ├── ent/                   # ドメインモデルやエンティティ
  │   └── task.go            # Taskエンティティ
  ├── repository/            # リポジトリ層 (データアクセス)
  │   └── task_repository.go # リポジトリインターフェース
  ├── service/               # サービス層 (ビジネスロジック)
  │   └── task_service.go    # タスクのビジネスロジック
  ├── handler/               # ハンドラー層 (プレゼンテーション層)
  │   └── task_handler.go    # HTTPハンドラー
  ├── router/                # ルーティング設定
  │   └── router.go          # ルーティング処理
  ├── config/                # 設定管理
  │   └── config.go          # 設定情報
  ├── test/                  # テスト用
  │
  ├── pkg/                   # 再利用可能なパッケージ
  │   ├── db/                # DB接続用パッケージ
  │   │   └── db.go              
  │   └── log/               # ロギングユーティリティ
  │       └── log.go
  ```

## 利用方法
[利用方法](docs/note.md)

## 開発メモ
[開発メモ](docs/note.md)