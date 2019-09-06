# Vending Machine

A TDD execercise.
Tasks are [here](https://gist.github.com/yattom/884741ecbd3c660fb393b2d7b116b4b2).

I stopped this execercise halfway, because I found that my TDD was a bad way and
I want to restart TDD from the beginning.

Why was my TDD bad?\
Because, I wrote many combination tests.\
Tests got heavy very quickly, and time which I took to add a new function
got slow likewise.

Why did I write combination tests?\
Because, I aimed at blackbox test extremely.\
If I've checked objects' encapsulated state in tests,
I could have easily got confident that
all effects of a method are tested sufficiently.
I checked effects of a method through another method's running result.
This is combination test, and it is necessary to test many situations
in order to be confident.
