CC = clang

all : test

.PHONY : clean test all

tuple_test.exe : tuple_test.o floats.o
	$(CC) -o tuple_test.exe tuple_test.o tuple.o floats.o

matrix_test.exe : matrix_test.o matrix.o floats.o tuple.o
	$(CC) -o matrix_test.exe matrix_test.o matrix.o floats.o tuple.o

canvas_test.exe : canvas_test.o canvas.o tuple.o floats.o
	$(CC) -o canvas_test.exe canvas_test.o canvas.o tuple.o floats.o

transformations_test.exe : transformations_test.o tuple.o matrix.o floats.o transformations.o
	$(CC) -o transformations_test.exe transformations_test.o tuple.o matrix.o floats.o transformations.o

ray_test.exe : ray_test.o ray.o tuple.o floats.o sphere.o
	$(CC) -o ray_test.exe ray_test.o ray.o tuple.o floats.o sphere.o

intersection_test.exe : intersection_test.o sphere.o intersection.o floats.o
	$(CC) -o intersection_test.exe intersection_test.o sphere.o intersection.o floats.o

test : matrix_test.exe tuple_test.exe canvas_test.exe transformations_test.exe ray_test.exe intersection_test.exe
	matrix_test.exe
	tuple_test.exe
	canvas_test.exe
	transformations_test.exe
	ray_test.exe
	intersection_test.exe


clean :
	rm *.o
	rm *.exe