## --- Day 13: Transparent Origami ---

In this puzzle folding is not very complicated but there is enough logistics to take care of that can take quite awhile. In the part two, the answer is a word, and while it is not strictly _required_ to [OCR](https://en.wikipedia.org/wiki/Optical_character_recognition) it from the dots I wanted to have that too. I did not consider this part of the challenge so I googled about Advent Of Code font, and OCR and I found this [repository](https://github.com/mstksg/advent-of-code-ocr). The main thing I used from there are the actual letters, coding up the OCR itself was not that
difficult.

The input gives us an array of points, so this is the data structure that I used. The fold functions manipulate the points coordinates and remove duplicates if any. For OCR purposes the finial array of points needs to be plotted onto an "image".
