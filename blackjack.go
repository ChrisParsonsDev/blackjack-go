package main

import (
    "fmt"
    "strings"
    "math/rand"
    "strconv"
    "time"
)

const (
  dealerThreshold = 16 //Determines when the dealer will stick
  bustThreshold = 21 //Determines when the player will bust
  randomThreshold = 9 //Determines max values for random number generation
)

//Global Variable Declarations
var playerName string
var anotherHand string
var playerScore int
var dealerScore int
var userWonHands int
var dealerWonHands int

//Random number generator
func newCardGenerator() (int){
  randomSeed := rand.NewSource(time.Now().UnixNano())
  seededRandomGen := rand.New(randomSeed)
  return (seededRandomGen.Intn(randomThreshold) + 1)
}

//Game Methods
func determineHandWinner(computerScore int, userScore int)  (bool){
  //returns true if the player wins
  if (computerScore < userScore){
    if (userScore < 22){
      return true
    } else if (computerScore == userScore){
      //tie dealer wins
      return false
    } else {
      //user was bust
      return false
    }
  } else if (computerScore > userScore){
    if (dealerScore < 22){
      return false
    } else {
      return true
    }
  }
  //If there is an error, dealer wins
  return false
}

func dealerTurn()  {
  for dealerRequestCard := true; dealerRequestCard; dealerRequestCard = (dealerScore < dealerThreshold) {
    currentCard := newCardGenerator()
    dealerScore += currentCard
    fmt.Print("The dealer has "+strconv.Itoa(dealerScore)+"\n")
  }
}


func playerTurn(){
  var twist = "t"
  for userRequestCard := true; userRequestCard; userRequestCard = (!strings.EqualFold(twist, "s")) {
    currentCard := newCardGenerator()
    playerScore += currentCard
    if (playerScore > 21){
      fmt.Print(playerName+" you've gone bust with "+ strconv.Itoa(playerScore)+"\n")
      twist = "s"
    }else{
      fmt.Print("You have "+strconv.Itoa(playerScore)+" would you like to (s)tick or (t)wist\n")
      fmt.Scan(&twist)
    }
  }
}

func playGame()  {
  playerTurn()
  dealerTurn()
  if(determineHandWinner(dealerScore, playerScore)){
    userWonHands += 1
    fmt.Print("Congratulations "+playerName+" you won the hand with "+strconv.Itoa(playerScore)+" vs the dealer's "+strconv.Itoa(dealerScore)+"\n")
  } else {
    dealerWonHands +=1
    fmt.Print("Unlucky "+playerName+" the dealer won the hand with "+strconv.Itoa(dealerScore)+" you had "+strconv.Itoa(playerScore)+"\n")
  }
  fmt.Print("The scores are: \n"+playerName+" "+strconv.Itoa(userWonHands)+" wins\n"+" Dealer "+strconv.Itoa(dealerWonHands)+" wins\n")
  continueGame()
}

func exitGame()  {
  var playerQuit string
  fmt.Print("Are you sure you would like to exit? (Y/N)\n")
  fmt.Scan(&playerQuit)
  if strings.EqualFold(playerQuit,"Y") {
    fmt.Print("Bye, thanks for playing\n")
  }else if strings.EqualFold(playerQuit,"N"){
    playBlackjack()
  } else {
    fmt.Print("Sorry, we didn't recognise that, please enter 'Y' or 'N'\n")
    exitGame()
  }
}

func continueGame()  {
  fmt.Print("Hello "+playerName+" are you ready to play Blackjack? (Y/N):\n ")
  fmt.Scan(&anotherHand)
  if strings.EqualFold(anotherHand,"Y") {
      playerScore = 0
      dealerScore = 0
      playGame()
  } else if strings.EqualFold(anotherHand,"N"){
    exitGame()
  } else {
      fmt.Print("Sorry, we didn't recognise that, please enter 'Y' or 'N'\n")
      continueGame()
  }
}

func showRules(){
  var userRuleChoice string
  fmt.Print(playerName + " Would you like to view the rules? (Y/N)\n")
  fmt.Scan(&userRuleChoice)
  if strings.EqualFold(userRuleChoice,"Y"){
    fmt.Printf("LMGTFY\n")
    showRules()
  }else if strings.EqualFold(userRuleChoice,"N"){
    fmt.Print("Ok, let's begin "+playerName+"\n")
    continueGame()
  }else{
    fmt.Print("Sorry, we didn't recognise that, please enter 'Y' or 'N'\n")
    showRules()
  }
}

func playBlackjack()  {
  showRules()
}


//Main Execution
func main()  {
  fmt.Printf("Please enter your name:\n")
  fmt.Scan(&playerName)
  playBlackjack()
}
