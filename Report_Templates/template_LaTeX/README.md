# CS328 Homework LaTeX Template

This template for CS328 is adapted from https://github.com/ziqin/LaTeX-SUSTechHomework

We thank the original authors for their effort.

## Introduction

This is a LaTeX template designed for assignments in SUSTech. This The design style is inspired by [Henry's work](https://github.com/Henrycobaltech/SUSTechHomeworkTemplate).

## Usage

1. Download `SUSTechHomework.cls`, and place it in the same directory with your LaTeX document;
2. Set the `documentclass` to `SUSTechHomework`.

## Example

```latex
% !TEX TS-program = xelatex
% !TEX encoding = UTF-8
\documentclass[onecolumn,oneside]{SUSTechHomework}

\author{Name}
\sid{11110000}
\title{Homework 1}
\coursecode{CS101}
\coursename{Introduction to Computer Science}

\begin{document}
  \maketitle
  
  \section{Introduction}
  
  Hello World!

\end{document}
```

## Tips

- Currently you need to compile your document with XeLaTeX;
- You may enjoy using [Overleaf](https://www.overleaf.com), [Papeeria](https://www.papeeria.com), or other online LaTeX writing systems.