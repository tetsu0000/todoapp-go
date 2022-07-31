# デプロイ手順

以下のコマンドを実行

```bash
go mod tidy
cdk deploy TodoappGoStack
```

# 動作確認

以下のコマンドを実行

```bash
export ENDPOINT="API Gatewayのエンドポイント"
bash test/test.sh
```

# アプリ削除方法

以下のコマンドを実行

```bash
cdk destroy TodoappGoStack
```
