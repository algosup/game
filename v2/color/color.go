package color

import col "image/color"

type Color col.Color

var Black = col.Black
var LightBlue = col.RGBA{0, 0xF0, 0xF0, 0xFF}
var Yellow = col.RGBA{240, 240, 0, 255}
var Red = col.RGBA{240, 0, 0, 255}
var Purple = col.RGBA{160, 0, 240, 255}
var Orange = col.RGBA{240, 160, 0, 255}
var Blue = col.RGBA{0, 0, 240, 255}
var Green = col.RGBA{0, 240, 0, 255}
var White = col.White
