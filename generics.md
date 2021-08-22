# A Summary of Go Generics Research

> Last updates: Aug 2021
>
> Author: Changkun Ou

This document summarizes a personal research on a selected list of
references and relevant discussions regarding the Go generics.

## Research Paper

Research is leading the world. Many good ideas are done by researchers.
One must read papers for digging the intrinsic properties of the generics design problem.

- [Cardelli and Wegner 1985] Cardelli, Luca, and Peter Wegner. "On understanding types, data abstraction, and polymorphism." ACM Computing Surveys (CSUR) 17.4 (1985): 471-523. http://lucacardelli.name/Papers/OnUnderstanding.A4.pdf
- [Stroustrup 1989] Bjarne Stroustrup. "Parameterized Types for C++." 1989. AT&T Bell Laboratories. https://www.usenix.org/legacy/publications/compsystems/1989/win_stroustrup.pdf
- [Morrisett 1995] Morrisett, Greg. "Compiling with types." CARNEGIE-MELLON UNIV PITTSBURGH PA SCHOOL OF COMPUTER SCIENCE, 1995. https://apps.dtic.mil/dtic/tr/fulltext/u2/a597747.pdf
- [Wickline et al 1998] Wickline, Philip, et al. "Modal types as staging specifications for run-time code generation." ACM Computing Surveys (CSUR) 30.3es (1998): 8-es. https://www.cs.cmu.edu/~fp/papers/sope98.pdf
- [Davies and Pfenning 2001] Davies, Rowan, and Frank Pfenning. "A modal analysis of staged computation." Journal of the ACM (JACM) 48.3 (2001): 555-604. https://www.cs.cmu.edu/~fp/papers/jacm00.pdf
- [Garcia and Ronald 2003] Garcia, Ronald, et al. "A comparative study of language support for generic programming." Proceedings of the 18th annual ACM SIGPLAN conference on Object-oriented programing, systems, languages, and applications. 2003. https://www.csie.ntu.edu.tw/~d00922011/java2/p115-garcia.pdf
- [Siek and Lumsdaine 2011] Siek, Jeremy G., and Andrew Lumsdaine. "A language for generic programming in the large." Science of Computer Programming 76.5 (2011): 423-465. https://www.sciencedirect.com/science/article/pii/S0167642308001123. Arxiv: https://arxiv.org/abs/0708.2255.
- [Greenman et al 2014] Greenman, Ben, Fabian Muehlboeck, and Ross Tate. "Getting F-bounded polymorphism into shape." *ACM SIGPLAN Notices* 49.6 (2014): 89-99. https://www.cs.cornell.edu/~fabianm/papers/shapes-pldi14-tr.pdf
- [Zhang et al 2015] Zhang, Yizhou, et al. "Lightweight, flexible object-oriented generics." Proceedings of the 36th ACM SIGPLAN Conference on Programming Language Design and Implementation. 2015. https://www.cs.cornell.edu/~yizhou/papers/genus-pldi2015.pdf
- [Zhang and Myers 2017] Zhang, Yizhou, and Andrew C. Myers. "Familia: unifying interfaces, type classes, and family polymorphism." *Proceedings of the ACM on Programming Languages* 1.OOPSLA (2017): 1-31. https://www.cs.cornell.edu/andru/papers/familia/familia.pdf
- [Griesemer et al 2020] Griesemer, Robert, et al. "Featherweight Go." arXiv preprint arXiv:2005.11710 (2020). https://arxiv.org/abs/2005.11710

## Go Proposals

