package models

import "testing"

func TestGetALLArticles(t *testing.T) {
	alist := GetAllArticles()

	if len(alist) != len(ArticleList) {
		t.Fail()
	}

	for i, v := range alist {
		if v.Content != ArticleList[i].Content ||
			v.ID != ArticleList[i].ID ||
			v.Title != ArticleList[i].Title {

			t.Fail()
			break
		}
	}
}
