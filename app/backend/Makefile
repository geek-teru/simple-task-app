.PHONY: ent-new generate migrate

# エンティティ作成
# 例: make ent-new name=Task
ent-new:
	go run -mod=mod entgo.io/ent/cmd/ent new $(name)

# アセット生成
generate:
	go generate ./ent

# マイグレーション実行
migrate:
	go run ./cmd/main.go migrate

# test
test:
	go test ./... -v