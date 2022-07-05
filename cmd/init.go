package cmd

func init() {
	initJsonCmd()

	rootCmd.AddCommand(readCmd)
}
