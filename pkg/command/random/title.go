package random

import (
	"errors"

	"github.com/Pallinder/go-randomdata"
	"github.com/spf13/cobra"
)

const (
	exampleTitle = `
		# Generate a random title
		!random title

		# Generate a random male title
		!random title --male

		# Generate a random female title
		!random title --female`
)

type (
	randomTitle struct {
		male   bool
		female bool
	}
)

func newRandomTitleCommand() *cobra.Command {
	n := newRandomTitle()

	c := &cobra.Command{
		Use:     "title",
		Short:   "Generate a random title",
		Example: exampleTitle,
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			text, err := n.generate()
			if err != nil {
				text = err.Error()
			}
			cmd.OutOrStdout().Write([]byte(text))
			n.reset()
		},
	}

	c.Flags().BoolVarP(&n.male, "male", "m", false, "Only male titles")
	c.Flags().BoolVarP(&n.female, "female", "f", false, "Only female titles")

	return c
}

func newRandomTitle() *randomTitle {
	return &randomTitle{}
}

func (n *randomTitle) generate() (string, error) {
	if n.male && n.female {
		return "", errors.New("--male and --female are mutually exclusives")
	}

	gender := randomdata.RandomGender
	if n.male {
		gender = randomdata.Male
	}
	if n.female {
		gender = randomdata.Female
	}

	return randomdata.Title(gender), nil
}

func (n *randomTitle) reset() {
	n.female = false
	n.male = false
}
