
# Quiz Application in Go

This is a command-line quiz application written in Go. The application reads quiz questions and answers from a CSV file, prompts the user to answer the questions, and provides a timed quiz experience.

## Features

- Reads questions and answers from a CSV file
- Timed quiz with a user-defined duration
- Command-line arguments for specifying the CSV file path and timer duration
- Concurrent handling of user input and timer expiration

## Installation
 Clone the repository or download the source code.

```bash
git clone https://github.com/Darshan016/go-quiz-app.git
cd go-quiz-app
```

## Usage

To run the quiz application, use the following command:

```bash
go run main.go
```




## How It Works

1. The application reads the CSV file.
2. It parses the questions and answers into a slice of `problem` structs.
3. It starts a timer for the duration specified by the `-t` flag.
4. It prompts the user to answer each question.
5. It checks the user's answers and keeps track of the correct answers.
6. The quiz ends when all questions are answered or when the timer expires.
7. The application prints the user's score.
