#include "tuple.h"

#pragma once

typedef struct canvas_s {
    int width;
    int height;
    tuple* pixels;
} canvas;

canvas newCanvas(int width, int height);

tuple getPixel(canvas c, int x, int y);

void setPixel(canvas c, int x, int y, tuple col);

char* canvasToPPM(canvas c);