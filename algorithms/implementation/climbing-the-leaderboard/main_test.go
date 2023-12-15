package main

import (
	"reflect"
	"testing"
)

func TestClimbingLeaderboard(t *testing.T) {

	t.Run("Climbing Leaderboard Sample", func(t *testing.T) {

		ranked := []int32{100, 90, 90, 80}
		player := []int32{70, 80, 105}

		actual := climbingLeaderboard(ranked, player)

		expected := []int32{4, 3, 1}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

	t.Run("Climbing Leaderboard Sample 1", func(t *testing.T) {

		ranked := []int32{100, 100, 50, 40, 40, 20, 10}
		player := []int32{5, 25, 50, 120}

		actual := climbingLeaderboard(ranked, player)

		expected := []int32{6, 4, 2, 1}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

	t.Run("Climbing Leaderboard Sample 2", func(t *testing.T) {

		ranked := []int32{100, 90, 90, 80, 75, 60}
		player := []int32{50, 65, 77, 90, 102}

		actual := climbingLeaderboard(ranked, player)

		expected := []int32{6, 5, 4, 2, 1}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

}
