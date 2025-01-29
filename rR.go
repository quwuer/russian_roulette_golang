package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
	"bufio"
	"strings"
)

var revolver int = 0
var c string

func createOp(){
	filePath := "computer.txt"
	content := "Hello, I am Alex, I'm your opponent. Why are we playing this game? Well... nevermind. Let's have some fun ;)"

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Opponent already exists, skipping delivering.")
		time.Sleep(2 * time.Second)
	} else {
		defer file.Close()
		_, err = file.WriteString(content)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		fmt.Println("Your opponent has been delivered.")
		time.Sleep(2 * time.Second)
	}
}

func chance(revolver int, c *string){	
	switch revolver {
	case 1:
		*c = "1 in 6"
	case 2:
		*c = "2 in 6"
	case 3:
		*c = "3 in 6"
	case 4:
		*c = "4 in 6"
	case 5:
		*c = "50/50"
	case 6:
		*c = "No way to escape the death"
	}
}

func respond(){
	
	random := rand.Intn(3)+1
	
	switch random{
	case 1:
		fmt.Println("Phew lucky me")
	case 2:
		fmt.Println("Hehe...")
	case 3:
		fmt.Println("That was pretty close")
	}

}

func start(){
	fmt.Println("Hi.")
	time.Sleep(2 * time.Second)
	fmt.Println("You are gonna play russian roulette.")
	time.Sleep(2 * time.Second)
	fmt.Println("I loaded the revolver with 1 bullet, you both will shoot in turns. Until the first death.")
	time.Sleep(3 * time.Second)
	fmt.Println("Start?(Y/N)")
}


func main() {
	
	createOp()
	rand.Seed(time.Now().UnixNano())
	xBullet := rand.Intn(6) +1
	turn := rand.Intn(2) +1
	ans := ""
	
	start()

	for{
		fmt.Scan(&ans)
		if ans == "N" || ans == "n"{
			fmt.Println("How sad.")
			return
		}else if ans == "Y" || ans == "y" {
			for{
				if turn == 2{ //programs turn
					time.Sleep(2 * time.Second)
					fmt.Println("My turn.")
					
					revolver++
					chance(revolver, &c)
					time.Sleep(2 * time.Second)
					fmt.Printf("Chances to die: %s\n", c)
					time.Sleep(2 * time.Second)
					if revolver != xBullet{
						fmt.Println("Chck")
					}
					time.Sleep(2 * time.Second)
					if revolver == xBullet{
						fmt.Println("Bang!")
						
						pcLife := "computer.txt"
						err :=  os.Remove(pcLife)
						if err != nil {
							fmt.Println("Error to killing Alex", err)
						} else {
							fmt.Println("Alex died. Thank you for game! See you later...")
						}

						return
					}else{
						respond()
						turn = 1
					}
				}else {
					
					time.Sleep(2 * time.Second)
					fmt.Println("Your turn.")
					revolver++
					chance(revolver, &c)
					time.Sleep(2 * time.Second)
					fmt.Printf("Chances to die: %s\n", c)
					time.Sleep(2 * time.Second)
					fmt.Println("Wish to continue?(Y/N)")
					fmt.Scan(&ans)
					time.Sleep(2 * time.Second)
					
					if ans == "Y" || ans == "y"{
						if revolver != xBullet{
							fmt.Println("Chck")
							time.Sleep(2 * time.Second)
						}
						
						if revolver == xBullet{
							fmt.Println("Bang! You are dead")
							return
						}else{
							fmt.Println("You are alive...")
							turn = 2
						}
					} else if ans =="N" || ans =="n"{
						fmt.Println("...")
						time.Sleep(2 * time.Second)
						fmt.Println("Okay coward. Just write 'I am loser' and we are done.")
						
						reader := bufio.NewReader(os.Stdin)
						reader.ReadString('\n')
						input, _ := reader.ReadString('\n') 
						input = strings.TrimSpace(input)

						if input == "I am loser" || input == "i am loser"{
							fmt.Println("Goodbye.")
							return
						} else {
							time.Sleep(2 * time.Second)
							fmt.Println("I see... ")
							revolver = revolver -1
						}
						
					}else{
						fmt.Println("Invalid input")
					}
				}
			}
		} else {
			fmt.Println("Invalid Input")
		}
	}
}
