package main

import (
    "testing"

//     "os"
)
func TestNewDeck (t *testing.T) {
    d := newDeck()
    if len(d) != 16 {
        t.Errorf("Expected a deal of 16 cards, but got ", len(d))
    }
}

// func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
//     os.Remove("_decktesting")
//
//     deck := newDeck()
//     deck.saveToFile("_decktesting")
//
//     loadedDeck := newDeckFromFile("_decktesting")
//     if len(loadedDeck) != 16 {
//         t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck)
//     }
//
//     os.Remove("_decktesting")
// }

// func TestShuffle (t *testing.T) {
//     d := newDeck()
//     s := t.shuffle()
//
//     dHand, dRemainingCards := deal(d, 5)
//     sHand, sRemainingCards := deal(s, 5)
//
// //     if dHand[0] == sHand[0] {
// //         t.ErrorF("The deck where not shuffled")
// //     }
//
//     for i, v := range dHand {
//         if v != sHand[i]
//         t.ErrorF("The deck where not shuffled")
//     }
// }

// Checking if two arrays are thr same
// func Equal(a, b []int) bool {
//     if len(a) != len(b) {
//         return false
//     }
//     for i, v := range a {
//         if v != b[i] {
//             return false
//         }
//     }
//     return true
// }
