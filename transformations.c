#include "transformations.h"

#include <math.h>

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

matrix rotateX(float rad) {
    return newMatrix(
        4, 4,
        (float *)(float[4][4]){1, 0, 0, 0, 0, cos(rad), -1 * sin(rad), 0, 0,
                               sin(rad), cos(rad), 0, 0, 0, 0, 1});
}

matrix rotateY(float rad) {
    return newMatrix(
        4, 4,
        (float *)(float[4][4]){cos(rad), 0, sin(rad), 0, 0, 1, 0, 0,
                               -1 * sin(rad), 0, cos(rad), 0, 0, 0, 0, 1});
}

matrix rotateZ(float rad) {
    return newMatrix(
        4, 4,
        (float *)(float[4][4]){cos(rad), -1 * sin(rad), 0, 0, sin(rad),
                               cos(rad), 0, 0, 0, 0, 1, 0, 0, 0, 0, 1});
}