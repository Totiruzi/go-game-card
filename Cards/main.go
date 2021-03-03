package main

 //import "fmt"

func main() {
// dealing a hand of cards
//     cards := newDeck()

    //hand, remainingCards := deal(cards, 5)
    //hand.print()
    //remainingCards.print()

    //fmt.Println(cards.toString())
//     cards.saveToFile("my_cards")

// Printing cards from the file
//     cards := newDeckFromFile("my_cards")
//     cards.print()

// Shuffling the deck of cards
    cards := newDeck()
    cards.shuffle()
    cards.print()
}

