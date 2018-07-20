package bible

import (
	"github.com/spf13/cobra"
	"fmt"
	"strings"
	"net/url"
	"github.com/disiqueira/ultraslackbot/pkg/command"
	"errors"
)

const (
	urlBible = "http://labs.bible.org/api/?passage=%s&type=json&formatting=plain"

	formatBible = "%s %d:%d - %s"
	example = `
		# Get John 3:16 bible verse
		!bible John 3:16`
)

type (
	bible struct {
		total int
		skip  int
		random bool
		votd bool
	}

	bibleResponse []struct {
		Bookname string `json:"bookname"`
		Chapter  string `json:"chapter"`
		Verse    string `json:"verse"`
		Text     string `json:"text"`
	}
)

func NewBibleCommand() *cobra.Command {
	b := newBible()

	c := &cobra.Command{
		Use:     "bible",
		Short:   "Ask a question to Bible",
		Example: example,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			r, err := b.search(strings.Join(args, " "))
			if err != nil {
				r = err.Error()
			}
			cmd.OutOrStdout().Write([]byte(r))
			b.reset()
		},
	}

	c.Flags().IntVarP(&b.total, "total", "t",1, "How many verses will be returned")
	c.Flags().IntVarP(&b.skip, "skip","s", 0, "How many verses should be skipped")
	c.Flags().BoolVarP(&b.random, "random", "r",false, "Return an random verse")
	c.Flags().BoolVarP(&b.votd, "verse-of-the-day","v", false, "Return the Bible.org Verse of the Day (VOTD)")

	return c
}

func newBible() *bible {
	return &bible{}
}

func (b *bible) search(q string) (string, error) {
	if err := b.validate(); err != nil {
		return "", err
	}

	if b.random {
		q = "random"
	}

	if b.votd {
		q = "votd"
	}

	data := &bibleResponse{}
	query := url.QueryEscape(q)
	u := fmt.Sprintf(urlBible, query)
	err := command.GetJSON(u, data)
	if err != nil {
		return "", err
	}

	var msgList []string
	skip := 0
	for _, item := range *data {
		if b.skip > skip {
			skip++
			continue
		}
		msgList = append(msgList, fmt.Sprintf(formatBible, item.Bookname, item.Chapter, item.Verse, item.Text))

		if len(msgList) == b.total {
			break
		}
	}

	return strings.Join(msgList, " "), nil
}

func (b *bible) validate() error {
	if b.votd && b.random {
		return errors.New("--random and --votd are mutually exclusive")
	}

	return nil
}

func (b *bible) reset() {
	b.skip = 0
	b.total = 1
	b.random = false
	b.votd = false
}
