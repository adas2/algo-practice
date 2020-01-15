You have developed a program in a new scripting language. Of course, it requires accurate parsing in order to perform as expected, and it is very cryptic. You want to determine how many valid commands can be made out of your lines of script. To do this, you count all of the substrings that make up a valid command. Each of the valid commands will be in the following format:

The first letter is a lowercase English letter.
Next, it contains a sequence of zero or more of the following characters: lowercase English letters, digits, and colons.
Next, it contains a forward slash '/'.
Next, it contains a sequence of one or more of the following characters: lowercase English letters and digits.
Next, it contains a backward slash '\'.
Next, it contains a sequence of one or more lowercase English letters.
 

Given some string, s, we define the following:

s[i..j] is a substring consisting of all the characters in the inclusive range between index i and index j.
Two substrings, s[i[1]..j[1]] and s[i[2]..j[2]], are said to be distinct if either i[1] ≠ i[2] or j[1] ≠ j[2].
