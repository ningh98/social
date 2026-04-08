package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/ningh98/social/internal/store"
)

var usernames = []string{
	"alice", "bob", "charlie", "dave", "eve",
	"frank", "grace", "heidi", "ivan", "judy",
	"karl", "laura", "mallory", "nancy", "oscar",
	"peggy", "quentin", "rachel", "sam", "trent",
	"ursula", "victor", "wendy", "xavier", "yvonne",
	"zack", "aaron", "bella", "cody", "diana",
	"edward", "fiona", "george", "hannah", "ian",
	"jasmine", "kevin", "linda", "mike", "nina",
	"oliver", "paula", "quincy", "robin", "steve",
	"tina", "uma", "vincent", "will", "zoe",
}

var titles = []string{
	"Getting Started with Go",
	"Understanding Pointers",
	"REST APIs in Go",
	"Intro to Microservices",
	"Docker Basics",
	"Clean Code Tips",
	"Debugging Like a Pro",
	"Concurrency in Go",
	"Error Handling Patterns",
	"Building Your First API",
	"Go vs Python",
	"Structs and Interfaces",
	"Testing in Go",
	"Database Connections 101",
	"Writing Better Functions",
	"Intro to System Design",
	"Logging Best Practices",
	"Handling JSON in Go",
	"Optimizing Performance",
	"Common Go Mistakes",
}

var content = []string{
	"Just started learning Go and it feels really clean so far.",
	"Spent the day debugging a weird bug, finally figured it out.",
	"Trying to build a REST API from scratch this week.",
	"Reading about system design and scalability patterns.",
	"Today I learned how pointers actually work in Go.",
	"Working on improving my problem solving skills with LeetCode.",
	"Exploring Docker and how to containerize my app.",
	"Struggling a bit with concurrency but getting there.",
	"Just deployed my first backend service, feels great.",
	"Learning how to write cleaner and more maintainable code.",
	"Experimenting with different database schemas.",
	"Trying to understand how caching improves performance.",
	"Built a small project to practice full stack development.",
	"Working on error handling patterns in Go.",
	"Looking into how microservices communicate with each other.",
	"Refactoring my old code to make it more readable.",
	"Learning about middleware and request handling.",
	"Trying to optimize response time for my API.",
	"Reading blogs about backend engineering best practices.",
	"Building something small every day to stay consistent.",
}

var tags = []string{
	"go", "backend", "api", "database", "performance",
	"docker", "devops", "testing", "debugging", "tools",
	"microservices", "architecture", "concurrency", "cloud", "aws",
	"security", "auth", "scalability", "caching", "redis",
}

var comments = []string{
	"Great post, really helped me understand this topic.",
	"I was confused about this before, thanks for the clear explanation.",
	"This is exactly what I was looking for.",
	"Nice work, looking forward to more content like this.",
	"I tried this approach and it worked perfectly.",
	"Can you explain more about the last part?",
	"This saved me a lot of time, appreciate it.",
	"Interesting perspective, I hadn’t thought about it this way.",
	"I think there might be a small bug in your example.",
	"Very well written and easy to follow.",
	"Thanks for sharing your experience.",
	"I’m going to try this in my own project.",
	"Do you have any resources for learning more about this?",
	"This is super helpful for beginners.",
	"I like how you broke down the problem step by step.",
	"Not sure I fully understand, but I’m getting there.",
	"This clarified a lot of things for me.",
	"Awesome explanation, keep it up!",
	"I ran into an issue when implementing this, any advice?",
	"Simple and effective, thanks!",
}

func Seed(store store.Storage){
	ctx := context.Background()

	users := generateUsers(100)
	for _, user := range users{
		if err := store.Users.Create(ctx, user); err != nil{
			log.Printf("Error creating user:", err)
			return 
		}
	}

	posts := generatePosts(200, users)
	for _, post := range posts{
		if err := store.Posts.Create(ctx, post); err != nil{
			log.Printf("Error creating post:", err)
			return 
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments{
		if err := store.Comments.Create(ctx, comment); err != nil{
			log.Printf("Error creating comment:", err)
			return 
		}
	}

	log.Printf("Seeding complete")

	
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d",i),
			Email:  usernames[i%len(usernames)] + fmt.Sprintf("%d",i) + "@example.com",
			Password: "123123",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post{
	posts := make([]*store.Post,num)
	for i := 0; i < num; i++{
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID: user.ID,
			Title: titles[rand.Intn(len(titles))],
			Content: content[rand.Intn(len(content))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}
	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment{
	cms := make([]*store.Comment, num)
	for i:= 0; i < num; i++{
		cms[i] = &store.Comment{
			PostID: posts[rand.Intn(len(posts))].ID,
			UserID: users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}
	return cms
}