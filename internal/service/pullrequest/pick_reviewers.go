package pullrequest

import "math/rand"

// PickReviewers выбирает reviewersCount случайных уникальных ревьюверов
// из массива activeTeamMember.
// Частично использует алгоритм Фишера–Йетса
func PickReviewers(activeTeamMembers []string, reviewersCount int) []string {
	n := len(activeTeamMembers)
	if n <= reviewersCount {
		return copySlice(activeTeamMembers)
	}

	for i := 0; i < reviewersCount; i++ {
		j := i + rand.Intn(n-i) //nolint:gosec
		swap(activeTeamMembers, i, j)
	}

	return activeTeamMembers[:reviewersCount]
}

func copySlice(activeTeamMembers []string) []string {
	return append([]string{}, activeTeamMembers...)
}

func swap(activeTeamMembers []string, i, j int) {
	activeTeamMembers[i], activeTeamMembers[j] = activeTeamMembers[j], activeTeamMembers[i]
}
