/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"image/png"
	"os"
	"strings"

	"github.com/coreyog/palette/pkg/gen"

	"github.com/spf13/cobra"
)

// genRefImageCmd represents the genRefImage command
var genRefImageCmd = &cobra.Command{
	Use:   "gen KIND",
	Short: "Used to generate images",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		kind := strings.ToLower(args[0])

		var err error

		switch kind {
		case "ref", "refimg", "refimage":
			err = saveRefImage()
		}

		if err != nil {
			fmt.Printf("something went wrong: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(genRefImageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genRefImageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genRefImageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func saveRefImage() (err error) {
	img := gen.ReferenceImage()

	ref, err := os.Create("ref.png")
	if err != nil {
		return err
	}
	defer ref.Close()

	err = png.Encode(ref, img)
	if err != nil {
		return err
	}

	return nil
}
