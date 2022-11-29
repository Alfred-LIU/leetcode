package UnionFind

import "sort"

/*
Given a list of accounts where each element accounts[i] is a list of strings, where the first element accounts[i][0] is a name, and the rest of the elements are emails representing emails of the account.

Now, we would like to merge these accounts. Two accounts definitely belong to the same person if there is some common email to both accounts. Note that even if two accounts have the same name, they may belong to different people as people could have the same name. A person can have any number of accounts initially, but all of their accounts definitely have the same name.

After merging the accounts, return the accounts in the following format: the first element of each account is the name, and the rest of the elements are emails in sorted order. The accounts themselves can be returned in any order.
*/

func accountsMerge(accounts [][]string) [][]string {
	uf := newUnionFind(len(accounts))

	emails := map[string]int{}
	for i, account := range accounts {
		for j := 1; j < len(account); j++ {
			email := account[j]
			if j, ok := emails[email]; ok {
				uf.union(i, j)
			} else {
				emails[email] = i
			}
		}
	}

	mergedAccounts := map[int][]string{}
	for email, account := range emails {
		p := uf.find(account)
		mergedAccounts[p] = append(mergedAccounts[p], email)
	}

	res := [][]string{}
	for user, account := range mergedAccounts {
		sort.Slice(account, func(i, j int) bool {
			return account[i] < account[j]
		})
		account = append([]string{accounts[user][0]}, account...)
		res = append(res, account)
	}

	return res
}

type UnionFind struct {
	Uids  []int
	Ranks []int
}

func newUnionFind(l int) *UnionFind {
	uids := make([]int, l)
	ranks := make([]int, l)

	for i := range uids {
		uids[i] = i
		ranks[i] = 1
	}
	return &UnionFind{
		Uids:  uids,
		Ranks: ranks,
	}
}

func (uf *UnionFind) find(i int) int {
	for uf.Uids[i] != i {
		i = uf.Uids[i]
	}
	return i
}

func (uf *UnionFind) union(i, j int) {
	pI := uf.find(i)
	pJ := uf.find(j)

	if uf.Ranks[pI] > uf.Ranks[pJ] {
		uf.Uids[pJ] = pI
		uf.Ranks[pI] += uf.Ranks[pJ]
	} else {
		uf.Uids[pI] = pJ
		uf.Ranks[pJ] += uf.Ranks[pI]
	}
}
