package google

import (
	"strings"

	"github.com/spf13/cobra"
	"fmt"
)

const (
	imageLong = `
		image make a custom search on google filtering by only images and
		returns the first result.`

	imageExample = `
		# Post a image of dogs
		!image dogs

		# Post a image of a hot dog
		!image hot dog

		# Post two images of hot dog
		!image hot dog --total 2

		# Post the second image of a hot dog
		!image hot dog --skip 1

		# Post the third and fourth image of a hot dog
		!image hot dog --skip 2 --total 2`
)

type (
	googleImage struct {
		google
	}
)

func NewGoogleImageCommand(key, cx string) *cobra.Command {
	i := newGoogleImage(key, cx)

	c := &cobra.Command{
		Use:     "image QUERY",
		Short:   "image search for images using Google",
		Long:    imageLong,
		Example: imageExample,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			r, err := i.Search(strings.Join(args, " "))
			if err != nil {
				r = err.Error()
			}
			cmd.OutOrStdout().Write([]byte(r))
			i.reset()
		},
		Aliases: []string{"gis", "pictures", "imagem", "imagens", "gi"},
	}

	c.Flags().IntVarP(&i.total, "total", "t",1, "How many images will be returned")
	c.Flags().IntVarP(&i.skip, "skip","s", 0, "How many images should be skipped")

	return c
}

func newGoogleImage(key, cx string) *googleImage {
	return &googleImage{
		google{
			key: key,
			cx : cx,
		},
	}
}

func (gi *googleImage) Search(q string) (string, error) {
	params := map[string]string{}
	params["searchType"] = "image"

	r, err := gi.search(q, params)
	if err != nil {
		return "", err
	}

	var msgList []string
	for _, item := range r {
		msgList = append(msgList, fmt.Sprintf("%s - %s", item.Title, item.Link))
	}

	return strings.Join(msgList, "\n"), nil
}

