# emojipicker
`emojipicker` is a simple cli tool to search for emojies and copy them to clipboard
[![asciicast](https://asciinema.org/a/opK9uatC93Ul05OQBJZyjpLA0.png)](https://asciinema.org/a/opK9uatC93Ul05OQBJZyjpLA0)
## install
```
go install github.com/transacid/emojipicker@latest
```
## usage
just hit `emojipicker` and start typing. select with hitting `Enter`. The emoji will be copied to your clipboard and echoed on the terminal. When you want to pass the emoji to another command use the `-n` switch to supress the trailing newline.

## motivation
This is just a fun project for me to learn some go :wink:
I'm making use of these libraries:
* 	[atotto/clipboard](https://github.com/atotto/clipboard)
*	[kenshaw/emoji](https://github.com/kenshaw/emoji)
*	[ktr0731/go-fuzzyfinder](https://github.com/ktr0731/go-fuzzyfinder)
