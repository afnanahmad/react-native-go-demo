# React Native <-> Go

This is just a barebone project just to test go module integration in Android with a React Native module

To run this project

`yarn install`

`yarn android`

Once Android emulator launches click `CALL NATIVE FUNCTION`, this is launches a web server writing in GO and embedded in compiled android library.

You can then open Google Chrome on the emulator and to this link: `localhost:8080` and it should display `Hello Http` message.
