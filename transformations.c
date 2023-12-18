#include "transformations.h"

matrix translation(float x, float y, float z) {
    return newMatrix(
        4, 4,
        (float *)(float[4][4]){1, 0, 0, x, 0, 1, 0, y, 0, 0, 1, z, 0, 0, 0, 1});
}

matrix scaling(float x, float y, float z) {
    return newMatrix(
        4, 4,
        (float *)(float[4][4]){x, 0, 0, 0, 0, y, 0, 0, 0, 0, z, 0, 0, 0, 0, 1});
}