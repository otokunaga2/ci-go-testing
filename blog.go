package main

type Article struct {
	Title string `json:"title"`
	Body string `json:"body"`
}

type Blog struct{
	Articles []Article
}

func New() *Blog{
	return &Blog{}
}

func (b *Blog) SaveArticle(artcle Article){
	b.Articles = append(b.Articles, artcle)
}

func (b *Blog) FetchAll() []Article{
	return b.Articles
}