package app

import (
	"flag"
	"fmt"
	"strings"

	"github.com/omeid/uconfig/flat"
	"github.com/rs/zerolog/log"
)

func mergeWithBeamFlags(config any) {
	fields, err := flat.View(config)
	if err != nil {
		panic(fmt.Errorf("extracting config fields: %w", err))
	}

	for _, field := range fields {
		name := flagName(field)
		if name == "-" {
			continue
		}

		if flg := flag.CommandLine.Lookup(name); flg != nil {
			overwriteFlag(flg, config, field)
		} else {
			usage, _ := field.Tag("usage")
			flag.CommandLine.Var(field, name, usage)
		}
	}
}

func flagName(field flat.Field) string {
	name, _ := field.Tag("flag")
	if name != "" {
		return name
	}

	return strings.ToLower(strings.ReplaceAll(field.Name(), ".", "-"))
}

func overwriteFlag(flg *flag.Flag, config any, field flat.Field) {
	switch v := nestedFieldValueString(config, field.Name()); v {
	case "", "0", "false", "[]":
	default:
		err := flg.Value.Set(v)
		if err != nil {
			panic(fmt.Errorf("setting command line flag '%s': %w", flg.Name, err))
		}
		log.Trace().
			Str("config_name", field.Name()).
			Str("flag_name", flg.Name).
			Str("flag_value", v).
			Msg("[arqbeam] Set beam flag.")
	}
}
