/*
Copyright © 2024 rymiyamoto
*/
package cmd

import (
	"log/slog"
	"os"
	"time"

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
	RunE: func(cmd *cobra.Command, args []string) error {
		// デフォルトフラグ
		dryRun, _ := rootCmd.Flags().GetBool("dry-run")
		concurrency, _ := rootCmd.Flags().GetUint("concurrency")
		waitTime, _ := rootCmd.Flags().GetUint("wait-time")

		// サブコマンド固有フラグ
		targetAt, _ := cmd.Flags().GetString("target-at")

		// 各引数を埋め込んでログ設定
		base := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
		logger := base.With("dry-run", dryRun, "concurrency", concurrency, "wait-time", waitTime, "target-at", targetAt)
		slog.SetDefault(logger)

		slog.Info("hello world!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
	helloCmd.Flags().StringP("target-at", "t", time.Now().In(time.FixedZone("Asia/Tokyo", 9*60*60)).Format(time.DateOnly), "対象日(e.g 2023-10-05)")
}
