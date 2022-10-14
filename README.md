# PiUrlSolution
Solution to the first two phases Sigmageek's Pi digits problem, along with additional material.

The `./python` subdirectory includes a solution for the first phase of the problem using Chudnovsky's algorithm optimized with binary splitting, its implementation was heavily inspired by [this one](https://www.craig-wood.com/nick/articles/pi-chudnovsky/).

`cmd/v1` tries to solve the second phase of the problem using calls to [pi.delivery's API](https://pi.delivery) that returns up to 1000 pi digits from an arbitrary starting point. But i quickly realized that this method was going to be way too slow since the maximum possible rate of calls was too small.

`cmd/v2` was my alternative way of doing this. It requires the user to download .ycd (very big) files containing 10^11 digits of pi each (those files come from Google's latest record breaking pi calculation) availible [here](https://storage.googleapis.com/pi100t/index.html). Uncompress them with [DigitViewer](https://github.com/Mysticial/DigitViewer) and then read them concurrently with the script in this subdirectory. This implementation unfortunately was not completed within the deadline of the challenge and still lacks a way to make the reading of the digit continuous between files.