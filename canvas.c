#include "canvas.h"

#include <malloc.h>
#include <math.h>
#include <stdio.h>
#include <string.h>

#include "tuple.h"

#define MIN(X, Y) (((X) < (Y)) ? (X) : (Y))
#define MAX(X, Y) (((X) < (Y)) ? (Y) : (X))

const int PPM_BUFFER_SIZE = 32 * 1024 * 1024;

canvas newCanvas(int width, int height) {
    tuple* pixels = malloc(width * height * sizeof(tuple));
    canvas c = {width, height, pixels};

    for (int x = 0; x < width; x++) {
        for (int y = 0; y < height; y++) {
            c.pixels[x * height + y] = newColor(0, 0, 0);
        }
    }

    return c;
}

tuple getPixel(canvas c, int x, int y) {
    tuple t = c.pixels[x * c.height + y];
    return t;
}

void setPixel(canvas c, int x, int y, tuple col) {
    c.pixels[x * c.height + y] = col;
}

tuple scale(tuple t) {
    t = multiplyTuple(t, 255);
    t.x = MIN(255, MAX(0, t.x));
    t.y = MIN(255, MAX(0, t.y));
    t.z = MIN(255, MAX(0, t.z));
    return t;
}

// TODO: using strlen and strcat is inefficient, I should keep my own accounting
char* canvasToPPM(canvas c) {
    char* buf = malloc(PPM_BUFFER_SIZE * sizeof(char));
    buf[0] = '\0';

    strcat_s(buf, PPM_BUFFER_SIZE, "P3\n");
    sprintf(buf + strlen(buf), "%d %d\n", c.width, c.height);
    strcat_s(buf, PPM_BUFFER_SIZE, "255\n");

    tuple p;
    for (int y = 0; y < c.height; y++) {
        for (int x = 0; x < c.width - 1; x++) {
            p = scale(getPixel(c, x, y));
            sprintf(buf + strlen(buf), "%d %d %d ", (int)p.x, (int)p.y,
                    (int)p.z);
        }
        p = scale(getPixel(c, c.width - 1, y));
        sprintf(buf + strlen(buf), "%d %d %d\n", (int)p.x, (int)p.y, (int)p.z);
    }

    int size = strlen(buf);

    char* ret = malloc(size + 1);
    strcpy_s(ret, size + 1, buf);
    free(buf);

    return ret;
}