# 819. Most Common Word

## Given a string paragraph and a string array of the banned words banned, return the most frequent word that is not banned.

It is guaranteed there is at least one word that is not banned, and that the answer is unique.

The words in paragraph are case-insensitive and the answer should be returned in lowercase.

 
```
Example 1:

Input: paragraph = "Bob hit a ball, the hit BALL flew far after it was hit.", banned = ["hit"]
Output: "ball"
```
```
Example 2:

Input: paragraph = "a.", banned = []
Output: "a"
``` 

Constraints:

* 1 <= paragraph.length <= 1000
* paragraph consists of English letters, space ' ', or one of the symbols: "!?',;.".
* 0 <= banned.length <= 100
* 1 <= banned[i].length <= 10
* banned[i] consists of only lowercase English letters.