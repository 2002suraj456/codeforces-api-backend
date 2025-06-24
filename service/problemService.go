package service

import "cf.practice.com/models"

/***

Implementation for storing the problem in-memory and for quering the in-memory db.

-- DB structure.

- For all the problems from codeforces just return the cached result from the API.

- For getting a particular rating or a range of rating.
	- First maybe store the problems based on the ratings.
	- Then store them seperately based on the tags.

		rating_map[rating] -> list of tags
		list of tags -> problems

		tags_map map[tag][]problem

		rating_map map[rating] []tags_maps



**/

type Problem struct {
	Id        int    `json:"-"`
	ContestID int    `json:"contestId"`
	Index     string `json:"index"`
	Name      string `json:"name"`
	Rating    int    `json:"rating"`
}

type TagsMap map[string][]int
type RatingMap map[int]TagsMap

type ProblemDB struct {
	tags_map   TagsMap
	rating_map RatingMap
	problems   []Problem
}

var problemdb ProblemDB

func init() {
	// Initialize the maps
	problemdb.tags_map = make(TagsMap)
	problemdb.rating_map = make(RatingMap)

}

func InsertProblems(cfProblems []models.Problem) {
	for idx, problem := range cfProblems {
		if problem.Rating != 0 && len(problem.Tags) > 0 {

			if _, exists := problemdb.rating_map[problem.Rating]; !exists {
				problemdb.rating_map[problem.Rating] = make(TagsMap)
			}

			newproblem := Problem{
				Id:        idx,
				ContestID: problem.ContestID,
				Index:     problem.Index,
				Name:      problem.Name,
				Rating:    problem.Rating,
			}

			problemdb.problems = append(problemdb.problems, newproblem)

			for _, problemTag := range problem.Tags {
				problemdb.rating_map[problem.Rating][problemTag] = append(problemdb.rating_map[problem.Rating][problemTag], newproblem.Id)
			}

		}
	}
}

func GetProblems(ratingStart, ratingEnd int, tags []string) []int {
	var result []int
	for rating, val := range problemdb.rating_map {
		if ratingStart <= rating && rating <= ratingEnd {
			problems := getProblems(val, tags)
			result = append(result, problems...)
		}
	}

	return result
}

func getProblems(tagmap TagsMap, tags []string) []int {
	var result []int

	if len(tags) > len(tagmap) {
		return result
	}

	for _, tag := range tags {
		problems := tagmap[tag]
		result = intersect(result, problems)
	}

	return result
}

func intersect(first []int, second []int) []int {
	if len(first) == 0 {
		return second
	}

	var result []int

	var firstlen = len(first)
	var secondlen = len(second)

	var (
		f = 0
		s = 0
	)

	for f < firstlen && s < secondlen {
		if problemdb.problems[first[f]].Id < problemdb.problems[second[s]].Id {
			f++
		} else if problemdb.problems[first[f]].Id > problemdb.problems[second[s]].Id {
			s++
		} else {
			result = append(result, problemdb.problems[first[f]].Id)
			f++
			s++
		}
	}

	return result

}
