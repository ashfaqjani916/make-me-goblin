package models

type Posts struct{
  id uint `bson:"id" json:"id"`
  title string `bson:"title" json:"title"`
  content string `bson:"content" json:"content"`
  authorId uint  `bson:"author_id" json:"author_id"`
} 
