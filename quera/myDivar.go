package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var users []string
var posts = make(map[int][]string)
var favorites = make(map[int][]string)

func isRegister(name string) string {
	if contain(users, name) {
		return "invalid username"
	} else {
		return "registered successfully"
	}
}

func contain(slice []string, value string) bool {
	result := false
	for _, item := range slice {
		if item == value {
			result = true
		}
	}
	return result
}

func findIndex(slice []string, item string) int {
	var result int
	for index, value := range slice {
		if value == item {
			result = index
		}
	}

	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	linesCount, _ := strconv.Atoi(scanner.Text())

	var inputs []string
	//users := []string{}

	for i := 0; i < linesCount; i++ {
		scanner.Scan()
		line := scanner.Text()

		inputs = append(inputs, line)
	}

	for _, input := range inputs {
		items := strings.Split(input, " ")
		switch items[0] {
		case "register":
			status := isRegister(items[1])
			if status == "registered successfully" {
				users = append(users, items[1])
			}
			fmt.Println(status)
		case "add_advertise":
			status := isRegister(items[1])
			if status != "invalid username" {
				fmt.Println("invalid username")
			} else {
				index := findIndex(users, items[1])
				userPosts := posts[index]
				if contain(userPosts, items[2]) {
					fmt.Println("invalid title")
				} else {
					posts[index] = append(posts[index], items[2])
					fmt.Println("posted successfully")
				}
			}
		case "rem_advertise":
			status := isRegister(items[1])
			if status != "invalid username" {
				fmt.Println("invalid username")
			} else {
				index := findIndex(users, items[1])
				userPosts := posts[index]
				postIndex := findIndex(posts[index], items[2])
				if contain(userPosts, items[2]) {
					posts[index] = append(posts[index][:postIndex], posts[index][postIndex+1:]...)
					fmt.Println("removed successfully")
				} else {
					isExist := false
					for _, usersPost := range posts {
						if contain(usersPost, items[2]) {
							isExist = true
						}
					}

					if isExist {
						fmt.Println("access denied")
					} else {
						fmt.Println("invalid title")
					}
				}
			}
		case "list_my_advertises":
			status := isRegister(items[1])
			if status != "invalid username" {
				fmt.Println("invalid username")
			} else {
				if len(items) > 2 {
					fmt.Println("")
				} else {
					index := findIndex(users, items[1])
					userPosts := posts[index]
					for _, post := range userPosts {
						fmt.Printf("%s", post)
					}
					fmt.Println("")
				}
			}

		case "add_favorite":
			status := isRegister(items[1])
			if status != "invalid username" {
				fmt.Println("invalid username")
			} else {
				index := findIndex(users, items[1])
				userFavs := favorites[index]

				if contain(userFavs, items[2]) {
					fmt.Println("already favorite")
				} else {
					isExist := false
					for _, usersFav := range favorites {
						if contain(usersFav, items[2]) {
							isExist = true
						}
					}

					if isExist {
						favorites[index] = append(favorites[index], items[2])
						fmt.Println("added successfully")
					} else {
						fmt.Println("invalid title")
					}
				}
			}

		default:
			fmt.Println("")
		}
	}
}
