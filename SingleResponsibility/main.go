package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//Notepad structure
type Notepad struct {
	notes []string //slice of strings to to keep notes
}

//AddNote get string and add it to notes
func (n *Notepad) AddNote(text string) {
	n.notes = append(n.notes, text)
}

//RemoveNote get index and remove it from notes
func (n *Notepad) RemoveNote(index int) {
	n.notes =
		append(n.notes[:index], n.notes[index+1:]...)
}

//ToString join all notes and return them as one string
func (n *Notepad) ToString() string {
	return strings.Join(n.notes, "\n")
}

//function to save notes to flile
func (n *Notepad) SaveToFile(path string) {
	_ = ioutil.WriteFile(path,
		[]byte(n.ToString()), 0644)
}

//function to save notes to file
func SaveToFile(input string, path string) {
	_ = ioutil.WriteFile(path, []byte(input), 0644)
}

//save structure
type Save struct {
	default_path string
	//...
}

//function to save notes to flile
func (s *Save) SaveToFile(input string, path string) {
	if path == "" {
		path = s.default_path
	}
	_ = ioutil.WriteFile(path, []byte(input), 0644)
}

func main() {
	mynotes := Notepad{} //create notepad
	//adding note to notepad
	mynotes.AddNote("congratulations")
	mynotes.AddNote("Breaking the first principle was done successfully!")
	/*
		//Wrong way
			//save notes to file
			mynotes.SaveToFile("SingleResponsibility/01_Wrong_way")
			//print notes
			fmt.Println(mynotes.ToString())

			//removing failed note
			mynotes.RemoveNote(1)
	*/
	/*
		//correct way
			//save notes to file
			SaveToFile(mynotes.ToString(), "SingleResponsibility/02_correct_way")
			//print notes
			fmt.Println(mynotes.ToString())
	*/
	//better way
	//create save
	s := Save{}
	//set dedault path
	s.default_path = "SingleResponsibility/03_better_way"
	//save notes to file
	s.SaveToFile(mynotes.ToString(), "")
	//print notes
	fmt.Println(mynotes.ToString())

}
