# go-reloaded

Welcome back. Congratulations on your admission. We knew you would make it. Now it is time to get into projects.

Objectives
In this project you will use some of your old functions made in your old repository. You will use them with the objective of making a simple text completion/editing/auto-correction tool.

One more detail. This time the project will be corrected by auditors. The auditors will be other students and you will be an auditor as well.

We advise you to create your own tests for yourself and for when you will correct your auditees.

Introduction
Your project must be written in Go.
The code should respect the good practices.
It is recommended to have test files for unit testing.
The tool you are about to build will receive as arguments the name of a file containing a text that needs some modifications (the input) and the name of the file the modified text should be placed in (the output). Next is a list of possible modifications that your program should execute:

Every instance of (hex) should replace the word before with the decimal version of the word (in this case the word will always be a hexadecimal number). (Ex: "1E (hex) files were added" -> "30 files were added")

Every instance of (bin) should replace the word before with the decimal version of the word (in this case the word will always be a binary number). (Ex: "It has been 10 (bin) years" -> "It has been 2 years")

Every instance of (up) converts the word before with the Uppercase version of it. (Ex: "Ready, set, go (up) !" -> "Ready, set, GO !")

Every instance of (low) converts the word before with the Lowercase version of it. (Ex: "I should stop SHOUTING (low)" -> "I should stop shouting")

Every instance of (cap) converts the word before with the capitalized version of it. (Ex: "Welcome to the Brooklyn bridge (cap)" -> "Welcome to the Brooklyn Bridge")

For (low), (up), (cap) if a number appears next to it, like so: (low, <number>) it turns the previously specified number of words in lowercase, uppercase or capitalized accordingly. (Ex: "This is so exciting (up, 2)" -> "This is SO EXCITING")
Every instance of the punctuations ., ,, !, ?, : and ; should be close to the previous word and with space apart from the next one. (Ex: "I was sitting over there ,and then BAMM !!" -> "I was sitting over there, and then BAMM!!").

Except if there are groups of punctuation like: ... or !?. In this case the program should format the text as in the following example: "I was thinking ... You were right" -> "I was thinking... You were right".
The punctuation mark ' will always be found with another instance of it and they should be placed to the right and left of the word in the middle of them, without any spaces. (Ex: "I am exactly how they describe me: ' awesome '" -> "I am exactly how they describe me: 'awesome'")

If there are more than one word between the two ' ' marks, the program should place the marks next to the corresponding words (Ex: "As Elton John said: ' I am the most well-known homosexual in the world '" -> "As Elton John said: 'I am the most well-known homosexual in the world'")
Every instance of a should be turned into an if the next word begins with a vowel or a h. (Ex: "There it was. A amazing rock!" -> "There it was. An amazing rock!").
