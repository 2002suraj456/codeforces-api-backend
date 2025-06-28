package models

type Tag string
type SeqID int
type Rating int

type TagsMap map[Tag][]SeqID
type RatingMap map[Rating]TagsMap

type ProblemDBInterface interface {
	Init()
	InsertProblem(contestID int, index string, problemName string, rating int, tags []string)
	Query(ratingStart Rating, ratingEnd Rating, tags []Tag) []SeqID
	GetProblem(id SeqID) Problem
}

type ProblemDB struct {
	RatingMap RatingMap
	Problems  []Problem
}

// Init initializes the ProblemDB
func (p *ProblemDB) Init() {
	p.RatingMap = make(RatingMap)
	p.Problems = make([]Problem, 0)
}

// InsertProblem adds a new problem to the database
func (p *ProblemDB) InsertProblem(contestID int, index string, problemName string, rating int, tags []string) {
	nextSeqID := len(p.Problems)

	problem := Problem{
		ContestID: contestID,
		Name:      problemName,
		Rating:    rating,
		Tags:      tags,
		SeqID:     nextSeqID,
		Index:     index,
	}

	if rating == 0 || len(tags) == 0 {
		return
	}

	if _, exists := p.RatingMap[Rating(rating)]; !exists {
		p.RatingMap[Rating(rating)] = make(TagsMap)
	}

	for _, tag := range tags {
		p.RatingMap[Rating(rating)][Tag(tag)] = append(p.RatingMap[Rating(rating)][Tag(tag)], SeqID(nextSeqID))
	}

	p.Problems = append(p.Problems, problem)
}

func (p *ProblemDB) Query(ratingStart Rating, ratingEnd Rating, tags []Tag) []SeqID {
	var result []SeqID

	for rating, val := range p.RatingMap {
		if ratingStart <= rating && rating <= ratingEnd {
			problemSeqIDs := getProblems(val, tags)
			result = append(result, problemSeqIDs...)
		}
	}

	return result
}

func (p *ProblemDB) GetProblem(id SeqID) Problem {
	if int(id) >= len(p.Problems) {
		return Problem{}
	}
	return p.Problems[id]
}

func getProblems(tagMap TagsMap, tags []Tag) []SeqID {
	var result []SeqID

	if len(tags) == 0 {
		for _, problems := range tagMap {
			result = append(result, problems...)
		}
		return result
	}

	if len(tags) > len(tagMap) {
		return result
	}

	if len(tags) > 0 {
		if problems, exists := tagMap[tags[0]]; exists {
			result = problems
		}
	}

	for i := 1; i < len(tags); i++ {
		if problems, exists := tagMap[tags[i]]; exists {
			result = intersectSeqIDs(result, problems)
		} else {
			return []SeqID{}
		}
	}

	return result
}

func intersectSeqIDs(first []SeqID, second []SeqID) []SeqID {
	if len(first) == 0 {
		return []SeqID{}
	}
	if len(second) == 0 {
		return []SeqID{}
	}

	var (
		firstLen  = len(first)
		secondLen = len(second)
		f         = 0
		s         = 0
	)

	var result []SeqID

	for f < firstLen && s < secondLen {
		if first[f] < second[s] {
			f++
		} else if first[f] > second[s] {
			s++
		} else {
			result = append(result, first[f])
			f++
			s++
		}
	}

	return result
}
