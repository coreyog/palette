/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/coreyog/palette/pkg/apply"
	"github.com/coreyog/palette/pkg/model"
	"github.com/coreyog/palette/pkg/util"

	"github.com/spf13/cobra"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply TEMPLATE PALETTE",
	Short: "Apply a palette to a template",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		err := applyPaletteToTemplate(args[0], args[1])

		if err != nil {
			fmt.Printf("something went wrong: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// applyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// applyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func applyPaletteToTemplate(templatePath, palettePath string) error {
	tmpl, err := util.ReadImage(templatePath)
	if err != nil {
		return fmt.Errorf("unable to read template: %w", err)
	}

	pal, err := util.ReadImage(palettePath)
	if err != nil {
		return fmt.Errorf("unable to read palette: %w", err)
	}

	p, err := model.LoadPalette(pal)
	if err != nil {
		return fmt.Errorf("unable to load palette: %w", err)
	}

	colored, err := apply.ApplyPaletteToTemplate(tmpl, p)
	if err != nil {
		return fmt.Errorf("unable to apply palette to template: %w", err)
	}

	_, palFile := filepath.Split(palettePath)
	palExt := filepath.Ext(palFile)
	palBaseFilename := strings.TrimSuffix(palFile, palExt)
	palBaseFilename = strings.TrimSuffix(palBaseFilename, "_palette")

	tmpDir, tmpFile := filepath.Split(templatePath)
	tmpExt := filepath.Ext(tmpFile)
	tmpBaseFilename := strings.TrimSuffix(tmpFile, tmpExt)
	tmpBaseFilename = strings.TrimSuffix(tmpBaseFilename, "_template")

	name := fmt.Sprintf("%s_colored_w_%s.png", tmpBaseFilename, palBaseFilename)
	outFilename := filepath.Join(tmpDir, name)

	return util.WriteImage(outFilename, colored)
}
