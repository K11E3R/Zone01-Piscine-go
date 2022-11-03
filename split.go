package piscine

func Split(s, sep string) []string {
	var r []string
	k := 0
	if len(sep) > len(s) {
		r = append(r, s)
	} else {
		for i := 0; i < len(s); i++ {
			if k+len(sep) > len(s) {
				break
			} else {
				if s[i] == sep[0] {
					for j := 1; j < len(sep); j++ {
						if s[i+j] != sep[j] {
							break
						}
						if j == len(sep)-1 {
							{
								var m string
								for k := k; k < i; k++ {
									m = m + string(s[k])
								}
								r = append(r, m)
								k = i + len(sep)
							}
						}
					}
				}
			}
		}
		var m string
		for j := k; j < len(s); j++ {
			m = m + string(s[j])
		}
		r = append(r, m)
	}
	return r
}
