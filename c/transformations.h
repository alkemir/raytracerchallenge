#include "matrix.h"

#pragma once

matrix translation(float x, float y, float z);
matrix scaling(float x, float y, float z);
matrix rotationX(float rad);
matrix rotationY(float rad);
matrix rotationZ(float rad);
matrix shearing(float xtoy, float xtoz, float ytox, float ytoz, float ztox,
                float ztoy);

// Fluent API
matrix translate(float x, float y, float z, matrix m);
matrix scale(float x, float y, float z, matrix m);
matrix rotateX(float rad, matrix m);
matrix rotateY(float rad, matrix m);
matrix rotateZ(float rad, matrix m);
matrix shear(float xtoy, float xtoz, float ytox, float ytoz, float ztox,
             float ztoy, matrix m);