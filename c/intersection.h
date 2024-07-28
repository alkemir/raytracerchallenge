#include "sphere.h"

#pragma once

typedef struct intersection_s {
    float t;
    sphere *s;
} intersection;

intersection newIntersection(float t, sphere *object);