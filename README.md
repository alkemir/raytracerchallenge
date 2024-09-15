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

Chapter 11 was about reflection and refraction. Reflection is a bit simpler, so I turned one of the balls into a mirror:

![Mirrored Scene](/go/examples/chapter11/example_reflection.png?raw=true "MirroredRender")

Then, as I moved in with refraction I thought I had a bug that had me scratching my head:

![Transparency Bugged Scene](/go/examples/chapter11/example_bug.png?raw=true "TransparentBugRender")

But it turned out it was just a weird interaction resulting of making a material highly reflective and refractive. So here is the refractive version only:

![Transparency Scene](/go/examples/chapter11/example_refraction.png?raw=true "TransparentRender")

Pretty cool huh? But still, something was missing. The Fresnel effect. We implemented an approximation of it by Schlick.

![Fresnel Scene](/go/examples/chapter11/example.png?raw=true "FresnelRender")

To break with the monotony and familiarity of my favourite scene, we added cubes in chapter 12. Also implemented multi-threaded rendering as I got bored of waiting:

![Cube Scene](/go/examples/chapter12/example_old.png?raw=true "CubeRender")

I got tired of the patterns in the background due to aliasing, so I implemented AA. My first attempt was doing 4x supersampling:

![Supersampled Scene](/go/examples/chapter12/example_4aa.png?raw=true "SupersampledRender")

It did help, but you can still see some patterning, so I went for stochastic raytracing, which worked, but I did not like that the borders seem noisier. Well, what can you do.

![Stochastic Scene](/go/examples/chapter12/example.png?raw=true "StochasticRender")

In chapter 13 we implemented cylinders and cones, both capped and uncapped:

![Cylinders and Cones Scene](/go/examples/chapter13/example.png?raw=true "CylinderConeRender")

To prove that I have absolutely no creativity nor sense of aesthetics, I created this perfectly symmetrical scene (except for AA noise) which is hard to interpret at first, but it is my rendition of hexane (ok, ok its just a mirrory hexagon) to showcase the groups that I implemented in chapter 14.

![Transluscent Hexane Scene](/go/examples/chapter14/example.png?raw=true "HexaneRender")

In chapter 15 we implement triangles, and I let my creativity go loose with this crystal:

![Transluscent Crystal Scene](/go/examples/chapter15/example.png?raw=true "CrystalRender")

Not much to see, but it's not over. Then we move into OBJ files; watch this beautiful teapot which took 4 hours to render:

![Teapot Scene](/go/examples/chapter15teapot/example.png?raw=true "TeapotRender")

