package Map

import (
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	m := make(map[string]int)

	for _, cpdomain := range cpdomains {
		subDomains, count := parseHelper(cpdomain)
		for _, subDomain := range subDomains {
			if v, ok := m[subDomain]; ok {
				m[subDomain] = v + count
			} else {
				m[subDomain] = count
			}
		}
	}

	res := make([]string, 0)

	for k, v := range m {
		res = append(res, strconv.Itoa(v)+" "+k)
	}

	return res
}

func parseHelper(input string) ([]string, int) {
	inputs := strings.Split(input, " ")
	count, _ := strconv.Atoi(inputs[0])

	domains := strings.Split(inputs[1], ".")

	res := make([]string, 0)
	preDomain := domains[len(domains)-1]
	res = append(res, preDomain)

	if len(domains) > 1 {
		for i := len(domains) - 2; i >= 0; i-- {
			preDomain = domains[i] + "." + preDomain
			res = append(res, preDomain)
		}
	}

	return res, count
}
