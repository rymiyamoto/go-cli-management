# go-cli-management

## 概要

cobra を使って CLI アプリケーションを作成しています。

go v1.21.4 で動かしていますが任意のバージョンに切り替えてもらって構いません。

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

## 使い方

### 初期値の拡張

cmd/root.go に初期値を追加します。

今回は並列処理の管理と dry-run モードを追加します。

```go
const (
	// concurrencyDefault デフォルトの並列数
	concurrencyDefault = 10
	// waitTimeDefault デフォルトの処理チャンク単位の待機時間
	waitTimeDefault = 1
)

// ...

func init() {
	rootCmd.PersistentFlags().Bool("dry-run", false, "Dry run mode")
	rootCmd.PersistentFlags().Uint("concurrency", concurrencyDefault, "並列更新数(1以上)")
	rootCmd.PersistentFlags().Uint("wait-time", waitTimeDefault, "処理チャンク単位の待機時間(秒)")
}
```

### サブコマンドの追加

### 作成

make コマンドを使ってサブコマンドを追加します。

```sh
$ make add-batch cmd=hello
```

実行すると cmd 配下に hello.go が作成されます。(以下参照)

```go
/*
Copyright © 2024 rymiyamoto

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello called")
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

```

### 引数やフラグの追加

引数やフラグの追加は root のときと同様に行います

```go

// ...

func init() {
	rootCmd.AddCommand(helloCmd)
	helloCmd.Flags().StringP("target-at", "t", time.Now().In(time.FixedZone("Asia/Tokyo", 9*60*60)).Format(time.DateOnly), "対象日(e.g 2023-10-05)")
}
```

### 処理の整形

Run の部分に処理を記載していきますが、Run だとエラーの検知ができないので、RunE にしてエラーを返すようにします。
またこのときデフォルト引数やフラグを slog で埋め込むことでログに出力できるようにします。

```go
// ...

RunE: func(cmd *cobra.Command, args []string) error {
    // デフォルトフラグ
	dryRun, _ := rootCmd.Flags().GetBool("dry-run")
    concurrency, _ := rootCmd.Flags().GetUint("concurrency")
	waitTime, _ := rootCmd.Flags().GetUint("wait-time")

	// サブコマンド固有フラグ
	targetAt, _ := cmd.Flags().GetString("target-at")

    base := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
    logger := base.With("dry-run", dryRun, "concurrency", concurrency, "wait-time", waitTime, "target-at", targetAt)
    slog.SetDefault(logger)

    slog.Info("hello world!")
    return nil
},
```

## サブコマンドの実行

meko コマンドを使ってサブコマンドを実行します。

```sh
# cmd部分はフラグがあれば記載できるように文字列で指定します
$ make run-batch cmd="hello"
```

実行すると以下のようなログが出力されます。

```sh
{"time":"2024-03-21T03:36:16.496879503Z","level":"INFO","msg":"hello world!","dry-run":false,"concurrency":10,"wait-time":1,"target-at":"2024-03-21"}
```
