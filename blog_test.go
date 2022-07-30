package main

import "testing"


func TestSaveArticle(t *testing.T){
		blog := New()
		testedTitle := "My title"
		blog.SaveArticle(Article{testedTitle, "My Post Body"})
		if blog.Articles[0].Title != testedTitle {
				t.Errorf("Test was not added successfully")
		}
}

func TestFetchAllArticles(t *testing.T){
		blog := New()
		blog.SaveArticle(Article{"My title", "My Post Body"})
		articles := blog.FetchAll()

		if len(articles) != 1 {
				t.Errorf("Failed to fetch artciles expected 1 but %d", len(articles))
		}

}
