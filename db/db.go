package db

import (
	"github.com/zuzuleinen/translator/values"
)

//Each language has a corresponding WordsTable
//LastID is used to make sure we have unique ids. In a real-world application a UUID library can be used
type Store struct {
	Tables map[string]WordsTable
	LastID int64
}

type WordsTable struct {
	WordsToID map[string]int64
	IDToWords map[int64]string
}

func NewStore() *Store {
	var s Store
	s.Tables = make(map[string]WordsTable)
	return &s
}

func (s *Store) Store(req values.StoreRequest) {
	if s.hasWord(req.First) && !s.hasWord(req.Second) {
		s.storeWord(req.Second, s.wordId(req.First))
	}

	if !s.hasWord(req.First) && s.hasWord(req.Second) {
		s.storeWord(req.First, s.wordId(req.Second))
	}

	//here neither exists
	s.LastID++
	id := s.LastID
	s.storeWord(req.First, id)
	s.storeWord(req.Second, id)
}

func (s *Store) Find(req values.GetRequest) string {
	if !s.hasLanguage(req.InLanguage) {
		return ""
	}

	if !s.hasWord(req.Word) {
		return ""
	}

	wordId := s.wordId(req.Word)

	return s.Tables[req.InLanguage].IDToWords[wordId]
}

func (s *Store) hasWord(w values.Word) bool {
	wordTable, hasLanguage := s.Tables[w.Lang]

	if !hasLanguage {
		return false
	}

	_, hasWord := wordTable.WordsToID[w.Value]

	return hasWord
}

func (s *Store) hasLanguage(lang string) bool {
	_, hasLanguage := s.Tables[lang]

	return hasLanguage
}

func (s *Store) wordId(w values.Word) int64 {
	wordTable, _ := s.Tables[w.Lang]

	return wordTable.WordsToID[w.Value]
}

func (s *Store) storeWord(w values.Word, id int64) int64 {
	wordTable, hasLanguage := s.Tables[w.Lang]

	if hasLanguage {
		_, hasWord := wordTable.WordsToID[w.Value]
		if !hasWord {
			wordTable.saveWord(w, id)
			return id
		}
		return s.Tables[w.Lang].WordsToID[w.Value]
	}

	wordTable = WordsTable{}
	wordTable.WordsToID = make(map[string]int64)
	wordTable.IDToWords = make(map[int64]string)
	wordTable.saveWord(w, id)

	s.Tables[w.Lang] = wordTable

	return s.LastID
}

func (w *WordsTable) saveWord(word values.Word, id int64) {
	w.WordsToID[word.Value] = id
	w.IDToWords[id] = word.Value
}
