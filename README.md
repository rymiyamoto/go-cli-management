# go-cli-management

## 概要

cobra を使って CLI アプリケーションを作成しています。

go v1.21.4 で動かしていますが任意の任意のバージョンに切り替えてもらって構いません。

## はじめに

以下のコマンドで cobra を使えるようにします。

```sh
$ make init-batch
```

その後以下のような高瀬になっていれば準備 OK です

```sh
$ tree
.
├── Dockerfile
├── LICENSE
├── Makefile
├── README.md
├── cmd
│   └── root.go
├── cobra.yml
├── docker-compose.yml
├── go.mod
├── go.sum
└── main.go
```

## 構成説明

- `cmd` ディレクトリ

  - `root.go` は cobra で生成されたファイルです
  - このファイルをベースにしてコマンドを追加していきます

- `cobra.yml` ファイル

  - cobra でコマンドを追加する際の設定ファイルです
  - このファイルを編集することでコマンドの中身を拡張できます

- `Dockerfile` ファイル

  - Docker イメージを作成するためのファイルです

- `docker-compose.yml` ファイル

  - Docker イメージを作成するためのファイルです
  - 本来 1 つしかないので不要ですが既存のサービスとの連携を考えて作成してます

- `go.mod` ファイル

  - go modules の設定ファイルです

- `go.sum` ファイル

  - go modules の設定ファイルです

- `main.go` ファイル

  - エントリーポイントです

- `Makefile` ファイル
  - 長くなりがちなので Makefile で管理しています
