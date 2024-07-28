#include "intersection.h"

#include <assert.h>
#include <stdio.h>

#include "floats.h"

void test_intersectionConstructor() {
    sphere s = newSphere();

    intersection i = newIntersection(3.5, &s);

    assert(equalFloat(i.t, 3.5));
    assert(i.s == &s);
}

int main() {
    test_intersectionConstructor();
    printf("Intersection module tests successful\n");
}