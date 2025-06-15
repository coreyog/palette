/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/coreyog/palette/pkg/gen"
	"github.com/coreyog/palette/pkg/model"

	"github.com/spf13/cobra"
)

// extractCmd represents the extract command
var extractCmd = &cobra.Command{
	Use:   "extract IMAGE",
	Short: "Extract a palette from an existing image",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := extractPalette(args[0])

		if err != nil {
			fmt.Printf("something went wrong: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(extractCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// extractCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// extractCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func extractPalette(path string) (err error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	fmt.Println("extracting palette from", path)

	img, err := png.Decode(bytes.NewReader(raw))
	if err != nil {
		return err
	}

	ref := gen.ReferencePalette()
	template := image.NewRGBA(img.Bounds())
	pal := model.NewPalette()

	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			c := img.At(x, y)
			_, pal = pal.Add(c)

			index := pal.IndexOf(c)
			if index != -1 {
				unmapped := ref[index]
				template.Set(x, y, unmapped)
			}
		}
	}

	dir, file := filepath.Split(path)
	ext := filepath.Ext(file)
	basefilename := strings.TrimSuffix(file, ext)

	palFilename := fmt.Sprintf("%s_palette.png", basefilename)
	outPath := filepath.Join(dir, palFilename)

	out, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer out.Close()

	p := pal.ToImage()

	err = png.Encode(out, p)
	if err != nil {
		return err
	}

	templateFilename := fmt.Sprintf("%s_template.png", basefilename)
	outPath = filepath.Join(dir, templateFilename)

	out, err = os.Create(outPath)
	if err != nil {
		return err
	}
	defer out.Close()

	err = png.Encode(out, template)
	if err != nil {
		return err
	}

	return nil
}
