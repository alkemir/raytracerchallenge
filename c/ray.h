#include "sphere.h"
#include "tuple.h"

#pragma once

typedef struct ray_s {
    tuple origin;
    tuple direction;
} ray;

ray newRay(tuple origin, tuple direction);

tuple project(ray r, float distance);
int intersect(ray r, sphere s, float *out);