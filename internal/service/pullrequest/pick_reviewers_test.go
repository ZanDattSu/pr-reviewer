package pullrequest

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPickReviewers_AllEdgeCases(t *testing.T) {
	tests := []struct {
		name string

		members []string

		reviewersCount int

		expectedLen int

		expectCopy bool
	}{
		{
			name: "если reviewersCount == 0 должен вернуть пустой список",

			members: []string{"a", "b", "c"},

			reviewersCount: 0,

			expectedLen: 0,

			expectCopy: false,
		},

		{
			name: "если reviewersCount >= len вернуть копию массива",

			members: []string{"a", "b"},

			reviewersCount: 5,

			expectedLen: 2,

			expectCopy: true,
		},

		{
			name: "если ровно reviewersCount == len вернуть копию массива",

			members: []string{"a", "b", "c"},

			reviewersCount: 3,

			expectedLen: 3,

			expectCopy: true,
		},

		{
			name: "обычный случай: выбираем reviewersCount элементов",

			members: []string{"a", "b", "c", "d", "e"},

			reviewersCount: 2,

			expectedLen: 2,

			expectCopy: false,
		},

		{
			name: "работает с одним элементом",

			members: []string{"x"},

			reviewersCount: 1,

			expectedLen: 1,

			expectCopy: true,
		},

		{
			name: "работает с пустым массивом",

			members: []string{},

			reviewersCount: 3,

			expectedLen: 0,

			expectCopy: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orig := append([]string{}, tt.members...)

			got := PickReviewers(tt.members, tt.reviewersCount)

			require.Equal(t, tt.expectedLen, len(got))

			if tt.expectCopy {
				require.Equal(t, orig, tt.members)
			} else {
				require.Len(t, tt.members, len(orig))
			}

			for _, r := range got {
				require.Contains(t, orig, r)
			}

			// уникальность (алгоритм не должен выбирать одинаковых людей)

			seen := map[string]bool{}

			for _, r := range got {

				require.False(t, seen[r], "duplicate reviewer found")

				seen[r] = true

			}
		})
	}
}

// Вероятностный тест — проверяем, что алгоритм НЕ всегда возвращает одинаковый результат

func TestPickReviewers_Randomness_MultipleRuns(t *testing.T) {
	members := []string{"a", "b", "c", "d", "e"}

	res1 := PickReviewers(append([]string{}, members...), 3)

	res2 := PickReviewers(append([]string{}, members...), 3)

	require.NotEqual(t, res1, res2, "expected different selections for different seeds")
}
