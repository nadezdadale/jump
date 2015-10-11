package cmd

import (
	"sort"
	"strconv"

	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
	"github.com/gsamokovarov/jump/scoring"
)

func hintCmd(args cli.Args, conf *config.Config) {
	var hints scoring.Entries

	count, err := strconv.Atoi(args.Value("--count", "0"))
	if err != nil {
		count = 0
	}

	entries, err := conf.ReadEntries()
	if err != nil {
		cli.Exitf(1, "%s\n", err)
	}

	if len(args) == 0 {
		// We usually keep them reversely sort to optimize the fuzzy search.
		sort.Sort(sort.Reverse(entries))

		hints = hintSliceEntries(entries, count)
	} else {
		fuzzyEntries := scoring.NewFuzzyEntries(entries, args.CommandName())
		fuzzyEntries.Sort()

		hints = hintSliceEntries(fuzzyEntries.Entries, count)
	}

	for _, entry := range hints {
		cli.Outf("%s\n", entry.Path)
	}
}

func hintSliceEntries(entries scoring.Entries, limit int) scoring.Entries {
	if limit > 0 && limit < len(entries) {
		return entries[0:limit]
	} else {
		return entries
	}
}

func init() {
	cli.RegisterCommand("hint", "Hints relevant paths for jumping.", hintCmd)
}
