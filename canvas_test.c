#include "canvas.h"

#include <assert.h>
#include <stdio.h>
#include <string.h>

#include "tuple.h"

void test_canvasIsInitialized() {
    canvas c = newCanvas(10, 20);

    assert(c.width == 10);
    assert(c.height == 20);
    for (int x = 0; x < 10; x++) {
        for (int y = 0; y < 20; y++) {
            assert(equalTuple(newColor(0, 0, 0), getPixel(c, x, y)));
        }
    }
}

void test_canvasWrite() {
    canvas c = newCanvas(10, 20);
    tuple red = newColor(1, 0, 0);

    setPixel(c, 2, 3, red);

    assert(equalTuple(getPixel(c, 2, 3), red));
}

void test_canvasToPPMHeader() {
    canvas c = newCanvas(5, 3);

    char* canvasPPM = canvasToPPM(c);
    char* header = "P3\n5 3\n255\n";

    assert(strncmp(canvasPPM, header, strlen(header)) == 0);
}

void test_canvasToPPMPixeldata() {
    canvas c = newCanvas(5, 3);

    setPixel(c, 0, 0, newColor(1.5, 0, 0));
    setPixel(c, 2, 1, newColor(0, 0.5, 0));
    setPixel(c, 4, 2, newColor(-0.5, 0, 1));

    char* canvasPPM = canvasToPPM(c);
    assert(strcmp(canvasPPM,
                  "P3\n5 3\n255\n255 0 0 0 0 0 0 0 0 0 0 0 0 0 0\n0 0 0 0 0 0 "
                  "0 127 0 0 0 0 0 0 0\n0 0 0 0 0 0 0 0 0 0 0 0 0 0 255\n") ==
           0);
}

void test_canvasToPPMPixeldataLimited() {
    canvas c = newCanvas(10, 2);

    for (int x = 0; x < 10; x++) {
        for (int y = 0; y < 2; y++) {
            setPixel(c, x, y, newColor(1, 0.8, 0.6));
        }
    }

    char* canvasPPM = canvasToPPM(c);

    assert(strcmp(canvasPPM,
                  "P3\n10 2\n255\n255 204 153 255 204 153 255 204 153 255 204 "
                  "153 255 204 153 255 204\n153 255 204 153 255 204 153 255 "
                  "204 153 255 204 153\n255 204 153 255 204 153 255 204 153 "
                  "255 204 153 255 204 153 255 204\n153 255 204 153 255 204 "
                  "153 255 204 153 255 204 153\n") == 0);
}

int main() {
    test_canvasIsInitialized();
    test_canvasWrite();
    test_canvasToPPMHeader();
    test_canvasToPPMPixeldata();
    test_canvasToPPMPixeldataLimited();
}