#include "ray.h"

#include <assert.h>
#include <stdio.h>

#include "sphere.h"
#include "tuple.h"

void test_rayConstructor() {
    tuple origin = newPoint(1, 2, 3);
    tuple direction = newVector(4, 5, 6);

    ray r = newRay(origin, direction);

    assert(equalTuple(r.origin, origin));
    assert(equalTuple(r.direction, direction));
}

void test_rayProject() {
    ray r = newRay(newPoint(2, 3, 4), newVector(1, 0, 0));

    tuple p1 = project(r, 0);
    tuple p2 = project(r, 1);
    tuple p3 = project(r, -1);
    tuple p4 = project(r, 2.5);

    assert(equalTuple(p1, newPoint(2, 3, 4)));
    assert(equalTuple(p2, newPoint(3, 3, 4)));
    assert(equalTuple(p3, newPoint(1, 3, 4)));
    assert(equalTuple(p4, newPoint(4.5, 3, 4)));
}

void test_rayIntersect() {
    ray r = newRay(newPoint(0, 0, -5), newVector(0, 0, 1));
    sphere s = newSphere();
    float xs[2];

    int points = intersect(r, s, xs);

    assert(points == 2);
    assert(equalFloat(xs[0], 4.0));
    assert(equalFloat(xs[1], 6.0));
}

void test_rayIntersect_tangent() {
    ray r = newRay(newPoint(0, 1, -5), newVector(0, 0, 1));
    sphere s = newSphere();
    float xs[2];

    int points = intersect(r, s, xs);

    assert(points == 2);
    assert(equalFloat(xs[0], 5.0));
    assert(equalFloat(xs[1], 5.0));
}

void test_rayIntersect_miss() {
    ray r = newRay(newPoint(0, 2, -5), newVector(0, 0, 1));
    sphere s = newSphere();
    float xs[2];

    int points = intersect(r, s, xs);

    assert(points == 0);
}

void test_rayIntersect_inside() {
    ray r = newRay(newPoint(0, 0, 0), newVector(0, 0, 1));
    sphere s = newSphere();
    float xs[2];

    int points = intersect(r, s, xs);

    assert(points == 2);
    assert(equalFloat(xs[0], -1.0));
    assert(equalFloat(xs[1], 1.0));
}

void test_rayIntersect_behind() {
    ray r = newRay(newPoint(0, 0, 5), newVector(0, 0, 1));
    sphere s = newSphere();
    float xs[2];

    int points = intersect(r, s, xs);

    assert(points == 2);
    assert(equalFloat(xs[0], -6.0));
    assert(equalFloat(xs[1], -4.0));
}

int main() {
    test_rayConstructor();
    test_rayProject();
    test_rayIntersect();
    test_rayIntersect_tangent();
    test_rayIntersect_inside();
    test_rayIntersect_behind();
    printf("Ray module tests successful\n");
}