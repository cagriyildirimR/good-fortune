package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var fortunes = []string{
	"Your determination will lead to success.",
	"Good things come to those who believe.",
	"Your talents will open doors of opportunity.",
	"Stay focused on your goals; they will become a reality.",
	"Your positive outlook will attract positive people and experiences.",
	"Embrace challenges as opportunities for growth.",
	"Your hard work will be recognized and rewarded.",
	"Believe in yourself, and others will believe in you too.",
	"Every day is a chance to make a difference.",
	"Your positive energy is contagious.",
	"Success is a journey, not a destination.",
	"Your dreams are worth pursuing; never give up.",
	"Your positive attitude will overcome any obstacle.",
	"Trust in yourself; you are capable of great things.",
	"Good fortune will find you wherever you go.",
	"Your actions today will shape your tomorrow.",
	"Embrace the unknown with courage and curiosity.",
	"Your resilience will lead you to triumph.",
	"Stay true to yourself; it is the key to happiness.",
	"Your kindness will create ripples of positivity.",
	"Every step forward brings you closer to your goals.",
	"Your optimism will brighten even the darkest days.",
	"Success is not measured by material wealth but by fulfillment.",
	"Your determination will inspire those around you.",
	"Believe in your abilities, and you will achieve greatness.",
	"Your dreams are limited only by your imagination.",
	"Embrace the present moment; it holds infinite possibilities.",
	"Your efforts are never in vain; they shape your future.",
	"Good luck is on your side; seize the opportunities that come your way.",
	"Your positivity will attract abundance into your life.",
	"Success starts with a vision; envision your path to success.",
	"Your journey may have obstacles, but they will make you stronger.",
	"Believe in the power of your dreams; they have the ability to come true.",
	"Your actions have the power to create positive change.",
	"Embrace the beauty of each day; it is a gift.",
	"Your determination will overcome any adversity.",
	"Stay focused on your goals, and you will achieve them.",
	"Your positive mindset will lead to extraordinary results.",
	"Good fortune favors the bold; take calculated risks.",
	"Your perseverance will lead you to victory.",
	"Believe in yourself, and others will believe in you too.",
	"Your passion will fuel your success.",
	"Success is not final; failure is not fatal. It is the courage to continue that counts.",
	"Your positive attitude will attract endless possibilities.",
	"Your dreams are the seeds of your future reality.",
	"Embrace change; it leads to personal growth and transformation.",
	"Your journey is unique; embrace it and make it extraordinary.",
	"Stay committed to your goals, even in the face of challenges.",
	"Your kindness will have a lasting impact on the world.",
	"Good fortune flows to those who radiate positivity.",
	"Good things come to those who wait.",
	"The future looks bright.",
	"Your hard work will pay off.",
	"A pleasant surprise awaits you.",
	"Believe in yourself and all that you are.",
	"Opportunities are plentiful for those who seek them.",
	"Your creativity will lead you to success.",
	"Your kindness will be repaid in unexpected ways.",
	"Take the leap of faith; you won't regret it.",
	"A smile is your best accessory.",
	"The best is yet to come.",
	"Your positive attitude will attract great things.",
	"Trust in yourself; you have the power to achieve anything.",
	"Good fortune will find you when you least expect it.",
	"Stay focused and never give up on your dreams.",
	"Embrace change; it leads to growth and new possibilities.",
	"Success comes to those who work hard and stay persistent.",
	"Your optimism brightens the lives of those around you.",
	"Great things take time; be patient and keep going.",
	"Your dedication will open doors of opportunity.",
	"Believe in yourself, and others will believe in you too.",
	"Your determination will overcome any obstacle.",
	"Positive thoughts attract positive outcomes.",
	"Take a chance; the rewards will be worth it.",
	"Your generosity will make a difference in someone's life.",
	"Embrace challenges; they are stepping stones to success.",
	"Your future is full of abundance and prosperity.",
	"Good luck and good fortune will follow you always.",
	"Your hard work will pay off in remarkable ways.",
	"Every day is a new chance to make a difference.",
	"Your positive energy will bring you closer to your goals.",
	"Your persistence will lead you to great achievements.",
	"Success is within your reach; keep reaching higher.",
	"Your courage will open doors you never thought possible.",
	"Embrace the unknown; it holds exciting opportunities.",
	"Your dreams are within your grasp; go after them.",
	"Your efforts will be rewarded in due time.",
	"Every setback is a setup for a comeback.",
	"Believe in yourself, and others will too.",
	"Your passion will drive you to extraordinary places.",
	"Good fortune comes to those who take action.",
	"Stay true to yourself; it is the key to success.",
	"Your positive attitude will bring you success and happiness.",
	"Your determination will overcome any challenge.",
	"Trust in yourself; you have all the answers within.",
	"Your kindness will make a lasting impact on others.",
	"Seize the day; it is filled with endless possibilities.",
	"Your optimism will attract positive experiences.",
	"Your dreams are the blueprint for your success.",
	"Stay focused on your goals; they are within reach.",
	"Your hard work and dedication will pay off.",
	"Your positive mindset will lead to great achievements.",
	"Success is just around the corner; keep going.",
}

func main() {
	http.HandleFunc("/", handleFortune)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func handleFortune(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		rand.Seed(time.Now().UnixNano())
		randomIndex := rand.Intn(len(fortunes))
		fortune := fortunes[randomIndex]
		_, err := fmt.Fprintf(w, "Random Good Fortune: %s", fortune)
		if err != nil {
			return
		}
		return
	}

	html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Random Good Fortunes</title>
		</head>
		<body>
			<h1>Random Good Fortunes</h1>
			<form method="POST">
				<button type="submit">Get Fortune</button>
			</form>
		</body>
		</html>
	`

	_, err := fmt.Fprint(w, html)
	if err != nil {
		return
	}
}
