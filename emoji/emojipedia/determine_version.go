package emojipedia

import "OwlGramServer/emoji/emojipedia/types"

func determineVersion(pd map[string]*types.ProviderDescriptor, versionsMap map[int][]string) {
	for _, v := range pd {
		var topVersion int
		for unicodeVersion, emojis := range versionsMap {
			for _, emoji := range emojis {
				if _, ok := v.Emojis[emoji]; ok {
					if unicodeVersion > topVersion {
						topVersion = unicodeVersion
					}
				}
			}
		}
		v.UnicodeVersion = topVersion
	}
}
