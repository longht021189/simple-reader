package cmd

import (
	"github.com/longht021189/simplereader/src"
	"github.com/spf13/cobra"
)

type ReadVar struct {
	File string
	Key  string
	Type string
}

var readVar = &ReadVar{}

var readCmd = &cobra.Command{
	Use: "read",
	Run: func(cmd *cobra.Command, args []string) {
		src.Read(readVar.File, readVar.Key, readVar.Type)
	},
}

func initJsonCmd() {
	readCmd.Flags().StringVarP(&readVar.File, "file", "f", "", "File Path")
	readCmd.Flags().StringVarP(&readVar.Key, "key", "k", "", "Key Name. Ex: key1.key2.key3")
	readCmd.Flags().StringVarP(&readVar.Type, "type", "t", "json", "File Type")
}
