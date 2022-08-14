package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(Hello("Trey"))
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Returns a greeting for the named person
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Name can't be blank")
	}

	message := fmt.Sprintf(randomGreeting(), name)
	return message, nil
}

// Returns a map with each passed in name with a greeting
func Hellos(names []string) (map[string]string, error) {
	//a map to associate names with people
	messages := make(map[string]string)

	//loop through names to get a greeting for the name
	for _, name := range names {
		message, err := Hello(name)
		//If error is not equal to NOT ERROR return an error
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func randomGreeting() string {
	greetings := []string{
		"Hello, %v. Welcome!",
		"Aye %v. What's good wit ya boi",
		"Hail %v, well Met",
	}

	randomNum := rand.Intn(len(greetings))

	return greetings[randomNum]
}
