# Drawings

Small Go programs drawing cool stuff

## sierpinski

Unusual way of drawing Sierpinski triangles.

Based on a probabilistic method, described by a Numberphile video: https://www.youtube.com/watch?v=kbKtFN71Lfs

Drawing with dots was convenient, since Go's standard library does't come with methods to draw lines or shapes.

I could also use a Pascal triangle, or the rule 60 elementary celular automation: http://mathworld.wolfram.com/Rule60.html

## line

Talking about drawing a line between points A and B: that's super interesting

## mandelbrot

Basic black and white Mandelbrot set tracing.
Did you know Go had built-in support for complex numbers?

Specify the complex plane region to draw from: `./mandelbrot -x1 -0.25 -x2 0.04 -y1 -1 -y2 -0.85`

## stegano

Not really drawing but sort of manipulating images
You can hide a text inside an image. Works for BMP, not sure about other formats.

Encode:
`./stegano in.bmp out.bmp '"Text to hide"'`

Decode:
`./stegano out.bmp`

Note: if you consider using Lenna as the support picture for this steganography example as sexist, please feel free to use any other picture such as a man, an animal or an object like a car or a landscape. Be sure to use the BMP (Bitmap) format though.

# Upcoming

- Dragon curve
- Barycentric subdivisions (triangles): http://drorbn.net/AcademicPensieve/2010-06/nb/BarycentricSubdivision.pdf
- Pythagoras tree