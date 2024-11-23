package main

const (
	csvFilePath    = "dictionary.csv"
	indexPath      = "index.dat"
	changeLogPath  = "changelog.csv"
	newCsvFilePath = "new_dictionary.csv"
	newIndexPath   = "new_index.dat"
)

func main() {
	createNewIndex()

	wordLookup("Apple")
	wordLookup("Dog")

	addWord("Banana", "Yellow fruit often eaten by monkeys and humans alike")
	addWord("Elephant", "Large mammal with a trunk and grey colored")
	addWord("Funny", "A trait many people lack.")
	addWord("Dog", "Bow bow bow wow wow bark")

	wordLookup("Elephant")
	wordLookup("Orange")
	wordLookup("Dog")
}