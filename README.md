# sbee-cheat

This allows you to cheat at the NY Times Spelling Bee game by extracting the answers from the game itself. It is just a fun toy and will be trivial for the nytimes game devs to render inoperable.  

*PLEASE NOTE*: I'm not a very good `go` programmer.  This isn't a very good `go` program. But it works.
If anyone wants to figure out the right way to extract the answers dict from the HTML source, rather than a regex, I welcome a PR.

## License
Apache 2.0.  Help yourself. Not my fault if you break stuff.

## Getting
Binaries are available in the [bin](bin/) directory for Linux (x86 and Raspberry Pi), Mac and Windows.  I don't have a Windows box so the Windows binary is a guess.

## Using

```
prog/sbee-cheat/bin$ ./sbee-cheat
pangrams [talkative]
answers [talkative alike alit alive alkali allele alleviate atilt avail elate elevate elite evil kale keel kettle kill kilt kvell lake late latke latte lava lave leak leave leek levee level levitate like lilt lite little live tail tale tali talk talkie tall tallit tattle tattletale teakettle teal tell telltale tile till tilt titillate title tittle vale valet valve veal veil velvet vial vile villa vital] (65)
```

The number in parentheses (in this case, 65) is the number of words in the answers section.
Calculating the number of words in the pangrams section is left as an exercise to the reader.

## Building
This is Go, it's pretty easy. Use the Makefile:

Windows (I think):
`make windows`

Intel Mac:
`make mintel`

M1 Mac:
`make m1`

Linux (x86):
`make linux_x86`

Linux (Raspberry pi):
`make linux_rpi`

There's also `make all` and `make clean`.

I tried building it for WASM and I could get a binary out but couldn't figure out how to run it in the browser.  The build variables are `GOOS=js GOARCH=wasm`.