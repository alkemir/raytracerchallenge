#include "matrix.h"

#pragma once

matrix translation(float x, float y, float z);
matrix scaling(float x, float y, float z);
matrix rotationX(float rad);
matrix rotationY(float rad);
matrix rotationZ(float rad);
matrix shearing(float xtoy, float xtoz, float ytox, float ytoz, float ztox,
                float ztoy);