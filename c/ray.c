#include "ray.h"

#include <math.h>

#include "tuple.h"

ray newRay(tuple origin, tuple direction) {
    ray r = {origin, direction};
    return r;
};

tuple project(ray r, float distance) {
    return addTuple(r.origin, multiplyTuple(r.direction, distance));
}

int intersect(ray r, sphere s, float *out) {
    tuple sphereToRay = substractTuple(r.origin, newPoint(0, 0, 0));

    float a = dot(r.direction, r.direction);
    float b = 2 * dot(r.direction, sphereToRay);
    float c = dot(sphereToRay, sphereToRay) - 1;

    float discriminant = b * b - 4 * a * c;

    if (discriminant < 0) {
        return 0;
    }

    out[0] = (-1 * b - sqrt(discriminant)) / (2 * a);
    out[1] = (-1 * b + sqrt(discriminant)) / (2 * a);
    return 2;
}