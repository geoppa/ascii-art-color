# ASCII Art Generator
A simple Go-based CLI tool that converts text strings into stylized ASCII art using a banner template (defaulting to standard.txt).  
## Features

Converts alphanumeric strings and symbols into 8-line tall ASCII art.  
Handles newline characters (\n) within the input string.  
Validates command-line arguments (supports single string input).  
Reads character maps from external .txt files.  

## Prerequisites

Go: Ensure you have Go installed (1.16+ recommended).  
Banner File: A file named standard.txt must be present in the root directory. This file should contain ASCII character templates, where each character is represented by 8 lines of art followed by 1 empty line.

## Installation

Clone this repository or save the code to main.go.  
Ensure standard.txt is in the same folder.

## Usage
Run the program using go run followed by the string you want to convert:  
```
go run . "Hello"
```

**Handling Newlines**  
You can include literal \n in your string to print art across multiple lines:  
```
go run . "Hello\nWorld"
```

**Argument Behavior**  
No arguments: Prints "No Arguments...".  
Empty string: Exits silently.  
Multiple arguments: Notifies the user and processes only the first argument.  

How It Works

    Input Parsing: The program takes the first command-line argument and splits it by the \n delimiter.
    Character Mapping: It calculates the starting line of a character in the banner file using the formula:
    startLine := int(char - 32) * 9 + 1
    Rendering: It iterates through 8 vertical layers, printing the corresponding line for each character in the input string before moving to the next layer.
