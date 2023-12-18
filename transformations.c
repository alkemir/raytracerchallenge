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

matrix rotationX(float rad) {
    return newMatrix(
        4, 4,
        (float *)(float[4][4]){1, 0, 0, 0, 0, cos(rad), -1 * sin(rad), 0, 0,
                               sin(rad), cos(rad), 0, 0, 0, 0, 1});
}

matrix rotationY(float rad) {
    return newMatrix(
        4, 4,
        (float *)(float[4][4]){cos(rad), 0, sin(rad), 0, 0, 1, 0, 0,
                               -1 * sin(rad), 0, cos(rad), 0, 0, 0, 0, 1});
}

matrix rotationZ(float rad) {
    return newMatrix(
        4, 4,
        (float *)(float[4][4]){cos(rad), -1 * sin(rad), 0, 0, sin(rad),
                               cos(rad), 0, 0, 0, 0, 1, 0, 0, 0, 0, 1});
}

matrix shearing(float xtoy, float xtoz, float ytox, float ytoz, float ztox,
                float ztoy) {
    return newMatrix(4, 4,
                     (float *)(float[4][4]){1, xtoy, xtoz, 0, ytox, 1, ytoz, 0,
                                            ztox, ztoy, 1, 0, 0, 0, 0, 1});
}

matrix translate(float x, float y, float z, matrix m) {
    matrix t = translation(x, y, z);
    return multiplyMatrices(t, m);
}

matrix scale(float x, float y, float z, matrix m) {
    matrix s = scaling(x, y, z);
    return multiplyMatrices(s, m);
}

matrix rotateX(float rad, matrix m) {
    matrix r = rotationX(rad);
    return multiplyMatrices(r, m);
}

matrix rotateY(float rad, matrix m) {
    matrix r = rotationY(rad);
    return multiplyMatrices(r, m);
}

matrix rotateZ(float rad, matrix m) {
    matrix r = rotationZ(rad);
    return multiplyMatrices(r, m);
}

matrix shear(float xtoy, float xtoz, float ytox, float ytoz, float ztox,
             float ztoy, matrix m) {
    matrix s = shearing(xtoy, xtoz, ytox, ytoz, ztox, ztoy);
    return multiplyMatrices(s, m);
}