- [Taylor 2010] Ian Lance Taylor. "Type Functions." golang/proposals, June 2010. https://github.com/golang/proposal/blob/master/design/15292/2010-06-type-functions.md
- [Taylor 2011] Ian Lance Taylor. "Generalized Types."golang/proposals, March 2011.  https://github.com/golang/proposal/blob/master/design/15292/2011-03-gen.md
- [Cox 2012] Russ Cox. "Alternatives to Dynamic Code Generation in Go." September 2012. https://docs.google.com/document/pub?id=1IXHI5Jr9k4zDdmUhcZImH59bOUK0G325J1FY6hdelcM
- [Taylor 2013a] Ian Lance Taylor. "Generalized Types In Go." golang/proposals, October 2013. https://github.com/golang/proposal/blob/master/design/15292/2013-10-gen.md
- [Taylor 2013b] Ian Lance Taylor. "Type Parameters in Go." golang/proposals, December 2013. https://github.com/golang/proposal/blob/master/design/15292/2013-12-type-params.md
- [Pike 2014] Rob Pike. "Go Generate." January 2014. http://golang.org/s/go1.4-generate
- [Mills 2016] Bryan C. Mills. "Compile-time Functions and First Class Types." golang/proposals, September 2016. https://github.com/golang/proposal/blob/master/design/15292/2016-09-compile-time-functions.md
- [Taylor 2016] Ian Lance Taylor. "Go should have generics." golang/proposals, January 2011. https://github.com/golang/proposal/blob/b571c3273d2c6988d24a22dd1c529387ff05962a/design/15292-generics.md Updated: April 2016. https://github.com/golang/proposal/blob/master/design/15292-generics.md
- [Cox 2018] Russ Cox. "Generics — Problem Overview." golang/proposals, August 27, 2018. https://github.com/golang/proposal/blob/master/design/go2draft-generics-overview.md
- [Taylor and Griesemer 2019] Ian Lance Taylor, Robert Griesemer. "Contracts — Draft Design." golang/proposals, August 27, 2018, Updated: July 31, 2019. https://github.com/golang/proposal/blob/master/design/go2draft-contracts.md
- [Taylor and Griesemer 2020] Ian Lance Taylor, Robert Griesemer. "Type Parameters - Draft Design
." golang/proposals, June 16, 2020, Updated: August 21, 2020. https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-type-parameters.md
- [Taylor and Griesemer 2020] Ian Lance Taylor, Robert Griesemer. "Type Switches on Type Parameters." golang/proposals, August 25, 2020. https://go-review.googlesource.com/c/proposal/+/250518
- [Taylor and Griesemer 2021] Ian Lance Taylor, Robert Griesemer. "Type Parameters Proposal." golang/proposals, March 19, 2021, Updated: August 20, 2021. https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md
- [Findley and Griesemer 2021] Rob Findley, Robert Griesemer. "Additions to go/ast and go/token to support parameterized functions and types." golang/proposals, Auguest 18, 2021. ttps://github.com/golang/proposal/blob/master/design/47781-parameterized-go-ast.md

## C++ Proposals

The proposals can be found in C++ Standards Committee Papers: http://www.open-std.org/jtc1/sc22/wg21/docs/papers/. Most of the proposals are done by Bjarne Stroustrup and can be found in http://www.stroustrup.com/WG21.html.

- [Stroustrup 1994] Stroustrup, Bjarne. The design and evolution of C++. Pearson Education India, 1994.
  + In this book, Bjarne discussed the decision of the origin template design.

C++ is also considered contracts but different than Go's contract design. Contracts are intended to communicate constraints on the values of runtime objects, with the parameters to a function being the most common case. A contract violation is always a bad thing, and the program exhibits undefined behavior.

- What Happened to C++20 Contracts? https://www.reddit.com/r/cpp/comments/cmk7ek/what_happened_to_c20_contracts/
- Simple Contracts for C++, http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2015/n4415.pdf

## Articles

- Russ Cox. "The Generic Dilemma." December 3, 2009. https://research.swtch.com/generic
- Go Community. Summary of Go Generics Discussions. *living document*. https://docs.google.com/document/d/1vrAy9gMpMoS3uaVphB32uVXX4pi-HnNjkMEgyAHX4N4/view#
- ExperienceReports https://github.com/golang/go/wiki/ExperienceReports#generics, https://github.com/golang/go/wiki/Go2GenericsFeedback
  + Many community experience reposrts can be found in this link. IMO: Most of the reports are trash.
- Jakub Cislo. C++ Concepts: Complete Overview

## Code base

There are some code base (not all) that is good for generics design:

- https://github.com/robpike/filter
- https://github.com/cheekybits/genny
- https://github.com/golang-collections

## Quotes

- Fancy algorithms are slow when n is small, and n is usually small. *- Rob Pike*

## Licnese

BSD-2-Clause

Copyright &copy; 2020-2021 [Changkun Ou](https://changkun.de)
