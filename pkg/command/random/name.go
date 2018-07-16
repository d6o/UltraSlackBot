package random

import (
	"errors"
	
	"github.com/spf13/cobra"
	"github.com/Pallinder/go-randomdata"
)

const (
	exampleName = `
		# Generate a random Name
		!random name

		# Generate a random male Name
		!random name --male

		# Generate a random female Name
		!random name --female

		# Generate a random first Name
		!random name --first

		# Generate a random first male Name
		!random name --first --male

		# Generate a random last Name
		!random name --last

		# Generate a random silly name
		!random name --silly`
)

type (
	randomName struct {
		male bool
		female bool
		first bool
		last bool
		silly bool
	}
)

func newRandomNameCommand() *cobra.Command {
	n := newRandomName()

	c := &cobra.Command{
		Use:     "name",
		Short:   "Generate a random name",
		Example: exampleName,
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

	c.Flags().BoolVarP(&n.male, "male", "m", false, "Only male names")
	c.Flags().BoolVarP(&n.female, "female", "f", false, "Only female names")

	c.Flags().BoolVarP(&n.first, "first", "i", false, "Only the first name")
	c.Flags().BoolVarP(&n.last, "last", "l", false, "Only the last name")

	c.Flags().BoolVarP(&n.silly, "silly", "s", false, "Silly name")

	return c
}

func newRandomName() *randomName {
	return &randomName{}
}

func (n *randomName) generate() (string, error) {
	if err := n.validate(); err != nil {
		return "", err
	}
	if n.silly {
		return randomdata.SillyName(), nil
	}

	gender := randomdata.RandomGender
	if n.male {
		gender = randomdata.Male
	}
	if n.female {
		gender = randomdata.Female
	}

	if n.first {
		return randomdata.FirstName(gender), nil
	}

	if n.last {
		return randomdata.LastName(), nil
	}

	return randomdata.FullName(gender), nil
}

func (n *randomName) validate() error {
	if n.silly && (n.male || n.female || n.first || n.last) {
		return errors.New("can not combine --silly with other flags")
	}
	if n.male && n.female {
		return errors.New("--male and --female are mutually exclusives")
	}
	if n.first && n.last {
		return errors.New("--first and --last are mutually exclusives")
	}
	if n.last && (n.male || n.female) {
		return errors.New("you can not specify a gender for a last name")
	}
	return nil
}

func (n *randomName) reset() {
	n.last = false
	n.first = false
	n.female = false
	n.male = false
	n.silly = false
}

