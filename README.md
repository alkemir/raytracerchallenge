Hello traveller. This repo contains my exploration of building a 3D renderer based off `The Ray Tracer Challenge` by Jamis Buck, a fun read.
It is not intended to be optimal, smart, clean, etc... just my ramblings as I learn something new, something old and something stolen. If
this is something that interests you, I invite you to follow along the commit history, it might save you having to do some typing (although
research has found it is essential for learning)

To show off the capabilities of the renderer, see this image which is the result of chapter 5:

![Basic rendering of a Sphere](/go/examples/chapter5/example.png?raw=true "BasicSphereRender")

Then, in chapter 6 we added light sources and their consequences according to the Phong model:

![Illuminated rendering of a Sphere](/go/examples/chapter6/example.png?raw=true "IlluminatedSphereRender")

In chapter 7 we created abstractions for cameras and worlds which makes setting up scenes way more easy:

![Scene of multiple Spheres](/go/examples/chapter7/example.png?raw=true "MultipledSpheresRender")

In chapter 8 we added shadows, which bring in a problem called acne:

![Scene with Acne](/go/examples/chapter8/example_fleas.png?raw=true "AcneSpheresRender")

But don't worry, we fixed it:

![Scene with Shadows](/go/examples/chapter8/example.png?raw=true "ShadowsSpheresRender")

In chapter 9 we implemented planes, so we no longer have to torture spheres out of shape to achieve this effect:

![Scene with Plane](/go/examples/chapter9/example.png?raw=true "PlaneSpheresRender")

Chapter 10 was all about adding patterns to leave behind the boring world of solid colors. But, of course, the first patterns are a bit boring:

![Boring Scene](/go/examples/chapter10/example_untransformed.png?raw=true "BoringPatternRender")

So we implemented support for patterns to be transformed, so there is some improvement, e.g. you are no longer confused by the middle sphere making a continuum with the floor.

![Transformed Scene](/go/examples/chapter10/example_transformed.png?raw=true "TransformedPatternRender")

Then we moved on to more interesting patterns

![Acned floor Scene](/go/examples/chapter10/example_fail.png?raw=true "FailedPatternRender")

Alas, I had a bug which reminded me of the acne from a couple of chapters ago. Had to use some colors to figure out whether it was a patterning issue or a shadowing one (the downsides of using black and white), so now thats taken care of. Lo, and behold:

![Patterned Scene](/go/examples/chapter10/example.png?raw=true "PatternedRender")

