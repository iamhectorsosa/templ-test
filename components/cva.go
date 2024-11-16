package components

import "strings"

type StyleConfig struct {
	Base            []string
	Variants        map[string]map[string]string
	DefaultVariants map[string]string
}

type StyleProps struct {
	Variants map[string]string
	Class    string
}

func CVA(options StyleConfig) func(props StyleProps) string {
	return func(props StyleProps) string {
		var classes strings.Builder
		classes.WriteString(strings.Join(options.Base, " "))

		if len(options.Variants) > 0 {
			for variant, variantOptions := range options.Variants {
				value := props.Variants[variant]
				if value == "" {
					value = options.DefaultVariants[variant]
				}
				if class, ok := variantOptions[value]; ok {
					classes.WriteString(" " + class)
				}
			}
		}

		if props.Class != "" {
			classes.WriteString(" " + props.Class)
		}

		return classes.String()
	}
}
