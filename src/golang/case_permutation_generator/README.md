# Case Permutation Generator

## Overview

The `CasePermutationGenerator` is a Go program designed to generate all possible permutations of a given string where each alphabetic character can appear in both uppercase and lowercase forms. Non-alphabetic characters (such as digits, spaces, or punctuation) are preserved in their original form and are not subject to case permutation.

## Features

- **Alphabetic Case Variations**: Generates permutations by altering the case (uppercase/lowercase) of alphabetic characters.
- **Non-Alphabetic Characters**: Maintains non-alphabetic characters unchanged, including digits and punctuation.
- **Recursive Approach**: Utilizes recursion to explore all possible combinations of letter cases.

## How It Works

1. **Input Handling**: The program takes a string as input.
2. **Recursive Permutation Generation**: For each character in the string:
   - If it is an alphabetic character, permutations are generated for both its lowercase and uppercase forms.
   - If it is a non-alphabetic character, it is included as-is in the permutations.
3. **Output**: The program outputs all possible permutations to the console.

## Example

Given the input string `"fB& 1"`, the program will output:
```
fb& 1
fB& 1
Fb& 1
FB& 1
```
